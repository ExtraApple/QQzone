package service

import (
	"QQZone/global"
	"QQZone/model"
	"errors"

	"gorm.io/gorm"
)

// 添加好友
func AddFriend(UserID, FriendID uint) error {
	// 检查是否已经是好友
	var existing model.UserFriend
	if err := global.DB.Where("user_id = ? AND friend_id = ?", UserID, FriendID).First(&existing).Error; err == nil {
		return errors.New("friend already exists")
	}

	// 创建好友关系
	friend := model.UserFriend{
		UserID:   UserID,
		FriendID: FriendID,
	}
	if err := global.DB.Create(&friend).Error; err != nil {
		return err
	}
	return nil
}

// 删除好友
func DeleteFriend(UserID, FriendID uint) error {
	// 查找是否存在好友关系
	var friend model.UserFriend
	if err := global.DB.Where("user_id = ? AND friend_id = ?", UserID, FriendID).First(&friend).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("friend not found")
		}
		return err
	}

	// 删除好友关系
	if err := global.DB.Delete(&friend).Error; err != nil {
		return err
	}
	return nil
}

// 获取好友列表
func ListFriend(UserID uint) ([]model.User, error) {
	var users []model.User
	err := global.DB.
		Table("users").
		Joins("JOIN user_friends ON user_friends.friend_id = users.id").
		Where("user_friends.user_id = ?", UserID).
		Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
