package response

type WeiXinLoginResponse struct {
	Response
	OpenID  string `json:"openid"`
	UnionID string `json:"unionid"`
	ErrCode int64  `json:"errcode" default:"0"`
	ErrMsg  string `json:"errmsg"`
	//自定义登录态，前端存入storage中,每次发起业务请求携带自定义登录态
	Token string `json:"token"`
}

type WeiXinLoginInfo struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
	ErrCode    int64  `json:"errcode" default:"0"`
	ErrMsg     string `json:"errmsg"`
}
