package bootstrap

import (
	"github.com/go/oauth2-server/rest/user"
)

func SetupSwaggerRouters(opts *Options) {
	//todo
}

func SetupUserRouters(opts *Options) {
	grp := opts.Engine.Group("/user")
	grp.GET("/list", user.List)
	grp.POST("/login", user.Login)
	grp.POST("/create", user.Create)
}

func SetupTokenRouters(opts *Options) {
	//grp := opts.Engine.Group("/token")
	//todo
}

func SetupClientRouters(opts *Options) {
	//grp := opts.Engine.Group("/client")
	//todo
}

func SetupPermissionRouters(opts *Options) {
	//grp := opts.Engine.Group("/permission")
	//todo
}
