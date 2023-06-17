package grade

import (
	"context"
)

type service struct {
	Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository,
	}
}

func (s *service) GetGrades(c context.Context, studentId int64) ([]Grade, error) {
	d, err := s.Repository.GetGrades(c, studentId)
	if err != nil {
		return d, err
	}
	return d, nil
}

func (s *service) CreateStudent(c context.Context, req *CreateGradeReq) (*Grade, error) {
	dish, err := s.Repository.CreateGrade(c, req)
	if err != nil {
		return &Grade{}, err
	}
	return dish, nil
}
