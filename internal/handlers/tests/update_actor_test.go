package handlers_test

import (
	"bytes"
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

func TestHandler_UpdateActor(t *testing.T) {
	type args struct {
		actor   models.Actor
		actorID int
	}

	tests := []struct {
		name                 string
		inputBody            string
		args                 args
		mockBehavior         func(s *mock_service.MockActors, a models.Actor, actorID int)
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "OK",
			inputBody: `{"name": "Alex"}`,
			args: args{
				actor: models.Actor{
					Name: "Alex",
				},
				actorID: 1,
			},
			mockBehavior: func(s *mock_service.MockActors, a models.Actor, actorID int) {
				s.EXPECT().UpdateActor(actorID, a).Return(nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "",
		},
		{
			name:      "EmptyRequestBody",
			inputBody: "",
			args: args{
				actor:   models.Actor{},
				actorID: 1,
			},
			mockBehavior: func(s *mock_service.MockActors, a models.Actor, actorID int) {
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
			testCase.mockBehavior(mockActorsService, testCase.args.actor, testCase.args.actorID)

			services := &service.Service{Actors: mockActorsService, Movies: mockMoviesService}
			handler := handlers.NewHandler(slog.Default(), services)

			w := httptest.NewRecorder()
			r := httptest.NewRequest(
				http.MethodPut,
				"/actors/"+fmt.Sprintf("%d", testCase.args.actorID),
				bytes.NewBufferString(testCase.inputBody),
			)

			handler.UpdateActor(w, r)
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
