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
	FACTORY
	PLATFORM
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
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`
	//注册时间
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli"`
	//注销选项
	DeletedAt gorm.DeletedAt
	//用户头像
	Pic string `gorm:""`
	//用户昵称
	Name     string `gorm:"notnull;uniqueindex;size:32;"`
	Password string `gorm:"notnull;size:32"`
	//用户身份标志符，注册时确认
	Role int `gorm:"notnull;type:int"`
	//用户等级，成长制度待定
	Level LEVEL `gorm:"notnull;type:int"`
}

//type Possesion struct {
//	UserName int64
//	Tickets  []Ticket
//}
