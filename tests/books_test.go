package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"neohub.asia/mod/databases/models"
	"neohub.asia/mod/di"
	"neohub.asia/mod/routes"
)

type ApiResponse[T any] struct {
	Success bool `json:"success"`
	Data    T    `json:"data"`
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	container := di.NewContainer()
	return routes.SetupRoutes(container)
}

func TestBookRoutes(t *testing.T) {
	router := setupRouter()

	// --- CREATE ---
	bookPayload := map[string]interface{}{
		"title":  "Integration Test Book",
		"author": "Test Author",
	}
	payloadBytes, _ := json.Marshal(bookPayload)

	reqCreate, _ := http.NewRequest(http.MethodPost, "/api/book/", bytes.NewBuffer(payloadBytes))
	reqCreate.Header.Set("Content-Type", "application/json")

	wCreate := httptest.NewRecorder()
	router.ServeHTTP(wCreate, reqCreate)

	assert.Equal(t, http.StatusCreated, wCreate.Code)

	var createResp ApiResponse[models.Book]
	err := json.Unmarshal(wCreate.Body.Bytes(), &createResp)
	assert.NoError(t, err)

	createdBook := createResp.Data
	assert.Equal(t, bookPayload["title"], createdBook.Title)
	assert.Equal(t, bookPayload["author"], createdBook.Author)
	assert.NotZero(t, createdBook.ID)

	bookID := fmt.Sprint(createdBook.ID)

	// --- LIST ---
	reqList, _ := http.NewRequest(http.MethodGet, "/api/book/", nil)
	wList := httptest.NewRecorder()
	router.ServeHTTP(wList, reqList)

	assert.Equal(t, http.StatusOK, wList.Code)

	var listResp ApiResponse[[]models.Book]
	err = json.Unmarshal(wList.Body.Bytes(), &listResp)
	assert.NoError(t, err)
	assert.True(t, len(listResp.Data) > 0)

	// --- GET ---
	reqGet, _ := http.NewRequest(http.MethodGet, "/api/book/"+bookID, nil)
	wGet := httptest.NewRecorder()
	router.ServeHTTP(wGet, reqGet)

	assert.Equal(t, http.StatusOK, wGet.Code)

	var getResp ApiResponse[models.Book]
	err = json.Unmarshal(wGet.Body.Bytes(), &getResp)
	assert.NoError(t, err)
	assert.Equal(t, createdBook.ID, getResp.Data.ID)

	// --- UPDATE ---
	updatePayload := map[string]interface{}{
		"title": "Updated Title",
	}
	updateBytes, _ := json.Marshal(updatePayload)

	reqUpdate, _ := http.NewRequest(http.MethodPatch, "/api/book/"+bookID, bytes.NewBuffer(updateBytes))
	reqUpdate.Header.Set("Content-Type", "application/json")

	wUpdate := httptest.NewRecorder()
	router.ServeHTTP(wUpdate, reqUpdate)

	assert.Equal(t, http.StatusOK, wUpdate.Code)

	var updateResp ApiResponse[models.Book]
	err = json.Unmarshal(wUpdate.Body.Bytes(), &updateResp)
	assert.NoError(t, err)
	assert.Equal(t, "Updated Title", updateResp.Data.Title)

	// --- DELETE ---
	reqDelete, _ := http.NewRequest(http.MethodDelete, "/api/book/"+bookID, nil)
	wDelete := httptest.NewRecorder()
	router.ServeHTTP(wDelete, reqDelete)

	assert.Equal(t, http.StatusOK, wDelete.Code)

	// --- VERIFY DELETION ---
	reqGetDeleted, _ := http.NewRequest(http.MethodGet, "/api/book/"+bookID, nil)
	wGetDeleted := httptest.NewRecorder()
	router.ServeHTTP(wGetDeleted, reqGetDeleted)

	assert.Equal(t, http.StatusNotFound, wGetDeleted.Code)
}
