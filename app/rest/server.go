package rest

import (
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/andrievsky/strawberry/app/store"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

// Server is a rest with store
type Server struct {
	Service *store.Service
	Version string
}

//Run the lister and request's router, activate rest server
func (s Server) Run() {
	log.Printf("[INFO] activate rest server")

	router := chi.NewRouter()
	router.Use(middleware.RealIP, Recoverer)
	router.Use(middleware.Throttle(1000), middleware.Timeout(60*time.Second))
	router.Use(AppInfo("strawberry", s.Version), Ping)

	router.Route("/api/v1", func(r chi.Router) {
		r.Use(Logger())
		r.Get("/resource/{id}", s.getResourceCtrl)
		r.Post("/resource", s.postResourceCtrl)
		r.Get("/list", s.getListCtrl)
	})

	router.Get("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		render.PlainText(w, r, "User-agent: *\nDisallow: /api/\n")
	})

	s.fileServer(router, "/", http.Dir(filepath.Join(".", "docroot")))

	log.Fatal(http.ListenAndServe(":8080", router))
}

// GET /v1/resource
func (s Server) getResourceCtrl(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	res, err := s.Service.Get(id)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, JSON{"error": err.Error()})
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, JSON{"resource": res})
}

// POST /v1/resource
func (s Server) postResourceCtrl(w http.ResponseWriter, r *http.Request) {
	res := store.Resource{}
	if err := render.DecodeJSON(r.Body, &res); err != nil {
		log.Printf("[WARN] can't bind resource %v", res)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, JSON{"error": err.Error()})
		return
	}

	resourceID, err := s.Service.Create(&res)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, JSON{"error": err.Error()})
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, JSON{"resourceID": resourceID})
}

// GET /v1/list
func (s Server) getListCtrl(w http.ResponseWriter, r *http.Request) {
	list, err := s.Service.List(0)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, JSON{"error": err.Error()})
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, JSON{"list": list})
}

// serves static files from ./docroot
func (s Server) fileServer(r chi.Router, path string, root http.FileSystem) {
	log.Printf("[INFO] run file server for %s", root)
	fs := http.StripPrefix(path, http.FileServer(root))
	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// don't show dirs, just serve files
		if strings.HasSuffix(r.URL.Path, "/") && len(r.URL.Path) > 1 && r.URL.Path != "/show/" {
			http.NotFound(w, r)
			return
		}
		fs.ServeHTTP(w, r)
	}))
}
