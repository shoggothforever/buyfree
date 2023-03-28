package driverapp

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type InventoryController struct {
	BaseDrController
}

// @Summary 车主设备库存
// @Description 展示车主拥有的设备库存数据
// @Tags Driver
// @Accept json
// @Produce json
// @Success 200 {object} response.InventoryResponse
// @Failure 400 {onject} response.Response
// @Router /dr/inventory [get]
func (i *InventoryController) Get(c *gin.Context) {
	greeting, ok := c.Get("hello")
	fmt.Println(greeting, ok)

}
