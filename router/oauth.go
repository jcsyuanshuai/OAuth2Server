package router

import (
	"github.com/gin-gonic/gin"
	oauth2 "oauth2server/rest/oauth"
)

func InitOAuthGrp(app *gin.Engine) {
	auth := app.Group("/api/v1/oauth2")
	auth.GET("/authorize", oauth2.Authorize)
	//
	auth.POST("/token", oauth2.Token)
	auth.GET("/register", oauth2.Register)
	auth.GET("/list", oauth2.List)
}
