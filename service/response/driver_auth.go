package response

type ScanResponse struct {
	Response
	DeviceID int64 `json:"device_id"`
}
type DriverAuthInfo struct {
	DriverID int64  `json:"driver_id"`
	DeviceID int64  `json:"device_id"`
	Name     string `json:"name,omitempty"`
	Mobile   string `json:"mobile,omitempty"`
	//IDCard   string `json:"id_card,omitempty"`
	//CarID    string `json:"car_id,omitempty"`
}

type BindDeviceResponse struct {
	Response
	Info *DriverAuthInfo `json:"info"`
}
