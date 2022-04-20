package delivery

type CreateTaskRequest struct {
	Name string `json:"name" form:"name" binding:"required"`
}

type UpdateTaskRequest struct {
	Name   string `json:"name" form:"name" binding:"required"`
	Status *uint8 `json:"status" form:"status" binding:"required,oneof=0 1"`
}
