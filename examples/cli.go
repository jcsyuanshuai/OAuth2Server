package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"oauth2server/router"
)

func main() {
	app := gin.New()
	router.InitCliGrp(app)
	_ = app.Run(":8080")
	fmt.Println("client running on port 8080")
}
