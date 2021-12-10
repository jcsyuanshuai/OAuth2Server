package user

import (
	"github.com/gin-gonic/gin"
	pb "github.com/go/oauth2-server/proto/user"
	"github.com/go/oauth2-server/service/user"
	"github.com/go/oauth2-server/util/output"
)

func Login(ctx *gin.Context) {
	cli := user.New()
	ret, err := cli.Login(ctx, &pb.LoginRequest{})
	if err != nil {
		return
	}
	output.Json(ctx, output.Success, ret)
}
