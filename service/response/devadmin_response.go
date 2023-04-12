package response

import (
	"buyfree/repo/model"
	"time"
)

// 设备以及设备拥有者的部分信息
type DevQueryInfo struct {
	//顺序编号
	Seq        int64  `json:"seq"`
	DevID      int64  `gorm:"id" json:"dev_id"`
	DriverName string `gorm:"name" json:"driver_name"`
	//车主电话号码
	Mobile string `gorm:"mobile" json:"mobile"`
	//今日营销额
	SalesOfToday float64 `json:"sales_of_today"`
	//地址信息
	Location string `json:"location"`
	//设备状态信息，state=1，2,3,4分别对应获取在线，离线,激活，未激活
	State string `json:"state"`
}

// 设备商品部分信息
type DevProductPartInfo struct {
	Sku  string `json:"sku"`
	Name string `json:"name"`
	//商品封面图片
	Pic string `json:"pic"`
	//批发价
	SupplyPrice  float64 `json:"supply_price"`
	MonthlySales float64 `json:"monthly_sales"`
	Inventory    int64   `json:"inventory"`
}

// 设备信息
type DevInfo struct {
	DevID int64 `json:"devID"`
	//设备激活时间
	ActivatedTime time.Time `json:"activated_time"`
	//设备更新时间
	UpdatedTime time.Time `json:"updated_time"`
	Location    string    `json:"location"`
	DriverName  string    `json:"driver_name"`
	Mobile      string    `json:"mobile"`
	ProductInfo []DevProductPartInfo
}
type DevResponse struct {
	Response
	DevQueryInfos []DevQueryInfo
}

type AddDevResponse struct {
	Response `json:"response"`
	DeviceQR string        `json:"device_QR,omitempty"`
	Devices  *model.Device `json:"devices,omitempty"`
}

type DevInfoResponse struct {
	Response
	model.SalesData
	DevInfo
}
