package handlers_test

import (
	"fmt"
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

func TestHandler_GetMovie(t *testing.T) {
	tests := []struct {
		name                 string
		input                int
		mockBehavior         func(s *mock_service.MockMovies, movieId int)
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:  "OK",
			input: 1,
			mockBehavior: func(s *mock_service.MockMovies, movieId int) {
				s.EXPECT().GetMovie(movieId).Return(
					models.MovieActors{
						Movie: models.Movie{
							ID:          movieId,
							Name:        "The Matrix",
							Date:        "1999-01-01",
							Description: "VHS film",
							Rating:      &[]float32{10}[0],
						},
						Actors: []int{1, 2},
					}, nil)
			},
			expectedStatusCode: 200,
			expectedResponseBody: strings.TrimSpace(`{"id":1,"name":"The Matrix",` +
				`"description":"VHS film","date":"1999-01-01","rating":10,"actors":[1,2]}`),
		},
		{
			name:  "Movie Not Found",
			input: 2,
			mockBehavior: func(s *mock_service.MockMovies, movieId int) {

				s.EXPECT().GetMovie(movieId).Return(models.MovieActors{}, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: strings.TrimSpace(`{"id":0,"name":"","description":"","date":"","rating":null,"actors":null}`),
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			mockMoviesService := mock_service.NewMockMovies(c)
			mockActorsService := mock_service.NewMockActors(c)
			testCase.mockBehavior(mockMoviesService, testCase.input)

			services := &service.Service{
				Actors: mockActorsService,
				Movies: mockMoviesService,
			}
			handler := handlers.NewHandler(slog.Default(), services)

			w := httptest.NewRecorder()
			r := httptest.NewRequest(
				http.MethodGet,
				"/movies/"+fmt.Sprintf("%d", testCase.input),
				nil,
			)

			handler.GetMovie(w, r, testCase.input)
			if w.Code != testCase.expectedStatusCode {
				t.Errorf("expected status code %d, got %d", testCase.expectedStatusCode, w.Code)
			}
			if strings.TrimSpace(w.Body.String()) != testCase.expectedResponseBody {
				t.Errorf("expected response body \n%s, got \n%s", testCase.expectedResponseBody, w.Body)
			}
		})
	}
}
