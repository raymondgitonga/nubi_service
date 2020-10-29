package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/raymondgitonga/nubi_service/email_service/dormain"
	"github.com/raymondgitonga/nubi_service/email_service/service"
	"net/http"
)

func AddUser(c *gin.Context) {
	var reqBody dormain.User

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   true,
			"message": "Invalid request body",
		})
		return
	}
	resp, err := service.UserService.AddUser(reqBody)

	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func GetUsers(c *gin.Context) {
	var users *[]dormain.User

	users, err := service.UserService.GetUsers()
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, users)
}

func GetUser(c *gin.Context) {
	email := c.Param("email")

	user, err := service.UserService.GetUser(email)

	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, user)
}
