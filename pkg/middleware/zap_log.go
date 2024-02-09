package middleware

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

func Logger() app.HandlerFunc {
	return func(_ context.Context, ctx *app.RequestContext) {

	}
}
