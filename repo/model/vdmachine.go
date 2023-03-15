package model

import "github.com/google/uuid"

type Chosen struct {
	Name     string
	Ischosen bool
}

//创建此表时还会创建VDProduct
type VdMachine struct {
	ID       uuid.UUID
	OwnerID  uuid.UUID
	Products []*VDProduct `gorm:"foreignKey:VDID"`

	Profit         float64
	Advertisements []*Advertisement
}
