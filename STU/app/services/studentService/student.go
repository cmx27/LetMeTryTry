package studentService

import (
	"STU/app/models"
	"STU/config/database"
)

func GetUserByUserID(user_id uint) (*models.User, error) {
	var user models.User
	result := database.DB.Where("id = ?", user_id).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func GetClassById(class_id uint) (*models.TeacherPostCourse, error) {
	var class models.TeacherPostCourse
	result := database.DB.Where("id = ?", class_id).First(&class)
	if result.Error != nil {
		return nil, result.Error
	}
	return &class, nil
}

func GetAllChoice(id uint) (*[]models.Choice, error) {
	var class []models.Choice
	result := database.DB.Where(&models.Choice{UserID: id}).Find(&class)
	if result.Error != nil {
		return nil, result.Error
	}
	return &class, nil
}

func ChoiceClass(course models.Choice) error {
	result := database.DB.Create(&course)
	return result.Error
}

func GetCourseList() ([]models.TeacherPostCourse, error) {
	//database.DB.Find(&models.TeacherPostCourse{})
	var courseList []models.TeacherPostCourse
	result := database.DB.Find(&courseList)
	if result.Error != nil {
		return nil, result.Error
	} else {
		return courseList, nil
	}

}

func GetTakenCourseList(user_id uint) ([]models.Choice, error) {

	result := database.DB.Where("user_id = ?", user_id).First(&models.Choice{})
	if result.Error != nil {
		return nil, result.Error
	}
	var courseList []models.Choice
	result = database.DB.Where("user_id = ?", user_id).Find(&courseList)
	if result.Error != nil {
		return nil, result.Error
	} else {
		return courseList, nil
	}

}

func CheckClass(class_id uint) error {
	result := database.DB.Where("class_id = ?", class_id).First(&models.Choice{})
	return result.Error
}

func GetUserByClassID(class_id uint) (*models.Choice, error) {
	var user models.Choice
	result := database.DB.Where("class_id = ?", class_id).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func DeleteCourse(class_id uint) error {
	result := database.DB.Where("class_id = ?", class_id).Delete(&models.Choice{})
	return result.Error
}

func CheckStudent(user_id uint) error {
	result := database.DB.Where("id = ?", user_id).First(&models.User{})
	return result.Error
}

// func CheckClassExist(class_id uint) error {
// 	result := database.DB.Where("class_id = ?", class_id).First(&models.Choice{})
// 	if result.Error != nil {
// 		return nil
// 	}
// 	return result.Error
// }
