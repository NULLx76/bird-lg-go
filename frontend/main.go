package main

import (
	"github.com/NULLx76/bird-lg-go/frontend/templates"
	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"strings"
)

//go:generate qtc -dir=templates
//go:generate npx tailwindcss build css/style.pcss -o build/style.css

var staticFilesHandler = fasthttp.FSHandler("/home/victor/src/bird-lg-go/frontend/build", 1)

func main() {
	log.Info("starting the server at http://localhost:8080 ...")

	log.Fatal(fasthttp.ListenAndServe(":8080", requestHandler))
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	path := string(ctx.Path())
	switch {
	case strings.HasPrefix(path, "/static"):
		staticFilesHandler(ctx)
	case path == "/":
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
