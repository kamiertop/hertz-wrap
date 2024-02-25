package middleware

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/keyauth"
)

// KeyAuth wrap from hertz
func KeyAuth() app.HandlerFunc {
	return keyauth.New()
}
