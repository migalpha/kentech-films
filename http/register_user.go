package http

import (
	"net/http"
	"time"
	"unicode"

	film "github.com/migalpha/kentech-films"
	"github.com/migalpha/kentech-films/config"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type registerUserRequest struct {
	Username string `json:"username" binding:"required,alphanum,min=5,max=75" example:"user21"`
	Password string `json:"password" binding:"required,alphanum,min=5,max=75" example:"secret"`
}

func (ru registerUserRequest) isStartWithDigit() bool {
	u := []rune(ru.Username)
	return unicode.IsDigit(u[0])
}

type RegisterUserHandler struct {
	Repo film.UserSaver
}

func (handler RegisterUserHandler) ServeHTTP(ctx *gin.Context) {
	body := registerUserRequest{}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if body.isStartWithDigit() {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "username must start with letter."})
		return
	}

	hashPassword, err := encodePasword(body.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = handler.Repo.Save(film.User{
		Username:  film.Username(body.Username),
		Password:  hashPassword,
		IsActive:  true,
		CreatedAt: time.Now(),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}

func encodePasword(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), config.Commons().BCryptCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
