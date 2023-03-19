package response

import (
	"buyfree/repo/model"
	"time"
)

type DevQueryInfo struct {
	Seq          int64
	DevID        int64
	DriverName   string
	Mobile       string
	SaleForToday float64
	Location     string
	State        string
}
type DevProductPartInfo struct {
	Sku         string
	Name        string
	Pic         string
	Prize       int64
	MonthlySold int64
	Inventory   int64
	//上架？
}
type DevInfo struct {
	DevID         int64
	ActivatedTime time.Time
	UpdatedTime   time.Time
	Location      string
	DriverName    string
	Mobile        string
	ProductInfo   []*DevProductPartInfo
}
type DevResponse struct {
	Response
	DevResponses []*DevQueryInfo
}

type AddDevResponse struct {
	Response
	Devices *model.Device
}

type DevInfoResponse struct {
	model.SalesData
	DevInfo
}
