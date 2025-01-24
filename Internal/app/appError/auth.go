package appError

import "errors"

var (
	ErrEmptyValue                  = errors.New("value: empty value")
	ErrBindBodyWithJSON            = errors.New("bind: error binding body with JSON")
	ErrUserNotFound                = errors.New("user: user not found")
	ErrUserIDNotFound              = errors.New("user: user ID not found")
	ErrSignup                      = errors.New("signup: error during signup")
	ErrSignupUserExists            = errors.New("signup: user already exists")
	ErrSignin                      = errors.New("signin: error during signin")
	ErrSigninSignInIncompleteForm  = errors.New("signin: error incomplete form")
	ErrSigninUserNotFound          = errors.New("signin: user not found")
	ErrSigninPasswordIncorrect     = errors.New("signin: incorrect password")
	ErrDelIncompleteForm           = errors.New("delAccount: error incomplete form")
	ErrDelAccount                  = errors.New("delAccount: error during account deletion")
	ErrDelAccountUserNotFound      = errors.New("delAccount: user not found")
	ErrDelAccountPasswordIncorrect = errors.New("delAccount: incorrect password")
)
