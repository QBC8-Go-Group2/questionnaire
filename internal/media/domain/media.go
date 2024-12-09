package domain

import "time"

type (
	MediaID   uint
	MediaType string
	MediaUUID string
)

type Media struct {
	ID        MediaID
	UUID      MediaUUID
	UserID    uint
	Path      string
	Type      MediaType
	Size      int64
	Name      string
	CreatedAt time.Time
}
