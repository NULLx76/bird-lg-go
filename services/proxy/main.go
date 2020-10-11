package main

import (
	"flag"
	logger "github.com/chi-middleware/logrus-logger"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

// Check if a byte is character for number
func isNumeric(b byte) bool {
	return b >= byte('0') && b <= byte('9')
}

// Default handler, returns 500 Internal Server Error
func invalidHandler(w http.ResponseWriter, _ *http.Request) {
	http.Error(w, "Invalid Request", http.StatusInternalServerError)
}

type settingType struct {
	birdSocket string
	listen     string
}

var setting settingType

// Wrapper of tracer
func main() {
	// Prepare default socket paths, use environment variable if possible
	var settingDefault = settingType{
		"/var/run/bird.ctl",
		":8000",
	}

	if birdSocketEnv := os.Getenv("BIRD_SOCKET"); birdSocketEnv != "" {
		settingDefault.birdSocket = birdSocketEnv
	}
	if listenEnv := os.Getenv("BIRDLG_LISTEN"); listenEnv != "" {
		settingDefault.listen = listenEnv
	}

	// Allow parameters to override environment variables
	birdParam := flag.String("bird", settingDefault.birdSocket, "socket file for bird, set either in parameter or environment variable BIRD_SOCKET")
	listenParam := flag.String("listen", settingDefault.listen, "listen address, set either in parameter or environment variable BIRDLG_LISTEN")
	flag.Parse()

	setting.birdSocket = *birdParam
	setting.listen = *listenParam

	r := chi.NewRouter()
	r.Use(logger.Logger("router", log.StandardLogger()))
	r.Use(middleware.Recoverer)

	// Start HTTP server
	r.HandleFunc("/", invalidHandler)
	r.HandleFunc("/bird", birdHandler)

	r.HandleFunc("/traceroute", tracerouteIPv4Wrapper)
	r.HandleFunc("/traceroute6", tracerouteIPv6Wrapper)

	log.Infof("Listening on %s", listenParam)

	log.Fatal(http.ListenAndServe(*listenParam, r))
}
