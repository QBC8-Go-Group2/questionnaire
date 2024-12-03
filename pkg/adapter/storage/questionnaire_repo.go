package storage

import (
	"context"
	"github.com/QBC8-Go-Group2/questionnaire/internal/questionnaire/domain"
	"github.com/QBC8-Go-Group2/questionnaire/internal/questionnaire/port"
	"github.com/QBC8-Go-Group2/questionnaire/pkg/adapter/storage/mapper"
	"github.com/QBC8-Go-Group2/questionnaire/pkg/adapter/storage/types"
	"gorm.io/gorm"
)

type questionnaireRepo struct {
	db *gorm.DB
}

func NewQuestionnaireRepo(db *gorm.DB) port.Repo {
	return &questionnaireRepo{
		db: db,
	}
}

func (q *questionnaireRepo) Create(ctx context.Context, questionnaire domain.Questionnaire) (domain.QuestionnaireDbID, error) {
	qStorage := mapper.QuestionnaireDomain2Storage(questionnaire)
	return domain.QuestionnaireDbID(qStorage.ID), q.db.Table("questionnaires").WithContext(ctx).Create(&qStorage).Error
}

func (q *questionnaireRepo) Update(ctx context.Context, questionnaire domain.Questionnaire) (domain.QuestionnaireID, error) {
	qStorage := mapper.QuestionnaireDomain2Storage(questionnaire)
	return domain.QuestionnaireID(qStorage.QuestionnaireID), q.db.Table("questionnaires").WithContext(ctx).Updates(&qStorage).Error
}

func (q *questionnaireRepo) FindWithQuestionnaireID(ctx context.Context, questionnaireId domain.QuestionnaireID) (domain.Questionnaire, error) {
	var qStorage *types.Questionnaire
	err := q.db.Table("questionnaires").WithContext(ctx).Where("questionnaire_id = ?", questionnaireId).First(qStorage).Error
	if err != nil {
		return domain.Questionnaire{}, err
	}
	return mapper.QuestionnaireStorage2Domain(qStorage), nil
}

func (q *questionnaireRepo) FindWithQuestionnaireDbID(ctx context.Context, questionnaireId domain.QuestionnaireDbID) (domain.Questionnaire, error) {
	var qStorage *types.Questionnaire
	err := q.db.Table("questionnaires").WithContext(ctx).Where("id = ?", questionnaireId).First(qStorage).Error
	if err != nil {
		return domain.Questionnaire{}, err
	}

	return mapper.QuestionnaireStorage2Domain(qStorage), nil
}

func (q *questionnaireRepo) DeleteWithQuestionnaireID(ctx context.Context, questionnaireId domain.QuestionnaireID) error {
	return q.db.Table("questionnaires").WithContext(ctx).Where("questionnaire_id = ?", questionnaireId).Delete(&types.Questionnaire{}).Error
}

func (q *questionnaireRepo) DeleteWithUserDbId(ctx context.Context, questionnaireId domain.QuestionnaireDbID) error {
	return q.db.Table("questionnaires").WithContext(ctx).Where("id = ?", questionnaireId).Delete(&types.Questionnaire{}).Error
}
