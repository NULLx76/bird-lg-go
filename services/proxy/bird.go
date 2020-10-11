package main

import (
	"bufio"
	"bytes"
	"github.com/pkg/errors"
	"io"
	"net"
	"net/http"
	"time"
)

// isNumeric check if a byte is a number
func isNumeric(b byte) bool {
	return b >= byte('0') && b <= byte('9')
}

// birdScanToEnd will scan the buffer to the end and write its content to w
func birdScanToEnd(s *bufio.Scanner, w io.Writer) error {
	if w == nil {
		return errors.New("nil writer")
	}

	for s.Scan() {
		c := s.Bytes()

		if len(c) > 4 && isNumeric(c[0]) && isNumeric(c[1]) && isNumeric(c[2]) && isNumeric(c[3]) {
			nc := c[5:]
			nc = bytes.Trim(nc, " ")

			if _, err := w.Write(nc); err != nil {
				return errors.Wrap(err, "error writing to writer")
			}
			if _, err := w.Write([]byte{'\n'}); err != nil {
				return errors.Wrap(err, "error writing newline to writer")
			}

			// 0xxx = success
			// 8xxx = failure
			// 9xxx = failure
			if c[0] == byte('0') || c[0] == byte('8') || c[0] == byte('9') {
				break
			}
		} else {
			c = bytes.Trim(c, " ")

			if _, err := w.Write(c); err != nil {
				return errors.Wrap(err, "error writing to writer")
			}
			if _, err := w.Write([]byte{'\n'}); err != nil {
				return errors.Wrap(err, "error writing to writer")
			}
		}
	}
	return nil
}

// birdWriteln writes a command to a bird socket
func birdWriteln(bird io.Writer, s string) error {
	_, err := bird.Write([]byte(s + "\n"))
	return errors.Wrap(err, "sending message to bird failed")
}

// birdHandler http handler that interacts with the bird socket
func birdHandler(socket string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("q")
		if query == "" {
			invalidHandler(w, r)
			return
		}

		// Connect to BIRDv4 socket
		birdS, err := net.Dial("unix", socket)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer birdS.Close()

		// Set deadline t
		if err = birdS.SetDeadline(time.Now().Add(time.Second * 30)); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Create scanner for the socket
		bird := bufio.NewScanner(birdS)

		// Advance beyond the welcome message
		bird.Scan()

		// Restrict access so no modifications can take place
		err = birdWriteln(birdS, "restrict")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Advance over reply
		bird.Scan()

		// Send query to bird
		err = birdWriteln(birdS, query)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// write reply to http socket
		err = birdScanToEnd(bird, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
