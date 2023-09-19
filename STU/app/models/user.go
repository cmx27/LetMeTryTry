package models

type User struct {
	Type     uint   `json:"type"`
	ID       uint   `json:"user_id"`
	Account  string `json:"-"  gorm:"type:varchar(20) check:Account > 8"`
	Password string `json:"-" gorm:"type:varchar(20) check:Password > 8"`
	Name     string `json:"-"`
}
