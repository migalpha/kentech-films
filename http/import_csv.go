package http

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	film "github.com/migalpha/kentech-films"
)

const MaxConcurrency = 10

type filmResult struct {
	ID   int
	Film film.Film
	Err  error
}
type ImportCSVHandler struct {
	Repo film.FilmSaver
}

// @BasePath /api/v1
// Import films godoc
// @Summary Import films data from csv file.
// @Description Import films data from csv file.
// @Tags Films
// @Accept plain
// @Produce json
// @Success 201 {object} emptyResponse
// @Failure 400 {object} errorResponse "error 400"
// @Failure 500 {object} errorResponse "error 500"
// @Router /api/v1/csv/films [post]
func (handler ImportCSVHandler) ServeHTTP(ctx *gin.Context) {
	reqctx := ctx.Request.Context()
	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userID, err := getUserIDFromRequest(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	file := form.File["file"]
	f, err := file[0].Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	r := csv.NewReader(f)

	records, err := r.ReadAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	wg := sync.WaitGroup{}
	limiter := make(chan struct{}, MaxConcurrency)
	results := make(chan filmResult, len(records))

	for i, r := range records[1:] {
		limiter <- struct{}{}
		wg.Add(1)
		go func(r []string, id int) {
			filmToSave, err := filmValidate(r)
			filmToSave.CreatedBy = userID
			<-limiter
			results <- filmResult{
				ID:   id,
				Film: filmToSave,
				Err:  err,
			}
			wg.Done()
		}(r, i)
	}

	wg.Wait()
	close(limiter)
	close(results)

	for res := range results {
		if res.Err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("row: %d, err:%s", res.ID+1, res.Err.Error())})
			return
		}
		_, err = handler.Repo.Save(reqctx, res.Film)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("row: %d, err:%s", res.ID+1, err.Error())})
		}
	}
	ctx.JSON(http.StatusCreated, gin.H{})
}

func filmValidate(f []string) (film.Film, error) {
	if f[1] == "" {
		return film.Film{}, film.ErrEmptyFilmDirector
	}
	if f[2] == "" {
		return film.Film{}, film.ErrEmptyFilmGenre
	}
	if f[4] == "" {
		return film.Film{}, film.ErrEmptyFilmTitle
	}
	if f[5] == "" {
		return film.Film{}, film.ErrEmptyFilmReleaseYear
	}
	year, err := strconv.Atoi(f[5])
	if err != nil {
		return film.Film{}, err
	}
	return film.Film{
		Starring:    f[0],
		Director:    f[1],
		Genre:       f[2],
		Sypnosis:    f[3],
		Title:       f[4],
		ReleaseYear: year,
	}, nil
}
