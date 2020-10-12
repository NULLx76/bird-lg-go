package main

import (
	logger "github.com/chi-middleware/logrus-logger"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"os"
)

type Config struct {
	ListenAddr string            `yaml:"listen"`
	Servers    map[string]string `yaml:"servers"`
}

func readConfigFromFile(filename string) (c Config, err error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return Config{}, err
	}

	err = yaml.Unmarshal(file, &c)
	if err != nil {
		return Config{}, err
	}

	return c, nil
}

func main() {
	file := "./config.yml"
	if fileEnv := os.Getenv("CONFIG_FILE"); fileEnv != "" {
		file = fileEnv
	}

	cfg, err := readConfigFromFile(file)
	if err != nil {
		log.Fatalf("Error encountered while reading config: %v", err)
	}

	s := NewRoutes(&cfg)

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
