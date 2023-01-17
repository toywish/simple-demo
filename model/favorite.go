package model

import "gorm.io/gorm"

// Favorite 点赞
type Favorite struct {
	gorm.Model
	UserID  int64
	VideoID int64
}
