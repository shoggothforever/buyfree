package model

import (
	"time"
)

type ROLE int

const (
	PASSENGER ROLE = iota
	DRIVER
	FACTORYADMIN
	PLATFORMADMIN
)

type LEVEL int

const (
	ZERO LEVEL = iota
	I
	II
	IIV
	IV
	V
	VI
)

type User struct {
	//唯一标志符
	ID int64 `gorm:"primaryKey;" json:"id" form:"id"`
	//注册时间
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli" json:"updated_at"`
	//注销选项	GORM.
	DeletedAt time.Time `json:"deleted_at"`
	//账户余额
	Balance float64 `gorm:"comment:账户余额" json:"balance"`
	//用户头像(需要添加修改头像功能）
	Pic string `gorm:"comment:用户头像" json:"pic"`
	//用户昵称
	Name     string `gorm:"notnull;unique;size:32;comment:用户昵称" json:"name"`
	Password string `gorm:"size:32:comment:用户密码" json:"password"`
	//密码盐
	PasswordSalt string `gorm:"comment:年销售量" json:"password_salt"`
	Mobile       string `gorm:"comment:手机号" json:"mobile"`
	IDCard       string `gorm:"comment:身份证" json:"id_card"`
	//用户身份标志符，注册时确认
	Role int `gorm:"notnull;type:int;comment:身份 0-乘客 1-司机 2-场站管理员 3-平台管理员 " json:"role"`
	//用户等级，成长制度待定
	Level LEVEL `gorm:"notnull;type:int;comment:用户等级" json:"level"`
}

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

type Platform struct {
	User
	//登记的司机
	AuthorizedDrivers []*Driver        `gorm:"foreignkey:PlatformID" json:"authorized_drivers"`
	Devices           []*Device        `gorm:"foreignkey:PlatformID" json:"devices"`
	Advertisements    []*Advertisement `gorm:"foreignkey:PlatformID" json:"advertisements"`
}

type Possesion struct {
	UserID  int64
	Balance float64
	Tickets []Ticket
}

type LoginInfo struct {
	UserID   int64  `json:"id"`
	Password string `json:"password"`
	Salt     string `gorm:"comment:加密盐" json:"salt"`
	Jwt      string `gorm:"comment:鉴权值" json:"jwt"`
}

type BankCardInfo struct {
	ID       int64   `gorm:"comment:用户ID" json:"id"`
	CardID   int64   `gorm:"unique" json:"card_id"`
	BankName string  `gorm:"银行名称" json:"bankName"`
	Password string  `json:"password"`
	Cash     float64 `gorm:"comment:账户余额" json:"cash"`
	BankFund float64 `gorm:"comment:银行资金" json:"bank_fund"`
}
