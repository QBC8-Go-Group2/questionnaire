package storage

import (
	"context"
	"github.com/QBC8-Go-Group2/questionnaire/internal/question/domain"
	"github.com/QBC8-Go-Group2/questionnaire/internal/question/port"
	"github.com/QBC8-Go-Group2/questionnaire/pkg/adapter/storage/mapper"
	"github.com/QBC8-Go-Group2/questionnaire/pkg/adapter/storage/types"
	"gorm.io/gorm"
)

type questionRepo struct {
	db *gorm.DB
}

func NewQuestionRepo(db *gorm.DB) port.Repo {
	return &questionRepo{
		db: db,
	}
}

func (q *questionRepo) Create(ctx context.Context, question domain.Question) (domain.QuestionDbID, error) {
	qStorage := mapper.QuestionDomain2Storage(question)
	return domain.QuestionDbID(qStorage.ID), q.db.Table("questions").WithContext(ctx).Create(&question).Error
}

func (q *questionRepo) Update(ctx context.Context, question domain.Question) (domain.QuestionDbID, error) {
	qStorage := mapper.QuestionDomain2Storage(question)
	return domain.QuestionDbID(qStorage.ID), q.db.Table("questions").WithContext(ctx).Updates(&question).Error
}

func (q *questionRepo) FindWithQuestionID(ctx context.Context, questionId domain.QuestionDbID) (domain.Question, error) {
	var question types.Question
	err := q.db.Table("questions").WithContext(ctx).Where("id = ?", questionId).First(&question).Error
	if err != nil {
		return domain.Question{}, err
	}
	return mapper.QuestionStorage2Domain(question), nil

}

func (q *questionRepo) DeleteWithQuestionID(ctx context.Context, questionId domain.QuestionDbID) error {
	return q.db.Table("questions").WithContext(ctx).Where("id = ?", questionId).Delete(&types.Question{}).Error
}
