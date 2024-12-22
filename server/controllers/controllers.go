package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"peninsula12/python_student/config"
	"peninsula12/python_student/models"
	"strconv"
)

// Login 登录
func Login(c *gin.Context) {
	var t struct {
		Number   string `json:"number"`
		Password string `json:"password"`
	}
	err := c.ShouldBindBodyWithJSON(&t)
	var teacher models.Teacher
	fmt.Println(t.Number, t.Password)
	err = config.DB.Where("id = ? AND password = ?", t.Number, t.Password).Find(&teacher).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "用户名或密码错误"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "登录成功", "number": teacher.ID, "name": teacher.Name})
	}
}

// ChangePassword 修改密码
func ChangePassword(c *gin.Context) {
	var request struct {
		Number      string `json:"number"`
		NewPassword string `json:"newPassword"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	number, _ := strconv.Atoi(request.Number)
	err := config.DB.Updates(&models.Teacher{ID: uint(number), Password: request.NewPassword}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "密码修改失败"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "密码修改成功"})
	}
}

// GetStudents 获取所有学生
func GetStudents(c *gin.Context) {
	number := c.Query("number")
	if number == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "缺少 number 参数"})
		return
	}
	num, err := strconv.Atoi(number)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的 number 参数"})
		return
	}
	var students []models.StudentTeacher
	err = config.DB.Where("teacher_id = ?", num).Find(&students).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "获取学生失败"})
		return
	}
	ans := make([]models.Student, 0, len(students))
	for _, val := range students {
		var student models.Student
		err := config.DB.Find(&student, val.StudentID).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "获取学生失败"})
			return
		}
		ans = append(ans, student)
	}

	c.JSON(http.StatusOK, gin.H{"data": ans})
}

// CreateStudent 创建学生
func CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&student)
	c.JSON(http.StatusOK, student)
}

// DeleteStudent 删除学生
func DeleteStudent(c *gin.Context) {

}

// GetLaboratories 获取所有实验室
func GetLaboratories(c *gin.Context) {
	number := c.Query("number")
	var laboratories []models.Laboratory
	err := config.DB.Where("teacher_num = ?", number).Find(&laboratories).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "获取实验室失败"})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": laboratories})
	}
}

// GetCourses 获取所有课程
func GetCourses(c *gin.Context) {
	number := c.Query("number")
	var courses []models.Course
	err := config.DB.Where("teacher_id = ?", number).Find(&courses).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "获取课程失败"})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": courses})
	}
}
