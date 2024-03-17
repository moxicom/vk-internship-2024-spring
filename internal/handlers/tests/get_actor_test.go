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

func TestHandler_GetActor(t *testing.T) {
	tests := []struct {
		name                 string
		input                int
		mockBehavior         func(s *mock_service.MockActors, actorId int)
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:  "OK",
			input: 1,
			mockBehavior: func(s *mock_service.MockActors, actorId int) {
				s.EXPECT().GetActor(actorId).Return(models.ActorFilms{
					Actor: models.Actor{
						ID:       actorId,
						Name:     "John",
						Gender:   "male",
						BirthDay: "1990-01-01",
					},
					Movies: []int{
						101,
					},
				}, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: strings.TrimSpace(`{"id":1,"name":"John",` + `"gender":"male","birthday":"1990-01-01","movies":[101]}`),
		},
		{
			name:  "Actor Not Found",
			input: 2,
			mockBehavior: func(s *mock_service.MockActors, actorId int) {
				s.EXPECT().GetActor(actorId).Return(models.ActorFilms{}, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "{}",
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			mockActorsService := mock_service.NewMockActors(c)
			mockMoviesService := mock_service.NewMockMovies(c)
			testCase.mockBehavior(mockActorsService, testCase.input)

			services := &service.Service{Actors: mockActorsService, Movies: mockMoviesService}
			handler := handlers.NewHandler(slog.Default(), services)

			w := httptest.NewRecorder()
			r := httptest.NewRequest(
				http.MethodGet,
				"/actors/"+fmt.Sprintf("%v", testCase.input),
				nil,
			)

			handler.GetActor(w, r, testCase.input)
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
