package main

import (
	"api/db"
	"api/internal/grade"
	"api/internal/student"
	"api/router"
	"log"
)

// hello world
func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("could not initialize database connection: %s", err)
	}
	// New Repo Instance

	studentRep := student.NewRepository(dbConn.GetDB())
	studentSvc := student.NewService(studentRep)
	studentHandler := student.NewHandler(studentSvc)

	gradeRep := grade.NewRepository(dbConn.GetDB())
	gradeSvc := grade.NewService(gradeRep)
	gradeHandler := grade.NewHandler(gradeSvc)

	router.InitRouter(studentHandler, gradeHandler)
	err = router.Start("0.0.0.0:8081")
	if err != nil {
		return
	}
}
