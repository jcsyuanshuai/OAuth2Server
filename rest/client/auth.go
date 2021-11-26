package client

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"net/http"
)

type AuthBody struct {
	State string `form:"state"`
	Code  string `form:"code"`
}

func Auth(ctx *gin.Context) {
	body := AuthBody{}
	err := ctx.ShouldBind(&body)
	if err != nil {
		return
	}
	print("state: ", body.State, "\n")
	print("code: ", body.Code, "\n")
	if body.State != "xyz" {
		ctx.JSON(http.StatusBadRequest, "state invalid")
		return
	}
	if body.Code == "" {
		ctx.JSON(http.StatusBadRequest, "code invalid")
		return
	}
	print(config.ClientSecret, "\n")
	token, err := config.Exchange(ctx, body.Code, oauth2.SetAuthURLParam("code_verifier", "s256example"))
	if err != nil {
		print(err.Error())
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, token)
}
