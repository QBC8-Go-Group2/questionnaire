package types

type Question struct {
	ID              uint   `gorm:"primaryKey;autoIncrement"`
	QuestionnaireID string `gorm:"not null;size:100"`
	Type            uint8  `gorm:"not null"`
	Number          uint   `gorm:"not null"`
	Count           uint   `gorm:"not null"`
	Title           string `gorm:"not null;size:255"`
	Media           string `gorm:"size:500"`
}
