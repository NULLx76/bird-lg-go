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

const (
	commandSummary  = "show protocols"
	commandDetails  = "show protocols all %s"
	commandRoute    = "show route for %s"
	commandRouteAll = "show route for %s all"
)

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

func Summary(server string) (SummaryTable, error) {
	query, err := queryBackend(server, commandSummary)
	if err != nil {
		return nil, err
	}

	tbl, err := parseSummaryTable(query)
	if err != nil {
		return nil, err
	}

	return tbl, nil
}

func Details(server, peer string) (*PeerDetails, error) {
	cmd := fmt.Sprintf(commandDetails, peer)
	query, err := queryBackend(server, cmd)
	if err != nil {
		return nil, err
	}

	details, err := parsePeerDetails(query)
	if err != nil {
		return nil, err
	}

	return details, nil
}

func Route(server, address string, all bool) (*RouteDetails, error) {
	var cmd string
	if all {
		cmd = fmt.Sprintf(commandRouteAll, address)
	} else {
		cmd = fmt.Sprintf(commandRoute, address)
	}

	query, err := queryBackend(server, cmd)
	if err != nil {
		return nil, err
	}

	details := RouteDetails{
		address: address,
		details: query,
	}

	return &details, nil
}
