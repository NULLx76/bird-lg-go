package main

import (
	"github.com/NULLx76/bird-lg-go/frontend/templates"
	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

//go:generate qtc -dir=templates

func main() {
	log.Info("starting the server at http://localhost:8080 ...")

	log.Fatal(fasthttp.ListenAndServe(":8080", requestHandler))
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "/":
		fallthrough
	default:
		mainPageHandler(ctx)
	}
	ctx.SetContentType("text/html; charset=utf-8")
}

func mainPageHandler(ctx *fasthttp.RequestCtx) {
	p := &templates.MainPage{
		CTX: ctx,
	}
	templates.WritePageTemplate(ctx, p)
}
