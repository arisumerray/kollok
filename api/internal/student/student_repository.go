package student

import (
	"context"
	"database/sql"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type repository struct {
	db DBTX
}

func NewRepository(db DBTX) Repository {
	return &repository{db: db}
}

func (r *repository) GetAll(ctx context.Context) ([]Student, error) {
	students := make([]Student, 0)
	query := "SELECT * FROM student"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return students, err
	}
	for rows.Next() {
		var dish Student
		err = rows.Scan(&dish.Id, &dish.Name, &dish.Age, &dish.Spec)
		if err != nil {
			return students, err
		}
		students = append(students, dish)
	}
	return students, nil
}

func (r *repository) CreateStudent(ctx context.Context, student *CreateStudentReq) (*Student, error) {
	d := Student{
		Name: student.Name,
		Age:  student.Age,
		Spec: student.Spec,
	}
	var lastInsertId int
	query := "INSERT INTO \"student\"(name, age, spec) VALUES ($1, $2, $3) returning id"
	err := r.db.QueryRowContext(ctx, query, student.Name, student.Age, student.Spec).Scan(&lastInsertId)
	if err != nil {
		return &Student{}, err
	}

	d.Id = int64(lastInsertId)
	return &d, nil
}
