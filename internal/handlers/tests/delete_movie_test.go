package handlers_test

import (
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/moxicom/vk-internship-2024-spring/internal/handlers"
	"github.com/moxicom/vk-internship-2024-spring/internal/service"
	mock_service "github.com/moxicom/vk-internship-2024-spring/internal/service/mocks"
)

func TestHandler_DeleteMovie(t *testing.T) {
	type args struct {
		movieID int
	}

	tests := []struct {
		name                 string
		args                 args
		idString             string
		mockBehavior         func(s *mock_service.MockMovies, movieID int)
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "OK",
			args: args{
				movieID: 1,
			},
			mockBehavior: func(s *mock_service.MockMovies, movieID int) {
				s.EXPECT().DeleteMovie(movieID).Return(nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "",
		},
		{
			idString: "notInt",
			name:     "BadRequest",
			args: args{
				movieID: 0,
			},
			mockBehavior:         func(s *mock_service.MockMovies, movieID int) {},
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
			testCase.mockBehavior(mockMoviesService, testCase.args.movieID)

			services := &service.Service{
				Actors: mockActorsService,
				Movies: mockMoviesService,
			}
			handler := handlers.NewHandler(slog.Default(), services)

			w := httptest.NewRecorder()
			r := httptest.NewRequest(
				http.MethodDelete,
				"/movies/"+fmt.Sprintf("%d", testCase.args.movieID)+testCase.idString,
				nil,
			)

			handler.DeleteMovie(w, r)
		})
	}
}
