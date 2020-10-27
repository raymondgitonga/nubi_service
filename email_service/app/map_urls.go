package app

import "github.com/raymondgitonga/nubi_service/email_service/controller"

func mapUrls() {
	router.POST("/users", controller.AddUser)
	router.GET("/users", controller.GetUsers)
	router.GET("/users/:email", controller.GetUser)
}