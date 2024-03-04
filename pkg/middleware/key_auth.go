package middleware

import (
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/keyauth"
)

// KeyAuth wrap from hertz, you can set it manually.
func KeyAuth() app.HandlerFunc {
	return keyauth.New(
		keyauth.WithContextKey("token"),
		keyauth.WithKeyLookUp(lookUp("header", "Authorization"), "Bearer"),
		// The middleware is skipped when true is returned.
		keyauth.WithFilter(func(c context.Context, ctx *app.RequestContext) bool {
			return true
		}),

		// It may be used to validate key.
		// If returns false or err != nil, then errorHandler is used.
		keyauth.WithValidator(func(ctx context.Context, requestContext *app.RequestContext, s string) (bool, error) {
			return false, keyauth.ErrMissingOrMalformedAPIKey
		}),

		// It may be used to define a custom error.
		keyauth.WithErrorHandler(func(ctx context.Context, requestContext *app.RequestContext, err error) {
			requestContext.AbortWithMsg("msg", http.StatusBadRequest)
		}),
		keyauth.WithSuccessHandler(func(c context.Context, ctx *app.RequestContext) {
			ctx.Next(c)
		}),
	)
}

// lookUp set key auth lookUp, panic can be avoided in keyauth.WithKeyLookUp.
func lookUp(key, value string) string {
	return key + ":" + value
}
