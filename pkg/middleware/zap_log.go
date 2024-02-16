package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"go.uber.org/zap"

	"hertz/pkg/log"
	"hertz/pkg/utils"
)

const _execTimeout = 1

// Logger middleware record request and response information.
func Logger() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		logMsg := []zap.Field{
			zap.String("route", ctx.FullPath()),
			zap.String("ip", ctx.RemoteAddr().String()),
			zap.String("agent", utils.BytesToString(ctx.UserAgent())),
		}

		switch string(ctx.Method()) {
		case http.MethodGet:
			if v := ctx.QueryArgs(); v.Len() != 0 {
				logMsg = append(logMsg, zap.String("query_rgs", utils.BytesToString(v.QueryString())))
			}
			if p := ctx.Params; len(p) != 0 {
				logMsg = append(logMsg, zap.Any("params", ctx.Params))
			}
		case http.MethodPost:
			ctx.Request.PostArgString()
			if v := ctx.Request.PostArgs().String(); v != "" {
				logMsg = append(logMsg, zap.String("post", v))
			}
		}

		s := time.Now()
		ctx.Next(c)
		cost := time.Since(s)
		logMsg = append(logMsg, zap.String("cost", cost.String()))

		switch ctx.Response.StatusCode() {
		case http.StatusOK:
			if cost.Seconds() > _execTimeout {
				log.Warn(string(ctx.Request.URI().Path()), logMsg...)
			} else {
				log.Info(string(ctx.Request.URI().Path()), logMsg...)
			}
		default:
			log.Error(string(ctx.Request.URI().Path()), logMsg...)
		}
	}
}
