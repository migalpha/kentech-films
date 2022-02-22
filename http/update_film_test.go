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

func Test_UpdateFilmHandler_ServeHTTP(t *testing.T) {
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
		mockUpdater := mocks.FilmUpdater{}
		mockUpdater.On("Update", mock.Anything, mock.Anything).Return(nil)
		handler := UpdateFilmHandler{Provider: &mockProvider, Updater: &mockUpdater}

		url := "/films/:id"
		var jsonData = []byte(`{
			"starring": "Brad Pitt, Christoph Waltz, Michael Fassbender, Julie Dreyfus",
			"release_year": 2010
		}`)

		w := httptest.NewRecorder()
		gin.SetMode(gin.TestMode)
		c, r := gin.CreateTestContext(w)
		r.Use(func(c *gin.Context) {
			c.Set("user_id", uint(1))
		})
		r.PATCH(url, handler.ServeHTTP)
		c.Request, _ = http.NewRequest(http.MethodPatch, "/films/2", bytes.NewBuffer(jsonData))
		r.ServeHTTP(w, c.Request)

		assert.Equal(t, http.StatusOK, w.Code)
	})
}
