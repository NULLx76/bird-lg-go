package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"runtime"
	"strings"
)

// Wrapper of traceroute, IPv4
func tracerouteIPv4Wrapper(httpW http.ResponseWriter, httpR *http.Request) {
	tracerouteRealHandler(false, httpW, httpR)
}

// Wrapper of traceroute, IPv6
func tracerouteIPv6Wrapper(httpW http.ResponseWriter, httpR *http.Request) {
	tracerouteRealHandler(true, httpW, httpR)
}

func tracerouteTryExecute(cmd []string, args [][]string) ([]byte, string) {
	var output []byte
	var errString = ""
	for i := range cmd {
		var err error
		var cmdCombined = cmd[i] + " " + strings.Join(args[i], " ")

		instance := exec.Command(cmd[i], args[i]...)
		output, err = instance.CombinedOutput()
		if err == nil {
			return output, ""
		}
		errString += fmt.Sprintf("+ (Try %d) %s\n%s\n\n", i+1, cmdCombined, output)
	}
	return nil, errString
}

// Real handler of traceroute requests
func tracerouteRealHandler(useIPv6 bool, httpW http.ResponseWriter, httpR *http.Request) {
	query := httpR.URL.Query().Get("q")
	query = strings.TrimSpace(query)
	if query == "" {
		invalidHandler(httpW, httpR)
	} else {
		var result []byte
		var errString string
		if runtime.GOOS == "freebsd" || runtime.GOOS == "netbsd" {
			if useIPv6 {
				result, errString = tracerouteTryExecute(
					[]string{
						"traceroute6",
						"traceroute",
					},
					[][]string{
						{"-q1", "-w1", query},
						{"-q1", "-w1", query},
					},
				)
			} else {
				result, errString = tracerouteTryExecute(
					[]string{
						"traceroute",
						"traceroute6",
					},
					[][]string{
						{"-q1", "-w1", query},
						{"-q1", "-w1", query},
					},
				)
			}
		} else if runtime.GOOS == "openbsd" {
			if useIPv6 {
				result, errString = tracerouteTryExecute(
					[]string{
						"traceroute6",
						"traceroute",
					},
					[][]string{
						{"-q1", "-w1", query},
						{"-q1", "-w1", query},
					},
				)
			} else {
				result, errString = tracerouteTryExecute(
					[]string{
						"traceroute",
						"traceroute6",
					},
					[][]string{
						{"-A", "-q1", "-w1", query},
						{"-A", "-q1", "-w1", query},
					},
				)
			}
		} else if runtime.GOOS == "linux" {
			if useIPv6 {
				result, errString = tracerouteTryExecute(
					[]string{
						"traceroute",
						"traceroute",
						"traceroute",
						"traceroute",
					},
					[][]string{
						{"-6", "-q1", "-N32", "-w1", query},
						{"-4", "-q1", "-N32", "-w1", query},
						// For Busybox traceroute which doesn't support simultaneous requests
						{"-6", "-q1", "-w1", query},
						{"-4", "-q1", "-w1", query},
					},
				)
			} else {
				result, errString = tracerouteTryExecute(
					[]string{
						"traceroute",
						"traceroute",
						"traceroute",
						"traceroute",
					},
					[][]string{
						{"-4", "-q1", "-N32", "-w1", query},
						{"-6", "-q1", "-N32", "-w1", query},
						// For Busybox traceroute which doesn't support simultaneous requests
						{"-4", "-q1", "-w1", query},
						{"-6", "-q1", "-w1", query},
					},
				)
			}
		} else {
			httpW.WriteHeader(http.StatusInternalServerError)
			_, _ = httpW.Write([]byte("traceroute not supported on this node.\n"))
			return
		}
		if errString != "" {
			httpW.WriteHeader(http.StatusInternalServerError)
			_, _ = httpW.Write([]byte("traceroute returned error:\n\n" + errString))
		}
		if result != nil {
			_, _ = httpW.Write(result)
		}
	}
}
