package utils

import (
	"buyfree/repo/model"
	"sync"
	"time"
)

const MAXBUFFER int = 1024

//购物通道|补货通道
var Orderchannel chan *model.OrderProductImplement = make(chan *model.OrderProductImplement, MAXBUFFER)

//退款通道
var Refundchannel chan *model.OrderProductImplement = make(chan *model.OrderProductImplement, MAXBUFFER)
var Lock sync.Locker

func GetDateKey() (day, month, year int64) {
	now := time.Now().In(time.Local)
	y, m, d := now.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, time.Local).Unix(),
		time.Date(y, m, 0, 0, 0, 0, 0, time.Local).Unix(),
		time.Date(y, 0, 0, 0, 0, 0, 0, time.Local).Unix()
}

//对redis数据库进行操作,考虑退款操作
func AddSale(p *model.OrderProductImplement, mode int) {
	//c := context.TODO()
	//rc := dal.Getrd()
	//y, m, d := GetDateKey()

}

func DetectSailSign() {

	select {
	case product := <-Orderchannel:
		Lock.Lock()
		AddSale(product, 1) //购物
		Lock.Unlock()
	case product := <-Refundchannel:
		Lock.Lock()
		AddSale(product, -1) //退款
		Lock.Unlock()
	}
}
