package dal

import (
	"errors"
	"time"
)

var ErrNotFound = errors.New("não encontrado")

type Task struct {
	ID          string    `json:"id" db:"id"`
	Subject     string    `json:"subject" db:"subject"`
	Description string    `json:"description" db:"description"`
	Status      string    `json:"status" db:"status"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
