package service

import (
	"github.com/raymondgitonga/nubi_service/email_service/dormain"
	"github.com/raymondgitonga/nubi_service/email_service/utils"
)

type userService struct{}

var (
	UserService userService
)

func (u *userService) GetUser(email string) (*dormain.User, *utils.AppError) {
	errChan := make(chan *utils.AppError, 1)
	userChan := make(chan *dormain.User, 1)

	go func() {
		user, err := dormain.UserDaoInterface.GetUser(email)
		errChan <- err
		userChan <- user
	}()
	user := <-userChan
	err := <-errChan

	if err != nil {
		return nil, err
	}
	close(userChan)
	close(errChan)

	return user, nil
}

func (u *userService) GetUsers() (*[]dormain.User, *utils.AppError) {
	errChan := make(chan *utils.AppError, 1)
	usersChan := make(chan *[]dormain.User, 1)

	go func() {
		users, err := dormain.UserDaoInterface.GetUsers()
		usersChan <- users
		errChan <- err
	}()

	err := <-errChan
	users := <-usersChan

	if err != nil {
		return nil, err
	}
	close(usersChan)
	close(errChan)

	return users, nil
}

func (u *userService) AddUser(user dormain.User) (*utils.SuccessResponse, *utils.AppError) {

	errChan := make(chan *utils.AppError, 1)
	respChan := make(chan *utils.SuccessResponse, 1)

	go func() {
		resp, err := dormain.UserDaoInterface.AddUser(user)
		respChan <- resp
		errChan <- err
	}()

	err := <-errChan
	resp := <-respChan

	if err != nil {
		return nil, err
	}
	close(respChan)
	close(errChan)

	return resp, nil
}
