package main

import log "github.com/sirupsen/logrus"

type config struct {
	birdServers []string
}

func main() {
	cfg := config{
		birdServers: []string{"http://dn42:8000"},
	}

	log.Infof("config %v", cfg)

	for server := range cfg.birdServers {
		log.Info(queryBackend(cfg.birdServers[server], "show protocols"))
	}
}
