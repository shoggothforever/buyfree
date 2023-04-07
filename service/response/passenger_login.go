package response

type WeiXinLoginResponse struct {
	Response
	OpenID  string `json:"openid"`
	UnionID string `json:"unionid"`
	ErrCode int64  `json:"errcode" default:"0"`
	ErrMsg  string `json:"errmsg"`
}

type WeiXinLoginInfo struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
	ErrCode    int64  `json:"errcode" default:"0"`
	ErrMsg     string `json:"errmsg"`
}
