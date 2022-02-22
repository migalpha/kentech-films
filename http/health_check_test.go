package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_HealthCheck(t *testing.T) {
	router := gin.New()
	router.GET("/health", HealthCheck)

	t.Run("Successful request", func(t *testing.T) {
		expResponse := `{"status":"OK"}`
		url := "/health"
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", url, nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, expResponse, w.Body.String())
	})
}
