package metrics

import (
	"net/http"
	"net/http/pprof"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/pprof/adaptor"
)

// Init register pprof, steal from Hertz
// Use middleware conveniently and better routing group control.
func Init(h *server.Hertz) {
	g := h.Group("debug/pprof")
	g.GET("/", wrapPprof(pprof.Index))
	g.GET("/cmdline", wrapPprof(pprof.Cmdline))
	g.GET("/profile", wrapPprof(pprof.Profile))
	g.POST("/symbol", wrapPprof(pprof.Symbol))
	g.GET("/symbol", wrapPprof(pprof.Symbol))
	g.GET("/trace", wrapPprof(pprof.Trace))

	g.GET("/allocs", wrapPprof(pprof.Handler("allocs").ServeHTTP))
	g.GET("/block", wrapPprof(pprof.Handler("block").ServeHTTP))
	g.GET("/goroutine", wrapPprof(pprof.Handler("goroutine").ServeHTTP))
	g.GET("/heap", wrapPprof(pprof.Handler("heap").ServeHTTP))
	g.GET("/mutex", wrapPprof(pprof.Handler("mutex").ServeHTTP))
	g.GET("/threadcreate", wrapPprof(pprof.Handler("threadcreate").ServeHTTP))
}

func wrapPprof(h http.HandlerFunc) app.HandlerFunc {
	return adaptor.NewHertzHTTPHandlerFunc(h)
}
