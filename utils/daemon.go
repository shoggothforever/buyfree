package utils

import (
	"buyfree/repo/model"
	"sync"
	"time"
)

const MAXBUFFER int = 1024

//热销榜，排行榜data struct ZSET  key:data    val: productname  score sale func:zrankbyscore,
//清理过去数据使用 ZREMRANGEBYLEX
//统计七天销售数据 data struct List key:date	val:sales 策略：每天0点 使用lpush操作添加
const DailySalesKey string = "Sales:Daily:"          //val:product
const Constantly7aysSalesKey string = "Sales:7days:" //要求连续 val:
const WeeklySalesKey string = "Sales:Weekly:"        //不要求连续 val:
const MonthSalesKey string = "Sales:Monthly:"        //+month
const AnnuallySalesKey string = "Sales:Annually:"    //+year

//购物通道|补货通道
var Orderchannel chan *model.OrderProductImplement = make(chan *model.OrderProductImplement, MAXBUFFER)

//退款通道
var Refundchannel chan *model.OrderProductImplement = make(chan *model.OrderProductImplement, MAXBUFFER)
var Lock sync.Locker

func GetDateKey() (y int, m time.Month, d, offset int) {
	now := time.Now().In(time.Local)
	y, m, d = now.Date()
	offset = int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	return y, m, d, offset
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
