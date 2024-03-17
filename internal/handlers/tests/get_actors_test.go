package handlers

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

func TestHandler_GetActors(t *testing.T) {
	tests := []struct {
		name                 string
		mockBehavior         func(s *mock_service.MockActors)
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "OK",
			mockBehavior: func(s *mock_service.MockActors) {
				s.EXPECT().GetActors().Return([]models.ActorFilms{
					{
						Actor: models.Actor{
							ID:       1,
							Name:     "John",
							Gender:   "male",
							BirthDay: "1990-01-01",
						},
						Movies: []int{
							101,
						},
					},
					{
						Actor: models.Actor{
							ID:       2,
							Name:     "Mary",
							Gender:   "female",
							BirthDay: "1999-01-01",
						},
						Movies: []int{
							103,
							105,
						},
					},
				}, nil)
			},
			expectedStatusCode: 200,
			expectedResponseBody: strings.TrimSpace(`[{"id":1,"name":"John","gender":"male","birthday":"1990-01-01","movies":[101]},` +
				`{"id":2,"name":"Mary","gender":"female","birthday":"1999-01-01","movies":[103,105]}]`),
		},
		{
			name: "Database error",
			mockBehavior: func(s *mock_service.MockActors) {
				s.EXPECT().GetActors().Return([]models.ActorFilms{}, errors.New("Server database exception"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: strings.TrimSpace("Server database exception\n[]"),
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			mockActorsService := mock_service.NewMockActors(c)
			mockMoviesService := mock_service.NewMockMovies(c)
			testCase.mockBehavior(mockActorsService)

			services := &service.Service{Actors: mockActorsService, Movies: mockMoviesService}
			handler := handlers.NewHandler(slog.Default(), services)

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/actors/", nil)

			handler.GetActors(w, r)
			actual := strings.TrimSpace(w.Body.String())
			if w.Code != testCase.expectedStatusCode {
				t.Errorf("expected status code %d, got %d", testCase.expectedStatusCode, w.Code)
			}

			if actual != testCase.expectedResponseBody {
				t.Errorf("expected response body\n'%s'\ngot\n'%s'", testCase.expectedResponseBody, actual)
			}
		})
	}
}
