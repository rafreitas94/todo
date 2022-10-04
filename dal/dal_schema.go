package dal

import "time"

type Task struct {
	ID          string
	Subject     string
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
