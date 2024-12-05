package port

import (
	"context"
	"github.com/QBC8-Go-Group2/questionnaire/internal/question/domain"
)

type Repo interface {
	Create(ctx context.Context, question domain.Question) (domain.QuestionDbID, error)
	Update(ctx context.Context, question domain.Question) (domain.QuestionDbID, error)
	FindWithQuestionID(ctx context.Context, questionId domain.QuestionDbID) (domain.Question, error)
	DeleteWithQuestionID(ctx context.Context, questionId domain.QuestionDbID) error
	// add filters
}
