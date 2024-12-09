package http

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/QBC8-Go-Group2/questionnaire/internal/questionnaire/domain"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockQuestionnaireService struct {
	mock.Mock
}

func (m *MockQuestionnaireService) CreateQuestionnaire(ctx context.Context, questionnaire domain.Questionnaire) (domain.QuestionnaireDbID, error) {
	args := m.Called(ctx, questionnaire)
	return args.Get(0).(domain.QuestionnaireDbID), args.Error(1)
}

func (m *MockQuestionnaireService) UpdateQuestionnaire(ctx context.Context, questionnaire domain.Questionnaire) (domain.QuestionnaireID, error) {
	args := m.Called(ctx, questionnaire)
	return args.Get(0).(domain.QuestionnaireID), args.Error(1)
}

func (m *MockQuestionnaireService) FindQuestionnaireWithQuestionnaireID(ctx context.Context, questionnaireId domain.QuestionnaireID) (domain.Questionnaire, error) {
	args := m.Called(ctx, questionnaireId)
	return args.Get(0).(domain.Questionnaire), args.Error(1)
}

func (m *MockQuestionnaireService) FindQuestionnaireWithQuestionnaireDbID(ctx context.Context, questionnaireId domain.QuestionnaireDbID) (domain.Questionnaire, error) {
	args := m.Called(ctx, questionnaireId)
	return args.Get(0).(domain.Questionnaire), args.Error(1)
}

func (m *MockQuestionnaireService) DeleteQuestionnaireWithQuestionnaireID(ctx context.Context, questionnaireId domain.QuestionnaireID) error {
	args := m.Called(ctx, questionnaireId)
	return args.Error(0)
}

func (m *MockQuestionnaireService) DeleteQuestionnaireWithUserDbId(ctx context.Context, questionnaireId domain.QuestionnaireDbID) error {
	args := m.Called(ctx, questionnaireId)
	return args.Error(0)
}

func setupTestApp() (*fiber.App, *MockQuestionnaireService) {
	mockService := new(MockQuestionnaireService)
	handler := NewQuestionnaireHandler(mockService)

	app := fiber.New()
	RegisterQuestionnaireRoutes(app, func(c *fiber.Ctx) error { return c.Next() }, handler)

	return app, mockService
}

func TestCreateQuestionnaire(t *testing.T) {
	app, mockService := setupTestApp()

	t.Run("Create Questionnaire - Success", func(t *testing.T) {
		mockService.On("CreateQuestionnaire", mock.Anything, mock.Anything).Return("12345", nil)

		request := map[string]interface{}{
			"title":       "Test Questionnaire",
			"description": "This is a test questionnaire",
			"duration":    60,
			"editable":    true,
			"randomable":  false,
			"validTo":     time.Now().Add(time.Hour).Unix(),
		}
		body, _ := json.Marshal(request)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/questionnaire", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusCreated, resp.StatusCode)
		mockService.AssertCalled(t, "CreateQuestionnaire", mock.Anything, mock.Anything)
	})

	t.Run("Create Questionnaire - Invalid User ID", func(t *testing.T) {
		request := map[string]interface{}{
			"title":       "Test Questionnaire",
			"description": "This is a test questionnaire",
			"duration":    60,
			"editable":    true,
			"randomable":  false,
		}
		body, _ := json.Marshal(request)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/questionnaire", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("id", "invalid")
		req.Header.Set("email", "test@example.com")

		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})
}

func TestCreateQuestions(t *testing.T) {
	app, mockService := setupTestApp()

	t.Run("Create Questions - Success", func(t *testing.T) {
		mockService.On("FindQuestionnaireWithQuestionnaireID", mock.Anything, mock.Anything).Return(&domain.Questionnaire{}, nil)

		request := map[string]interface{}{
			"questionnaireId": "12345",
			"questions": []map[string]interface{}{
				{
					"type":    1,
					"title":   "Sample question",
					"options": []map[string]interface{}{{"text": "Option 1", "isAnswer": true}},
				},
			},
		}
		body, _ := json.Marshal(request)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/questionnaire/questions", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusCreated, resp.StatusCode)
		mockService.AssertCalled(t, "FindQuestionnaireWithQuestionnaireID", mock.Anything, mock.Anything)
	})

	t.Run("Create Questions - Invalid Questionnaire", func(t *testing.T) {
		mockService.On("FindQuestionnaireWithQuestionnaireID", mock.Anything, mock.Anything).Return(nil, errors.New("not found"))

		// Request body
		request := map[string]interface{}{
			"questionnaireId": "invalid-id",
			"questions": []map[string]interface{}{
				{
					"type":    1,
					"title":   "Sample question",
					"options": []map[string]interface{}{{"text": "Option 1", "isAnswer": true}},
				},
			},
		}
		body, _ := json.Marshal(request)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/questionnaire/questions", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("id", "1")
		req.Header.Set("email", "test@example.com")

		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})
}
