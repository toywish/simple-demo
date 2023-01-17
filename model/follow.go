package model

import "gorm.io/gorm"

// Follow 关注
type Follow struct {
	gorm.Model
	// 粉丝 Follower 关注 Follow
	FollowerID int64
	// 被关注者
	FollowID int64
}
