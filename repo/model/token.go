package model

import "github.com/google/uuid"

// UserToken 用户令牌儿表 /*
type UserToken struct {
	Token  string    `gorm:"not null"`
	UserID uuid.UUID `gorm:"primaryKey"`
	Role   string    `gorm:"not null;default:0"` // 用户角色
}

func (u *UserToken) Auth() bool {
	//TODO AdminEnd and DriverEnd Login
	return true
}
