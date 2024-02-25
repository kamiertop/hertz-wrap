package middleware

import "github.com/hertz-contrib/jwt"

func JwtAuth() {
	_, _ = jwt.New(&jwt.HertzJWTMiddleware{})
}
