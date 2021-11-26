package router

import (
	"github.com/gin-gonic/gin"
	"oauth2server/rest/server"
)

func InitServerGrp(app *gin.Engine) {
	//app.LoadHTMLGlob("templates/*.tmpl")
	app.LoadHTMLGlob("static/*.html")
	serverGrp := app.Group("")
	serverGrp.GET("/server/login", server.Login)
	serverGrp.POST("/server/login", server.Login)
	serverGrp.GET("/server/auth", server.Auth)
	serverGrp.POST("/server/auth", server.Auth)
}
