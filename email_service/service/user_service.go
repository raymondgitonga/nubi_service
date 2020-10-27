package service

import (
	"github.com/raymondgitonga/nubi_service/email_service/dormain"
	"github.com/raymondgitonga/nubi_service/email_service/utils"
)

type userService struct{}

var (
	UserService userService
)

func (u *userService)GetUser(email string) (*dormain.User, *utils.AppError) {
	user, err := dormain.GetUser(email)

	if err!= nil {
		return nil, err
	}

	return user,nil
}

func (u *userService)GetUsers() (*[]dormain.User, *utils.AppError) {
	users, err := dormain.GetUsers()

	if err!= nil {
		return nil, err
	}

	return users, nil
}

func (u *userService)AddUser(user dormain.User) (*utils.SuccessResponse, *utils.AppError) {
	resp, err := dormain.AddUser(user)

	if err!=nil {
		return nil, err
	}

	return resp, nil
}