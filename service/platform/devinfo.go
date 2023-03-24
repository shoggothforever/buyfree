package platform

import (
	"buyfree/dal"
	"buyfree/repo/gen"
	"buyfree/repo/model"
	"buyfree/service/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

type DevinfoController struct {
	BaseController
}

//TODO: 利用redis对车主端的数据进行统计，
func (d *DevinfoController) AnaSales(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}

//TODO:swagger
// @Summary 展示设备详情信息
// @Description	输入设备的ID以查看对应设备的销量,绑定车主以及库存的信息
// @Tags	Platform/Device
// @Accept json
// @Accept mpfd
// @Produce json
// @Param id path int true "use to identify Device"
// @Success 200 {object} response.DevInfoResponse
// @Failure 400 {object} response.Response
// @Router /pt/dev-admin/infos/{id} [get]
func (d *DevinfoController) LsInfo(c *gin.Context) {
	//TODO:分析数据的服务
	var err error
	dev := model.Device{}
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	dev, err = gen.Device.GetByID(id)
	if err != nil {
		c.JSON(200, response.Response{
			400,
			"查找对应设备信息失败,请检查输入ID是否正确",
		})
		return
	}
	driver, err := gen.Driver.GetByID(dev.OwnerID)
	if err != nil {
		c.JSON(200, response.Response{
			400,
			"查找对应设备绑定的车主信息失败,请检查设备是否合法",
		})
		return
	}
	//products, err := gen.DeviceProduct.GetAllDeviceProduct(dev.ID)

	var prinfos []response.DevProductPartInfo
	dal.Getdb().Raw("SELECT * FROM device_products where device_id="+
		"(SELECT id from devices where id=?)", dev.ID).Find(&prinfos)
	if err != nil {
		c.JSON(200, response.Response{
			400,
			"查找对应设备的商品信息失败",
		})
		return
	}
	//n := len(products)
	//prinfos := make([]response.DevProductPartInfo, n, n)
	//for i := 0; i < n; i++ {
	//	prinfos[i].Sku = products[i].Sku
	//	prinfos[i].Name = products[i].Name
	//	prinfos[i].Pic = products[i].Pic
	//	prinfos[i].SupplyPrice = products[i].SupplyPrice
	//	prinfos[i].MonthlySales = products[i].MonthlySales
	//	prinfos[i].Inventory = products[i].Inventory
	//}
	if err == nil {
		c.JSON(200, response.DevInfoResponse{
			response.Response{
				200,
				"ok"},
			//TODO:添加数据分析的响应信息
			model.SalesData{
				0, 0, 0, 0, 0,
			}, response.DevInfo{
				dev.ID,
				dev.ActivatedTime,
				dev.UpdatedTime,
				driver.Location,
				driver.Name,
				driver.Mobile,
				prinfos,
			},
		})
	} else {
		c.JSON(200, response.Response{
			400,
			"加载页面失败",
		})
	}
}

//TODO:swagger
func (d *DevinfoController) TakeDown(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}
