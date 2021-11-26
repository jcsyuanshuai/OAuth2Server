package client

import (
	"crypto/sha256"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"net/http"
)

func Home(ctx *gin.Context) {
	u := config.AuthCodeURL(
		"xyz",
		oauth2.SetAuthURLParam(
			"code_challenge",
			func(s string) string {
				s256 := sha256.Sum256([]byte(s))
				return base64.URLEncoding.EncodeToString(s256[:])
			}("s256example"),
		),
		oauth2.SetAuthURLParam(
			"code_challenge_method",
			"S256",
		),
	)
	ctx.Redirect(http.StatusFound, u)
}
