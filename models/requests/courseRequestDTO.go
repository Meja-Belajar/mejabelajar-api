package requests

type GetCourseRequestDTO struct {
	ID string `json:"id" form: "id" binding:"required"`
}

type AddCourseRequestDTO struct {
	Name      string `json:"name" form: "name" binding:"required"`
	Detail    string `json:"detail" form: "detail" binding:"required"`
	IsActive  bool   `json:"isactive" form: "isactive" binding:"required"`
	CreatedBy string `json:"createdby" form: "createdby" binding:"required"`
}

type UpdateCourseRequestDTO struct {
	ID        string `json:"id" form: "id" binding:"required"`
	Name      string `json:"name" form: "name" binding:"omitempty"`
	Detail    string `json:"detail" form: "detail" binding:"omitempty"`
	IsActive  bool   `json:"isactive" form: "isactive" binding:"omitempty"`
	UpdatedBy string `json:"updatedby" form: "updatedby" binding:"required"`
}
