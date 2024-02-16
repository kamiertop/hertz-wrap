package router

import (
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"

	"hertz/pkg/log"
	"hertz/pkg/middleware"
)

func Init() *server.Hertz {
	h := server.New()
	// use middleware
	h.Use(
		recovery.Recovery(),
		middleware.Logger(),
	)

	initRouter(h)

	return h
}

// initRouter init all routers
func initRouter(h *server.Hertz) {
	h.GET("/ping", func(_ context.Context, ctx *app.RequestContext) {
		ctx.String(http.StatusOK, "pong")
	})
	h.GET("/ping/:id", func(_ context.Context, ctx *app.RequestContext) {
		ctx.String(http.StatusOK, ctx.Param("id"))
		log.Info("id")
	})
	h.GET("/ping/{name}", func(_ context.Context, ctx *app.RequestContext) {
		ctx.String(http.StatusOK, ctx.Param("name"))
		log.Info("name")
	})
}
