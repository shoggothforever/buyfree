package model

import (
	"time"
)

type Chosen struct {
	Name     string
	Ischosen bool
}

//创建此表时还会创建DeviceProduct
type Device struct {
	ID            int64            `gorm:"primaryKey" json:"id"`
	OwnerID       int64            `gorm:"comment:车主ID" json:"owner_id"`
	PlatformID    int64            `json:"platform_id"`
	Products      []*DeviceProduct `gorm:"foreignKey:DeviceID;comment:供货情况"`
	IsActivated   bool             `gorm:"comment:1为激活，0为未激活" json:"is_activated"`
	ActivatedTime time.Time        `gorm:"comment:激活时间;autocreatetime"`
	UpdatedTime   time.Time        `gorm:"comment:更新时间" json:"updated_time"`
	IsOnline      bool             `gorm:"comment:1为上线，0为未上线" json:"is_online"`
	Profit        float64          `gorm:"comment:收益额" json:"profit"`
}
