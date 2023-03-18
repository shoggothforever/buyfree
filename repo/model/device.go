package model

import (
	"github.com/google/uuid"
	"time"
)

type Chosen struct {
	Name     string
	Ischosen bool
}

//创建此表时还会创建DeviceProduct
type DEVICE struct {
	ID            uuid.UUID
	OwnerID       uuid.UUID `gorm:"comment:车主ID"`
	PlatformID    uuid.UUID
	Products      []*DeviceProduct `gorm:"foreignKey:DeviceID;comment:供货情况"`
	IsActivated   bool             `gorm:"comment:1为激活，0为未激活"`
	ActivatedTime time.Time        `gorm:"comment:激活时间"`
	UpdatedTime   time.Time        `gorm:"comment:更新时间"`
	IsOnline      bool             `gorm:"comment:1为上线，0为未上线"`
	Profit        float64          `gorm:"comment:收益额"`
}
