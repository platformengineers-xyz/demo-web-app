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

func SetUpRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	return router
}

func TestGetAlbums(t *testing.T) {
	r := SetUpRouter()
	r.GET("/albums", getAlbums)
	req, _ := http.NewRequest("GET", "/albums", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	var albums []album
	json.Unmarshal(w.Body.Bytes(), &albums)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, albums)
}

func TestGetAlbumByID(t *testing.T) {
	r := SetUpRouter()
	r.GET("/albums/:id", getAlbumByID)
	req, _ := http.NewRequest("GET", "/albums/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	var album album
	json.Unmarshal(w.Body.Bytes(), &album)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, album)
}

func TestGetAlbumByIDNotFound(t *testing.T) {
	r := SetUpRouter()
	r.GET("/albums/:id", getAlbumByID)
	req, _ := http.NewRequest("GET", "/albums/9", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestPostAlbums(t *testing.T) {
	r := SetUpRouter()
	r.POST("/albums", postAlbums)
	album := album{
		ID:     "4",
		Title:  "Babylon By Bus",
		Artist: "Bob Marley & The Wailers",
		Price:  34.95,
	}
	jsonValue, _ := json.Marshal(album)
	req, _ := http.NewRequest("POST", "/albums", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}
