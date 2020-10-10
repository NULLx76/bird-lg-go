package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Config struct {
	birdServers map[string]string
}

func main() {
	cfg := &Config{
		birdServers: map[string]string{
			"xirion": "http://dn42:8000",
		},
	}

	s := NewRoutes(cfg)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/servers", s.GetServers)
	r.Get("/server/{server}", s.GetSummary)
	r.Get("/server/{server}/details/{peer}", s.GetDetails)
	r.Get("/server/{server}/route/{route}", s.GetRoute)

	log.Fatal(http.ListenAndServe(":8000", r))
}
