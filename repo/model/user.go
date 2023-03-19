package model

import (
	"gorm.io/gorm"
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
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli"`
	//注销选项	GORM.
	DeletedAt gorm.DeletedAt
	//账户余额
	Balance float64 `gorm:"comment:账户余额" json:"balance"`
	//用户头像(需要添加修改头像功能）
	Pic string `gorm:"comment:用户头像" json:"pic"`
	//用户昵称
	Name     string `gorm:"notnull;unique;size:32;comment:用户昵称" json:"name"`
	Password string `gorm:"notnull;size:32:comment:用户密码" json:"password"`
	Mobile   string `gorm:"comment:手机号" json:"mobile"`
	IDCard   string `gorm:"comment:身份证" json:"id_card"`
	//用户身份标志符，注册时确认
	Role int `gorm:"notnull;type:int;comment:身份 0-乘客 1-司机 2-场站管理员 3-平台管理员 " json:"role"`
	//用户等级，成长制度待定
	Level LEVEL `gorm:"notnull;type:int;comment:用户等级" json:"level"`
}

type Admin struct {
	User
	//密码盐
	PasswordSalt string `gorm:"comment:年销售量" json:"password_salt"`
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
