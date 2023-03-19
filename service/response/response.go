package response

type PingResponse struct {
	Msg string `json:"msg" form:"msg"`
}
type Response struct {
	Code int64  `json:"code" form:"code"`
	Msg  string `json:"msg" form:"msg"`
}
