package service

import (
	"QQZone/global"
	"QQZone/model"
	"errors"

	"gorm.io/gorm"
)

// CreateComment: userID 从中间件获取，parentID 可为 nil（顶级评论）
func CreateComment(userID uint, articleID uint, content string, parentID *uint) (*model.Comment, error) {
	// 1. 检查文章是否存在
	var art model.Article
	if err := global.DB.First(&art, articleID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("article not found")
		}
		return nil, err
	}

	// 2. 如果 parentID 不为空，该评论为自评论，校验父评论存在且属于同一篇文章
	if parentID != nil {
		var parent model.Comment
		if err := global.DB.First(&parent, *parentID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.New("parent comment not found")
			}
			return nil, err
		}
		if parent.ArticleID != articleID {
			return nil, errors.New("parent comment does not belong to this article")
		}
	}

	// 3. 创建评论
	comment := model.Comment{
		ArticleID: articleID,
		UserID:    userID,
		Content:   content,
		ParentID:  parentID,
	}
	if err := global.DB.Create(&comment).Error; err != nil {
		return nil, err
	}

	return &comment, nil
}

// buildCommentTree: 从扁平列表构建父子树
func buildCommentTree(comments []model.Comment) []model.Comment {
	commentMap := make(map[uint]*model.Comment)
	var roots []model.Comment

	// 建立映射（注意用 &comments[i] 安全取地址）
	for i := range comments {
		commentMap[comments[i].ID] = &comments[i]
	}
	for i := range comments {
		c := comments[i]
		if c.ParentID != nil {
			if parent, ok := commentMap[*c.ParentID]; ok {
				parent.Replies = append(parent.Replies, c)
			} else {
				// 父评论找不到，视为顶级
				roots = append(roots, c)
			}
		} else {
			roots = append(roots, c)
		}
	}
	return roots
}

func GetCommentsByArticle(articleID uint) ([]model.Comment, error) {
	var comments []model.Comment
	if err := global.DB.
		Where("article_id = ?", articleID).
		Order("created_at asc").
		Find(&comments).Error; err != nil {
		return nil, err
	}

	return buildCommentTree(comments), nil
}

// DeleteComment: 删除评论及其所有子孙（执行者必须是评论作者）
func DeleteComment(commentID uint, userID uint) error {
	var c model.Comment
	if err := global.DB.First(&c, commentID).Error; err != nil {
		return err
	}
	// 权限判断，仅允许评论作者删除
	if c.UserID != userID {
		return errors.New("no permisson to delete this comment")
	}
	// 广度优先收集所有子孙ID
	ids := []uint{commentID}
	for i := 0; i < len(ids); i++ {
		var children []model.Comment
		if err := global.DB.Where("parent_id = ?", ids[i]).Find(&children).Error; err != nil {
			return err
		}
		for _, ch := range children {
			ids = append(ids, ch.ID)
		}
	}
	// 批量删除
	if err := global.DB.Delete(&model.Comment{}, ids).Error; err != nil {
		return err
	}
	return nil
}
