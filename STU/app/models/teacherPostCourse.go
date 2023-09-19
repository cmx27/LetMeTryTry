package models

type TeacherPostCourse struct {
	ID          uint   `json:"class_id"`
	UserID      uint   `json:"user_id"`
	ClassName   string `json:"class_name"`
	Time        uint   `json:"time"`
	Weekday     uint   `json:"weekday"`
	Type        uint   `json:"type"`
	Number      uint   `json:"number"`
	Total       uint   `json:"total"`
	TeacherName string `json:"teacher_name"`
}
