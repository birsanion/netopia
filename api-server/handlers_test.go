package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/birsanion/netopia/api-server/models/responses"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupTestRouter() *gin.Engine {
	router := gin.Default()
	return router
}

type TestClient struct {
	Router *gin.Engine
}

func (client *TestClient) PerformRequest(method, path string, body interface{}, headers map[string]string) *httptest.ResponseRecorder {
	var reqBody *bytes.Reader
	if body != nil {
		jsonBody, _ := json.Marshal(body)
		reqBody = bytes.NewReader(jsonBody)
	} else {
		reqBody = bytes.NewReader([]byte{})
	}

	req, _ := http.NewRequest(method, path, reqBody)

	req.Header.Set("Content-Type", "application/json")

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	w := httptest.NewRecorder()
	client.Router.ServeHTTP(w, req)
	return w
}

func NewTestClient() *TestClient {
	router := SetupTestRouter()
	client := &TestClient{Router: router}
	return client
}

func TestHealthCheckHandler(t *testing.T) {
	client := NewTestClient()

	if err := LoadConfig(); err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	db, err := setupDB(AppConfig)
	if err != nil {
		t.Fatalf("Failed to setup DB: %v", err)
	}

	q, err := setupMessageQueue(AppConfig)
	if err != nil {
		t.Fatalf("Failed to setup message queue: %v", err)
	}

	RegisterHealthRoute(client.Router, db, q)

	response := client.PerformRequest("GET", "/health", nil, nil)

	assert.Equal(t, http.StatusOK, response.Code)

	var actualBody responses.HealthCheckRespose
	err = json.Unmarshal(response.Body.Bytes(), &actualBody)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	assert.Equal(t, actualBody.Status.IsOk(), true)
}
