package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	UserName  string    `gorm:"index;unique"`
	CreatedAt time.Time `gorm:"default:now()"`
}
