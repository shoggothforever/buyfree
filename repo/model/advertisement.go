package model

import "time"

type Advertisement struct {
	ID                int64
	ExpectedPlayTimes int64
	NowPlayTimes      int64
	//投资金额
	InvestFund int64
	video      string
	AdOwner    string
	PlayUrl    string
	ExpireAt   time.Time
}
