package model

import (
	"time"
)

type ROLE int

const (
	PASSENGER int64 = iota
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
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at" form:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli" json:"updated_at" form:"updated_at"`
	//注销选项	GORM.
	DeletedAt time.Time `json:"deleted_at" form:"deleted_at"`
	//账户余额
	Balance float64 `gorm:"comment:账户余额" json:"balance" form:"balance"`
	//用户头像(需要添加修改头像功能）
	Pic string `gorm:"comment:用户头像" json:"pic" form:"pic"`
	//用户昵称
	Name     string `gorm:"notnull;unique;size:32;comment:用户昵称" json:"name" form:"name"`
	Password string `gorm:"size:32:comment:用户密码" json:"password" form:"password"`
	//密码盐
	PasswordSalt string `gorm:"comment:年销售量" json:"password_salt" form:"password_salt"`
	Mobile       string `gorm:"comment:手机号" json:"mobile" form:"mobile"`
	IDCard       string `gorm:"comment:身份证" json:"id_card" form:"id_card"`
	//用户身份标志符，注册时确认
	Role int `gorm:"notnull;type:int;comment:身份 0-乘客 1-司机 2-场站管理员 3-平台管理员 " json:"role" form:"role"`
	//用户等级，成长制度待定
	Level LEVEL `gorm:"notnull;type:int;comment:用户等级" json:"level" form:"level"`
	//BankCardInfos BankCardInfo `gorm:"foreignkey:id" json:"bank_card_infos" form:"bank_card_infos"`
}

//创建此表时还会创建用户的购物车表以及订单表
type Passenger struct {
	//积分
	User
	Score int64 `gorm:"comment:用户积分" json:"score" form:"score"`
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
	UserID   int64  `json:"id" form:"user_id"`
	ROLE     int64  `gorm:"commenr:0代表乘客，1代表司机，2代表场站，3代表平台" json:"role" form:"role"`
	Password string `json:"password" form:"password"`
	Salt     string `gorm:"comment:加密盐" json:"salt" form:"salt"`
	Jwt      string `gorm:"comment:鉴权值" json:"jwt" form:"jwt"`
}

//TODO:用户银行卡信息
type BankCardInfo struct {
	//外键
	ID       int64   `gorm:"comment:用户ID" json:"id" form:"id"`
	CardID   int64   `gorm:"unique" json:"card_id" form:"card_id"`
	BankName string  `gorm:"银行名称" json:"bankName" form:"bank_name"`
	Password string  `json:"password" form:"password"`
	Cash     float64 `gorm:"comment:账户余额" json:"cash" form:"cash"`
	BankFund float64 `gorm:"comment:银行资金" json:"bank_fund" form:"bank_fund"`
}
