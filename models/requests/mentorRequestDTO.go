package requests

type RegisterMentorRequestDTO struct {
	UserID         string    `json:"user_id" binding:"required"`
	Username       string    `json:"user_name" binding:"required"`
	University     string    `json:"university" binding:"required"`
	Email          string    `json:"email" binding:"required"`
	PhoneNumber    string    `json:"phone_number" binding:"required"`
	Description    string    `json:"description" binding:"omitempty"`
	ProfilePicture string    `json:"profile_picture" binding:"omitempty"`
	BOD            string `json:"bod" binding:"required"`
	Courses        []string  `json:"courses" binding:"required"`
}
