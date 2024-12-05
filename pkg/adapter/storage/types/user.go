package types

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time `gorm:"autoCreateTime;not null"`
	Email     string    `gorm:"unique;not null;size:255"`
	Password  string    `gorm:"not null;size:255"`
	NatID     string    `gorm:"unique;not null"`
	Role      uint8     `gorm:"default:0;not null"`
}
