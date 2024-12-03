package domain

import "time"

type (
	RoleType uint8
	UserDbID uint
	UserID   string
)

const (
	AdminRole RoleType = iota
	UserRole
)

type User struct {
	ID        UserDbID
	CreatedAT time.Time
	UserID    UserID
	Email     string
	Password  string
	NatId     string
	Role      RoleType
}
