package app

import (
	"github.com/ggowtham2745/users_api/controllers/ping"
	"github.com/ggowtham2745/users_api/controllers/user"
)

func urlMaps() {
	router.GET("/", ping.Welcome)
	router.GET("/ping", ping.Ping)
	router.GET("/users", user.GetUserList)
}
