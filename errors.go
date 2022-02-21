package film

import "errors"

var (
	ErrorAuthCheckTokenFail    = errors.New("invalid token")
	ErrorAuthCheckTokenTimeout = errors.New("token expired")
	ErrBlacklistToken          = errors.New("token is blacklisted")
	ErrBlacklistCheckToken     = errors.New("can't check if token is blacklisted")
	ErrCreateFile              = errors.New("can't create file")
	ErrInvalidFilmID           = errors.New("invalid film id")
	ErrorMissingToken          = errors.New("missing token")
	ErrUserNotFound            = errors.New("user not found")
	ErrUserNotFoundToken       = errors.New("user not found in token")
	ErrUserForbiddenDestroy    = errors.New("this user can't delete this film")
	ErrUserForbiddenUpdate     = errors.New("this user can't update this film")
	ErrWrongCredentials        = errors.New("wrong credentials")
)
