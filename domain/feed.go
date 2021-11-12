package domain

import "time"

type Feed struct {
	ID          string
	Title       string
	Description string
	Link        string
	UpdatedAT   *time.Time
	CreatedAT   *time.Time
}
