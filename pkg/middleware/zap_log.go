package middleware

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"go.uber.org/zap"

	"hertz/pkg/log"
	"hertz/pkg/utils"
)

const _execTimeout = 1

// Logger middleware record request and response information.
func Logger() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		logMsg := []zap.Field{
			zap.String("method", utils.BytesToString(c.Request.Method())),
			zap.String("route", c.FullPath()),
		}
		logMsg = requestArg(c, logMsg)
		s := time.Now()
		c.Next(ctx)
		cost := time.Since(s)

		logMsg = append(logMsg,
			zap.Int("status", c.Response.StatusCode()),
			zap.String("cost", cost.String()),
			zap.String("handler", c.FullPath()),
			zap.String("ip", c.RemoteAddr().String()),
			zap.String("agent", utils.BytesToString(c.UserAgent())))
		switch c.Response.StatusCode() {
		case http.StatusOK:
			if cost.Seconds() > _execTimeout {
				log.Warn(string(c.Request.URI().Path()), logMsg...)
			} else {
				log.Info(string(c.Request.URI().Path()), logMsg...)
			}
		default:
			if e := c.Errors.Last(); e != nil {
				logMsg = append(logMsg, zap.Error(e.Err), zap.String("err_msg", e.Meta.(string)))
			}
			log.Error(string(c.Request.URI().Path()), logMsg...)
		}
	}
}

func requestArg(c *app.RequestContext, logMsg []zap.Field) []zap.Field {
	switch string(c.Method()) {
	case http.MethodGet:
		query, err := url.QueryUnescape(c.QueryArgs().String())
		if err == nil && query != "" {
			logMsg = append(logMsg, zap.String("query_rgs", query))
		}
		if p := c.Params; len(p) != 0 {
			logMsg = append(logMsg, zap.Any("params", c.Params))
		}
	case http.MethodPost:
		if body := c.Request.Body(); body != nil {
			ct := utils.BytesToString(c.GetHeader(consts.HeaderContentType))
			if ct == consts.MIMEApplicationJSON || ct == consts.MIMEApplicationJSONUTF8 {
				var b map[string]any
				if err := json.Unmarshal(body, &b); err == nil {
					logMsg = append(logMsg, zap.Any("body", b))
				}
			}
		}
	case http.MethodDelete, http.MethodPatch:
	case http.MethodPut:
	default:

	}

	return logMsg
}
