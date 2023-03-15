package model

//创建此表时还会创建用户的购物车表以及订单表
type Passenger struct {
	//积分

	User
	Score int
	//用户的购物车(如果一次只能买一件，可以不用）
	Cart PassengerCart `gorm:"foreignkey:PassengerID"`
	//订单
	OrderForms PassengerOrderForm `gorm:"foreignkey:UserKey"`
	//购物券
}
