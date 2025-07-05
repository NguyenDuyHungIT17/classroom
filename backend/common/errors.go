package common

const (
	SUCCESS_CODE = 0
	SUCCESS_MESS = "DONE"

	INVALID_REQUEST_CODE = 1000
	INVALID_REQUEST_MESS = "invalid request"

	DB_ERROR_CODE = 1001
	DB_ERROR_MESS = "db error"
)

// user error
const (
	USER_EXISTED_CODE = 1200
	USER_EXISTED_MESS = "user existed"

	SEND_EMAIL_ERROR_CODE = 1201
	SEND_EMAIL_ERROR_MESS = "send email error"

	USER_IS_NOT_EXIST_CODE = 1202
	USER_IS_NOT_EXIST_MESS = "user is not exist"

	PASSWORD_IS_WRONG_CODE = 1203
	PASSWORD_IS_WRONG_MESS = "password is wrong"

	INVALID_TOKEN_CODE = 1204
	INVALID_TOKEN_MESS = "forbidden access"

	INVALID_SESSION_USER_CODE = 1205
	INVALID_SESSION_USER_MESS = "invalid session"
)

// teacher error
const (
	ADD_TEACHER_ERROR_CODE = 1300
	ADD_TEACHER_ERROR_MESS = "teacher can not add"
)
