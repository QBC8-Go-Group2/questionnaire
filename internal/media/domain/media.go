package domain

import "time"

type (
	MediaID   uint
	MediaType string
)

type Media struct {
	ID        MediaID
	UserID    uint
	Path      string
	Type      MediaType
	Size      int64
	Name      string
	CreatedAt time.Time
}
