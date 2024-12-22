package models

type Student struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
}

type Teacher struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Laboratory struct {
	ID         uint   `json:"id" gorm:"primary_key"`
	Location   string `json:"location"`
	Capacity   string `json:"capacity"`
	TeacherNum string `json:"teacher_id"`
}

type Course struct {
	ID         uint   `json:"id" gorm:"primary_key"`
	CourseName string `json:"name"`
	TeacherID  uint   `json:"teacher_id"`
	Location   string `json:"location"`
	ClassTime  string `json:"class_time"`
}

type StudentTeacher struct {
	StudentID int `json:"student_id"`
	TeacherID int `json:"teacher_id"`
}
