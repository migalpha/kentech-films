package http

import (
	"net/http"

	film "github.com/migalpha/kentech-films"

	"github.com/gin-gonic/gin"
)

type GetFilmDetailsHandler struct {
	Repo film.FilmProvider
}

func (handler GetFilmDetailsHandler) ServeHTTP(ctx *gin.Context) {
	ID := ctx.Param("id")

	filmID, err := film.NewFilmID(ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	filmDB, err := handler.Repo.FilmbyID(filmID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"film": filmDB,
	})
}
