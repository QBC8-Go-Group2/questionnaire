package domain

import (
	userDomian "github.com/QBC8-Go-Group2/questionnaire/internal/user/domain"
	"github.com/dchest/siphash"
	"time"
)

type (
	QuestionnaireDbID uint
	QuestionnaireID   string
)

type Questionnaire struct {
	ID              QuestionnaireDbID
	QuestionnaireID QuestionnaireID
	OwnerID         userDomian.UserDbID
	Title           string
	Description     string
	Duration        uint // minutes
	Editable        bool
	Randomable      bool
	CreatedAt       time.Time
	ValidTo         time.Time
}

const KEY = "Quera"

func CreateQuestionnaireID(word string) (QuestionnaireID, error) {
	h := siphash.New([]byte(KEY))
	_, err := h.Write([]byte(word))
	if err != nil {
		return "", err
	}
	return QuestionnaireID(h.Sum(nil)), nil
}
