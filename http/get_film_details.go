package http

import (
	"net/http"

	film "github.com/migalpha/kentech-films"

	"github.com/gin-gonic/gin"
)

type GetFilmDetailsHandler struct {
	Repo film.FilmProvider
}

// @BasePath /api/v1
// Get film godoc
// @Summary Get a specific film.
// @Description Given a film_id returns this film.
// @Tags Films
// @Accept json
// @Produce json
// @Param        film_id  path      int  true  "Film id"
// @Success 200 {object} film.Film "Returns film"
// @Failure 400 {object} errorResponse "error 400"
// @Failure 500 {object} errorResponse "error 500"
// @Router /films/{film_id} [get]
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
