package response

import "time"

//显示支付成功或者支付失败
type PayResponse struct {
	Response
}

type SubmitReqs []*SubmitReq
type SubmitReq struct {
	FactoryDistanceReq
	Common  string    `json:"common,omitempty"`
	GetTime time.Time `json:"get_time"`
}
type PayReqs []PayReq
type PayReq struct {
	Comment  string `json:"comment,omitempty"`
	Mobile   string `json:"mobile,omitempty"`
	OrderIDs int64  `json:"order_ids,omitempty"`
}
