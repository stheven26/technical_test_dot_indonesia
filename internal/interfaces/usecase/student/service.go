package student

import "context"

type StudentService interface {
	CreateStudent(ctx context.Context, req StudentRequest) (res StudentResponse, err error)
	GetAll(ctx context.Context) (res StudentResponse, err error)
	GetOneById(ctx context.Context, id string) (res StudentResponse, err error)
	UpdateById(ctx context.Context, id string, req StudentRequest) (res StudentResponse, err error)
	PatchById(ctx context.Context, id string, req PatchStudentNameRequest) (res StudentResponse, err error)
	DeleteById(ctx context.Context, id string) (res StudentResponse, err error)
}
