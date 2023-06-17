package grade

import (
	"context"
)

type Grade struct {
	Id        int64  `json:"id"`
	StudentId int64  `json:"student_id"`
	Subject   string `json:"subject"`
	Grade     int32  `json:"grade"`
}

type CreateGradeReq struct {
	StudentId int64  `json:"student_id"`
	Subject   string `json:"subject"`
	Grade     int32  `json:"grade"`
}

type Repository interface {
	CreateGrade(ctx context.Context, dish *CreateGradeReq) (*Grade, error)
	GetGrades(ctx context.Context, studentId int64) ([]Grade, error)
}

type Service interface {
	CreateGrade(ctx context.Context, dish *CreateGradeReq) (*Grade, error)
	GetGrades(ctx context.Context, studentId int64) ([]Grade, error)
}
