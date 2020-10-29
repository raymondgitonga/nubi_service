package service

import (
	"github.com/raymondgitonga/nubi_service/email_service/dormain"
	"github.com/raymondgitonga/nubi_service/email_service/utils"
)

type userService struct{}

var (
	UserService userService
	errChan     = make(chan *utils.AppError, 1)
	userChan    = make(chan *dormain.User, 1)
	usersChan   = make(chan *[]dormain.User, 1)
	respChan    = make(chan *utils.SuccessResponse, 1)
)

func (u *userService) GetUser(email string) (*dormain.User, *utils.AppError) {
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

	return user, nil
}

func (u *userService) GetUsers() (*[]dormain.User, *utils.AppError) {

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

	return users, nil
}

func (u *userService) AddUser(user dormain.User) (*utils.SuccessResponse, *utils.AppError) {

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

	return resp, nil
}
