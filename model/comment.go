package model

import "gorm.io/gorm"

// Comment 评论
type Comment struct {
	gorm.Model
	UserID  int64
	VideoID int64
	Content string
}
