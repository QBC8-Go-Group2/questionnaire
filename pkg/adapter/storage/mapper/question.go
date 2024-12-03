package mapper

import (
	"github.com/QBC8-Go-Group2/questionnaire/internal/question/domain"
	questionnaireDomain "github.com/QBC8-Go-Group2/questionnaire/internal/questionnaire/domain"
	"github.com/QBC8-Go-Group2/questionnaire/pkg/adapter/storage/types"
)

func QuestionDomain2Storage(question domain.Question) types.Question {
	return types.Question{
		ID:              uint(question.ID),
		QuestionnaireID: string(question.QuestionnaireID),
		Type:            uint8(question.Type),
		Number:          question.Number,
		Count:           question.Count,
		Title:           question.Title,
		Media:           string(question.Media),
	}
}

func QuestionStorage2Domain(question types.Question) domain.Question {
	return domain.Question{
		ID:              domain.QuestionDbID(question.ID),
		QuestionnaireID: questionnaireDomain.QuestionnaireID(question.QuestionnaireID),
		Type:            domain.QuestionType(question.Type),
		Number:          question.Number,
		Count:           question.Count,
		Title:           question.Title,
		Media:           domain.MediaPath(question.Media),
	}
}
