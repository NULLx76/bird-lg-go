package main

import (
	"github.com/NULLx76/bird-lg-go/api/proxy"
	"github.com/NULLx76/bird-lg-go/frontend/templates"
	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"strings"
)

//go:generate qtc -dir=templates
//go:generate npx tailwindcss build css/style.pcss -o build/style.css

var staticFilesHandler = fasthttp.FSHandler("/home/victor/src/bird-lg-go/frontend/build", 1)

func main() {
	addr := ":8080"

	log.Infof("starting the server at %s ...", addr)

	log.Fatal(fasthttp.ListenAndServe(addr, requestHandler))
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
	// TODO: cache
	servers, err := GetServers()
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	summaries := make(map[string]proxy.SummaryTable)
	for i := range servers {
		sum, err := GetSummary(servers[i])
		if err != nil {
			ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
			return
		}

		summaries[servers[i]] = sum
	}

	p := &templates.MainPage{
		CTX:       ctx,
		Summaries: summaries,
	}
	templates.WritePageTemplate(ctx, p)
}
