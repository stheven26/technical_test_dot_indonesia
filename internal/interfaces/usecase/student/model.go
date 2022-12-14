package student

type StudentRequest struct {
	Name   string `json:"name" validate:"required"`
	Age    int    `json:"age" validate:"required,number"`
	School string `json:"school" validate:"required"`
	Grade  int    `json:"grade" validate:"required,number"`
}

type PatchStudentNameRequest struct {
	Name string `json:"name" validate:"required"`
}

type StudentResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type StudentResponseData struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Age       int    `json:"age,omitempty"`
	School    string `json:"school,omitempty"`
	Grade     int    `json:"grade,omitempty" `
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
