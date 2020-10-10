package main

import (
	"github.com/fasthttp/router"
	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

//go:generate qtc -dir=templates
//go:generate yarn build
//go:generate yarn minify

func main() {
	addr := ":8080"

	log.Infof("starting the server at %s ...", addr)

	r := router.New()
	r.GET("/", html(mainPageHandler))
	r.GET("/{server}/details/{peer}", html(peerPageHandler))

	// Static files
	r.ServeFiles("/static/{filepath:*}", "frontend/build")

	log.Fatal(fasthttp.ListenAndServe(addr, r.Handler))
}

func html(h fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.SetContentType("text/html; charset=utf-8")
		h(ctx)
	}
}
