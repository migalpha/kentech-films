package http

import (
	"net/http"

	film "github.com/migalpha/kentech-films"
	"github.com/migalpha/kentech-films/jwt"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginHandler struct {
	Repo film.UserProvider
}

type loginRequest struct {
	Username string `json:"username" binding:"required,alphanum,min=5,max=75" example:"user21"`
	Password string `json:"password" binding:"required,alphanum,min=5,max=75" example:"secret"`
}

type loginResponse struct {
	Token string `jaon:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"`
}

// @BasePath /api/v1
// Login godoc
// @Summary Get valid credentials.
// @Description Returns a valid token if credentials are right.
// @Tags Users
// @Accept json
// @Produce json
// @Param        login  body      loginRequest  true  "username & password"
// @Success 200 {object} loginResponse "Returns a token"
// @Failure 400 {object} errorResponse "error 400"
// @Failure 500 {object} errorResponse "error 500"
// @Router /login [post]
func (handler LoginHandler) ServeHTTP(ctx *gin.Context) {
	body := loginRequest{}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := handler.Repo.UserbyUsername(film.Username(body.Username))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = checkEncodePassword(body.Password, user.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": film.ErrWrongCredentials.Error()})
		return
	}

	token, err := jwt.GenerateToken(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, loginResponse{Token: token})
}

func checkEncodePassword(passReq, passDB string) error {
	return bcrypt.CompareHashAndPassword([]byte(passDB), []byte(passReq))
}
