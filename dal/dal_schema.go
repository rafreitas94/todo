package dal

import "time"

// domain
// storage

// contrato da dal -> interface
// CRUD

// create
// read
// update
// delete
// list

type Task struct {
	ID          string
	Subject     string
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
