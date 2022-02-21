package http

import (
	"net/http"

	film "github.com/migalpha/kentech-films"

	"github.com/gin-gonic/gin"
)

type GetFilmsHandler struct {
	Repo film.FilmProvider
}

func (handler GetFilmsHandler) ServeHTTP(ctx *gin.Context) {
	criteria := []string{"genre", "release_year", "title"}
	filters := make(map[string]interface{})

	for _, c := range criteria {
		if ctx.Query(c) != "" {
			filters[c] = ctx.Query(c)
		}
	}

	films, err := handler.Repo.GetFilms(filters)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": films,
	})
}
