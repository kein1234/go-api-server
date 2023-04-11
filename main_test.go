package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Test helper functions
func performRequest(r http.Handler, method, path string, body []byte) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, bytes.NewReader(body))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestGetUsers(t *testing.T) {
	router := gin.Default()
	router.GET("/users", func(c *gin.Context) {
		users := []User{
			{ID: 1, Name: "Alice", Email: "alice@example.com"},
			{ID: 2, Name: "Bob", Email: "bob@example.com"},
		}
		c.JSON(http.StatusOK, gin.H{"users": users})
	})

	w := performRequest(router, "GET", "/users", nil)

	var response map[string][]User
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Nil(t, err)

	users := response["users"]
	assert.Equal(t, 2, len(users))
	assert.Equal(t, "Alice", users[0].Name)
	assert.Equal(t, "Bob", users[1].Name)
}
