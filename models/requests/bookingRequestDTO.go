package requests

type NewBookingRequestDTO struct {
	UserID            string `json:"user_id"`
	MentorID          string `json:"mentor_id"`
	CourseID          string `json:"course_id"`
	ScheduleTime      string `json:"scheduled_at"`
	ScheduledLocation string `json:"scheduled_location"`
	Invoice           struct {
		Payment_method string  `json:"payment_method"`
		Payment_name   string  `json:"payment_name"`
		Payment_status string  `json:"payment_status"`
		Payment_amount float64 `json:"payment_amount"`
	} `json:"invoice"`
}
