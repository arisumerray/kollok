package router

import (
	"api/internal/grade"
	"api/internal/student"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter(studentHandler *student.Handler, gradeHandler *grade.Handler) {
	r = gin.Default()

	r.Use(CORS())

	r.GET("/students", studentHandler.GetStudents)
	r.POST("/students", studentHandler.CreateStudent)
	r.GET("/grades/:studentId", gradeHandler.GetGrades)
	r.POST("/grades", gradeHandler.CreateGrade)
}

func Start(addr string) error {
	return r.Run(addr)
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
