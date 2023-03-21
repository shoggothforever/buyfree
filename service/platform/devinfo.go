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

func (d *DevinfoController) LsInfo(c *gin.Context) {
	//TODO:分析数据的服务

	var err error
	dev := model.Device{}
	//只需要传递ID
	//c.ShouldBind(&dev)
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	//err := dal.Getdb().Model(&model.Device{}).Where("id = ?", dev.ID).First(&dev).Error
	dev, err = gen.Device.GetByID(id)
	if err != nil {
		c.JSON(200, response.Response{
			400,
			"加载页面失败 1",
		})
		return
	}
	driver, err := gen.Driver.GetByID(dev.OwnerID)
	if err != nil {
		c.JSON(200, response.Response{
			400,
			"加载页面失败 2",
		})
		return
	}
	products, err := gen.DeviceProduct.GetAllDeviceProduct(dev.ID)
	if err != nil {
		c.JSON(200, response.Response{
			400,
			"加载页面失败 3",
		})
		return
	}
	n := len(products)
	prinfos := make([]response.DevProductPartInfo, n, n)
	//var prinfo *response.DevProductPartInfo
	for i := 0; i < n; i++ {
		prinfos[i].Sku = products[i].Sku
		prinfos[i].Name = products[i].Name
		prinfos[i].Pic = products[i].Pic
		prinfos[i].Price = products[i].SupplyPrice
		prinfos[i].MonthlySales = products[i].MonthlySales
		prinfos[i].Inventory = products[i].Inventory
	}
	prinfos = append(prinfos, prinfos...)
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

//此处应该知道设备的ID
func (d *DevinfoController) LsDev(c *gin.Context) {
	dev := model.Device{}
	c.ShouldBind(&dev)
	// id:=c.PostForm("id")
	err := dal.Getdb().Model(&model.Device{}).Where("id = ?", dev.ID).First(&dev).Error
	if err == nil {
		c.JSON(200, response.Response{
			200,
			"ok"})
	}
}

func (d *DevinfoController) LsDriver(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}

//从数据库获取相关信息
func (d *DevinfoController) LsDevProduct(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}

func (d *DevinfoController) TakeDown(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}
