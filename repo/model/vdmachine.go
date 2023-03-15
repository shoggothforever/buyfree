package model

import "github.com/google/uuid"

type Chosen struct {
	Name     string
	Ischosen bool
}

//创建此表时还会创建VDProduct
type VdMachine struct {
	OwnerID  uuid.UUID
	Products []*Product `gorm:"foreignKey:VDID"`

	IsChosen       map[string]int64
	Profit         float64 `gorm:"default:0.0"`
	Advertisements []*Advertisement
}
