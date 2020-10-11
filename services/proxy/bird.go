package main

import (
	"bufio"
	"io"
	"net"
	"net/http"
	"time"
)

// Read a line from bird socket, removing preceding status number, output it.
// Returns if there are more lines.
func birdReadln(bird io.Reader, w io.Writer) bool {
	// Read from socket byte by byte, until reaching newline character
	c := make([]byte, 1024)
	pos := 0
	for {
		if pos >= 1024 {
			break
		}
		_, err := bird.Read(c[pos : pos+1])
		if err != nil {
			panic(err)
		}
		if c[pos] == byte('\n') {
			break
		}
		pos++
	}
	c = c[:pos+1]

	// Remove preceding status number, different situations
	if pos < 4 {
		// Line is too short to have a status number
		if w != nil {
			pos = 0
			for c[pos] == byte(' ') {
				pos++
			}
			if _, err := w.Write(c[pos:]); err != nil {
				panic(err)
			}
		}
		return true
	} else if isNumeric(c[0]) && isNumeric(c[1]) && isNumeric(c[2]) && isNumeric(c[3]) {
		// There is a status number at beginning, remove first 5 bytes
		if w != nil && pos > 6 {
			pos = 5
			for c[pos] == byte(' ') {
				pos++
			}
			if _, err := w.Write(c[pos:]); err != nil {
				panic(err)
			}
		}
		return c[0] != byte('0') && c[0] != byte('8') && c[0] != byte('9')
	} else {
		// There is no status number, only remove preceding spaces
		if w != nil {
			pos = 0
			for c[pos] == byte(' ') {
				pos++
			}
			if _, err := w.Write(c[pos:]); err != nil {
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
