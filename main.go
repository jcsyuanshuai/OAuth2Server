package main

import (
	"github.com/go/oauth2-server/bootstrap"
)

func main() {

	app := bootstrap.NewApp()

	if err := app.Init(); err != nil {
		panic(err)
	}

	if err := app.Start(); err != nil {
		panic(err)
	}

	err := <-app.Options().ErrorCh
	_ = app.Stop()
	panic(err)

}
