package port

import (
	"context"
	"github.com/QBC8-Go-Group2/questionnaire/internal/question/domain"
	questionnaireDomain "github.com/QBC8-Go-Group2/questionnaire/internal/questionnaire/domain"
	userDomian "github.com/QBC8-Go-Group2/questionnaire/internal/user/domain"
)

type Service interface {
	CreateQuestion(ctx context.Context, questionnaire domain.Question) (domain.Question, error)
	UpdateQuestion(ctx context.Context, questionnaire domain.Question) (domain.Question, error)
	FindQuestionWithQuestionDbID(ctx context.Context, questionnaireId domain.QuestionDbID) (domain.Question, error)
	FindQuestionWithQuestionnaireDBID(ctx context.Context, questionnaireId questionnaireDomain.QuestionnaireDbID) ([]domain.Question, error)
	DeleteQuestionWithQuestionID(ctx context.Context, questionnaireId domain.QuestionDbID) error
	DeleteQuestionWithUserDbId(ctx context.Context, userID userDomian.UserDbID) error
	// add filter
}
