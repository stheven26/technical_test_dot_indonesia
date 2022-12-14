package class

type ClassRequest struct {
	Subject       string `json:"subject" validate:"required"`
	DurationClass int    `json:"durationClass" validate:"required,number"`
	StudentName   string `json:"studentName" validate:"required"`
}

type ClassResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ClassResponseData struct {
	Id            int64  `json:"id"`
	Subject       string `json:"subject"`
	DurationClass int    `json:"durationClass"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}
