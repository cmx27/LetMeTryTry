package studentControllers

import (
	//"STU/app/models"
	"STU/app/models"
	"STU/app/services/studentService"

	"STU/app/services/teacherService"
	"STU/app/utils"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 选课
type ChoiceClassData struct {
	UserID  uint `json:"user_id" binding:"required"`
	ClassID uint `json:"class_id" binding:"required"`
}

// 添加
func ChoiceClass(c *gin.Context) {
	var data ChoiceClassData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}

	//判断学生身份
	err = studentService.CheckStudent(data.UserID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.JsonErrorResponse(c, 401, "你谁啊你")
		} else {
			utils.JsonInternalServerErrorResponse(c)
		}
		return
	}

	//获取相应课程
	var class *models.TeacherPostCourse
	class, err = studentService.GetClassById(data.ClassID)
	if err != nil {
		utils.JsonErrorResponse(c, 200502, "参数错误")
		return
	}

	//判断是否重复选课

	// err = teacherService.CheckClass(data.ClassID)
	// if err != nil {
	// 	if err == gorm.ErrRecordNotFound {
	// 		utils.JsonErrorResponse(c, 401, "课程已经选过")
	// 	} else {
	// 		utils.JsonInternalServerErrorResponse(c)
	// 	}
	// 	return
	// }

	//totol+1
	if class.Total <= class.Number {
		class.Total++
	} else {
		utils.JsonErrorResponse(c, 200501, "不能选课")
		return
	}

	//加入时间
	now := time.Now()      //获取当前时间
	year := now.Year()     //年
	month := now.Month()   //月
	day := now.Day()       //日
	hour := now.Hour()     //小时
	minute := now.Minute() //分钟
	second := now.Second() //秒
	t := "%d-%02d-%02d %02d:%02d:%02d"
	var target_url = fmt.Sprintf(t, year, month, day, hour, minute, second)

	err = studentService.ChoiceClass(models.Choice{
		UserID:      data.UserID,
		ClassID:     data.ClassID,
		ClassName:   class.ClassName,
		Time:        class.Time,
		Weekday:     class.Weekday,
		Type:        class.Type,
		TeacherName: class.TeacherName,
		Date:        target_url,
	})
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	utils.JsonSuccessResponse(c, nil)

}

// 获取课表
func GetList(c *gin.Context) {

	var courseList []models.TeacherPostCourse
	courseList, _ = studentService.GetCourseList()
	utils.JsonSuccessResponse(c, gin.H{
		"class_list": courseList,
	})
}

type GetTakenListData struct {
	UserID uint `form:"user_id" binding:"required"`
}

// 获取已选列表
func GetTakenList(c *gin.Context) {
	var data GetTakenListData
	err := c.ShouldBindQuery(&data)
	var courseList []models.Choice
	courseList, err = studentService.GetTakenCourseList(data.UserID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.JsonErrorResponse(c, 200506, "选课为空")
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

type DeleteClassData struct {
	UserID  uint `json:"user_id" binding:"required"`
	ClassID uint `json:"class_id" binding:"required"`
}

// 删除课程
// 删除
func DeleteClass(c *gin.Context) {
	var data DeleteClassData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}

	//判断课程是否存在
	err = studentService.CheckClass(data.ClassID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.JsonErrorResponse(c, 401, "课程不存在")
		} else {
			utils.JsonInternalServerErrorResponse(c)
		}
		return
	}

	// 获取学生信息
	var user *models.Choice
	user, err = studentService.GetUserByClassID(data.ClassID)
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

	err = studentService.DeleteCourse(data.ClassID)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}
