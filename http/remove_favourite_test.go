package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	film "github.com/migalpha/kentech-films"
	"github.com/migalpha/kentech-films/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_RemoveFavouriteHandler_ServeHTTP(t *testing.T) {
	filmMock := film.Film{
		ID:          2,
		Starring:    "test",
		Director:    "test",
		Genre:       "test",
		Sypnosis:    "test",
		Title:       "test",
		ReleaseYear: 1234,
		CreatedBy:   1,
	}

	t.Run("Happy path", func(t *testing.T) {
		mockProvider := mocks.FilmProvider{}
		mockProvider.On("FilmbyID", mock.Anything, mock.Anything).Return(filmMock, nil)
		mockDestroyer := mocks.FavouriteDestroyer{}
		mockDestroyer.On("Destroy", mock.Anything, mock.Anything, mock.Anything).Return(nil)
		handler := RemoveFavouriteHandler{Provider: &mockProvider, Destroyer: &mockDestroyer}

		url := "/favourites/:id"

		w := httptest.NewRecorder()
		gin.SetMode(gin.TestMode)
		c, r := gin.CreateTestContext(w)
		r.Use(func(c *gin.Context) {
			c.Set("user_id", uint(1))
		})
		r.DELETE(url, handler.ServeHTTP)
		c.Request, _ = http.NewRequest(http.MethodDelete, "/favourites/2", nil)
		r.ServeHTTP(w, c.Request)

		assert.Equal(t, http.StatusOK, w.Code)
	})
}
