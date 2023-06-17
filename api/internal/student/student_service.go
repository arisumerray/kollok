package student

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

func (s *service) GetAll(c context.Context) ([]Student, error) {
	d, err := s.Repository.GetAll(c)
	if err != nil {
		return d, err
	}
	return d, nil
}

func (s *service) CreateStudent(c context.Context, req *CreateStudentReq) (*Student, error) {
	dish, err := s.Repository.CreateStudent(c, req)
	if err != nil {
		return &Student{}, err
	}
	return dish, nil
}
