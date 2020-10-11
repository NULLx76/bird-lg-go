package main

import (
	"bufio"
	"bytes"
	"io"
	"net"
	"net/http"
	"time"
)

// isNumeric check if a byte is a number
func isNumeric(b byte) bool {
	return b >= byte('0') && b <= byte('9')
}

func birdScanToEnd(s *bufio.Scanner, w io.Writer) {
	for s.Scan() {
		c := s.Bytes()

		if len(c) > 4 && isNumeric(c[0]) && isNumeric(c[1]) && isNumeric(c[2]) && isNumeric(c[3]) {
			if w != nil {
				nc := c[5:]
				nc = bytes.Trim(nc, " ")

				if _, err := w.Write(nc); err != nil {
					panic(err)
				}
			}

			// 0xxx = success
			// 8xxx = failure
			// 9xxx = failure
			if c[0] == byte('0') || c[0] == byte('8') || c[0] == byte('9') {
				return
			}
		} else {
			if w != nil {
				c = bytes.Trim(c, " ")
				if _, err := w.Write(c); err != nil {
					panic(err)
				}
			}
		}
	}
}

// Read a line from bird socket, removing preceding status number, output it.
// Returns if there are more lines.
func birdReadln(bird *bufio.Reader, w io.Writer) bool {
	c, err := bird.ReadBytes('\n')
	if err != nil {
		panic(err)
	}

	// Check if it starts with a status code (4 digits)
	if len(c) > 4 && isNumeric(c[0]) && isNumeric(c[1]) && isNumeric(c[2]) && isNumeric(c[3]) {
		if w != nil {
			nc := c[5:]
			nc = bytes.Trim(nc, " ")

			if _, err := w.Write(nc); err != nil {
				panic(err)
			}
		}

		// 0xxx = success
		// 8xxx = failure
		// 9xxx = failure
		return c[0] != byte('0') && c[0] != byte('8') && c[0] != byte('9')
	} else {
		if w != nil {
			c = bytes.Trim(c, " ")
			if _, err := w.Write(c); err != nil {
				panic(err)
			}
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
			return
		}
		// Initialize BIRDv4 socket
		sock, err := net.Dial("unix", socket)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer sock.Close()
		if err = sock.SetDeadline(time.Now().Add(time.Second * 30)); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		bird := bufio.NewScanner(sock)

		bird.Scan()
		birdWriteln(sock, "restrict")
		bird.Scan()
		birdWriteln(sock, query)
		birdScanToEnd(bird, w)
	}
}
