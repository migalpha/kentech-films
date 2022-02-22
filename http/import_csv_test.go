package http

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	film "github.com/migalpha/kentech-films"
	"github.com/migalpha/kentech-films/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_ImportCSVHandler_ServeHTTP(t *testing.T) {
	t.Run("Happy path", func(t *testing.T) {
		mockFilm := film.Film{
			ID: 0,
		}
		mockSaver := mocks.FilmSaver{}
		mockSaver.On("Save", mock.Anything, mock.Anything).Return(mockFilm.ID, nil)
		handler := ImportCSVHandler{Repo: &mockSaver}

		filePath := "./testdata/films.csv"
		fieldName := "file"
		body := new(bytes.Buffer)

		mw := multipart.NewWriter(body)

		file, err := os.Open(filePath)
		if err != nil {
			t.Fatal(err)
		}

		wr, err := mw.CreateFormFile(fieldName, filePath)
		if err != nil {
			t.Fatal(err)
		}

		if _, err := io.Copy(wr, file); err != nil {
			t.Fatal(err)
		}
		mw.Close()

		url := "/csv/films"
		w := httptest.NewRecorder()
		gin.SetMode(gin.TestMode)
		c, r := gin.CreateTestContext(w)
		r.Use(func(c *gin.Context) {
			c.Set("user_id", uint(1))
		})
		r.POST(url, handler.ServeHTTP)
		c.Request, _ = http.NewRequest(http.MethodPost, url, body)
		c.Request.Header.Add("Content-Type", mw.FormDataContentType())
		r.ServeHTTP(w, c.Request)

		assert.Equal(t, http.StatusCreated, w.Code)
	})
}
