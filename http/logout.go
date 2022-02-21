package http

import (
	"net/http"

	film "github.com/migalpha/kentech-films"
	"github.com/migalpha/kentech-films/jwt"

	"github.com/gin-gonic/gin"
)

type LogoutHandler struct {
	Repo film.TokenSaver
}

func (handler LogoutHandler) ServeHTTP(ctx *gin.Context) {
	reqctx := ctx.Request.Context()
	headerToken := ctx.Request.Header["Authorization"]
	signature := jwt.GetTokenSignature(headerToken)

	err := handler.Repo.Save(reqctx, signature)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"signature": signature,
	})
}
