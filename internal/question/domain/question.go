package domain

import questionnaireDomain "github.com/QBC8-Go-Group2/questionnaire/internal/questionnaire/domain"

type (
	QuestionDbID uint
	QuestionType uint8
	MediaPath    string
)

const (
	QuestionTypeAnatomical QuestionType = iota + 1
	QuestionTypeOptional
)

type Question struct {
	ID              QuestionDbID
	QuestionnaireID questionnaireDomain.QuestionnaireID
	Type            QuestionType
	Number          uint
	Count           uint
	Title           string
	Media           MediaPath
}
