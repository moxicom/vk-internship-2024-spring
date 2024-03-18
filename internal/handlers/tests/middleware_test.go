package handlers_test

import (
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/moxicom/vk-internship-2024-spring/internal/handlers"
	"github.com/moxicom/vk-internship-2024-spring/internal/service"
	mock_service "github.com/moxicom/vk-internship-2024-spring/internal/service/mocks"
	"github.com/stretchr/testify/assert"
)

func TestHandler_IsAdminAuth(t *testing.T) {
	// Mock the necessary dependencies
	c := gomock.NewController(t)
	mockUserService := mock_service.NewMockUsers(c) // Implement MockUserService
	service := &service.Service{Users: mockUserService}
	mockUserService.EXPECT().CheckUser("admin", "password", true).Return(true, nil).AnyTimes()

	handler := handlers.NewHandler(slog.Default(), service)

	// Mock a request with valid credentials
	req := httptest.NewRequest("GET", "/example", nil)
	req.SetBasicAuth("admin", "password")

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the IsAdminAuth function
	handler.IsAdminAuth(rr, req, func(w http.ResponseWriter, r *http.Request) {
		// This function should be called for successful authentication
		t.Logf("Next handler called for successful authentication")
	})

	// Check if the response status code is StatusOK (200)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestHandler_IsBasicUserAuth(t *testing.T) {
	// Mock the necessary dependencies
	c := gomock.NewController(t)
	mockUserService := mock_service.NewMockUsers(c) // Implement MockUserService
	service := &service.Service{Users: mockUserService}
	mockUserService.EXPECT().CheckUser("user", "password", false).Return(true, nil).AnyTimes()

	handler := handlers.NewHandler(slog.Default(), service)

	// Mock a request with valid credentials
	req := httptest.NewRequest("GET", "/example", nil)
	req.SetBasicAuth("user", "password")

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the IsBasicUserAuth function
	handler.IsBasicUserAuth(rr, req, func(w http.ResponseWriter, r *http.Request) {
		// This function should be called for successful authentication
		t.Logf("Next handler called for successful authentication")
	})

	// Check if the response status code is StatusOK (200)
	assert.Equal(t, http.StatusOK, rr.Code)
}
