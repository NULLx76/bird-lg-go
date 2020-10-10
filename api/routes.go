package main

import (
	"bird-lg-go/api/comm"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
	"strings"
)

type Routes struct {
	cfg *Config
}

func NewRoutes(cfg *Config) *Routes {
	return &Routes{
		cfg,
	}
}

func (s *Routes) GetSummary(w http.ResponseWriter, r *http.Request) {
	server := chi.URLParam(r, "server")

	bird := s.cfg.birdServers[server]
	if bird == "" {
		http.Error(w, "Invalid server", http.StatusBadRequest)
		return
	}

	summ, err := comm.Summary(bird)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.Render(w, r, summ)
}

func (s *Routes) GetDetails(w http.ResponseWriter, r *http.Request) {
	server := chi.URLParam(r, "server")
	peer := chi.URLParam(r, "peer")

	bird := s.cfg.birdServers[server]
	if bird == "" {
		http.Error(w, "Invalid server", http.StatusBadRequest)
		return
	}

	details, err := comm.Details(bird, peer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.Render(w, r, details)
}

func (s *Routes) GetRoute(w http.ResponseWriter, r *http.Request) {
	server := chi.URLParam(r, "server")
	route := chi.URLParam(r, "route")

	allParam := r.URL.Query().Get("all")
	all := false
	if allParam == "1" || strings.ToLower(allParam) == "true" {
		all = true
	}

	bird := s.cfg.birdServers[server]
	if bird == "" {
		http.Error(w, "Invalid server", http.StatusBadRequest)
		return
	}

	routeDetails, err := comm.Route(bird, route, all)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.Render(w, r, routeDetails)
}
