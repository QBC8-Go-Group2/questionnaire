package domain

import (
	userDomian "github.com/QBC8-Go-Group2/questionnaire/internal/user/domain"
	"time"
)

type (
	QuestionnaireDbID uint
	QuestionnaireID   string
)

type Questionnaire struct {
	ID              QuestionnaireDbID
	QuestionnaireID QuestionnaireID
	OwnerID         userDomian.UserID
	Title           string
	Description     string
	Duration        time.Time
	Editable        bool
	Randomable      bool
	CreatedAt       time.Time
	ValidTo         time.Time
}
