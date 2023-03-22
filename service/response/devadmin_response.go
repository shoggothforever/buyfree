package response

import (
	"buyfree/repo/model"
	"time"
)

type DevQueryInfo struct {
	Seq          int64   `json:"seq"`
	DevID        int64   `gorm:"id" json:"dev_id"`
	DriverName   string  `gorm:"name" json:"driver_name"`
	Mobile       string  `gorm:"mobile" json:"mobile"`
	SalesOfToday float64 `json:"sales_of_today"`
	Location     string  `json:"location"`
	State        string  `json:"state"`
}
type DevProductPartInfo struct {
	FactoryName  string  `json:"factory_name"`
	Sku          string  `json:"sku"`
	Name         string  `json:"name"`
	Pic          string  `json:"pic"`
	Price        float64 `json:"price"`
	MonthlySales float64 `json:"monthly_sales"`
	Inventory    int64   `json:"inventory"`
}
type DevInfo struct {
	DevID         int64     `json:"devID"`
	ActivatedTime time.Time `json:"activated_time"`
	UpdatedTime   time.Time `json:"updated_time"`
	Location      string    `json:"location"`
	DriverName    string    `json:"driver_name"`
	Mobile        string    `json:"mobile"`
	ProductInfo   []DevProductPartInfo
}
type DevResponse struct {
	Response
	DevResponses []DevQueryInfo
}

type AddDevResponse struct {
	Response
	Devices *model.Device
}

type DevInfoResponse struct {
	Response
	model.SalesData
	DevInfo
}
