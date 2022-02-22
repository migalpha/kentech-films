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

func Test_ExportCSVHandler_ServeHTTP(t *testing.T) {
	filmMock := film.Films{
		Films: []film.Film{
			{
				ID:          1,
				Starring:    "test",
				Director:    "test",
				Genre:       "test",
				Sypnosis:    "test",
				Title:       "test",
				ReleaseYear: 1234,
				CreatedBy:   1,
			},
			{
				ID:          2,
				Starring:    "test2",
				Director:    "test2",
				Genre:       "test2",
				Sypnosis:    "test2",
				Title:       "test2",
				ReleaseYear: 4321,
				CreatedBy:   2,
			},
		},
	}

	t.Run("Happy path", func(t *testing.T) {
		mockProvider := mocks.FilmProvider{}
		mockProvider.On("GetFilms", mock.Anything).Return(filmMock, nil)
		handler := ExportCSVHandler{Repo: &mockProvider}

		url := "/csv/films"
		w := httptest.NewRecorder()
		gin.SetMode(gin.TestMode)
		c, r := gin.CreateTestContext(w)
		r.Use(func(c *gin.Context) {
			c.Set("user_id", uint(1))
		})
		r.GET(url, handler.ServeHTTP)
		c.Request, _ = http.NewRequest(http.MethodGet, url, nil)
		r.ServeHTTP(w, c.Request)

		assert.Equal(t, http.StatusOK, w.Code)
	})
}
