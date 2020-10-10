package main

import (
	"fmt"
	"github.com/NULLx76/bird-lg-go/api/proxy"
	"github.com/NULLx76/bird-lg-go/frontend/templates"
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

func peerPageHandler(ctx *fasthttp.RequestCtx) {
	server := fmt.Sprintf("%s", ctx.UserValue("server"))
	peer := fmt.Sprintf("%s", ctx.UserValue("peer"))

	details, err := GetDetails(server, peer)
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	p := &templates.PeerPage{
		CTX:  ctx,
		Peer: details,
	}
	templates.WritePageTemplate(ctx, p)
}
