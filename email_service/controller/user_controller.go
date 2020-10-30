package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/raymondgitonga/nubi_service/email_service/dormain"
	"github.com/raymondgitonga/nubi_service/email_service/service"
	"net/http"
)

type userController struct{}

type userControllerInterface interface {
	AddUser(c *gin.Context)
	GetUsers(c *gin.Context)
	GetUser(c *gin.Context)
}

var (
	UserControllerInterface userControllerInterface
)

func init() {
	UserControllerInterface = &userController{}
}

func (u *userController) AddUser(c *gin.Context) {
	var reqBody dormain.User

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   true,
			"message": "Invalid request body",
		})
		return
	}
	resp, err := service.UserServiceInterface.AddUser(reqBody)

	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func (u *userController) GetUsers(c *gin.Context) {
	var users *[]dormain.User

	users, err := service.UserServiceInterface.GetUsers()
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, users)
}

func (u *userController) GetUser(c *gin.Context) {
	email := c.Param("email")

	user, err := service.UserServiceInterface.GetUser(email)

	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, user)
}
