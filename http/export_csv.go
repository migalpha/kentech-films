package http

import (
	"encoding/csv"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	film "github.com/migalpha/kentech-films"
)

type ExportCSVHandler struct {
	Repo film.FilmProvider
}

func (handler ExportCSVHandler) ServeHTTP(ctx *gin.Context) {
	films, err := handler.Repo.GetFilms(map[string]interface{}{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fileName := "films.csv"
	f, err := os.Create(fileName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer f.Close()

	var records [][]string
	records = append(records, []string{"ID", "STARRING", "DIRECTOR", "GENRE", "SYPNOSIS", "TITLE", "RELEASE_YEAR", "CREATED_BY"})
	for _, f := range films.Films {
		records = append(records, []string{strconv.Itoa(int(f.ID.Uint())), f.Starring, f.Director, f.Genre, f.Sypnosis, f.Title, strconv.Itoa(f.ReleaseYear), strconv.Itoa(int(f.CreatedBy))})
	}

	w := csv.NewWriter(f)
	err = w.WriteAll(records)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.File(fileName)
}
