package handlers

import (
	"errors"
	"log"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/moxicom/vk-internship-2024-spring/internal/models"
	"github.com/moxicom/vk-internship-2024-spring/internal/service"
	mock_service "github.com/moxicom/vk-internship-2024-spring/internal/service/mocks"
)

func TestHandler_getActors(t *testing.T) {
	type fields struct {
		service service.Actors
		log     log.Logger
	}

	tests := []struct {
		name                 string
		mockBehavior         func(s *mock_service.MockActors)
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "OK",
			mockBehavior: func(s *mock_service.MockActors) {
				s.EXPECT().GetActors().Return([]models.ActorFilm{
					models.ActorFilm{
						Actor: models.Actor{
							ID:       1,
							Name:     "John",
							IsMale:   true,
							BirthDay: "1990-01-01",
						},
						Movies: []int{
							101,
						},
					},
					models.ActorFilm{
						Actor: models.Actor{
							ID:       2,
							Name:     "Mary",
							IsMale:   false,
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
			expectedResponseBody: strings.TrimSpace(`[{"id":1,"name":"John","is_male":true,"birthday":"1990-01-01","movies":[101]},` +
				`{"id":2,"name":"Mary","is_male":false,"birthday":"1999-01-01","movies":[103,105]}]`),
		},
		{
			name: "Database error",
			mockBehavior: func(s *mock_service.MockActors) {
				s.EXPECT().GetActors().Return([]models.ActorFilm{}, errors.New("Server database exception"))
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
			testCase.mockBehavior(mockActorsService)

			services := &service.Service{Actors: mockActorsService}
			handler := NewHandler(slog.Default(), services)

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/actors/", nil)

			handler.getActors(w, r)
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

func TestHandler_addActor(t *testing.T) {
	type fields struct {
		service service.Actors
		log     log.Logger
	}

	tests := []struct {
		name                 string
		input                models.Actor
		mockBehavior         func(s *mock_service.MockActors, a models.Actor)
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "OK",
			mockBehavior: func(s *mock_service.MockActors, a models.Actor) {
				s.EXPECT().AddActor(a).Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: strings.TrimSpace(``),
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			mockActorsService := mock_service.NewMockActors(c)
			testCase.mockBehavior(mockActorsService, testCase.input)

			services := &service.Service{Actors: mockActorsService}
			handler := NewHandler(slog.Default(), services)

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/actors/", nil)

			handler.addActor(w, r)
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
