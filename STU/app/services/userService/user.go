package userService

import (
	"STU/app/models"
	"STU/app/utils"
	"STU/config/database"
)

// func CheckUser(Password string) (*models.User, error) {
// 	pass := utils.Encryrpt(Password)
// 	println(pass)
// 	result := database.DB.Where("password = ?", Password).First(&models.User{})
// 	return nil, result.Error
// }

func CheckUserExistByAccount(account string) error {
	result := database.DB.Where("account = ?", account).First(&models.User{})
	return result.Error
}

func GetUserByAccount(account string) (*models.User, error) {
	var user models.User
	result := database.DB.Where("account = ?", account).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func ComparePwd(pwd1 string, pwd2 string) bool {
	return pwd1 == pwd2
}

func Register(user models.User) error {
	result := database.DB.Create(&user)
	return result.Error
}

func UpdateUserPasswordByAccount(account, password string) error {
	user, _ := GetUserByAccount(account)
	pass := utils.Encryrpt(password)
	user.Password = pass
	err := database.DB.Model(models.User{}).Where(
		models.User{
			Account: user.Account,
		}).Updates(user).Error
	if err != nil {
		return err
	} else {
		return nil
	}
}

func CheckUserBYAccountAndPassword(Account, password string) bool {
	pass := utils.Encryrpt(password)
	println(pass)
	user := models.User{}
	result := database.DB.Where(
		models.User{
			Account:  Account,
			Password: pass,
		}).First(&user)
	return result.Error == nil
}
