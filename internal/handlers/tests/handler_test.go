package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/moxicom/vk-internship-2024-spring/internal/handlers"
	"github.com/moxicom/vk-internship-2024-spring/internal/service"
	mock_service "github.com/moxicom/vk-internship-2024-spring/internal/service/mocks"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	c := gomock.NewController(t)
	mockUserService := mock_service.NewMockUsers(c) // Implement MockUserService
	mockMoviesService := mock_service.NewMockMovies(c)
	mockActorsService := mock_service.NewMockActors(c)
	service := &service.Service{
		Users:  mockUserService,
		Movies: mockMoviesService,
		Actors: mockActorsService,
	}
	// Mock service, logger not needed for this test

	// Create a test server
	testServer := httptest.NewServer(handlers.NewHandler(nil, service).InitRoutes())
	defer testServer.Close()

	// Create an HTTP client
	client := &http.Client{
		Timeout: time.Second * 5, // Set a timeout for the client
	}

	// Make a sample request to the test server
	resp, err := client.Get(testServer.URL + "/example/") // Adjust URL as needed
	assert.NoError(t, err, "Error making request")
	defer resp.Body.Close()

	// Check the response status code
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	// Add more assertions as needed
}

func TestHandler_InitRoutes(t *testing.T) {
	// Create a Handler
	handler := handlers.NewHandler(nil, nil)

	// Call the InitRoutes function
	mux := handler.InitRoutes()

	// Create a test server
	testServer := httptest.NewServer(mux)
	defer testServer.Close()

	// Create an HTTP client
	client := &http.Client{
		Timeout: time.Second * 5, // Set a timeout for the client
	}

	// Make a sample request to the test server
	resp, err := client.Get(testServer.URL + "/example/") // Adjust URL as needed
	assert.NoError(t, err, "Error making request")
	defer resp.Body.Close()

	// Check the response status code
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	// Add more assertions as needed
}

func TestHandler_actorsMainHandler(t *testing.T) {
	c := gomock.NewController(t)
	mockUserService := mock_service.NewMockUsers(c) // Implement MockUserService
	mockMoviesService := mock_service.NewMockMovies(c)
	mockActorsService := mock_service.NewMockActors(c)
	mockService := &service.Service{
		Users:  mockUserService,
		Movies: mockMoviesService,
		Actors: mockActorsService,
	}

	handler := handlers.NewHandler(nil, mockService)

	// Test GET request
	req, err := http.NewRequest(http.MethodGet, "/actors", nil)
	assert.NoError(t, err, "Error creating GET request")

	recorder := httptest.NewRecorder()
	handler.ActorsMainHandler(recorder, req)

	assert.Equal(t, http.StatusUnauthorized, recorder.Code)
	// Add more assertions as needed

	// Test POST request
	req, err = http.NewRequest(http.MethodPost, "/actors", nil)
	assert.NoError(t, err, "Error creating POST request")

	recorder = httptest.NewRecorder()
	handler.ActorsMainHandler(recorder, req)

	assert.Equal(t, http.StatusUnauthorized, recorder.Code)
	// Add more assertions as needed

	// Test PUT request
	req, err = http.NewRequest(http.MethodPut, "/actors", nil)
	assert.NoError(t, err, "Error creating PUT request")

	recorder = httptest.NewRecorder()
	handler.ActorsMainHandler(recorder, req)

	assert.Equal(t, http.StatusUnauthorized, recorder.Code)
	// Add more assertions as needed

	// Test DELETE request
	req, err = http.NewRequest(http.MethodDelete, "/actors", nil)
	assert.NoError(t, err, "Error creating DELETE request")

	recorder = httptest.NewRecorder()
	handler.ActorsMainHandler(recorder, req)

	assert.Equal(t, http.StatusUnauthorized, recorder.Code)
	// Add more assertions as needed

	// Test invalid method
	req, err = http.NewRequest(http.MethodPatch, "/actors", nil)
	assert.NoError(t, err, "Error creating PATCH request")

	recorder = httptest.NewRecorder()
	handler.ActorsMainHandler(recorder, req)

	assert.Equal(t, http.StatusUnauthorized, recorder.Code)
	// Add more assertions as needed
}

