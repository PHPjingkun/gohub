package user

import "gohub/pkg/database"

// IsEmailExist 判断 Email 已被注册
func IsEmailExist(Email string) bool {
	var count int64
	database.DB.Model(User{}).Where("email = ?", Email).Count(&count)
	return count > 0
}

// IsPhoneExist 判断手机号已被注册
func IsPhoneExist(phone string) bool {
	var count int64
	database.DB.Model(User{}).Where("phone = ?", phone).Count(&count)
	return count > 0
}
