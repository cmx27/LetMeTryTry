package userControllers

import (
	"STU/app/models"
	"STU/app/services/userService"
	"STU/app/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginData struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// 登录
func Login(c *gin.Context) {
	var data LoginData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}

	// 判断用户是否存在
	err = userService.CheckUserExistByAccount(data.Account)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.JsonErrorResponse(c, 401, "账号不存在")
		} else {
			utils.JsonInternalServerErrorResponse(c)
		}
		return
	}

	// 获取用户信息
	var user *models.User
	user, err = userService.GetUserByAccount(data.Account)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	// 判断密码是否正确
	// flag := userService.ComparePwd(data.Password, user.Password)
	// if !flag {
	// 	utils.JsonErrorResponse(c, 200503, "密码错误")
	// 	return
	// }

	flag := userService.CheckUserBYAccountAndPassword(data.Account, data.Password)
	if !flag {
		utils.JsonResponse(c, 405, 400, "密码错误", nil)
		return
	}

	// 返回用户信息//只返回type,id
	utils.JsonSuccessResponse(c, user)
}
