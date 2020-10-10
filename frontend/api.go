package main

import (
	"encoding/json"
	"fmt"
	"github.com/NULLx76/bird-lg-go/api/proxy"
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"
)

var baseURL = "http://localhost:8000"

func get(path string) (int, []byte, error) {
	status, body, err := fasthttp.Get(nil, baseURL+path)
	if err != nil {
		return 0, nil, errors.Wrap(err, "error in get from api")
	}
	if status != fasthttp.StatusOK {
		return 0, nil, errors.Errorf("status: %d, msg: %s", status, string(body))
	}

	return status, body, nil
}

func GetServers() ([]string, error) {
	_, body, err := get("/servers")
	if err != nil {
		return nil, errors.Wrap(err, "error getting servers")
	}

	var servers []string
	if err := json.Unmarshal(body, &servers); err != nil {
		return nil, err
	}

	return servers, nil
}

const getSummaryFmt string = "/server/%s"

func GetSummary(server string) (proxy.SummaryTable, error) {
	url := fmt.Sprintf(getSummaryFmt, server)
	_, body, err := get(url)
	if err != nil {
		return nil, errors.Wrap(err, "error getting summary")
	}

	var summ proxy.SummaryTable
	if err := json.Unmarshal(body, &summ); err != nil {
		return nil, err
	}

	return summ, nil
}

const getDetailsFmt string = "/server/%s/details/%s"

func GetDetails(server, peer string) (*proxy.PeerDetails, error) {
	url := fmt.Sprintf(getDetailsFmt, server, peer)
	_, body, err := get(url)
	if err != nil {
		return nil, errors.Wrap(err, "error getting summary")
	}

	var details proxy.PeerDetails
	if err := json.Unmarshal(body, &details); err != nil {
		return nil, err
	}

	return &details, nil
}
