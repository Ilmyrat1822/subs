package models

import (
	"time"

	"github.com/google/uuid"
)

type Subscription struct {
	ID          int       `json:"id" example:"1"`
	ServiceName string    `json:"service_name" binding:"required" example:"Yandex Plus"`
	Price       int       `json:"price" binding:"required,min=0" example:"400"`
	UserID      uuid.UUID `json:"user_id" binding:"required" example:"60601fee-2bf1-4721-ae6f-7636e79a0cba"`
	StartDate   string    `json:"start_date" binding:"required" example:"07-2025"`
	EndDate     *string   `json:"end_date,omitempty" example:"12-2025"`
	CreatedAt   time.Time `json:"created_at" example:"2025-01-12T10:00:00Z"`
	UpdatedAt   time.Time `json:"updated_at" example:"2025-01-12T10:00:00Z"`
}
