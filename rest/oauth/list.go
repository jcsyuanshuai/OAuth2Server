package oauth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func List(ctx *gin.Context) {

	cs := GetClientStore()
	info, err := cs.GetByID(ctx, "1024")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}
	ctx.JSON(200, info)
}
