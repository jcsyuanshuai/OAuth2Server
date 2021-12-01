package bootstrap

import (
	"github.com/go/oauth2-server/rest"
)

func SetupSwaggerRouters(opts *Options) {
	grp := opts.Engine.Group("/user")
	grp.GET("/list", rest.ListUser)
}

func SetupUserRouters(opts *Options) {
	grp := opts.Engine.Group("/user")
	grp.GET("/list", rest.ListUser)
}

func SetupTokenRouters(opts *Options) {
	grp := opts.Engine.Group("/token")
	grp.GET("/list", rest.ListUser)
}

func SetupClientRouters(opts *Options) {
	grp := opts.Engine.Group("/client")
	grp.GET("/list", rest.ListUser)
}

func SetupPermissionRouters(opts *Options) {
	grp := opts.Engine.Group("/permission")
	grp.GET("/list", rest.ListUser)
}
