package domain

import "github.com/QBC8-Go-Group2/questionnaire/internal/question/domain"

type (
	OptionID uint
)

type Option struct {
	ID         OptionID
	QuestionID domain.QuestionDbID
	Text       string
	IsAnswer   bool
}
