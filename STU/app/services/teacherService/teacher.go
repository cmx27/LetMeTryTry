package teacherService

import (
	"STU/app/models"
	"STU/config/database"
)

func CreateCourse(course models.TeacherPostCourse) error {
	result := database.DB.Create(&course)
	return result.Error
}

func UpdateCourse(course models.TeacherPostCourse) error {
	// 与课上不同，不多接收 owner_id 参数，选择使用 omit 忽略该字段的更新
	result := database.DB.Omit("").Save(&course)
	return result.Error
}

func DeleteCourse(class_id uint) error {
	result := database.DB.Where("id = ?", class_id).Delete(&models.TeacherPostCourse{})
	return result.Error
}

func GetCourseList(user_id uint) ([]models.TeacherPostCourse, error) {
	result := database.DB.Where("user_id = ?", user_id).First(&models.TeacherPostCourse{})
	if result.Error != nil {
		return nil, result.Error
	}
	var courseList []models.TeacherPostCourse
	result = database.DB.Where("user_id = ?", user_id).Find(&courseList)
	if result.Error != nil {
		return nil, result.Error
	} else {
		return courseList, nil
	}

}

func CheckClass(class_id uint) error {
	result := database.DB.Where("id = ?", class_id).First(&models.TeacherPostCourse{})
	return result.Error
}

func CompareUser(user1 uint, user2 uint) bool {
	return user1 == user2
}

func GetUserByClassID(id uint) (*models.TeacherPostCourse, error) {
	var user models.TeacherPostCourse
	result := database.DB.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
