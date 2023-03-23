package utils

import (
	"buyfree/repo/model"
	"fmt"
	"sync"
	"time"
)

const MAXBUFFER int = 8192

//热销榜，排行榜data struct ZSET  key:data    val: productname  score sale func:zrankbyscore,
//清理过去数据使用 ZREMRANGEBYLEX
//统计七天销售数据 data struct List key:date	val:sales 策略：每天0点 使用lpush操作添加
const DailySalesKey string = "Sales:Daily:"          //val:product
const Constantly7aysSalesKey string = "Sales:7days:" //要求连续 val:
const WeeklySalesKey string = "Sales:Weekly:"        //不要求连续 val:
const MonthSalesKey string = "Sales:Monthly:"        //+month
const AnnuallySalesKey string = "Sales:Annually:"    //+year

type DealWithOrderForm func(o *model.OrderFormRequest)

var sem = make(chan bool, MAXBUFFER)
var TimeOut = time.Duration(500 * time.Millisecond)

/*
设计订单接口，使用装饰模式鉴别是购货订单还是退货订单
*/
//购物通道|补货通道
var Orderchannel chan *model.OrderFormRequest = make(chan *model.OrderFormRequest)

//var Orderchannel chan *model.OrderFormRequest = make(chan *model.OrderFormRequest, MAXBUFFER)

//退款通道
var Refundchannel chan *model.OrderFormRequest = make(chan *model.OrderFormRequest)

//var Refundchannel chan *model.OrderFormRequest = make(chan *model.OrderFormRequest, MAXBUFFER)
var Lock sync.Locker

//获得一天开头的确切时间
func GetBeginningOfTheDay() string {
	y, m, d := time.Now().In(time.Local).Date()
	return fmt.Sprintf("%d-%d-%d 00:00:00:", y, m, d)
}

//获取每周第一天（周一）的日期
func GetFirstDayOfWeek() (int, int, int) {
	now := time.Now().In(time.Local)

	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	y, m, d := time.Now().In(time.Local).AddDate(0, 0, offset).Date()
	return y, int(m), d
}

//获取每月第一天的日期
func GetFirstDayOfMonth() (int, int, int) {
	now := time.Now().In(time.Local)
	y, m, _ := now.Date()
	return y, int(m), 1
}

//获取每年第一天的日期
func GetFirstDayOfYear() (int, int, int) {
	now := time.Now().In(time.Local)
	y, _, _ := now.Date()
	return y, 1, 1
}

//对redis数据库进行操作,考虑退款操作
func HandleOrderForm(p *model.OrderFormRequest) {
	//c := context.TODO()
	//rc := dal.Getrd()
	//y, m, d := GetDateKey()
	p.ReplySign <- true
	<-sem
}

func DetectSaleSign(work DealWithOrderForm, OrderChannel chan *model.OrderFormRequest, ofsquit chan bool) {
	var ticker = time.NewTicker(time.Duration(TimeOut))
	defer ticker.Stop()
	for {
		sem <- true
		select {
		case orderformrequest := <-OrderChannel:
			go work(orderformrequest) //购物
		case <-ofsquit:
			return
		case <-ticker.C:
			ofsquit <- true
		default:

		}
	}
}
func StartServer(work DealWithOrderForm) (service chan *model.OrderFormRequest, quit chan bool) {
	service = make(chan *model.OrderFormRequest)
	quit = make(chan bool)
	go DetectSaleSign(work, service, quit)
	return service, quit
}
func init() {
	//TODO:迁移到main.go中去
	ordqueue, ofsquit := StartServer(HandleOrderForm)
	defer close(ordqueue)
	defer close(ofsquit)
	//ordqueue作为服务端暴露给客户端使用

	//服务结束
	ofsquit <- true
}

//func ListenOrderFromRequest(){
//	var ticker = time.NewTicker(time.Duration(time.Millisecond * 500))
//	defer ticker.Stop()
//	for {
//		select {
//		case product:=<-Orderchannel:{
//		//TODO:处理司机发来的订单信息
//		}
//		case product:=<-Refundchannel:{
//		//TODO:处理司机发来的信息
//		}
//		case <-ticker.C:{
//			fmt.Println("服务超时")
//		}
//		default:
//
//		}
//
//	}
//}
