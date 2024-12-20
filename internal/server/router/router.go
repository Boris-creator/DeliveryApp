package router

import (
	"fmt"
	"html/template"
	"log"

	"playground/internal/server/api/address_suggest"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func MakeRouter() *router.Router {
	r := router.New()
	webRoutes(r)
	apiRoutes(r)
	return r
}

func apiRoutes(r *router.Router) {
	api := r.Group("/api/v1")
	api.POST("/suggest-address", address_suggest.HandleSuggest)
}

func webRoutes(r *router.Router) {
	r.GET("/", func(ctx *fasthttp.RequestCtx) {
		tmpl, err := template.ParseFiles("web/views/index.html")
		if err != nil {
			log.Println(err)
			return
		}
		ctx.SetContentType("text/html")
		tmpl.Execute(ctx, struct {
			Title   string
			Content string
		}{"Hello World!", "Test"})
	})
	r.ServeFilesCustom("/{filepath:*}", &fasthttp.FS{
		Root:               "web/public",
		IndexNames:         []string{"index.html", "index.js"},
		GenerateIndexPages: true,
		AcceptByteRange:    true,
		PathNotFound: func(ctx *fasthttp.RequestCtx) {
			fmt.Fprintf(ctx, "404 not found")
		},
	})
}
