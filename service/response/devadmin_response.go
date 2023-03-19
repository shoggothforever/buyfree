package response

type DevQueryInfo struct {
	Seq          int64
	DevID        int64
	DriverName   string
	Mobile       string
	SaleForToday float64
	Location     string
	State        string
}

type DevResponse struct {
	Response
	DevResponses []*DevQueryInfo
}
