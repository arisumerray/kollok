package student

import (
	"context"
)

type Student struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Age  int32  `json:"age"`
	Spec string `json:"spec"`
}

type CreateStudentReq struct {
	Name string `json:"name"`
	Age  int32  `json:"age"`
	Spec string `json:"spec"`
}

type Repository interface {
	CreateStudent(ctx context.Context, dish *CreateStudentReq) (*Student, error)
	GetAll(ctx context.Context) ([]Student, error)
}

type Service interface {
	CreateStudent(ctx context.Context, dish *CreateStudentReq) (*Student, error)
	GetAll(ctx context.Context) ([]Student, error)
}
