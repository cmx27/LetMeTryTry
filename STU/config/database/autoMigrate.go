package database

import (
	"STU/app/models"

	"gorm.io/gorm"
)

func autoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.User{},
		&models.TeacherPostCourse{},
		&models.Choice{},
	)

	return err
}
