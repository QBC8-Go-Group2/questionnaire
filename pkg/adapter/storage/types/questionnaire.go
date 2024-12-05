package types

import (
	"database/sql"
	"time"
)

type Questionnaire struct {
	ID              uint         `gorm:"primaryKey;autoIncrement"`
	QuestionnaireID string       `gorm:"unique;not null;size:100"`
	OwnerID         uint         `gorm:"not null;size:100"`
	Title           string       `gorm:"not null;size:255"`
	Description     string       `gorm:"size:500"`
	Duration        uint         `gorm:"not null"`
	Editable        bool         `gorm:"not null"`
	Randomable      bool         `gorm:"not null"`
	CreatedAt       time.Time    `gorm:"autoCreateTime"`
	ValidTo         sql.NullTime `gorm:"default:null"`
}
