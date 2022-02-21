package http

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	film "github.com/migalpha/kentech-films"
)

type ImportCSVHandler struct {
	Repo film.FilmSaver
}

func (handler ImportCSVHandler) ServeHTTP(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("%+v", form)
	file := form.File["file"]
	fmt.Printf("%+v", file)
	f, _ := file[0].Open()
	r := csv.NewReader(f)

	records, err := r.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(records)
	ctx.JSON(http.StatusCreated, gin.H{
		"distances": "resp",
	})
}
