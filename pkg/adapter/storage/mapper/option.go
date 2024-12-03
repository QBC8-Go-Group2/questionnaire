package mapper

import (
	"github.com/QBC8-Go-Group2/questionnaire/internal/option/domain"
	questionDomain "github.com/QBC8-Go-Group2/questionnaire/internal/question/domain"
	"github.com/QBC8-Go-Group2/questionnaire/pkg/adapter/storage/types"
)

func OptionDomain2Storage(option domain.Option) types.Option {
	return types.Option{
		ID:         uint(option.ID),
		QuestionID: uint(option.QuestionID),
		Text:       option.Text,
		IsAnswer:   option.IsAnswer,
	}
}
func OptionStorage2Domain(option types.Option) domain.Option {
	return domain.Option{
		ID:         domain.OptionID(option.ID),
		QuestionID: questionDomain.QuestionDbID(option.QuestionID),
		Text:       option.Text,
		IsAnswer:   option.IsAnswer,
	}
}
