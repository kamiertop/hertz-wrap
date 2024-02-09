package router

import (
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"

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

}
