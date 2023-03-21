package response

type OrderInfostruct struct {
	FactoryName string `json:"factory_name"`
	Sku         string `json:"sku"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	//State       int     `json:"state"`
	Pic       string  `json:"pic"`
	Sales     float64 `json:"sales"`
	Inventory int64   `json:"inventory"`
}

type OrderResponse struct {
	Response
	OrderInfostructs []OrderInfostruct
}
