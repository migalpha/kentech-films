package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	film "github.com/migalpha/kentech-films"
)

type RemoveFavouriteHandler struct {
	Provider  film.FilmProvider
	Destroyer film.FavouriteDestroyer
}

// @BasePath /api/v1
// Remove Favourite godoc
// @Summary Remove a film from favourites
// @Description Allow to remove a film from favourites list.
// @Tags Favourites
// @Accept json
// @Produce json
// @Param        film_id  path      int  true  "Remove favourites"
// @Success 200 {object} emptyResponse
// @Failure 400 {object} errorResponse "error 400"
// @Failure 500 {object} errorResponse "error 500"
// @Router /api/v1/favourites/{film_id} [delete]
func (handler RemoveFavouriteHandler) ServeHTTP(ctx *gin.Context) {
	reqctx := ctx.Request.Context()
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

	filmFromDB, err := handler.Provider.FilmbyID(reqctx, filmID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if userID != filmFromDB.CreatedBy {
		ctx.JSON(http.StatusForbidden, gin.H{"error": film.ErrUserForbiddenDestroy.Error()})
		return
	}

	err = handler.Destroyer.Destroy(reqctx, filmID, film.UserID(userID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, emptyResponse{})
}

type emptyResponse struct {
}
