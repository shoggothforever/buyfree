package platform

import (
	"buyfree/dal"
	"buyfree/repo/model"
	"buyfree/service/response"
	"buyfree/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	err = dal.Getdb().Model(&model.Device{}).Where("id=?", id).First(&dev).Error
	if err != gorm.ErrRecordNotFound && err != nil {
		c.JSON(200, response.Response{
			400,
			"查找对应设备信息失败,请检查输入ID是否正确",
		})
		return
	}
	var driver model.Driver
	err = dal.Getdb().Model(&model.Driver{}).Where("id=?", dev.OwnerID).First(&driver).Error
	if err != gorm.ErrRecordNotFound && err != nil {
		c.JSON(200, response.Response{
			400,
			"查找对应设备绑定的车主信息失败,请检查设备是否合法",
		})
		return
	}
	var prinfos []response.DevProductPartInfo
	err = dal.Getdb().Raw("SELECT * FROM device_products where device_id="+
		"(SELECT id from devices where id=?)", dev.ID).Find(&prinfos).Error
	if err != gorm.ErrRecordNotFound && err != nil {
		c.JSON(200, response.Response{
			400,
			"查找对应设备的商品信息失败",
		})
		return
	}
	iadmin, ok := c.Get("admin")
	name := iadmin.(model.Platform).Name
	if ok != true {
		d.Error(c, 400, "获取用户信息失败")
		return
	}
	rdb := dal.Getrdb()
	info, err := utils.GetSalesInfo(c, rdb, name)
	if err != nil {
		c.JSON(200, response.Response{
			400,
			"获取销量数据失败",
		})
		return
	}
	var salesinfo model.SalesData
	salesinfo.DailySales = info[0]
	salesinfo.WeeklySales = info[1]
	salesinfo.MonthlySales = info[2]
	salesinfo.AnnuallySales = info[3]
	salesinfo.TotalSales = info[4]
	if err == nil {
		c.JSON(200, response.DevInfoResponse{
			response.Response{
				200,
				"获取设备详细信息成功"},
			//TODO:添加数据分析的响应信息
			salesinfo,
			response.DevInfo{
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
