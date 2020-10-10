package main

import (
	"fmt"
	"github.com/NULLx76/bird-lg-go/api/proxy"
	"github.com/NULLx76/bird-lg-go/frontend/templates"
	"github.com/valyala/fasthttp"
)

func mainPageHandler(ctx *fasthttp.RequestCtx) {
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
