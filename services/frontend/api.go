package main

import (
	"encoding/json"
	"fmt"
	"github.com/NULLx76/bird-lg-go/pkg/proxy"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

func get(path string) ([]byte, error) {
	resp, err := http.Get(cfg.ApiUrl + path)
	if err != nil {
		return nil, errors.Wrap(err, "error in get from api")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't read return body")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.Errorf("%d: %s", resp.StatusCode, string(body))
	}

	return body, nil
}

func GetServers() ([]string, error) {
	body, err := get("/servers")
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
	body, err := get(url)
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
	body, err := get(url)
	if err != nil {
		return nil, errors.Wrap(err, "error getting summary")
	}

	var details proxy.PeerDetails
	if err := json.Unmarshal(body, &details); err != nil {
		return nil, err
	}

	return &details, nil
}

const getRouteFmt string = "/server/%s/route/%s?all=%t"

func GetRoute(server, address string, all bool) (*proxy.RouteDetails, error) {
	url := fmt.Sprintf(getRouteFmt, server, address, all)
	body, err := get(url)
	if err != nil {
		return nil, errors.Wrap(err, "error getting summary")
	}

	var route proxy.RouteDetails
	if err := json.Unmarshal(body, &route); err != nil {
		return nil, err
	}

	return &route, nil
}
