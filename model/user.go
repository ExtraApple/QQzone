package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"not null" json:"username"`
	Password string `json:"-"`
	Article  []Article
	Role     string `gorm:"size:32;default:user" json:"role"` // user or admin
	Friend   []User `gorm:"many2many:user_friends;joinForeignKey:UserID;joinReferences:FriendID"`
}

type UserFriend struct {
	gorm.Model
	UserID   uint
	FriendID uint
}
