package models

type Choice struct {
	ID          uint   `json:"-"`
	ClassID     uint   `json:"class_id"`
	UserID      uint   `json:"user_id"`
	ClassName   string `json:"class_name"`
	Time        uint   `json:"time"`
	Weekday     uint   `json:"weekday"`
	Type        uint   `json:"type"`
	TeacherName string `json:"teacher_name"`
	Date        string `json:"date"`
}
