package mapper

import (
	"database/sql"
	"github.com/QBC8-Go-Group2/questionnaire/internal/questionnaire/domain"
	userDomian "github.com/QBC8-Go-Group2/questionnaire/internal/user/domain"
	"github.com/QBC8-Go-Group2/questionnaire/pkg/adapter/storage/types"
)

func QuestionnaireDomain2Storage(q domain.Questionnaire) types.Questionnaire {
	return types.Questionnaire{
		ID:              uint(q.ID),
		QuestionnaireID: string(q.QuestionnaireID),
		OwnerID:         string(q.OwnerID),
		Title:           q.Title,
		Description:     q.Description,
		Duration:        q.Duration,
		Editable:        q.Editable,
		Randomable:      q.Randomable,
		CreatedAt:       q.CreatedAt,
		ValidTo: sql.NullTime{
			Time:  q.ValidTo,
			Valid: !q.ValidTo.IsZero(),
		},
	}
}
func QuestionnaireStorage2Domain(q *types.Questionnaire) domain.Questionnaire {
	return domain.Questionnaire{
		ID:              domain.QuestionnaireDbID(q.ID),
		QuestionnaireID: domain.QuestionnaireID(q.QuestionnaireID),
		OwnerID:         userDomian.UserID(q.OwnerID),
		Title:           q.Title,
		Description:     q.Description,
		Duration:        q.Duration,
		Editable:        q.Editable,
		Randomable:      q.Randomable,
		CreatedAt:       q.CreatedAt,
		ValidTo:         q.ValidTo.Time,
	}
}
