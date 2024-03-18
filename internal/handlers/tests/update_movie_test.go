package handlers_test

import (
	"errors"
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

func TestHandler_UpdateMovie(t *testing.T) {
	type args struct {
		movie   models.Movie
		movieID int
	}

	tests := []struct {
		name                 string
		inputBody            string
		args                 args
		mockBehavior         func(s *mock_service.MockMovies, a models.Movie, movieID int)
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "OK",
			inputBody: `{"name": "The Matrix"}`,
			args: args{
				movie: models.Movie{
					Name: "The Matrix",
				},
				movieID: 1,
			},
			mockBehavior: func(s *mock_service.MockMovies, m models.Movie, movieID int) {
				s.EXPECT().
					UpdateMovie(movieID, m).
					Return(nil).
					AnyTimes()
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: "",
		},
		{
			name:      "EmptyRequestBody",
			inputBody: "",
			args: args{
				movie:   models.Movie{},
				movieID: 1,
			},
			mockBehavior: func(s *mock_service.MockMovies, m models.Movie, movieID int) {
				s.EXPECT().UpdateMovie(movieID, m).Return(errors.New("empty request body")).AnyTimes()
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: handlers.JsonParseErr,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			mockMoviesService := mock_service.NewMockMovies(c)

			testCase.mockBehavior(mockMoviesService, testCase.args.movie, testCase.args.movieID)

			services := &service.Service{Movies: mockMoviesService}
			handler := handlers.NewHandler(slog.Default(), services)

			w := httptest.NewRecorder()
			r := httptest.NewRequest(
				http.MethodPut,
				fmt.Sprintf("/movies/%d", testCase.args.movieID),
				strings.NewReader(testCase.inputBody),
			)

			testCase.mockBehavior(mockMoviesService, testCase.args.movie, testCase.args.movieID)

			handler.UpdateMovie(w, r)
			if w.Code != testCase.expectedStatusCode {
				t.Errorf("expected status code %d, got %d", testCase.expectedStatusCode, w.Code)
			}
		})
	}
}
