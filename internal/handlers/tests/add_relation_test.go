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

func TestHandler_AddRelation(t *testing.T) {
	tests := []struct {
		name         string
		inputBody    string
		input        models.RelationMoviesActors
		mockBehavior func(s *mock_service.MockRelations, input models.RelationMoviesActors)
		expectedCode int
		expectedBody string
	}{
		{
			name:      "OK",
			inputBody: `{"movie_id": "1", "actor_id": "1"}`,
			input: models.RelationMoviesActors{
				MovieID: "1",
				ActorID: "1",
			},
			mockBehavior: func(s *mock_service.MockRelations, input models.RelationMoviesActors) {
				s.EXPECT().AddRelation(input).Return(nil).AnyTimes()
			},
			expectedCode: http.StatusOK,
			expectedBody: "",
		},
		{
			name:      "EmptyRequestBody",
			inputBody: "",
			input:     models.RelationMoviesActors{},
			mockBehavior: func(s *mock_service.MockRelations, input models.RelationMoviesActors) {
				s.EXPECT().AddRelation(input).Return(errors.New("empty request body")).AnyTimes()
			},
			expectedCode: http.StatusBadRequest,
			expectedBody: handlers.JsonParseErr,
		},
		{
			name:      "InvalidRequestBody",
			inputBody: `{"movie_id": "1", "actor_id": "2as"}`,
			input: models.RelationMoviesActors{
				MovieID: "1",
				ActorID: "2asd",
			},
			mockBehavior: func(s *mock_service.MockRelations, input models.RelationMoviesActors) {
				s.EXPECT().AddRelation(input).Return(errors.New("invalid request body")).AnyTimes()
			},
			expectedCode: http.StatusBadRequest,
			expectedBody: handlers.JsonParseErr,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			mockRelationSerivce := mock_service.NewMockRelations(c)
			testCase.mockBehavior(mockRelationSerivce, testCase.input)

			services := &service.Service{
				Relations: mockRelationSerivce,
			}
			handler := handlers.NewHandler(slog.Default(), services)

			w := httptest.NewRecorder()
			r := httptest.NewRequest(
				http.MethodPost,
				"/relations/",
				strings.NewReader(testCase.inputBody),
			)

			handler.AddRelation(w, r)

			if w.Code != testCase.expectedCode {
				t.Errorf("expected status code %d, got %d", testCase.expectedCode, w.Code)
			}
		})
	}
}
