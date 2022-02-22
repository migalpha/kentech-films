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

func Test_RegisterUserHandler_ServeHTTP(t *testing.T) {
	user := film.User{
		ID: 1,
	}

	t.Run("Happy path", func(t *testing.T) {
		mockSaver := mocks.UserSaver{}
		mockSaver.On("Save", mock.Anything, mock.Anything).Return(user.ID, nil)
		handler := RegisterUserHandler{Repo: &mockSaver}

		url := "/register"
		var jsonData = []byte(`{"username": "test12345", "password": "test1234"}`)

		w := httptest.NewRecorder()
		gin.SetMode(gin.TestMode)
		c, r := gin.CreateTestContext(w)

		r.POST(url, handler.ServeHTTP)
		c.Request, _ = http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
		r.ServeHTTP(w, c.Request)

		assert.Equal(t, http.StatusCreated, w.Code)
	})
}
