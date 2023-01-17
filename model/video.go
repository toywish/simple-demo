package model

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	Title         string
	AuthorID      int64
	PlayUrl       string
	CoverUrl      string
	FavoriteCount int64 `gorm:"DEFAULT:0"`
	CommentCount  int64 `gorm:"DEFAULT:0"`
}
