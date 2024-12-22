package main

import (
	"github.com/gin-gonic/gin"
	"peninsula12/python_student/config"
	"peninsula12/python_student/controllers"
	"peninsula12/python_student/models"
)

func main() {
	r := gin.Default()

	// 自动迁移数据库
	err := config.DB.AutoMigrate(&models.Student{}, &models.Teacher{}, &models.Laboratory{}, &models.Course{},
		&models.StudentTeacher{})
	if err != nil {
		panic(err)
	}

	// 路由设置
	r.POST("/login", controllers.Login)
	r.PUT("/password", controllers.ChangePassword)

	r.GET("/students", controllers.GetStudents)
	r.POST("/students", controllers.CreateStudent)
	r.DELETE("/students/:id", controllers.DeleteStudent)

	r.GET("/labs", controllers.GetLaboratories)

	r.GET("/courses", controllers.GetCourses)

	// 启动服务器
	_ = r.Run(":8080")
}
