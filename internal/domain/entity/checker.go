package entity

import "time"

type Checker struct {
	id        string
	kind      string
	createdAt time.Time
	updatedAt time.Time
}
