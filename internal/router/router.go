package router

import (
	"context"
	"flag"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/pprof/adaptor"

	"hertz/internal/metrics"
	"hertz/pkg/config"
	"hertz/pkg/consts"
	"hertz/pkg/middleware"
	"hertz/web"
)

var enablePprof = flag.Bool("pprof", false, "open/close pprof")

func Init() *server.Hertz {
	flag.Parse()
	hlog.SetSilentMode(true)
	h := server.New(
		server.WithHostPorts(config.Conf.System.Addr),
	)
	// use middleware
	if *enablePprof {
		metrics.Init(h)
	}
	h.Use(
		recovery.Recovery(),
		middleware.Logger(),
	)

	initRouter(h)

	return h
}

// initRouter init all routers
func initRouter(h *server.Hertz) {
	h.NoMethod(noMethod)
	h.NoRoute(noRoute)

	h.GET("/ping", func(_ context.Context, c *app.RequestContext) {
		c.String(http.StatusOK, "pong")
	})

	h.GET("/", header, adaptor.NewHertzHTTPHandler(http.FileServer(http.FS(web.IndexHtml))))
	h.GET("/static/*filepath", header, adaptor.NewHertzHTTPHandler(http.FileServerFS(web.Dist)))
	h.POST("/ping", func(_ context.Context, c *app.RequestContext) {
		var m map[string]any
		if err := c.BindJSON(&m); err != nil {
			consts.BadRequest(c, err)
			return
		}

		consts.SuccessData(c, m)
	})
}

// noRoute Custom implementations if don't want use default
func noRoute(_ context.Context, c *app.RequestContext) {}

// noMethod Custom implementations if don't want use default
func noMethod(_ context.Context, c *app.RequestContext) {}

func header(_ context.Context, c *app.RequestContext) {
	c.Request.Header.Add("Content-Type", "text/html; charset=utf-8")
	c.Header("Content-Type", "text/html; charset=utf-8")
}
