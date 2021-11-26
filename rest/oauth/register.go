package oauth

import (
	"github.com/gin-gonic/gin"
	"github.com/go-oauth2/oauth2/v4/models"
	"strconv"
)

var counter int = 1100

func Register(ctx *gin.Context) {
	counter = counter + 1
	clientId := strconv.Itoa(counter)

	info := &models.Client{
		ID:     clientId,
		Secret: clientId,
		Domain: "http://localhost:8080/",
	}
	cs := GetClientStore()
	err := cs.Set(clientId, info)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	ctx.JSON(200, info)
}
