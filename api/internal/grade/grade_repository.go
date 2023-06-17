package grade

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

func (r *repository) GetGrades(ctx context.Context, studentId int64) ([]Grade, error) {
	grades := make([]Grade, 0)
	query := "SELECT * FROM grade WHERE studentId = $1"
	rows, err := r.db.QueryContext(ctx, query, studentId)
	if err != nil {
		return grades, err
	}
	for rows.Next() {
		var grade Grade
		err = rows.Scan(&grade.Id, &grade.StudentId, &grade.Grade, &grade.Subject)
		if err != nil {
			return grades, err
		}
		grades = append(grades, grade)
	}
	return grades, nil
}

func (r *repository) CreateGrade(ctx context.Context, grade *CreateGradeReq) (*Grade, error) {
	d := Grade{
		StudentId: grade.StudentId,
		Grade:     grade.Grade,
		Subject:   grade.Subject,
	}
	var lastInsertId int
	query := "INSERT INTO \"grade\"(studentId, grade, subject) VALUES ($1, $2, $3) returning id"
	err := r.db.QueryRowContext(ctx, query, grade.StudentId, grade.Grade, grade.Subject).Scan(&lastInsertId)
	if err != nil {
		return &Grade{}, err
	}

	d.Id = int64(lastInsertId)
	return &d, nil
}
