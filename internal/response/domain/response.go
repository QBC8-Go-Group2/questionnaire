package domain

import (
	optionDomain "github.com/QBC8-Go-Group2/questionnaire/internal/option/domain"
	questionDomain "github.com/QBC8-Go-Group2/questionnaire/internal/question/domain"
	userDomian "github.com/QBC8-Go-Group2/questionnaire/internal/user/domain"
	"time"
)

type (
	ResponseID uint
)

type Response struct {
	ID         ResponseID
	Type       questionDomain.QuestionType
	UserID     userDomian.UserDbID
	QuestionID questionDomain.QuestionDbID
	Data       string
	OptionID   optionDomain.OptionID
	CreatedAt  time.Time
}
