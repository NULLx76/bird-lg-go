package main

import (
	"github.com/NULLx76/bird-lg-go/pkg/proxy"
	"github.com/NULLx76/bird-lg-go/services/frontend/templates"
	"github.com/go-chi/chi"
	"net/http"
)

func mainPageHandler(w http.ResponseWriter, _ *http.Request) {
	servers, err := GetServers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	summaries := make(map[string]proxy.SummaryTable)
	for i := range servers {
		sum, err := GetSummary(servers[i])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		summaries[servers[i]] = sum
	}

	p := &templates.MainPage{
		Summaries: summaries,
	}
	templates.WritePageTemplate(w, p)
}

func peerPageHandler(w http.ResponseWriter, r *http.Request) {
	server := chi.URLParam(r, "server")
	peer := chi.URLParam(r, "peer")

	details, err := GetDetails(server, peer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	p := &templates.PeerPage{
		Peer: details,
	}
	templates.WritePageTemplate(w, p)
}

func routePageHandler(w http.ResponseWriter, r *http.Request) {
	server := chi.URLParam(r, "server")
	ip := chi.URLParam(r, "ip")

	allQ := r.URL.Query().Get("all")
	all := false
	if allQ == "1" || allQ == "true" {
		all = true
	}

	route, err := GetRoute(server, ip, all)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	p := &templates.RoutePage{
		Route: route,
	}

	templates.WritePageTemplate(w, p)
}
