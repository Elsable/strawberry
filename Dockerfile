ARG TZ=America/Chicago

FROM golang:1.9-alpine as build-backend

ARG TZ

RUN go version

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN \
    apk add --no-cache --update tzdata git &&\
    cp /usr/share/zoneinfo/$TZ /etc/localtime &&\
    go get -u gopkg.in/alecthomas/gometalinter.v1 && \
    ln -s /go/bin/gometalinter.v1 /go/bin/gometalinter && \
    gometalinter --install --force

WORKDIR /go/src/github.com/andrievsky/strawberry

ADD app /go/src/github.com/andrievsky/strawberry/app
ADD vendor /go/src/github.com/andrievsky/strawberry/vendor
ADD .git /go/src/github.com/andrievsky/strawberry/.git

RUN cd app && go test ./...

RUN gometalinter --disable-all --deadline=300s --vendor --enable=vet --enable=vetshadow --enable=golint \
    --enable=staticcheck --enable=ineffassign --enable=goconst --enable=errcheck --enable=unconvert \
    --enable=deadcode  --enable=gosimple --exclude=test --exclude=mock --exclude=vendor ./...

RUN \
    version=$(git rev-parse --abbrev-ref HEAD)-$(git describe --abbrev=7 --always --tags)-$(date +%Y%m%d-%H:%M:%S) && \
    echo "version $version" && \
    go build -o strawberry -ldflags "-X main.revision=${version} -s -w" ./app

# FROM node:9.4-alpine as build-frontend
#
# ADD webapp /srv/webapp
# RUN apk add --no-cache --update git python make g++
# RUN \
#     cd /srv/webapp && \
#     npm i --production && npm run build


FROM alpine:3.7

ARG TZ

COPY --from=build-backend /go/src/github.com/andrievsky/strawberry/strawberry /srv/

RUN \
    apk add --update --no-cache tzdata && \
    cp /usr/share/zoneinfo/$TZ /etc/localtime &&\
    adduser -s /bin/bash -D -u 1001 app && \
    chown -R app:app /srv

EXPOSE 8080
USER app
WORKDIR /srv
VOLUME ["/srv/docroot"]

ENTRYPOINT ["/srv/strawberry"]