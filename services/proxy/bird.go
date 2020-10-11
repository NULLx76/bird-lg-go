package main

import (
	"bufio"
	"bytes"
	"io"
	"net"
	"net/http"
	"time"
)

// Read a line from bird socket, removing preceding status number, output it.
// Returns if there are more lines.
func birdReadln(bird *bufio.Reader, w io.Writer) bool {
	c, err := bird.ReadBytes('\n')
	if err != nil {
		panic(err)
	}

	if len(c) > 4 && isNumeric(c[0]) && isNumeric(c[1]) && isNumeric(c[2]) && isNumeric(c[3]) {
		c = c[5:]

		bytes.TrimSpace(c)
		if _, err := w.Write(c); err != nil {
			panic(err)
		}

		return c[0] != byte('0') && c[0] != byte('8') && c[0] != byte('9')
	} else {
		bytes.TrimSpace(c)
		if _, err := w.Write(c); err != nil {
			panic(err)
		}

		return true
	}
}

// Write a command to a bird socket
func birdWriteln(bird io.Writer, s string) {
	_, err := bird.Write([]byte(s + "\n"))
	if err != nil {
		panic(err)
	}
}

func birdHandler(socket string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("q")
		if query == "" {
			invalidHandler(w, r)
		} else {
			// Initialize BIRDv4 socket
			sock, err := net.Dial("unix", socket)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer sock.Close()
			err = sock.SetDeadline(time.Now().Add(time.Second * 30))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			bird := bufio.NewReader(sock)

			birdReadln(bird, nil)
			birdWriteln(sock, "restrict")
			birdReadln(bird, nil)
			birdWriteln(sock, query)
			for birdReadln(bird, w) {
			}
		}
	}
}
