package port

import (
	"context"
	"github.com/QBC8-Go-Group2/questionnaire/internal/questionnaire/domain"
)

type Repo interface {
	Create(ctx context.Context, questionnaire domain.Questionnaire) (domain.QuestionnaireDbID, error)
	Update(ctx context.Context, questionnaire domain.Questionnaire) (domain.QuestionnaireID, error)
	FindWithQuestionnaireID(ctx context.Context, questionnaireId domain.QuestionnaireID) (domain.Questionnaire, error)
	FindWithQuestionnaireDbID(ctx context.Context, questionnaireId domain.QuestionnaireDbID) (domain.Questionnaire, error)
	DeleteWithQuestionnaireID(ctx context.Context, questionnaireId domain.QuestionnaireID) error
	DeleteWithUserDbId(ctx context.Context, questionnaireId domain.QuestionnaireDbID) error
	// add filters
}
