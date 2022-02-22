package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	film "github.com/migalpha/kentech-films"
)

type createFilmRequest struct {
	Starring    string `json:"starring" binding:"required,max=255" example:"Brad Pitt, Christoph Waltz, Michael Fassbender"`
	Director    string `json:"director" binding:"required,max=255" example:"Quentin Tarantino"`
	Genre       string `json:"genre" binding:"required,max=255" example:"action, comedy, war"`
	Sypnosis    string `json:"sypnosis"`
	Title       string `json:"title" binding:"required,max=255" example:"Inglourious Basterds"`
	ReleaseYear int    `json:"release_year" binding:"required" example:"2009"`
	CreatedBy   uint   `json:"created_by"`
}

type createFilmResponse struct {
	ID          uint   `json:"id"`
	Starring    string `json:"starring"`
	Director    string `json:"director"`
	Genre       string `json:"genre"`
	Sypnosis    string `json:"sypnosis"`
	Title       string `json:"title"`
	ReleaseYear int    `json:"release_year"`
	CreatedBy   uint   `json:"created_by"`
}

type CreateFilmHandler struct {
	Repo film.FilmSaver
}

// @BasePath /api/v1
// Create film godoc
// @Summary Create a new film.
// @Description Allow to register a new film .
// @Tags Films
// @Accept json
// @Produce json
// @Param        film  body      createFilmRequest  true  "create film"
// @Success 200 {object} createFilmResponse
// @Failure 400 {object} errorResponse "error 400"
// @Failure 500 {object} errorResponse "error 500"
// @Router /films [post]
func (handler CreateFilmHandler) ServeHTTP(ctx *gin.Context) {
	reqctx := ctx.Request.Context()
	body := createFilmRequest{}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := getUserIDFromRequest(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": film.ErrUserNotFound.Error()})
		return
	}

	filmID, err := handler.Repo.Save(reqctx, film.Film{
		Starring:    body.Starring,
		Director:    body.Director,
		Genre:       body.Genre,
		Sypnosis:    body.Sypnosis,
		Title:       body.Title,
		ReleaseYear: body.ReleaseYear,
		CreatedBy:   userID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"film": createFilmResponse{
		ID:          filmID.Uint(),
		Starring:    body.Starring,
		Director:    body.Director,
		Genre:       body.Genre,
		Sypnosis:    body.Sypnosis,
		Title:       body.Title,
		ReleaseYear: body.ReleaseYear,
		CreatedBy:   userID,
	}})
}
