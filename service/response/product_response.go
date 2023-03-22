package response

import "buyfree/repo/model"

type FactoryGoodsResponse struct {
	Response
	model.FactoryProduct
}
