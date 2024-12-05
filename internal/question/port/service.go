package port

import (
	"context"
	"github.com/QBC8-Go-Group2/questionnaire/internal/questionnaire/domain"
)

type Service interface {
	CreateQuestionnaire(ctx context.Context, questionnaire domain.Questionnaire) (domain.QuestionnaireDbID, error)
	UpdateQuestionnaire(ctx context.Context, questionnaire domain.Questionnaire) (domain.QuestionnaireID, error)
	FindQuestionnaireWithQuestionnaireID(ctx context.Context, questionnaireId domain.QuestionnaireID) (domain.Questionnaire, error)
	FindQuestionnaireWithQuestionnaireDbID(ctx context.Context, questionnaireId domain.QuestionnaireDbID) (domain.Questionnaire, error)
	DeleteQuestionnaireWithQuestionnaireID(ctx context.Context, questionnaireId domain.QuestionnaireID) error
	DeleteQuestionnaireWithUserDbId(ctx context.Context, questionnaireId domain.QuestionnaireDbID) error
	// add filter
}
