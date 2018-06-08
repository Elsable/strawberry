package main

import (
	"fmt"
	"github.com/andrievsky/strawberry/app/rest"
	"github.com/andrievsky/strawberry/app/store"
	"github.com/jessevdk/go-flags"
	"log"
	"os"
)

var opts struct {
	Engine string `short:"e" long:"engine" env:"ENGINE" description:"storage engine" choice:"MEMORY" choice:"BOLT" default:"MEMORY"`
	BoltDB string `long:"bolt" env:"BOLT_FILE" default:"/tmp/secrets.bd" description:"boltdb file"`
	Dbg    bool   `long:"dbg" description:"debug mode"`
}

var revision string

func main() {

	if _, err := flags.Parse(&opts); err != nil {
		os.Exit(1)
	}
	fmt.Printf("secrets %s\n", revision)

	log.SetFlags(log.Ldate | log.Ltime)
	if opts.Dbg {
		log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
	}

	engine := getEngine(opts.Engine, opts.BoltDB)
	service := store.Service{
		Engine: engine,
	}
	server := rest.Server{
		Service: &service,
		Version: revision,
	}
	server.Run()
}

func getEngine(engineType string, boltFile string) store.Engine {
	switch engineType {
	case "MEMORY":
		return store.New()
		/*case "BOLT":
		boltStore, err := store.NewBolt(boltFile, time.Minute*5)
		if err != nil {
			log.Fatalf("[ERROR] can't open db, %v", err)
		}
		return boltStore
		*/
	}
	log.Fatalf("[ERROR] unknown engine type %s", engineType)
	return nil
}
