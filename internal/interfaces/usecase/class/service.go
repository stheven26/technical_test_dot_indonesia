package class

import "context"

type ClassService interface {
	CreateClass(ctx context.Context, req ClassRequest) (res ClassResponse, err error)
	GetAll(ctx context.Context) (res ClassResponse, err error)
	GetOneById(ctx context.Context, id string) (res ClassResponse, err error)
	DeleteById(ctx context.Context, id string) (res ClassResponse, err error)
}
