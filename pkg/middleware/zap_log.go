package middleware

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"go.uber.org/zap"

	"hertz/pkg/log"
	"hertz/pkg/utils"
)

const _execTimeout = 1

// Logger middleware record request and response information.
func Logger() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		logMsg := []zap.Field{
			zap.String("ip", ctx.RemoteAddr().String()),
			zap.String("agent", utils.BytesToString(ctx.UserAgent())),
		}
		s := time.Now()
		ctx.Next(c)
		cost := time.Since(s)
		logMsg = append(logMsg, zap.String("cost", cost.String()))
		if string(ctx.Method()) == consts.MethodPost {
			ctx.Request.PostArgString()
			if v := ctx.Request.PostArgs().String(); v != "" {
				logMsg = append(logMsg, zap.String("post", v))
			}
		}
		switch ctx.Response.StatusCode() {
		case consts.StatusOK:
			if cost.Seconds() > _execTimeout {
				log.Warn(ctx.Request.URI().String(), logMsg...)
			} else {
				log.Info(ctx.Request.URI().String(), logMsg...)
			}
		default:
			log.Error(ctx.Request.URI().String(), logMsg...)
		}
	}
}
