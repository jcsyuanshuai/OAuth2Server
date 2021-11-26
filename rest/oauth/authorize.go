package oauth

import (
	"github.com/gin-gonic/gin"
)

type AuthorizeRequestBody struct {
	ResponseType        string `json:"response_type" form:"response_type" binding:"required"`
	ClientID            string `json:"client_id" form:"client_id" binding:"required"`
	RedirectUri         string `json:"redirect_uri" form:"redirect_uri"`
	CodeChallenge       string `json:"code_challenge" form:"code_challenge"`
	CodeChallengeMethod string `json:"code_challenge_method" form:"code_challenge_method"`
	Scopes              string `json:"scopes" form:"scopes"`
	State               string `json:"state" form:"state"`
}

func Authorize(ctx *gin.Context) {
	srv := GetServer()
	err := srv.HandleAuthorizeRequest(ctx.Writer, ctx.Request)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
}
