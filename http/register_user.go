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

type registerResponse struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
}

type RegisterUserHandler struct {
	Repo film.UserSaver
}

// @BasePath /api/v1
// Register user godoc
// @Summary Allow users to register to consume this API.
// @Description Register user inside API.
// @Tags Users
// @Accept json
// @Produce json
// @Param        register  body      registerUserRequest  true  "Register user"
// @Success 200 {object} registerResponse
// @Failure 400 {object} errorResponse "error 400"
// @Failure 500 {object} errorResponse "error 500"
// @Router /register [post]
func (handler RegisterUserHandler) ServeHTTP(ctx *gin.Context) {
	reqctx := ctx.Request.Context()
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

	now := time.Now()
	userID, err := handler.Repo.Save(reqctx, film.User{
		Username:  film.Username(body.Username),
		Password:  hashPassword,
		IsActive:  true,
		CreatedAt: now,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"user": registerResponse{
		ID:        userID.Uint(),
		Username:  body.Username,
		IsActive:  true,
		CreatedAt: now,
	}})
}

func encodePasword(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), config.Commons().BCryptCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
