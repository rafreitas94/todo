package dal

import (
	"errors"
	"time"
)

// domain
// storage

// contrato da dal -> interface
// CRUD

// create
// read
// update
// delete
// list

var ErrorNotFound = errors.New("não encontrado")

type Task struct {
	ID          string    `json:"id"`
	Subject     string    `json:"subject"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
