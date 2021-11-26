package secret

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Patient struct {
	Pid      string
	Name     string
	Gender   string
	Age      string
	DoctorId string
	IdCard   string
	Birthday string
}

func Get(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Patient{
		Pid:      "1",
		Name:     "san",
		Gender:   "1",
		Age:      "20",
		DoctorId: "10",
		IdCard:   "100000xxxx1010101",
		Birthday: "2018-11-09",
	})
}
