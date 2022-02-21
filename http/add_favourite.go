package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	film "github.com/migalpha/kentech-films"
)

type AddFavouriteRequest struct {
	FilmID uint `json:"film_id"`
}

type AddFavouriteHandler struct {
	Saver    film.FavouriteSaver
	Provider film.FilmProvider
}

func (handler AddFavouriteHandler) ServeHTTP(ctx *gin.Context) {
	body := AddFavouriteRequest{}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := getUserIDFromRequest(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": film.ErrUserNotFound.Error()})
		return
	}

	filmID := film.FilmID(body.FilmID)

	_, err = handler.Provider.FilmbyID(filmID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = handler.Saver.Save(film.Favourite{
		FilmID: filmID,
		UserID: film.UserID(userID),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}
