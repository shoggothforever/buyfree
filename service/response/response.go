package response

type PingResponse struct {
	Msg string `json:"msg" form:"msg"`
}
type Response struct {
	Code int64  `json:"code" form:"code"`
	Msg  string `json:"msg" form:"msg"`
}

type LoginResponse struct {
	Response
	UserID int64  `json:"user_id"`
	Jwt    string `json:"jwt"`
}
