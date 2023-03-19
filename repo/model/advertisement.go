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
	ID                int64  `gorm:"primarykey"`
	Description       string `gorm:"comment:广告描述"`
	PlatformID        int64
	ExpectedPlayTimes int64     `gorm:"comment:预期播放次数"`
	NowPlayTimes      int64     `gorm:"comment:已经播放金额"`
	InvestFund        float64   `gorm:"comment:投资金额"`
	Profie            float64   `gorm:"comment:产生收入"`
	videoCover        string    `gorm:"comment:广告封面地址"`
	ADOwner           string    `gorm:"comment:广告商"`
	PlayUrl           string    `gorm:"comment:广告播放地址"`
	ExpireAt          time.Time `gorm:"comment:截止日期"`
	ADState           ADSTATE   `gorm:"comment:广告状态 1上线 ， 0下线"`
}
