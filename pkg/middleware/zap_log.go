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
		if ctx.Response.StatusCode() == consts.StatusOK {
			if cost.Seconds() > _execTimeout {
				log.Warn(ctx.Request.URI().String(), logMsg...)
			} else {
				log.Info(ctx.Request.URI().String(), logMsg...)
			}
		} else {
			log.Error(ctx.Request.URI().String(), logMsg...)
		}
	}
}
