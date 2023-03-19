package platform

import (
	"buyfree/dal"
	"buyfree/repo/model"
	"buyfree/service/response"
	"github.com/gin-gonic/gin"
)

func GetState(state bool) string {
	if state == true {

		return "在线"
	} else {
		return "离线"
	}
}
func GetAlldev(c *gin.Context) {
	var devs []*model.DEVICE
	var driver model.Driver
	//var devres []*response.DevQueryInfo
	dal.Getdb().Model(&model.DEVICE{}).Find(&devs)
	var size int = len(devs)
	devres := make([]*response.DevQueryInfo, size)
	for k := 0; k < size; k++ {
		devres[k].Seq = int64(k) + 1
		devres[k].DevID = devs[k].ID
		devres[k].State = GetState(devs[k].IsOnline)
		dal.Getdb().Model(&model.Driver{}).Select("id", "name", "mobile").Where("id = ?", devs[k].OwnerID).First(&driver)
		devres[k].DriverName = driver.Name
		devres[k].Mobile = driver.Mobile
		devres[k].Location = driver.Location

		//TODO
	}

}

func GetOndev(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}

func GetOffdev(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}

//从数据库获取相关信息
func GetActivated(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}
func GetNotActivated(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}
func AddDev(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}
