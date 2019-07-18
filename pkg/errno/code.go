package errno

var (
	// OK Common errors
	OK = &Errno{Code: 0, Message: "OK"}
	// InternalServerError internal server error Code
	InternalServerError = &Errno{Code: 10001, Message: "Internal Server Error."}
	// ErrBind binding error
	ErrBind = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}

	ErrValidation = &Errno{Code: 20001, Message: "Validation failed."}
	ErrDatabase   = &Errno{Code: 20002, Message: "Database error."}
	ErrToken      = &Errno{Code: 20003, Message: "Error occurred while signing the JSON web token."}

	// user errors

	ErrEncrypt           = &Errno{Code: 20101, Message: "Error occurred while encrypting the user password."}
	ErrUserNotFound      = &Errno{Code: 20102, Message: "The user was not found."}
	ErrUsersNone         = &Errno{Code: 20102, Message: "Cannot get users."}
	ErrTokenInvalid      = &Errno{Code: 401, Message: "The token was invalid."}
	ErrPasswordIncorrect = &Errno{Code: 20104, Message: "The password was incorrect."}
	ErrUpdatePassword     = &Errno{Code: 20105, Message: "Error occurred while updating the password."}

	// role errors
	ErrAccessLevel = &Errno{Code: 30001, Message: "You don't have the privilege."}
)
