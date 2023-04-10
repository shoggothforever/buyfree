package response

type ScanResponse struct {
	Response
	DeviceID int64 `json:"device_id"`
}
type QRUrlInfo struct {
	DeviceID int64  `json:"device_id,omitempty"`
	QRUrl    string `json:"qr_url,omitempty"`
}
type QRCodeResponse struct {
	Response
	QRUrlInfos []QRUrlInfo `json:"qr_url_infos"`
}
type DriverAuthInfo struct {
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
