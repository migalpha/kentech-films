package film

import "errors"

var (
	ErrAuthCheckTokenFail    = errors.New("invalid token")
	ErrAuthCheckTokenTimeout = errors.New("token expired")
	ErrBlacklistToken        = errors.New("token is blacklisted")
	ErrBlacklistCheckToken   = errors.New("can't check if token is blacklisted")
	ErrCreateFile            = errors.New("can't create file")
	ErrInvalidFilmID         = errors.New("invalid film id")
	ErrMissingToken          = errors.New("missing token")
	ErrEmptyFilmDirector     = errors.New("empty film director")
	ErrEmptyFilmGenre        = errors.New("empty film genre")
	ErrEmptyFilmTitle        = errors.New("empty film title")
	ErrEmptyFilmReleaseYear  = errors.New("empty film release year")
	ErrUserNotFound          = errors.New("user not found")
	ErrUserNotFoundToken     = errors.New("user not found in token")
	ErrUserForbiddenDestroy  = errors.New("this user can't delete this film")
	ErrUserForbiddenUpdate   = errors.New("this user can't update this film")
	ErrWrongCredentials      = errors.New("wrong credentials")
)
