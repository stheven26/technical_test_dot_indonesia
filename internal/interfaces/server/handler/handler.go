package handler

import "technical-test/internal/interfaces/container"

type Handler struct {
	studentHandler *studentHandler
	classHandler   *classHandler
}

func (h *Handler) Validate() *Handler {
	if h.studentHandler == nil {
		panic("studentHandler is nil")
	}
	if h.classHandler == nil {
		panic("classHandler is nil")
	}
	return h
}

func SetupHandler(container container.Container) *Handler {
	return &Handler{
		studentHandler: NewStudentHandler().SetStudentService(container.StudentService).Validate(),
		classHandler:   NewClassHandler().SetClassService(container.ClassService).Validate(),
	}
}
