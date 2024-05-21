package responses

import (
	"time"

	"github.com/google/uuid"
)

type NotificationReponse struct {
	ID          uuid.UUID `json:"notification_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}
