package questionnaire

import (
	"context"
	"github.com/QBC8-Go-Group2/questionnaire/internal/questionnaire/domain"
	"github.com/QBC8-Go-Group2/questionnaire/internal/questionnaire/port"
)

type service struct {
	repo port.Repo
}

func (s *service) CreateQuestionnaire(ctx context.Context, questionnaire domain.Questionnaire) (domain.QuestionnaireDbID, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) UpdateQuestionnaire(ctx context.Context, questionnaire domain.Questionnaire) (domain.QuestionnaireID, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) FindQuestionnaireWithQuestionnaireID(ctx context.Context, questionnaireId domain.QuestionnaireID) (domain.Questionnaire, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) FindQuestionnaireWithQuestionnaireDbID(ctx context.Context, questionnaireId domain.QuestionnaireDbID) (domain.Questionnaire, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) DeleteQuestionnaireWithQuestionnaireID(ctx context.Context, questionnaireId domain.QuestionnaireID) error {
	//TODO implement me
	panic("implement me")
}

func (s *service) DeleteQuestionnaireWithUserDbId(ctx context.Context, questionnaireId domain.QuestionnaireDbID) error {
	//TODO implement me
	panic("implement me")
}
