package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/go/oauth2-server/util"
)

func ListUser(ctx *gin.Context) {

}

func Login(ctx *gin.Context) {
	data := ""
	util.Json(ctx, util.Success, data)
}

func RegisterUser(ctx *gin.Context) {
	data := ""
	util.Json(ctx, util.Success, data)
}

func CreateUser(ctx *gin.Context) {
	data := ""
	util.Json(ctx, util.Success, data)
}
