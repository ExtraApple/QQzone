package service

import (
	"QQZone/global"
	"QQZone/model"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(username, password string) (*model.User, error) {
	//检查用户是否存在
	var exists model.User
	if err := global.DB.Where("username = ?", username).First(&exists).Error; err == nil {
		return nil, errors.New("user is exits")
	}
	//密码哈希
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := model.User{Username: username, Password: string(hash), Role: "user"}
	if err := global.DB.Create(&user).Error; err != nil {
		return nil, err
	}
	user.Password = ""
	return &user, nil
}

func Authenticate(username, password string) (*model.User, error) {
	var user model.User
	if err := global.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, errors.New("username not found")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("password is wrong")
	}
	user.Password = ""
	return &user, nil
}
