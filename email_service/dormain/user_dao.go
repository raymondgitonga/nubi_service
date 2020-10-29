package dormain

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/raymondgitonga/nubi_service/email_service/utils"
	"net/http"
)

type userDao struct{}

type userDaoInterface interface {
	GetUser(email string) (*User, *utils.AppError)
	GetUsers() (*[]User, *utils.AppError)
	AddUser(user User) (*utils.SuccessResponse, *utils.AppError)
}

var (
	UserDaoInterface userDaoInterface
)

func init() {
	UserDaoInterface = &userDao{}
}

func (u userDao) GetUser(email string) (*User, *utils.AppError) {
	var singleUser User

	if err := utils.DBClient.Get(&singleUser, "SELECT name, email FROM nubi_email WHERE email = ?;", email); err != nil {
		if err == sql.ErrNoRows {
			return nil, &utils.AppError{
				Message:    utils.UserNotFound,
				StatusCode: http.StatusNotFound,
			}
		}
		return nil, &utils.AppError{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	return &singleUser, nil
}

func (u userDao) GetUsers() (*[]User, *utils.AppError) {
	var users []User

	if err := utils.DBClient.Select(&users, "SELECT name, email from nubi_email;"); err != nil {
		return nil, &utils.AppError{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	return &users, nil
}

func (u userDao) AddUser(user User) (*utils.SuccessResponse, *utils.AppError) {

	res, err := utils.DBClient.Exec("INSERT INTO nubi_email (name, email) VALUES (?,?);",
		user.Name, user.Email)

	if err != nil {
		me, _ := err.(*mysql.MySQLError)
		if me.Number == 1062 {
			return nil, &utils.AppError{
				Message:    "email already registered",
				StatusCode: http.StatusBadRequest,
			}
		}
		return nil, &utils.AppError{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	id, e := res.RowsAffected()

	if e != nil {
		return nil, &utils.AppError{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	} else if id > 0 {
		return &utils.SuccessResponse{
			Message:    utils.UserSaved,
			StatusCode: http.StatusOK,
		}, nil
	}
	return nil, &utils.AppError{
		Message:    err.Error(),
		StatusCode: http.StatusInternalServerError,
	}
}
