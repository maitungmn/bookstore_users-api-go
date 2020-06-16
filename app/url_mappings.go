package app

import (
	"github.com/maitungmn/bookstore_users-api-go/controllers/ping"
	"github.com/maitungmn/bookstore_users-api-go/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)
}
