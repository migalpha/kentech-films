package http

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	film "github.com/migalpha/kentech-films"
	"github.com/migalpha/kentech-films/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_LoginHandler_ServeHTTP(t *testing.T) {
	user := film.User{
		ID:        1,
		Username:  "test12345",
		Password:  "$2a$14$fEIRjfJaRCANVw9f3ibjtu35u4HS48h1AZIFopYdG/pzXvx2ssORi",
		IsActive:  true,
		CreatedAt: time.Now(),
	}

	t.Run("Happy path", func(t *testing.T) {
		mockProvider := mocks.UserProvider{}
		mockProvider.On("UserbyUsername", mock.Anything).Return(user, nil)
		handler := LoginHandler{Repo: &mockProvider}

		url := "/register"
		var jsonData = []byte(`{"username": "test12345", "password": "test1234"}`)

		w := httptest.NewRecorder()
		gin.SetMode(gin.TestMode)
		c, r := gin.CreateTestContext(w)

		r.POST(url, handler.ServeHTTP)
		c.Request, _ = http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
		r.ServeHTTP(w, c.Request)

		assert.Equal(t, http.StatusOK, w.Code)
	})
}
