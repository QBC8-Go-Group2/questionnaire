package types

import "time"

type Response struct {
	ID         uint  `gorm:"primary_key;auto_increment;unique"`
	Type       uint8 `gorm:"not null"`
	UserID     uint  `gorm:"not null"`
	QuestionID uint  `gorm:"not null"`
	Data       string
	OptionID   uint
	CreatedAt  time.Time `gorm:"autoCreateTime;not null"`
}
