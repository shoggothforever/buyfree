package model

type Response struct {
	Code int    `form:"code" json:"code"`
	Msg  string `form:"msg" json:"msg"`
}
