package main

import (
	logger "github.com/chi-middleware/logrus-logger"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Config struct {
	ListenAddr  string
	BirdServers map[string]string
}

func main() {
	cfg := &Config{
		ListenAddr: ":8000",
		BirdServers: map[string]string{
			"xirion": "http://dn42:8000",
		},
	}

	s := NewRoutes(cfg)

	r := chi.NewRouter()
	r.Use(logger.Logger("router", log.StandardLogger()))
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/servers", s.GetServers)
	r.Get("/server/{server}", s.GetSummary)
	r.Get("/server/{server}/details/{peer}", s.GetDetails)
	r.Get("/server/{server}/route/{route}", s.GetRoute)

	log.Infof("starting the server at %s ...", cfg.ListenAddr)

	log.Fatal(http.ListenAndServe(cfg.ListenAddr, r))
}
