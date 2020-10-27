package utils

type AppError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

type SuccessResponse struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

var (
	UserNotFound = "user not found"
	SaveFail = "failed to save user"
	UserSaved = "successfully added user"
)