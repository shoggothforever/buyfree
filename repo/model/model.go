package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Model struct {
	//唯一标志符
	Uuid uuid.UUID `gorm:"type:uuid;primaryKey"`
	//注册时间
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli"`
	//注销选项
	DeletedAt gorm.DeletedAt
}
