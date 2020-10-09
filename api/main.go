package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

type config struct {
	birdServers []string
}

func main() {
	cfg := config{
		birdServers: []string{"http://dn42:8000"},
	}

	log.Tracef("config %v", cfg)

	for server := range cfg.birdServers {
		fmt.Println(Summary(cfg.birdServers[server]))
	}
}
