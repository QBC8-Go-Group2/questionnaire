package mapper

import (
	optionDomain "github.com/QBC8-Go-Group2/questionnaire/internal/option/domain"
	questionDomain "github.com/QBC8-Go-Group2/questionnaire/internal/question/domain"
	"github.com/QBC8-Go-Group2/questionnaire/internal/response/domain"
	userDomian "github.com/QBC8-Go-Group2/questionnaire/internal/user/domain"
	"github.com/QBC8-Go-Group2/questionnaire/pkg/adapter/storage/types"
)

func ResponseDomain2Storage(response domain.Response) types.Response {
	return types.Response{
		ID:         uint(response.ID),
		Type:       uint8(response.Type),
		UserID:     uint(response.UserID),
		QuestionID: uint(response.QuestionID),
		Data:       response.Data,
		OptionID:   uint(response.OptionID),
		CreatedAt:  response.CreatedAt,
	}
}

func ResponseStorage2Domain(response types.Response) domain.Response {
	return domain.Response{
		ID:         domain.ResponseID(response.ID),
		Type:       questionDomain.QuestionType(response.Type),
		UserID:     userDomian.UserDbID(response.UserID),
		QuestionID: questionDomain.QuestionDbID(response.QuestionID),
		Data:       response.Data,
		OptionID:   optionDomain.OptionID(response.OptionID),
		CreatedAt:  response.CreatedAt,
	}
}
