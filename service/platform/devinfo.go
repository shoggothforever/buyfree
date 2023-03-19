package platform

import (
	"buyfree/dal"
	"buyfree/repo/model"
	"buyfree/service/response"
	"github.com/gin-gonic/gin"
)

//TODO: 利用redis对车主端的数据进行统计，
func AnaSales(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}
func LsInfo(c *gin.Context) {
	dev := model.Device{}
	c.ShouldBind(&dev)
	// id:=c.PostForm("id")
	err := dal.Getdb().Model(&model.Device{}).Where("id = ?", dev.ID).First(&dev).Error

	//driver_id:=dev.OwnerID
	//err=
	if err == nil {
		c.JSON(200, response.Response{
			200,
			"ok"})
	}
}

//此处应该知道设备的ID
func LsDev(c *gin.Context) {
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

func LsDriver(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}

//从数据库获取相关信息
func LsDevProduct(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}

func TakeDown(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}
