package requests

type RegisterUserRequestDTO struct {
	UserName        string `json:"user_name" form:"username" binding:"required"`
	Email           string `json:"email" form:"email" binding:"required"`
	PhoneNumber     string `json:"phone_number" form:"phone_number" binding:"required"`
	BOD             string `json:"bod" form:"bod" binding:"required"`
	ProfilePicture  string `json:"profile_picture" form:"profile_picture" binding:"omitempty"`
	Password        string `json:"password" form:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password" binding:"required"`
	IsActive     bool      `json:"is_active" form:"is_available" binding:"omitempty"`
	CreatedBy       string `json:"created_by" form:"created_by" binding:"required"`
}
