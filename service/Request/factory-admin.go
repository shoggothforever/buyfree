package Request

type FactoryRegisterReq struct {
	Name    string `json:"name" binding:"required" `
	Address string `json:"address" binding:"required"`
	//经度
	Longitude string `json:"longitude" binding:"required"`
	//纬度
	Latitude string `json:"latitude" binding:"required"`
	//描述(选填)
	Description string `json:"description,omitempty"`
}
