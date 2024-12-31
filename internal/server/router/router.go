package router

import (
	"fmt"
	"html/template"
	"log"

	"playground/internal/server/api/address_suggest"
	"playground/internal/server/api/orders"
	"playground/internal/server/middleware"

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
	api.POST("/order", middleware.Validate[orders.SaveOrderRequest](orders.SaveOrder))
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
		}{"Main Page", "Order form"})
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
