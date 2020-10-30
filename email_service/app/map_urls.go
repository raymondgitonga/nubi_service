package app

import "github.com/raymondgitonga/nubi_service/email_service/controller"

func mapUrls() {
	router.POST("/users", controller.UserControllerInterface.AddUser)
	router.GET("/users", controller.UserControllerInterface.GetUsers)
	router.GET("/users/:email", controller.UserControllerInterface.GetUser)
}
