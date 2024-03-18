package handlers_test

import (
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/moxicom/vk-internship-2024-spring/internal/handlers"
	"github.com/moxicom/vk-internship-2024-spring/internal/models"
	"github.com/moxicom/vk-internship-2024-spring/internal/service"
	mock_service "github.com/moxicom/vk-internship-2024-spring/internal/service/mocks"
)

func TestHandler_AddMovie(t *testing.T) {
	tests := []struct {
		name                 string
		inputBody            string
		input                models.Movie
		mockBehavior         func(s *mock_service.MockMovies, m models.Movie)
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "OK",
			inputBody: `{"name": "The Matrix", "date": "1999-01-01", "description": "VHS film", "rating": 10}`,
			input: models.Movie{
				ID:          0,
				Name:        "The Matrix",
				Date:        "1999-01-01",
				Description: "VHS film",
				Rating:      &[]float32{10}[0],
			},
			mockBehavior: func(s *mock_service.MockMovies, m models.Movie) {
				s.EXPECT().AddMovie(m).Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: strings.TrimSpace(`{"id":1}`),
		},
		{
			name:      "EmptyRequestBody",
			inputBody: "",             // Пустое тело запроса
			input:     models.Movie{}, // Пустая структура актера
			mockBehavior: func(s *mock_service.MockMovies, m models.Movie) {
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: handlers.JsonParseErr,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			mockActorsService := mock_service.NewMockActors(c)
			mockMoviesService := mock_service.NewMockMovies(c)
			testCase.mockBehavior(mockMoviesService, testCase.input)

			services := &service.Service{Actors: mockActorsService, Movies: mockMoviesService}
			handler := handlers.NewHandler(slog.Default(), services)

			w := httptest.NewRecorder()
			r := httptest.NewRequest(
				http.MethodPost,
				"/movies/",
				strings.NewReader(testCase.inputBody),
			)

			handler.AddMovie(w, r)

			if w.Code != testCase.expectedStatusCode {
				t.Errorf("expected status code %d, got %d", testCase.expectedStatusCode, w.Code)
			}
		})
	}
}
