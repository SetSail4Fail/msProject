package entity

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	Id        uuid.UUID `db:"id"`
	Name      string    `db:"name"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
}
