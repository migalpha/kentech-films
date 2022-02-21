package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	film "github.com/migalpha/kentech-films"
)

type RemoveFavoriteHandler struct {
	Provider  film.FilmProvider
	Destroyer film.FavouriteDestroyer
}

func (handler RemoveFavoriteHandler) ServeHTTP(ctx *gin.Context) {
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

	err = handler.Destroyer.Destroy(filmID, film.UserID(userID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"distances": "resp",
	})
}
