package router

import (
	"fmt"
	"html/template"

	"github.com/fasthttp/router"
	fastHttpSwagger "github.com/swaggo/fasthttp-swagger"
	"github.com/valyala/fasthttp"
	_ "playground/api/docs"
	"playground/internal/logger"
)

func webRoutes(r *router.Router) {
	r.GET("/", func(ctx *fasthttp.RequestCtx) {
		tmpl, err := template.ParseFiles("web/views/index.html")
		if err != nil {
			logger.Error(err)
			return
		}

		ctx.SetContentType("text/html")

		err = tmpl.Execute(ctx, struct {
			Title   string
			Content string
		}{"Main Page", "Order form"})
		if err != nil {
			logger.Error(err)
		}
	})

	r.GET("/swagger/{filepath:*}", fastHttpSwagger.WrapHandler())

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
