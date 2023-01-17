package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Password string
	Name     string `gorm:"unique;DEFAULT:'未定义'"`
	// 关注
	FollowCount int64 `gorm:"DEFAULT:0"`
	// 粉丝
	FollowerCount int64 `gorm:"DEFAULT:0"`
}
