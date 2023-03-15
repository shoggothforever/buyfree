package model

import "github.com/google/uuid"

type Platform struct {
	//登记的司机
	AuthorizedDrivers map[uuid.UUID]*Driver
	Advertisements    []Advertisement
}
