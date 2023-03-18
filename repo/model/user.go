package model

import (
	"github.com/google/uuid"
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
	ID uuid.UUID `gorm:"type:uuid;primaryKey;"`
	//注册时间
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli"`
	//注销选项
	DeletedAt gorm.DeletedAt
	//账户余额
	Balance float64 `gorm:"comment:账户余额"`
	//用户头像
	Pic string `gorm:"comment:用户头像"`
	//用户昵称
	Name     string `gorm:"notnull;size:32;comment:用户昵称"`
	Password string `gorm:"notnull;size:32:comment:用户密码"`
	Mobile   string `gorm:"comment:手机号"`
	IDCard   string `gorm:"comment:身份证"`
	//用户身份标志符，注册时确认
	Role int `gorm:"notnull;type:int;comment:身份 0-乘客 1-司机 2-场站管理员 3-平台管理员 "`
	//用户等级，成长制度待定
	Level LEVEL `gorm:"notnull;type:int;comment:用户等级"`
}

type Admin struct {
	User
	//密码盐
	PasswordSalt string `gorm:"comment:年销售量"`
}

type Possesion struct {
	UserID  uuid.UUID
	Balance float64
	Tickets []Ticket
}

type LoginInfo struct {
	UserID   uuid.UUID
	Password string
	Salt     string `gorm:"comment:加密盐"`
	Jwt      string `gorm:"comment:鉴权值"`
}

type Cookies struct {
	UserID    uuid.UUID
	JWT       string
	CreatedAt time.Time
}
