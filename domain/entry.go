package domain

import "time"

type FeedEntry struct {
	ID          string
	Title       string
	Description string
	Link        string
	FeedID      string
	CreatedAt   *time.Time
}
