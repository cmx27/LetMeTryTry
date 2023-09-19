package userControllers

import (
	"STU/app/models"
	"STU/app/services/userService"
	"STU/app/utils"

	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RegisterData struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
	Type     uint   `json:"type" binding:"required"`
	Name     string `json:"name" binding:"required"`
}

// 注册
func Register(c *gin.Context) {
	var data RegisterData
	err := c.ShouldBindJSON(&data)

	//判断type
	if data.Type != 1 && data.Type != 2 {
		utils.JsonErrorResponse(c, 200504, "type错误")
		return
	}

	// 判断号是否已经注册
	err = userService.CheckUserExistByAccount(data.Account)
	if err == nil {
		utils.JsonErrorResponse(c, 200504, "账号已注册")
		return
	} else if err != nil && err != gorm.ErrRecordNotFound {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	//限制密码长度
	// if len(data.Password) < 8 || len(data.Password) > 20 {
	// 	utils.JsonResponse(c, 200, 401, "密码长度必须在8~20位之间", nil)
	// 	return
	// }
	// 注册用户
	err = userService.Register(models.User{
		Account:  data.Account,
		Password: data.Password,
		Type:     data.Type,
		Name:     data.Name,
	})
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	//加密
	err = userService.UpdateUserPasswordByAccount(data.Account, data.Password)
	if err != nil {
		log.Println(err)
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}
