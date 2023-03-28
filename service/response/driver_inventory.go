package response

import "buyfree/repo/model"

type InventoryResponse struct {
	Response
	Products []model.DeviceProduct
}
