package user

import (
	"github.com/gin-gonic/gin"
	pb "github.com/go/oauth2-server/proto/user"
	"github.com/go/oauth2-server/service/user"
	"github.com/go/oauth2-server/util/output"
)

type CreateParams struct {
	Email    string
	Name     string
	Password string
	PhoneNo  string
}

func Create(ctx *gin.Context) {
	p := &CreateParams{}
	err := ctx.ShouldBindJSON(p)
	if err != nil {
		return
	}

	cli := user.New()
	ret, err := cli.Create(ctx, &pb.User{
		Name:        p.Name,
		Gender:      "",
		Email:       p.Email,
		Password:    p.Password,
		PhoneNumber: p.PhoneNo,
	})
	if err != nil {
		return
	}

	output.Json(ctx, output.Success, ret)
}
