package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"oauth2server/rest/client"
)

func InitCliGrp(app *gin.Engine) {
	cliGrp := app.Group("/client")
	cliGrp.GET("/refresh", client.Refresh)
	cliGrp.GET("/oauth2", client.Auth)
	cliGrp.GET("/home", client.Home)
	cliGrp.GET("/test", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusFound, "/api/v1/secret/get")
	})
}
