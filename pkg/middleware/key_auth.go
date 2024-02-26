package middleware

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/keyauth"
)

// KeyAuth wrap from hertz
func KeyAuth() app.HandlerFunc {
	return keyauth.New(
		keyauth.WithKeyLookUp(lookUp("header", "Authorization"), "Bearer"),
	)
}

// lookUp set key auth lookUp, panic can be avoided in keyauth.WithKeyLookUp.
func lookUp(key, value string) string {
	return key + ":" + value
}
