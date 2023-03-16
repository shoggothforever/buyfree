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
	ID                string
	ExpectedPlayTimes int64
	NowPlayTimes      int64
	//投资金额
	InvestFund float64
	video      string
	AdOwner    string
	PlayUrl    string
	ExpireAt   time.Time
	Line       ADSTATE
}
