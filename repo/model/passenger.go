package model

//创建此表时还会创建用户的购物车表以及订单表
type Passenger struct {
	//积分
	User
	Score int64 `gorm:"comment:用户积分" json:"score"`
	//用户的购物车(如果一次只能买一件，可以不用）
	Cart *PassengerCart `gorm:"foreignKey:PassengerID" json:"cart"`
	//订单
	OrderForms *PassengerOrderForm `gorm:"foreignKey:PassengerID" json:"order_forms"`
	//购物券
	//Tickets []*Ticket `gorm:"foreignKey:PassengerID"`
}
