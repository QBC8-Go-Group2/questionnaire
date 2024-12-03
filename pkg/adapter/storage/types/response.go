package types

import "time"

type Response struct {
	ID         uint      `gorm:"primary_key;auto_increment;unique"`
	Type       uint8     `gorm:"not null"`
	UserID     uint      `gorm:"not null"`
	QuestionID uint      `gorm:"not null"`
	Data       string    `gorm:"not null"`
	OptionID   uint      `gorm:"not null"`
	CreatedAt  time.Time `gorm:"not null"`
}
