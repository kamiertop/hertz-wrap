// The files in basic_auth package are forked from gin[github.com/gin-gonic/gin],
// and we keep the original Copyright[Copyright 2014 gin authors] and License of gin for those files.
// We also need to modify as we need, the modifications are Copyright of 2022 CloudWeGo Authors.
// Thanks for gin authors! Below is the source code information:
// 		Repo: github.com/gin-gonic/gin
//		Forked Version: v1.9.1

package middleware

import (
	"context"
	"crypto/subtle"
	"encoding/base64"
	"net/http"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"

	"hertz/pkg/utils"
)

// AuthUserKey is the cookie name for user credential in basic auth.
const AuthUserKey = "user"

// Accounts defines a key/value for user/pass list of authorized logins.
type Accounts map[string]string

type authPair struct {
	value string
	user  string
}

type authPairs []authPair

func (a authPairs) searchCredential(authValue string) (string, bool) {
	if authValue == "" {
		return "", false
	}
	for _, pair := range a {
		if subtle.ConstantTimeCompare(utils.StringToBytes(pair.value), utils.StringToBytes(authValue)) == 1 {
			return pair.user, true
		}
	}
	return "", false
}

// BasicAuthForRealm returns a Basic HTTP Authorization middleware. It takes as arguments a map[string]string where
// the key is the user name and the value is the password, as well as the name of the Realm.
// If the realm is empty, "Authorization Required" will be used by default.
// (see http://tools.ietf.org/html/rfc2617#section-1.2)
func BasicAuthForRealm(accounts Accounts, realm string) app.HandlerFunc {
	if realm == "" {
		realm = "Authorization Required"
	}
	realm = "Basic realm=" + strconv.Quote(realm)
	pairs := processAccounts(accounts)
	return func(_ context.Context, c *app.RequestContext) {
		// Search user in the slice of allowed credentials
		user, found := pairs.searchCredential(c.Request.Header.Get("Authorization"))
		if !found {
			// Credentials doesn't match, we return 401 and abort handlers chain.
			c.Header("WWW-Authenticate", realm)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// The user credentials was found, set user's id to key AuthUserKey in this context, the user's id can be read later using
		// c.MustGet(gin.AuthUserKey).
		c.Set(AuthUserKey, user)
	}
}

// BasicAuth returns a Basic HTTP Authorization middleware. It takes as argument a map[string]string where
// the key is the user name and the value is the password.
func BasicAuth(accounts Accounts) app.HandlerFunc {
	return BasicAuthForRealm(accounts, "")
}

func processAccounts(accounts Accounts) authPairs {
	pairs := make(authPairs, 0, len(accounts))
	for user, password := range accounts {
		value := authorizationHeader(user, password)
		pairs = append(pairs, authPair{
			value: value,
			user:  user,
		})
	}
	return pairs
}

func authorizationHeader(user, password string) string {
	base := user + ":" + password
	return "Basic " + base64.StdEncoding.EncodeToString(utils.StringToBytes(base))
}
