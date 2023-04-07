package Request

type Code2SessionReqest struct {
	Appid     string `json:"appid,omitempty"`
	Secret    string `json:"secret,omitempty"`
	JsCode    string `json:"js_code,omitempty"`
	GrantType string `json:"grant_type,omitempty" `
}

func NewC2Req(a, s, j string) *Code2SessionReqest {
	return &Code2SessionReqest{a, s, j, "authorization_code"}
}
