package main

import (
	"github.com/caarlos0/env/v6"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)
import "github.com/chi-middleware/logrus-logger"

//go:generate qtc -dir=templates
//go:generate yarn build
//go:generate yarn minify

type Config struct {
	ApiUrl     string `env:"API_URL" envDefault:"http://0.0.0.0:8000"`
	ListenAddr string `env:"ADDRESS" envDefault:":8080"`
}

var cfg Config

func main() {
	cfg = Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("%+v\n", err)
	}

	r := chi.NewRouter()
	r.Use(logger.Logger("router", log.StandardLogger()))
	r.Use(middleware.GetHead)
	r.Use(middleware.Recoverer)

	r.Get("/", mainPageHandler)
	r.Get("/{server}/details/{peer}", peerPageHandler)
	r.Get("/{server}/route/{ip}", routePageHandler)
	r.Post("/{server}/route", routeFormHandler)

	// Static files
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "./build"))
	FileServer(r, "/static", filesDir)

	log.Infof("starting the server at %s ...", cfg.ListenAddr)

	log.Fatal(http.ListenAndServe(cfg.ListenAddr, r))
}

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		log.Fatal("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