func TestHandler_MoviesMainHandler(t *testing.T) {
	c := gomock.NewController(t)
	mockUserService := mock_service.NewMockUsers(c) // Implement MockUserService
	mockMoviesService := mock_service.NewMockMovies(c)
	mockActorsService := mock_service.NewMockActors(c)
	mockService := &service.Service{
		Users:  mockUserService,
		Movies: mockMoviesService,
		Actors: mockActorsService,
	}

	handler := handlers.NewHandler(nil, mockService)

	// Test GET request
	req, err := http.NewRequest(http.MethodGet, "/movies/", nil)
	assert.NoError(t, err, "Error creating GET request")

	recorder := httptest.NewRecorder()
	handler.MoviesMainHandler(recorder, req)

	assert.Equal(t, http.StatusUnauthorized, recorder.Code)
	// Add more assertions as needed

	// Test POST request
	req, err = http.NewRequest(http.MethodPost, "/movies/", nil)
	assert.NoError(t, err, "Error creating POST request")

	recorder = httptest.NewRecorder()
	handler.MoviesMainHandler(recorder, req)

	assert.Equal(t, http.StatusUnauthorized, recorder.Code)
	// Add more assertions as needed

	// Test PUT request
	req, err = http.NewRequest(http.MethodPut, "/movies/", nil)
	assert.NoError(t, err, "Error creating PUT request")

	recorder = httptest.NewRecorder()
	handler.MoviesMainHandler(recorder, req)

	assert.Equal(t, http.StatusUnauthorized, recorder.Code)
	// Add more assertions as needed

	// Test DELETE request
	req, err = http.NewRequest(http.MethodDelete, "/movies/", nil)
	assert.NoError(t, err, "Error creating DELETE request")

	recorder = httptest.NewRecorder()
	handler.MoviesMainHandler(recorder, req)

	assert.Equal(t, http.StatusUnauthorized, recorder.Code)
	// Add more assertions as needed

	// Test invalid method
	req, err = http.NewRequest(http.MethodPatch, "/movies/", nil)
	assert.NoError(t, err, "Error creating PATCH request")

	recorder = httptest.NewRecorder()
	handler.MoviesMainHandler(recorder, req)

	assert.Equal(t, http.StatusUnauthorized, recorder.Code)
	// Add more assertions as needed
}

func TestHandler_RelationsMainHandler(t *testing.T) {
	c := gomock.NewController(t)
	mockUserService := mock_service.NewMockUsers(c) // Implement MockUserService
	mockMoviesService := mock_service.NewMockMovies(c)
	mockActorsService := mock_service.NewMockActors(c)
	mockService := &service.Service{
		Users:  mockUserService,
		Movies: mockMoviesService,
		Actors: mockActorsService,
	}

	handler := handlers.NewHandler(nil, mockService)

	// Test DELETE request
	req, err := http.NewRequest(http.MethodDelete, "/relations/", nil)
	assert.NoError(t, err, "Error creating DELETE request")

	recorder := httptest.NewRecorder()
	handler.RelationsMainHandler(recorder, req)

	assert.Equal(t, http.StatusUnauthorized, recorder.Code)
	// Add more assertions as needed

	// Test POST request
	req, err = http.NewRequest(http.MethodPost, "/relations/", nil)
	assert.NoError(t, err, "Error creating POST request")

	recorder = httptest.NewRecorder()
	handler.RelationsMainHandler(recorder, req)

	assert.Equal(t, http.StatusUnauthorized, recorder.Code)
	// Add more assertions as needed

	// Test invalid method
	req, err = http.NewRequest(http.MethodPatch, "/relations/", nil)
	assert.NoError(t, err, "Error creating PATCH request")

	recorder = httptest.NewRecorder()
	handler.RelationsMainHandler(recorder, req)

	assert.Equal(t, http.StatusMethodNotAllowed, recorder.Code)
	// Add more assertions as needed
}
