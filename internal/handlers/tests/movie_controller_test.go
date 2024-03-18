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
)

func TestHandler_GetMoviesController(t *testing.T) {
	tests := []struct {
		name               string
		path               string
		expectedStatusCode int
	}{
		{
			name:               "Bad endpoint",
			path:               "/movies/12gg",
			expectedStatusCode: 400,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			mockMoviesService := mock_service.NewMockMovies(c)
			mockActorsService := mock_service.NewMockActors(c)
			services := &service.Service{Actors: mockActorsService, Movies: mockMoviesService}
			handler := handlers.NewHandler(slog.Default(), services)

			w := httptest.NewRecorder()
			r := httptest.NewRequest(
				http.MethodGet,
				testCase.path,
				nil,
			)

			handler.GetMoviesController(w, r)
			if w.Code != testCase.expectedStatusCode {
				t.Errorf("expected status code %d, got %d", testCase.expectedStatusCode, w.Code)
			}
		})
	}
}
