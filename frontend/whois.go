package main

import (
	"io/ioutil"
	"net"
)

// Send a whois request
func whois(s string) string {
	conn, err := net.Dial("tcp", setting.whoisServer+":43")
	if err != nil {
		return err.Error()
	}
	defer conn.Close()

	_, err = conn.Write([]byte(s + "\r\n"))
	if err != nil {
		return err.Error()
	}

	result, err := ioutil.ReadAll(conn)
	if err != nil {
		return err.Error()
	}
	return string(result)
}
