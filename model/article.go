package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title    string    `gorm:"size:255;not null" json:"title"`
	Content  string    `gorm:"type:text;not null" json:"content"`
	UserID   uint      `gorm:"not null" json:"user_id"`
	Media    []Media   `json:"media"`
	Comments []Comment `json:"comments" gorm:"foreignKey:ArticleID;constraint:OnDelete:CASCADE;"`
}

type Media struct {
	gorm.Model
	ArticleID uint   `gorm:"index;not null" json:"article_id"` // 外键
	Type      string `json:"type"`
	URL       string `json:"url"`
	ObjectKey string `json:"objectKey"`
	Size      int64  `json:"size"`
	Duration  int    `json:"duration"`
}

type Comment struct {
	gorm.Model
	ArticleID uint      `gorm:"index;not null" json:"article_id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	Content   string    `gorm:"type:text;not null" json:"content"`
	ParentID  *uint     `json:"parent_id"`        // nil 表示顶级评论
	Replies   []Comment `json:"replies" gorm:"-"` // 在 service 层组装，不持久化到 DB
}
