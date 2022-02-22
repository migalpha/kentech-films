package http

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	film "github.com/migalpha/kentech-films"
	"github.com/migalpha/kentech-films/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_AddFavouriteHandler_ServeHTTP(t *testing.T) {
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

	favouriteMock := film.Favourite{
		ID: 1,
	}

	t.Run("Happy path", func(t *testing.T) {
		mockProvider := mocks.FilmProvider{}
		mockProvider.On("FilmbyID", mock.Anything, mock.Anything).Return(filmMock, nil)
		mockSaver := mocks.FavouriteSaver{}
		mockSaver.On("Save", mock.Anything, mock.Anything).Return(favouriteMock.ID, nil)
		handler := AddFavouriteHandler{Provider: &mockProvider, Saver: &mockSaver}

		url := "/favourites"
		var jsonData = []byte(`{"film_id": 2}`)

		w := httptest.NewRecorder()
		gin.SetMode(gin.TestMode)
		c, r := gin.CreateTestContext(w)
		r.Use(func(c *gin.Context) {
			c.Set("user_id", uint(1))
		})
		r.POST(url, handler.ServeHTTP)
		c.Request, _ = http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
		r.ServeHTTP(w, c.Request)

		assert.Equal(t, http.StatusNoContent, w.Code)
	})
}
