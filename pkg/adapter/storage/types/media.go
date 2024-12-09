package types

import "time"

type Media struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	UUID      string    `gorm:"unique;not null;type:char(36)"`
	UserID    uint      `gorm:"not null"`
	Path      string    `gorm:"not null;size:255"`
	Type      string    `gorm:"not null;size:50"`
	Size      int64     `gorm:"not null"`
	Name      string    `gorm:"not null;size:255"`
	CreatedAt time.Time `gorm:"autoCreateTime;not null"`
}
