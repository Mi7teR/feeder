package domain

import "time"

type Subscriber struct {
	ID                  string
	FeedID              string
	OutputDestinationID string
	CreatedAT           *time.Time
	UpdatedAT           *time.Time
}
