package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	film "github.com/migalpha/kentech-films"
)

type DeleteFilmHandler struct {
	Provider  film.FilmProvider
	Destroyer film.FilmDestroyer
}

func (handler DeleteFilmHandler) ServeHTTP(ctx *gin.Context) {
	ID := ctx.Param("id")

	userID, err := getUserIDFromRequest(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filmID, err := film.NewFilmID(ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	filmFromDB, err := handler.Provider.FilmbyID(filmID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if userID != filmFromDB.CreatedBy {
		ctx.JSON(http.StatusForbidden, gin.H{"error": film.ErrUserForbiddenDestroy.Error()})
		return
	}

	err = handler.Destroyer.Destroy(filmID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}

func getUserIDFromRequest(ctx *gin.Context) (uint, error) {
	userID, ok := ctx.Get("user_id")
	if !ok {
		return 0, film.ErrUserNotFoundToken
	}

	return userID.(uint), nil
}
