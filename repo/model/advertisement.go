package model

import (
	"time"
)

type ADSTATE int

const (
	ONLINE ADSTATE = iota
	OFFLINE
)

type Advertisement struct {
	ID                int64     `gorm:"primarykey" json:"id" form:"id"`
	Description       string    `gorm:"comment:广告描述" json:"description" form:"description"`
	PlatformID        int64     `json:"platform_id"`
	ExpectedPlayTimes int64     `gorm:"comment:预期播放次数"  json:"expected_play_times" form:"expected_play_times"`
	PlayTimes         int64     `gorm:"comment:已经播放金额" json:"play_times" form:"play_times"`
	InvestFund        float64   `gorm:"comment:投资金额" json:"invest_fund" form:"invest_fund"`
	Profit            float64   `gorm:"comment:产生收入" json:"profit" form:"profit"`
	VideoCover        string    `gorm:"comment:广告封面地址" json:"video_cover" form:"video_cover"`
	ADOwner           string    `gorm:"comment:广告商" json:"ad_owner" form:"ad_owner"`
	PlayUrl           string    `gorm:"comment:广告播放地址" json:"play_url" form:"play_url"`
	ExpireAt          time.Time `gorm:"comment:截止日期" json:"expire_at" form:"expire_at"`
	ADState           int64     `gorm:"comment:广告状态 t上线 ， f下线" json:"ad_state" form:"ad_state"`
	//在投放广告的时候需要注意
	Devices []*Device `gorm:"many2many:Ad_Device"`
}
type ADurl struct {
	ID         int64  ` json:"id" form:"id"`
	VideoCover string `json:"video_cover" form:"video_cover"`
	PlayUrl    string `json:"play_url" form:"play_url"`
}
type Ad_Device struct {
	AdvertisementID int64   `gorm:"primarykey;column:advertisement_id" json:"advertisement_id" form:"advertisement_id"`
	DeviceID        int64   `gorm:"primarykey;column:device_id" json:"device_id" form:"device_id"`
	PlayTimes       int64   `json:"play_times;column:play_times" form:"play_times"`
	Profit          float64 `json:"profit;column:profit" form:"profit"`
}
type ID int64
type IDList struct {
	IDS []ID `json:"IDS,omitempty"`
}
