package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Original commands:
// 		"summary":         "show protocols",
//		"detail":          "show protocols all %s",
//		"route":           "show route for %s",
//		"route_all":       "show route for %s all",
//		"route_where":     "show route where net ~ [ %s ]",
//		"route_where_all": "show route where net ~ [ %s ] all",
//		"route_generic":   "show route %s",
//		"generic":         "show %s",
//		"traceroute":      "%s",

func queryBackend(host, command string) (string, error) {
	uri := host + "/bird" + "?q=" + url.QueryEscape(command)
	log.Tracef("Querying: %v", uri)

	res, err := http.Get(uri)
	if err != nil {
		return "", err
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	str := string(b)

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("host: %v, status: %v, command: %v, msg: %v", host, res.StatusCode, command, str)
	}

	return str, nil
}

func Summary(server string) (ProtocolTable, error) {
	query, err := queryBackend(server, "show protocols")
	if err != nil {
		return nil, err
	}

	tbl, err := parseProtocolTable(query)
	if err != nil {
		return nil, err
	}

	return tbl, nil
}
