package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"oauth2server/rest/oauth"
	"oauth2server/rest/secret"
	"time"
)

func InitPatientGrp(app *gin.Engine) {
	patient := app.Group("/api/v1/secret")
	patient.Use(ValidateToken)
	patient.GET("get", secret.Get)
}

func ValidateToken(ctx *gin.Context) {
	srv := oauth.GetServer()
	token, err := srv.ValidationBearerToken(ctx.Request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		ctx.Abort()
	} else {
		data := map[string]interface{}{
			"expires_in": int64(token.GetAccessCreateAt().Add(token.GetAccessExpiresIn()).Sub(time.Now()).Seconds()),
			"client_id":  token.GetClientID(),
			"user_id":    token.GetUserID(),
		}
		print(data)
	}
	ctx.Next()
}
