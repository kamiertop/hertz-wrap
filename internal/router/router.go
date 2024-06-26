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
	"hertz/pkg/middleware"
	"hertz/pkg/resp"
	"hertz/web"
)

var enablePprof = flag.Bool("pprof", false, "open/close pprof")

type Engine struct {
	*server.Hertz
}

func (e *Engine) wrapHandleFunc(httpMethod, relativePath, routeName string, handlers ...app.HandlerFunc) {
	setRouteName(httpMethod, relativePath, routeName)
	e.Handle(httpMethod, relativePath, handlers...)
}

func Init() *Engine {
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

	return &Engine{Hertz: h}
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
			resp.BadRequest(c, err)
			return
		}

		resp.SuccessData(c, m)
	})
}

// noRoute Custom implementations if don't want use default
func noRoute(_ context.Context, _ *app.RequestContext) {}

// noMethod Custom implementations if don't want use default
func noMethod(_ context.Context, _ *app.RequestContext) {}

func header(_ context.Context, c *app.RequestContext) {
	c.Request.Header.Add("Content-Type", "text/html; charset=utf-8")
	c.Header("Content-Type", "text/html; charset=utf-8")
}
