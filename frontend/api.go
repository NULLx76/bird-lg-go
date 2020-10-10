package main

import (
	"encoding/json"
	"github.com/NULLx76/bird-lg-go/api/proxy"
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"
)

var baseURL string = "http://localhost:8000"

func get(path string) (int, []byte, error) {
	status, body, err := fasthttp.Get(nil, baseURL+path)
	if err != nil {
		return 0, nil, errors.Wrap(err, "error in get from api")
	}
	if status != fasthttp.StatusOK {
		return 0, nil, errors.Errorf("status: %s, msg: %s", status, string(body))
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

func GetSummary(server string) (proxy.SummaryTable, error) {
	_, body, err := get("/server/" + server)
	if err != nil {
		return nil, errors.Wrap(err, "error getting summary")
	}

	var summ proxy.SummaryTable
	if err := json.Unmarshal(body, &summ); err != nil {
		return nil, err
	}

	return summ, nil
}
