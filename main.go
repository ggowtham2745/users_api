package main

import (
	"github.com/ggowtham2745/users_api/app"
	"github.com/ggowtham2745/users_api/config"
)

func main() {
	config.Connect()

	app.StartApplication()

}
