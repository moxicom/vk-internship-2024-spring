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
	"github.com/moxicom/vk-internship-2024-spring/internal/service"
	mock_service "github.com/moxicom/vk-internship-2024-spring/internal/service/mocks"
)

func TestHandler_DeleteActor(t *testing.T) {
	type args struct {
		actorID int
	}

	tests := []struct {
		name                 string
		args                 args
		mockBehavior         func(s *mock_service.MockActors, actorID int)
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "OK",
			args: args{
				actorID: 1,
			},
			mockBehavior: func(s *mock_service.MockActors, actorID int) {
				s.EXPECT().DeleteActor(actorID).Return(nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "",
		},
		{
			name: "UnexpectedError",
			args: args{
				actorID: 1,
			},
			mockBehavior: func(s *mock_service.MockActors, actorID int) {
				s.EXPECT().DeleteActor(actorID).Return(fmt.Errorf("error"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: "error",
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			mockActorsService := mock_service.NewMockActors(c)
			mockMoviesService := mock_service.NewMockMovies(c)
			testCase.mockBehavior(mockActorsService, testCase.args.actorID)

			services := &service.Service{Actors: mockActorsService, Movies: mockMoviesService}
			handler := handlers.NewHandler(slog.Default(), services)

			w := httptest.NewRecorder()
			r := httptest.NewRequest(
				http.MethodDelete,
				"/actors/"+fmt.Sprintf("%d", testCase.args.actorID),
				nil,
			)

			handler.DeleteActor(w, r)
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
