package http

import (
	context2 "context"
	"github.com/QBC8-Go-Group2/questionnaire/internal/option"
	optionDomain "github.com/QBC8-Go-Group2/questionnaire/internal/option/domain"
	optionPort "github.com/QBC8-Go-Group2/questionnaire/internal/option/port"
	questionnaire "github.com/QBC8-Go-Group2/questionnaire/internal/question"
	questionDomain "github.com/QBC8-Go-Group2/questionnaire/internal/question/domain"
	questionPort "github.com/QBC8-Go-Group2/questionnaire/internal/question/port"
	"github.com/QBC8-Go-Group2/questionnaire/internal/questionnaire/domain"
	questionnairePort "github.com/QBC8-Go-Group2/questionnaire/internal/questionnaire/port"
	userDomain "github.com/QBC8-Go-Group2/questionnaire/internal/user/domain"
	"github.com/QBC8-Go-Group2/questionnaire/pkg/adapter/storage"
	"github.com/QBC8-Go-Group2/questionnaire/pkg/context"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"time"
)

type QuestionnaireHandler struct {
	questionnaireService questionnairePort.Service
	questionService      questionPort.Service
}

func NewQuestionnaireHandler(questionnaireService questionnairePort.Service) *QuestionnaireHandler {
	return &QuestionnaireHandler{
		questionnaireService: questionnaireService,
	}
}

func RegisterQuestionnaireRoutes(app fiber.Router, transaction fiber.Handler, handler *QuestionnaireHandler) {
	group := app.Group("/questionnaire", SetUserContext)
	group.Post("/", handler.createQuestionnaire())
	group.Post("/questions", transaction, handler.createQuestions())
}

func (h *QuestionnaireHandler) createQuestionnaire() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var questionnaire CreateQuestionnaireRequest
		if err := c.BodyParser(&questionnaire); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		if questionnaire.ValidTo != 0 {
			t := time.Unix(questionnaire.ValidTo, 0)
			if t.Sub(time.Now().Local()) > 0 {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"message": "questionnaire too old",
				})
			}
		}

		userId := c.Get("id")
		ownerId, err := strconv.Atoi(userId)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "invalid userId",
			})
		}
		email := c.Get("email")

		qID, err := domain.CreateQuestionnaireID(email + time.Now().String())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "error in creating questionnaire",
			})
		}
		questionnaireId, err := h.questionnaireService.CreateQuestionnaire(c.UserContext(), domain.Questionnaire{
			QuestionnaireID: qID,
			OwnerID:         userDomain.UserDbID(ownerId),
			Title:           questionnaire.Title,
			Description:     questionnaire.Description,
			Duration:        questionnaire.Duration,
			Editable:        questionnaire.Editable,
			Randomable:      questionnaire.Randomable,
			ValidTo:         time.Unix(questionnaire.ValidTo, 0),
		})
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"questionnaireID": questionnaireId,
		})
	}
}

func (h *QuestionnaireHandler) createQuestions() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var questionsReq CreateQuestionsRequest
		if err := c.BodyParser(&questionsReq); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		_, err := h.questionnaireService.FindQuestionnaireWithQuestionnaireID(c.UserContext(), domain.QuestionnaireID(questionsReq.QuestionnaireID))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "questionnaire doesn't exists",
			})
		}

		db := context.GetDB(c.UserContext())
		if db == nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		questionRepo := storage.NewQuestionRepo(db)
		questionService := questionnaire.NewService(questionRepo)

		optionRepo := storage.NewOptionRepo(db)
		optionService := option.NewService(optionRepo)

		for _, v := range questionsReq.Questions {
			switch v.Type {
			case uint8(questionDomain.QuestionTypeOptional):
				if len(v.Options) == 0 {
					return c.SendStatus(fiber.StatusBadRequest)
				}
				question, err := createQuestion(questionService, c.UserContext(), v.toDomain(questionsReq.QuestionnaireID))
				if err != nil {
					return c.SendStatus(fiber.StatusInternalServerError)
				}
				for _, opt := range v.Options {
					_, err := createOptions(optionService, c.UserContext(), opt.toDomain(question.ID))
					if err != nil {
						return c.SendStatus(fiber.StatusInternalServerError)
					}
				}

			case uint8(questionDomain.QuestionTypeAnatomical):
				if len(v.Options) != 0 {
					return c.SendStatus(fiber.StatusBadRequest)
				}
				_, err := createQuestion(questionService, c.UserContext(), v.toDomain(questionsReq.QuestionnaireID))
				if err != nil {
					return c.SendStatus(fiber.StatusInternalServerError)
				}

			default:
				return c.SendStatus(fiber.StatusBadRequest)
			}
		}

		return c.SendStatus(fiber.StatusCreated)
	}
}

func createQuestion(questionService questionPort.Service, c context2.Context, q questionDomain.Question) (questionDomain.Question, error) {
	return questionService.CreateQuestion(c, q)
}

func createOptions(optionService optionPort.Service, ctx context2.Context, option optionDomain.Option) (optionDomain.OptionID, error) {
	return optionService.CreateOption(ctx, option)
}

type CreateQuestionnaireRequest struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Duration    uint   `json:"duration,omitempty"`
	Editable    bool   `json:"editable,omitempty"`
	Randomable  bool   `json:"randomable,omitempty"`
	ValidTo     int64  `json:"validTo,omitempty"`
}

type CreateQuestionsRequest struct {
	QuestionnaireID string           `json:"questionnaireId"`
	Questions       []CreateQuestion `json:"questions"`
}

type CreateQuestion struct {
	Type    uint8                  `json:"type,omitempty"`
	Title   string                 `json:"title,omitempty"`
	Number  uint8                  `json:"number,omitempty"`
	Count   uint8                  `json:"count,omitempty"`
	Media   string                 `json:"media,omitempty"`
	Options []CreateOptionQuestion `json:"options,omitempty"`
}

func (q *CreateQuestion) toDomain(qnnID string) questionDomain.Question {
	return questionDomain.Question{
		QuestionnaireID: domain.QuestionnaireID(qnnID),
		Type:            questionDomain.QuestionType(q.Type),
		Number:          uint(q.Number),
		Count:           uint(q.Count),
		Title:           q.Title,
		Media:           questionDomain.MediaPath(q.Media),
	}
}

func (q *CreateQuestion) IsValid() bool {
	ty := questionDomain.QuestionType(q.Type)
	if ty == questionDomain.QuestionTypeOptional && len(q.Options) == 0 {
		return false
	}
	if ty == questionDomain.QuestionTypeAnatomical && len(q.Options) != 0 {
		return false
	}
	return true
}

type CreateOptionQuestion struct {
	Text     string `json:"text,omitempty"`
	IsAnswer bool   `json:"isAnswer,omitempty"`
}

func (opt *CreateOptionQuestion) toDomain(questionID questionDomain.QuestionDbID) optionDomain.Option {
	return optionDomain.Option{
		QuestionID: questionID,
		Text:       opt.Text,
		IsAnswer:   opt.IsAnswer,
	}
}
