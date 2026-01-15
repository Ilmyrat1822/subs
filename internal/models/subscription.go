package models

import (
	"time"

	"github.com/google/uuid"
)

type Subscription struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	ServiceName string    `gorm:"type:varchar(255);not null"`
	Price       int       `gorm:"not null;check:price >= 0"`
	UserID      uuid.UUID `gorm:"type:uuid;not null"`
	StartDate   string    `gorm:"type:varchar(7);not null"`
	EndDate     *string   `gorm:"type:varchar(7)"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
