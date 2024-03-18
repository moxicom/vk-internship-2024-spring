package handlers_test

import (
	"errors"
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

func TestHandler_GetMovies(t *testing.T) {
	type input struct {
		params string
		sort   models.SortParams
		search models.SearchParams
	}
	tests := []struct {
		name                 string
		input                input
		mockBehavior         func(s *mock_service.MockMovies, args input)
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "OK",
			input: input{
				params: "?sort=rating&search=matr",
				sort:   models.SortParams{Sort: "rating", Order: "desc"},
				search: models.SearchParams{MovieName: "matr"},
			},
			mockBehavior: func(s *mock_service.MockMovies, args input) {
				s.EXPECT().GetMovies(
					models.SortParams{},
					models.SearchParams{},
				).Return(
					[]models.MovieActors{
						{
							Movie: models.Movie{
								ID:          1,
								Name:        "The Matrix",
								Date:        "1999-01-01",
								Description: "VHS film",
								Rating:      &[]float32{10}[0],
							},
							Actors: []int{1, 2},
						},
					}, nil)
			},
			expectedStatusCode: 200,
			expectedResponseBody: strings.TrimSpace(`[{"id":1,"name":"The Matrix",` +
				`"description":"VHS film","date":"1999-01-01","rating":10,"actors":[1,2]}]`),
		},
		{
			name: "InvalidSortParam",
			input: input{
				params: "?sort=invalid",
				sort:   models.SortParams{Sort: "invalid", Order: "asc"},
				search: models.SearchParams{},
			},
			mockBehavior: func(s *mock_service.MockMovies, args input) {
				s.EXPECT().GetMovies(
					models.SortParams{},
					models.SearchParams{},
				).Return(nil, errors.New("Invalid sort parameter"))
			},
			expectedStatusCode:   http.StatusInternalServerError,
			expectedResponseBody: "Failed to get movies",
		},
		{
			name: "NoMoviesFound",
			input: input{
				params: "?sort=rating&search=nonexistent",
				sort:   models.SortParams{Sort: "rating", Order: "desc"},
				search: models.SearchParams{MovieName: "nonexistent"},
			},
			mockBehavior: func(s *mock_service.MockMovies, args input) {
				s.EXPECT().GetMovies(
					models.SortParams{},
					models.SearchParams{},
				).Return(nil, nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: "[]",
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
				"/movies/",
				nil,
			)

			handler.GetMovies(w, r)
			actual := strings.TrimSpace(w.Body.String())
			if w.Code != testCase.expectedStatusCode {
				t.Errorf("test:"+testCase.name+"expected status code %d, got %d", testCase.expectedStatusCode, w.Code)
			}

			if actual != testCase.expectedResponseBody {
				t.Errorf("expected response body\n'%s'\ngot\n'%s'", testCase.expectedResponseBody, actual)
			}
		})
	}
}
