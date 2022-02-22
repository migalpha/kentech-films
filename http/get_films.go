package http

import (
	"net/http"

	film "github.com/migalpha/kentech-films"

	"github.com/gin-gonic/gin"
)

type GetFilmsHandler struct {
	Repo film.FilmProvider
}

// @BasePath /api/v1
// Get films godoc
// @Summary Get all films.
// @Description Allow to get all films by some filters.
// @Tags Films
// @Accept json
// @Produce json
// @Success 200 {array} film.Film "Returns films"
// @Failure 400 {object} errorResponse "error 400"
// @Failure 500 {object} errorResponse "error 500"
// @Router /films [get]
func (handler GetFilmsHandler) ServeHTTP(ctx *gin.Context) {
	reqctx := ctx.Request.Context()
	criteria := []string{"genre", "release_year", "title"}
	filters := make(map[string]interface{})

	for _, c := range criteria {
		if ctx.Query(c) != "" {
			filters[c] = ctx.Query(c)
		}
	}

	films, err := handler.Repo.GetFilms(reqctx, filters)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": films,
	})
}
