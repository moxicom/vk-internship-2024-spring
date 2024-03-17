package handlers_test

import (
	"bytes"
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

func TestHandler_AddActor(t *testing.T) {
	tests := []struct {
		name                 string
		inputBody            string
		input                models.Actor
		mockBehavior         func(s *mock_service.MockActors, a models.Actor)
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "OK",
			inputBody: `{"name": "Alex", "gender": "male", "birthday": "2004-12-14"}`,
			input: models.Actor{
				ID:       0,
				Name:     "Alex",
				Gender:   "male",
				BirthDay: "2004-12-14",
			},
			mockBehavior: func(s *mock_service.MockActors, a models.Actor) {
				s.EXPECT().AddActor(a).Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: strings.TrimSpace(`{"id":1}`),
		},
		{
			name:      "EmptyRequestBody",
			inputBody: "",             // Пустое тело запроса
			input:     models.Actor{}, // Пустая структура актера
			mockBehavior: func(s *mock_service.MockActors, a models.Actor) {
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
			testCase.mockBehavior(mockActorsService, testCase.input)

			services := &service.Service{Actors: mockActorsService, Movies: mockMoviesService}
			handler := handlers.NewHandler(slog.Default(), services)

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/actors/", bytes.NewBufferString(testCase.inputBody))

			handler.AddActor(w, r)
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
