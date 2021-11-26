package oauth

import "github.com/gin-gonic/gin"

type TokenRequestBody struct {
	ResponseType string `json:"response_type" form:"response_type" binding:"required"`
	ClientID     string `json:"client_id" form:"client_id" binding:"required"`
	RedirectUri  string `json:"redirect_uri" form:"redirect_uri"`
	Scopes       string `json:"scopes" form:"scopes"`
	State        string `json:"state" form:"state"`
}

func Token(ctx *gin.Context) {
	srv := GetServer()
	err := srv.HandleTokenRequest(ctx.Writer, ctx.Request)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
}
