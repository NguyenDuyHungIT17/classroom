package common

const (
	SUCCESS_CODE = 0
	SUCCESS_MESS = "DONE"

	INVALID_REQUEST_CODE = 1000
	INVALID_REQUEST_MESS = "invalid request"

	DB_ERROR_CODE = 1001
	DB_ERROR_MESS = "db error"

	SEND_EMAIL_ERROR_CODE = 1001
	SEND_EMAIL_ERROR_MESS = "send email error"
)

// user error
const (
	USER_EXISTED_CODE = 1200
	USER_EXISTED_MESS = "user existed"

	USER_IS_NOT_EXIST_CODE = 1201
	USER_IS_NOT_EXIST_MESS = "user is not exist"

	PASSWORD_IS_WRONG_CODE = 1202
	PASSWORD_IS_WRONG_MESS = "password is wrong"
)
