package teacherControllers

import (
	"STU/app/models"
	"STU/app/services/studentService"
	"STU/app/services/teacherService"
	"STU/app/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CreateCourseData struct {
	UserID      uint   `json:"user_id" binding:"required"`
	ClassName   string `json:"class_name" binding:"required"`
	Time        uint   `json:"time" binding:"required"`
	Weekday     uint   `json:"weekday" binding:"required"`
	Type        uint   `json:"type" binding:"required"`
	Number      uint   `json:"number" binding:"required"`
	TeacherName string `json:"teacher_name"`
}

// 添加
func CreateCourse(c *gin.Context) {
	var data CreateCourseData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}

	//判断time
	if data.Time < 1 || data.Time > 4 {
		utils.JsonErrorResponse(c, 200504, "time错误")
		return
	}

	//判断weekday
	if data.Weekday < 1 || data.Weekday > 7 {
		utils.JsonErrorResponse(c, 200504, "Weekday错误")
		return
	}

	//判断type
	if data.Type != 1 && data.Type != 2 && data.Type != 3 {
		utils.JsonErrorResponse(c, 200504, "type错误")
		return
	}

	//加入教师名称
	var user *models.User
	user, err = studentService.GetUserByUserID(data.UserID)
	if err != nil {
		utils.JsonErrorResponse(c, 200502, "参数错误")
		return
	}

	err = teacherService.CreateCourse(models.TeacherPostCourse{
		UserID:      data.UserID,
		ClassName:   data.ClassName,
		Time:        data.Time,
		Weekday:     data.Weekday,
		Type:        data.Type,
		Number:      data.Number,
		TeacherName: user.Name,
	})
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	utils.JsonSuccessResponse(c, nil)

}

type UpdateCourseData struct {
	ID          uint   `json:"class_id" binding:"required"`
	UserID      uint   `json:"user_id" binding:"required"`
	ClassName   string `json:"class_name,omitempty"`
	Time        uint   `json:"time,omitempty"`
	Weekday     uint   `json:"weekday,omitempty"`
	Type        uint   `json:"type,omitempty"`
	Number      uint   `json:"number,omitempty"`
	Total       uint   `json:"total,omitempty"`
	TeacherName string `json:"teacher_name,omitempty"`
}

// 更新信息
func UpdateCourse(c *gin.Context) {
	var data UpdateCourseData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}

	//判断课程是否存在
	err = teacherService.CheckClass(data.ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.JsonErrorResponse(c, 401, "课程不存在")
		} else {
			utils.JsonInternalServerErrorResponse(c)
		}
		return
	}

	// 获取课程信息
	var user *models.TeacherPostCourse
	user, err = teacherService.GetUserByClassID(data.ID)
	if err != nil {
		utils.JsonErrorResponse(c, 401, "加油啊！写代码的小姐姐！")
		return
	}

	//判断是否是同一个人
	flag := teacherService.CompareUser(data.UserID, user.UserID)
	if !flag {
		utils.JsonErrorResponse(c, 200503, "不是同一个人")
		return
	}

	err = teacherService.UpdateCourse(models.TeacherPostCourse{
		ID:          data.ID,
		UserID:      data.UserID,
		ClassName:   data.ClassName,
		Time:        data.Time,
		Weekday:     data.Weekday,
		Type:        data.Type,
		Number:      data.Number,
		Total:       data.Total,
		TeacherName: user.TeacherName,
	})
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}

type DeleteCourseData struct {
	ID     uint `json:"class_id" binding:"required"`
	UserID uint `json:"user_id" binding:"required"`
}

// 删除
func DeleteCourse(c *gin.Context) {
	var data DeleteCourseData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}

	//判断课程是否存在
	err = teacherService.CheckClass(data.ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.JsonErrorResponse(c, 401, "课程不存在")
		} else {
			utils.JsonInternalServerErrorResponse(c)
		}
		return
	}

	// 获取教师信息
	var user *models.TeacherPostCourse
	user, err = teacherService.GetUserByClassID(data.ID)
	if err != nil {
		utils.JsonErrorResponse(c, 401, "加油啊！写代码的小姐姐！")
		return
	}

	//判断是否是同一个人
	flag := teacherService.CompareUser(data.UserID, user.UserID)
	if !flag {
		utils.JsonErrorResponse(c, 200503, "不是同一个人")
		return
	}

	err = teacherService.DeleteCourse(data.ID)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}

type GetCourseData struct {
	UserID uint `form:"user_id" binding:"required"`
}

// 获取列表
func GetCourse(c *gin.Context) {
	var data GetCourseData
	err := c.ShouldBindQuery(&data)

	var courseList []models.TeacherPostCourse
	courseList, err = teacherService.GetCourseList(data.UserID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.JsonErrorResponse(c, 200506, "课表为空")
			return
		} else {
			utils.JsonInternalServerErrorResponse(c)
			return
		}
	}
	utils.JsonSuccessResponse(c, gin.H{
		"class_list": courseList,
	})
}
