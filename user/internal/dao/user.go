package dao

import (
	"errors"
	"gorm.io/gorm"
	"user/internal/service"
	utils "user/pkg/util"
)

type User struct {
	UserID         uint   `gorm:"primarykey"`
	UserName       string `gorm:"unique"`
	NickName       string
	PasswordDigest string
}

const PasswordConst = 12

func CheckUserExists(username string) (user *User, err error) {
	err = DB.Model(&User{}).Where("user_name=?", username).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("User not found")
		}
		return nil, err
	}
	return user, nil
}

//序列化User对象
func BuildUser(user *User) *service.UserModel {
	userId := uint32(user.UserID)
	userModel := &service.UserModel{
		UserID:   userId,
		UserName: user.UserName,
		NickName: user.NickName,
	}
	return userModel
}

func UserCreate(req *service.UserRequest) (*User, error) {
	var count int64
	err := DB.Model(&User{}).Where("user_name=?", req.UserName).Count(&count).Error
	if count != 0 {
		return nil, errors.New("UserName exist")
	}
	user := &User{
		UserName: req.UserName,
		NickName: req.NickName,
	}
	user.PasswordDigest = utils.GenerateHashFromPassword(req.Password)
	if err = DB.Model(&User{}).Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
