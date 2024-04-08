package responses

import (
	"time"
)

type LogResponseDTO struct {
	IsActive       bool      `json:"is_active"`
	CreatedBy      string    `json:"created_by"`
	UpdatedBy      string    `json:"updated_by"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}