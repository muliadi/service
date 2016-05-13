package handler

import (
	"testing"
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"net/http"
)

func TestGetAllCategories(t *testing.T) {
	gin.SetMode(gin.TestMode);
	router := gin.Default();
	router.GET("/categories", GetAllCategories)
	req, err := http.NewRequest("GET", "/categories", nil)
	if err != nil {
		fmt.Println(err)
	}
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	fmt.Println(resp.Body)
	assert.Equal(t, resp.Code, 200)
}