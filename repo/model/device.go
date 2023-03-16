package model

import "github.com/google/uuid"

type Chosen struct {
	Name     string
	Ischosen bool
}

//创建此表时还会创建DeviceProduct
type DEVICE struct {
	ID      uuid.UUID
	OwnerID uuid.UUID
	roducts []*DeviceProduct `gorm:"foreignKey:DeviceID"`

	Profit         float64
	Advertisements []*Advertisement
}
