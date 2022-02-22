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

func Test_CreateFilmHandler_ServeHTTP(t *testing.T) {
	mockFilm := film.Film{
		ID: 1,
	}

	t.Run("Happy path", func(t *testing.T) {
		mockSaver := mocks.FilmSaver{}
		mockSaver.On("Save", mock.Anything, mock.Anything).Return(mockFilm.ID, nil)
		handler := CreateFilmHandler{Repo: &mockSaver}

		url := "/films"
		var jsonData = []byte(`{
			"starring": "Brad Pitt, Christoph Waltz, Michael Fassbender",
			"director": "Quentin Tarantino",
			"genre": "action, comedy, war",
			"title": "Inglourious Basterds",
			"release_year": 2009
		}`)

		w := httptest.NewRecorder()
		gin.SetMode(gin.TestMode)
		c, r := gin.CreateTestContext(w)
		r.Use(func(c *gin.Context) {
			c.Set("user_id", uint(1))
		})
		r.POST(url, handler.ServeHTTP)
		c.Request, _ = http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
		r.ServeHTTP(w, c.Request)

		assert.Equal(t, http.StatusCreated, w.Code)
	})
}
