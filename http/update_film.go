package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	film "github.com/migalpha/kentech-films"
)

type UpdateFilmRequest struct {
	Starring    string `json:"starring,omitempty" binding:"max=255" example:"Brad Pitt, Christoph Waltz, Michael Fassbender"`
	Director    string `json:"director,omitempty" binding:"max=255" example:"Quentin Tarantino"`
	Genre       string `json:"genre,omitempty" binding:"max=255" example:"action, comedy, war"`
	Sypnosis    string `json:"sypnosis,omitempty"`
	Title       string `json:"title,omitempty" binding:"max=255" example:"Inglourious Basterds"`
	ReleaseYear int    `json:"release_year,omitempty" example:"2009"`
}

type UpdateFilmHandler struct {
	Provider film.FilmProvider
	Updater  film.FilmUpdater
}

func (handler UpdateFilmHandler) ServeHTTP(ctx *gin.Context) {
	body := UpdateFilmRequest{}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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
		ctx.JSON(http.StatusForbidden, gin.H{"error": film.ErrUserForbiddenUpdate.Error()})
		return
	}

	err = handler.Updater.Update(film.Film{
		ID:          filmID,
		Starring:    body.Starring,
		Director:    body.Director,
		Genre:       body.Genre,
		Sypnosis:    body.Sypnosis,
		Title:       body.Title,
		ReleaseYear: body.ReleaseYear,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"distances": "resp",
	})
}
