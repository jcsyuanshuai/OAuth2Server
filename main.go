package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"oauth2server/router"
)

func main() {
	app := gin.New()
	initRouterGroup(
		app,
		router.InitOAuthGrp,
		router.InitPatientGrp,
		router.InitServerGrp,
	)
	_ = app.Run(":8090")
	fmt.Println("server running on port 8090")

}

type InitGroupFunc func(app *gin.Engine)

func initRouterGroup(app *gin.Engine, igs ...InitGroupFunc) {
	for _, grp := range igs {
		grp(app)
	}
}
