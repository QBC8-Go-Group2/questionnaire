package types

type Option struct {
	ID         uint   `gorm:"primary_key;auto_increment"`
	QuestionID uint   `gorm:"not null"`
	Text       string `gorm:"not null; size:255"`
	IsAnswer   bool   `gorm:"not null;default:false"`
}
