package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	film "github.com/migalpha/kentech-films"
)

type DeleteFilmHandler struct {
	Provider  film.FilmProvider
	Destroyer film.FilmDestroyer
}

// @BasePath /api/v1
// Delete Film godoc
// @Summary Destroy a film from records.
// @Description Allow to remove a film from records.
// @Tags Films
// @Accept json
// @Produce json
// @Param        film_id  path      int  true  "Destroy film"
// @Success 200 {object} emptyResponse
// @Failure 400 {object} errorResponse "error 400"
// @Failure 500 {object} errorResponse "error 500"
// @Router /films/{film_id} [delete]
func (handler DeleteFilmHandler) ServeHTTP(ctx *gin.Context) {
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
		ctx.JSON(http.StatusForbidden, gin.H{"error": film.ErrUserForbiddenDestroy.Error()})
		return
	}

	err = handler.Destroyer.Destroy(filmID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}

func getUserIDFromRequest(ctx *gin.Context) (uint, error) {
	userID, ok := ctx.Get("user_id")
	if !ok {
		return 0, film.ErrUserNotFoundToken
	}

	return userID.(uint), nil
}
