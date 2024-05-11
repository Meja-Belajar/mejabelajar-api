package requests

import "github.com/google/uuid"

type NewBookingRequestDTO struct {
	UserID   uuid.UUID
	MentorID uuid.UUID
	CourseID uuid.UUID

	// Data invoice

}
