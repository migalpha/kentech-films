package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	film "github.com/migalpha/kentech-films"
)

type addFavouriteRequest struct {
	FilmID uint `json:"film_id" example:"1"`
}

type addFavouriteResponse struct {
	ID     uint `json:"id" example:"1"`
	FilmID uint `json:"film_id" example:"2"`
	UserID uint `json:"user_id" example:"7"`
}

type AddFavouriteHandler struct {
	Saver    film.FavouriteSaver
	Provider film.FilmProvider
}

// @BasePath /api/v1
// Add Favourite godoc
// @Summary Add film to favourites
// @Description Allow to add a film to favourites list.
// @Tags Favourites
// @Accept json
// @Produce json
// @Param        film_id  body      addFavouriteRequest  true  "Add favourites"
// @Success 200 {object} addFavouriteResponse "Returns object with favourite pk, film_id and user_id"
// @Failure 400 {object} errorResponse "error 400"
// @Failure 500 {object} errorResponse "error 500"
// @Router /favourites [post]
func (handler AddFavouriteHandler) ServeHTTP(ctx *gin.Context) {
	reqctx := ctx.Request.Context()
	body := addFavouriteRequest{}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := getUserIDFromRequest(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": film.ErrUserNotFound.Error()})
		return
	}

	filmID := film.FilmID(body.FilmID)

	_, err = handler.Provider.FilmbyID(reqctx, filmID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	favouriteID, err := handler.Saver.Save(reqctx, film.Favourite{
		FilmID: filmID,
		UserID: film.UserID(userID),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{"favourite": addFavouriteResponse{
		ID:     favouriteID.Uint(),
		FilmID: filmID.Uint(),
		UserID: userID,
	}})
}

type errorResponse struct {
	Error string `json:"error"`
}
