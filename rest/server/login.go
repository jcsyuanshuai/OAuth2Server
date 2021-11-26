package server

import (
	"github.com/gin-gonic/gin"
	"github.com/go-session/session"
	"net/http"
)

type LoginBody struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func Login(ctx *gin.Context) {
	if ctx.Request.Method == "POST" {
		body := LoginBody{}
		err := ctx.ShouldBind(&body)
		if err != nil {
			return
		}
		print("username: " + body.Password + "\n")
		print("password: " + body.Username + "\n")
		store, err := session.Start(nil, ctx.Writer, ctx.Request)
		if err != nil {
			print(err)
			return
		}
		store.Set("LoggedInUserID", "0000 0000")
		ctx.Redirect(http.StatusFound, "/server/auth")
		return
	}

	ctx.HTML(http.StatusOK, "login.html", gin.H{
		"title": "Login website",
	})

}
