package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

    "github.com/lfernando-silva/sgra-rest-go/util/types"
    "github.com/lfernando-silva/sgra-rest-go/util/formatters"
)

func SetUpRouter() *gin.Engine{
    router := gin.Default()
    return router
}

func TestHealthcheckHandler(t *testing.T) {
    mockResponse := types.THealthCheck {
        Message: "OK",
        Timestamp: 123,
    }
    var response types.THealthCheck
    r := SetUpRouter()
    r.GET("/", HealthCheck)
    req, _ := http.NewRequest("GET", "/", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    formatters.ParseTestResponse(w, &response)
    assert.Equal(t, mockResponse.Message, response.Message)
    assert.Equal(t, http.StatusOK, w.Code)
}