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
	ID                int64  `gorm:"primarykey" json:"id"`
	Description       string `gorm:"comment:广告描述" json:"description"`
	PlatformID        int64
	ExpectedPlayTimes int64     `gorm:"comment:预期播放次数"  json:"expected_play_times"`
	NowPlayTimes      int64     `gorm:"comment:已经播放金额" json:"now_play_times"`
	InvestFund        float64   `gorm:"comment:投资金额" json:"invest_fund"`
	Profit            float64   `gorm:"comment:产生收入" json:"profit"`
	VideoCover        string    `gorm:"comment:广告封面地址" json:"video_cover"`
	ADOwner           string    `gorm:"comment:广告商" json:"ad_owner"`
	PlayUrl           string    `gorm:"comment:广告播放地址" json:"play_url"`
	ExpireAt          time.Time `gorm:"comment:截止日期" json:"expire_at"`
	ADState           int64     `gorm:"comment:广告状态 1上线 ， 0下线" json:"ad_state"`
	Devices           []*Device `gorm:"many2many:Ad_Device"`
}
