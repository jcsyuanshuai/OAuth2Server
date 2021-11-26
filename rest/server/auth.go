package server

import (
	"github.com/gin-gonic/gin"
	"github.com/go-session/session"
	"net/http"
	"net/url"
)

func Auth(ctx *gin.Context) {
	if ctx.Request.Method == "POST" {
		store, err := session.Start(ctx.Request.Context(), ctx.Writer, ctx.Request)
		if err != nil {
			print(err)
			return
		}
		if _, ok := store.Get("LoggedInUserID"); !ok {
			ctx.Redirect(http.StatusFound, "/server/login")
			return
		}

		var form url.Values
		if v, ok := store.Get("ReturnUri"); ok {
			form = v.(url.Values)
		}
		u := new(url.URL)
		u.Path = "/api/v1/oauth2/authorize"
		u.RawQuery = form.Encode()
		ctx.Writer.Header().Set("Location", u.String())
		ctx.Writer.WriteHeader(http.StatusFound)
		store.Delete("Form")

		if v, ok := store.Get("LoggedInUserID"); ok {
			store.Set("UserID", v)
		}
		store.Save()
		return
	}
	ctx.HTML(http.StatusOK, "auth.html", gin.H{
		"title": "Auth website",
	})

}
