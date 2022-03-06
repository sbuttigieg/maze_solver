package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/sbuttigieg/maze_solver/app_errors"
	"github.com/sbuttigieg/maze_solver/database"
	"github.com/stretchr/testify/assert"
)

// Test that /levels GET function getLevels is returning the correct data
func TestGetLevels(t *testing.T) {
	database.ConnectDB()
	router := InitialiseApi()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/levels", nil)
	router.ServeHTTP(w, req)
	var typeToCompareTo, result map[int]database.LevelTableFields
	if err := json.Unmarshal(w.Body.Bytes(), &result); err != nil {
		panic(fmt.Sprintf("Failed to unmarshal levels. =>\n%v", err))
	}
	assert.Equal(t, 200, w.Code)
	assert.IsTypef(t, typeToCompareTo, result, "Types do not match")
}

// Test that /levels/[id] GET function getLevelById is returning the correct data
func TestGetLevelById(t *testing.T) {
	database.ConnectDB()
	router := InitialiseApi()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/levels/1", nil)
	router.ServeHTTP(w, req)
	var typeToCompareTo, result map[int]database.LevelTableFields
	if err := json.Unmarshal(w.Body.Bytes(), &result); err != nil {
		panic(fmt.Sprintf("Failed to unmarshal levels. =>\n%v", err))
	}
	assert.Equal(t, 200, w.Code)
	assert.IsTypef(t, typeToCompareTo, result, "Types do not match")
}

// Test that /levels/[id] GET function getLevelById is returning the correct error code
func TestGetLevelByIdFail(t *testing.T) {
	var errMsg app_errors.ErrorStruct
	database.ConnectDB()
	router := InitialiseApi()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/levels/a", nil)
	router.ServeHTTP(w, req)
	jsonErr := json.Unmarshal(w.Body.Bytes(), &errMsg)
	if jsonErr != nil {
		t.Fatalf("Incorrect JSON format")
	}
	assert.Equal(t, 400, w.Code)
	assert.Equal(t, 1005, errMsg.ErrorCode)
	assert.Equal(t, "Incorrect level ID type", errMsg.Description)
}

// Test that /[id] DELETE function deleteLevelById is returning the correct data
func TestDeleteLevelById(t *testing.T) {
	database.ConnectDB()
	router := InitialiseApi()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/1", nil)
	router.ServeHTTP(w, req)
	var typeToCompareTo, result string
	if err := json.Unmarshal(w.Body.Bytes(), &result); err != nil {
		panic(fmt.Sprintf("Failed to unmarshal levels. =>\n%v", err))
	}
	assert.Equal(t, 200, w.Code)
	assert.IsTypef(t, typeToCompareTo, result, "Types do not match")
}

// Test that /[id] DELETE function deleteLevelById is returning the correct error
func TestDeleteLevelByIdFail(t *testing.T) {
	var errMsg app_errors.ErrorStruct
	database.ConnectDB()
	router := InitialiseApi()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/a", nil)
	router.ServeHTTP(w, req)
	jsonErr := json.Unmarshal(w.Body.Bytes(), &errMsg)
	if jsonErr != nil {
		t.Fatalf("Incorrect JSON format")
	}
	assert.Equal(t, 400, w.Code)
	assert.Equal(t, 1005, errMsg.ErrorCode)
	assert.Equal(t, "Incorrect level ID type", errMsg.Description)
}

// Test that /submit POST function postLevel is returning the correct levelId
func TestPostSubmit(t *testing.T) {
	database.ConnectDB()
	router := InitialiseApi()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/level", bytes.NewBuffer([]byte("[[1, 1, 0, 1, 1], [1, 0, 2, 0, 1], [1, 1, 1, 1, 1]]")))
	router.ServeHTTP(w, req)
	var typeToCompareTo int
	assert.Equal(t, 201, w.Code)
	levelId, _ := strconv.Atoi(w.Body.String())
	assert.IsTypef(t, typeToCompareTo, levelId, "Types do not match")
}

// Test that /submit POST function post
func TestPostSubmitFail(t *testing.T) {
	database.ConnectDB()
	var errMsg app_errors.ErrorStruct
	router := InitialiseApi()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/level", bytes.NewBuffer([]byte("[[1, 1, 0, 1], [1, 0, 2, 0, 1], [1, 1, 1, 1, 1]]")))
	router.ServeHTTP(w, req)
	jsonErr := json.Unmarshal(w.Body.Bytes(), &errMsg)
	if jsonErr != nil {
		t.Fatalf("Incorrect JSON format")
	}
	assert.Equal(t, 400, w.Code)
	assert.Equal(t, 1000, errMsg.ErrorCode)
	assert.Equal(t, "Level Shape not rectangular", errMsg.Description)
}

// Test that /[id] PATCH function updateLevelById is returning the correct data
func TestPatchLevelById(t *testing.T) {
	database.ConnectDB()
	router := InitialiseApi()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PATCH", "/1", bytes.NewBuffer([]byte("[[1, 1, 0, 1, 1], [1, 0, 2, 0, 1], [1, 1, 1, 1, 1]]")))
	router.ServeHTTP(w, req)
	// var typeToCompareTo string
	var typeToCompareTo, result string
	if err := json.Unmarshal(w.Body.Bytes(), &result); err != nil {
		panic(fmt.Sprintf("Failed to unmarshal levels. =>\n%v", err))
	}
	assert.Equal(t, 200, w.Code)
	assert.IsTypef(t, typeToCompareTo, result, "Types do not match")
}

// Test that /[id] PATCH function updateLevelById is returning the correct error
func TestPatchLevelByIdFail(t *testing.T) {
	database.ConnectDB()
	var errMsg app_errors.ErrorStruct
	router := InitialiseApi()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PATCH", "/a", bytes.NewBuffer([]byte("[[1, 1, 0, 1, 1], [1, 0, 2, 0, 1], [1, 1, 1, 1, 1]]")))
	router.ServeHTTP(w, req)
	jsonErr := json.Unmarshal(w.Body.Bytes(), &errMsg)
	if jsonErr != nil {
		t.Fatalf("Incorrect JSON format")
	}
	assert.Equal(t, 400, w.Code)
	assert.Equal(t, 1005, errMsg.ErrorCode)
	assert.Equal(t, "Incorrect level ID type", errMsg.Description)
}
