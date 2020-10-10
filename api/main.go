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
		//fmt.Println(Summary(cfg.birdServers[server]))
		//fmt.Println(Details(cfg.birdServers[server], "icez"))
		fmt.Println(Route(cfg.birdServers[server], "172.20.0.53", true))
	}
}
