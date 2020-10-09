package main

import (
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
)

func queryBackend(host, command string) (string, error) {
	uri := host + "/bird" + "?q=" + url.QueryEscape(command)

	res, err := http.Get(uri)
	if err != nil {
		return "", err
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	str := string(b)

	log.Tracef("Backend result: %v", str)

	return str, nil
}
