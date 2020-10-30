package service

import (
	"github.com/raymondgitonga/nubi_service/email_service/dormain"
	"github.com/raymondgitonga/nubi_service/email_service/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

type userDaoMock struct{}

var (
	getUserFunction  func(email string) (*dormain.User, *utils.AppError)
	getUsersFunction func() (*[]dormain.User, *utils.AppError)
	addUserFunction  func(user dormain.User) (*utils.SuccessResponse, *utils.AppError)
	user             = dormain.User{Name: "Raymond", Email: "raytosh95@gmail.com"}
	err              = utils.AppError{Message: "user not found", StatusCode: 500}
)

func init() {
	dormain.UserDaoInterface = &userDaoMock{}
}

func (u userDaoMock) GetUser(email string) (*dormain.User, *utils.AppError) {
	return getUserFunction(email)
}
func (u userDaoMock) GetUsers() (*[]dormain.User, *utils.AppError) {
	return getUsersFunction()
}

func (u userDaoMock) AddUser(user dormain.User) (*utils.SuccessResponse, *utils.AppError) {
	return addUserFunction(user)
}

func TestUserService_GetUser(t *testing.T) {
	getUserFunction = func(email string) (*dormain.User, *utils.AppError) {
		return &user, nil
	}

	user, err := UserServiceInterface.GetUser("raytosh95@gmail.com")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "raytosh95@gmail.com", user.Email)
}

func TestUserService_GetUserError(t *testing.T) {
	getUserFunction = func(email string) (*dormain.User, *utils.AppError) {
		return nil, &err
	}

	user, err := UserServiceInterface.GetUser("raytosh95@gmail.com")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.Equal(t, 500, err.StatusCode)
}

func TestUserService_GetUsers(t *testing.T) {
	getUsersFunction = func() (*[]dormain.User, *utils.AppError) {
		return &[]dormain.User{
			user, user,
		}, nil
	}

	users, err := UserServiceInterface.GetUsers()

	assert.NotNil(t, users)
	assert.Nil(t, err)
	assert.Equal(t, "Raymond", (*users)[0].Name)
}

func TestUserService_GetUsersError(t *testing.T) {
	getUsersFunction = func() (*[]dormain.User, *utils.AppError) {
		return nil, &err
	}

	users, err := UserServiceInterface.GetUsers()

	assert.NotNil(t, err)
	assert.Nil(t, users)
	assert.Equal(t, 500, err.StatusCode)
}

func TestUserService_AddUser(t *testing.T) {
	addUserFunction = func(user dormain.User) (*utils.SuccessResponse, *utils.AppError) {
		return &utils.SuccessResponse{
			Message:    "successfully added user",
			StatusCode: 200,
		}, nil
	}

	user = user
	resp, err := UserServiceInterface.AddUser(user)

	assert.NotNil(t, resp)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUserService_AddUserError(t *testing.T) {
	addUserFunction = func(user dormain.User) (*utils.SuccessResponse, *utils.AppError) {
		return nil, &err
	}
	user := dormain.User{
		Name:  " ",
		Email: " ",
	}
	resp, err := UserServiceInterface.AddUser(user)

	assert.NotNil(t, err)
	assert.Nil(t, resp)
	assert.Equal(t, 500, err.StatusCode)
}
