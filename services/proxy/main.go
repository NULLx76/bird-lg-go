package main

import (
	logger "github.com/chi-middleware/logrus-logger"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

// Check if a byte is a character or a number
func isNumeric(b byte) bool {
	return b >= byte('0') && b <= byte('9')
}

// Default handler, returns 500 Internal Server Error
func invalidHandler(w http.ResponseWriter, _ *http.Request) {
	http.Error(w, "Invalid Request", http.StatusInternalServerError)
}

type Config struct {
	birdSocket string
	listen     string
}

func main() {
	var conf = Config{
		"/var/run/bird.ctl",
		":8000",
	}

	if birdSocketEnv := os.Getenv("BIRD_SOCKET"); birdSocketEnv != "" {
		conf.birdSocket = birdSocketEnv
	}
	if listenEnv := os.Getenv("BIRDLG_LISTEN"); listenEnv != "" {
		conf.listen = listenEnv
	}

	r := chi.NewRouter()
	r.Use(logger.Logger("router", log.StandardLogger()))
	r.Use(middleware.Recoverer)

	// Start HTTP server
	r.Get("/", invalidHandler)
	r.Get("/bird", birdHandler(conf.birdSocket))

	r.HandleFunc("/traceroute", tracerouteIPv4Wrapper)
	r.HandleFunc("/traceroute6", tracerouteIPv6Wrapper)

	log.Infof("Listening on %s", conf.listen)

	log.Fatal(http.ListenAndServe(conf.listen, r))
}
