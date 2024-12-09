package questionnaire

import (
	"context"
	domain2 "github.com/QBC8-Go-Group2/questionnaire/internal/question/domain"
	"github.com/QBC8-Go-Group2/questionnaire/internal/question/port"
	"github.com/QBC8-Go-Group2/questionnaire/internal/questionnaire/domain"
	userDomain "github.com/QBC8-Go-Group2/questionnaire/internal/user/domain"
)

type service struct {
	repo port.Repo
}

func NewService(repo port.Repo) port.Service {
	return &service{repo}
}

func (s *service) CreateQuestion(ctx context.Context, questionnaire domain2.Question) (domain2.Question, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) UpdateQuestion(ctx context.Context, questionnaire domain2.Question) (domain2.Question, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) FindQuestionWithQuestionDbID(ctx context.Context, questionnaireId domain2.QuestionDbID) (domain2.Question, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) FindQuestionWithQuestionnaireDBID(ctx context.Context, questionnaireId domain.QuestionnaireDbID) ([]domain2.Question, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) DeleteQuestionWithQuestionID(ctx context.Context, questionnaireId domain2.QuestionDbID) error {
	//TODO implement me
	panic("implement me")
}

func (s *service) DeleteQuestionWithUserDbId(ctx context.Context, userID userDomain.UserDbID) error {
	//TODO implement me
	panic("implement me")
}
