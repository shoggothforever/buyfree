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
	var devs []*model.Device
	var driver model.Driver
	dal.Getdb().Model(&model.Device{}).Find(&devs)
	var size int = len(devs)
	devres := make([]*response.DevQueryInfo, size)
	for k := 0; k < size; k++ {
		devres[k].Seq = int64(k) + 1
		devres[k].DevID = devs[k].ID
		devres[k].State = GetState(devs[k].IsOnline)
		dal.Getdb().Model(&model.Driver{}).Select("id", "name", "mobile").Where("id = ?", devs[k].OwnerID).First(&driver)
		devres[k].DriverName = driver.Name
		devres[k].Mobile = driver.Mobile
		//TODO GET DRIVER LOCATION USING API
		devres[k].Location = driver.Location

		//TODO: GET SALES DATA
		//devres[k].SaleForToday=

	}
	c.JSON(200, response.DevResponse{
		response.Response{
			200,
			"查询全部设备数据",
		},
		devres,
	})

}

func GetOndev(c *gin.Context) {
	var devs []*model.Device
	var driver model.Driver
	dal.Getdb().Model(&model.Device{}).Where("is_online = ?", true).Find(&devs)
	var size int = len(devs)
	devres := make([]*response.DevQueryInfo, size)
	for k := 0; k < size; k++ {
		devres[k].Seq = int64(k) + 1
		devres[k].DevID = devs[k].ID
		devres[k].State = GetState(devs[k].IsOnline)
		dal.Getdb().Model(&model.Driver{}).Select("id", "name", "mobile").Where("id = ?", devs[k].OwnerID).First(&driver)
		devres[k].DriverName = driver.Name
		devres[k].Mobile = driver.Mobile
		//TODO GET DRIVER LOCATION USING API
		devres[k].Location = driver.Location

		//TODO: GET SALES DATA
		//devres[k].SaleForToday=

	}
	c.JSON(200, response.DevResponse{
		response.Response{
			200,
			"查询全部设备数据",
		},
		devres,
	})
}

func GetOffdev(c *gin.Context) {
	var devs []*model.Device
	var driver model.Driver
	dal.Getdb().Model(&model.Device{}).Where("is_online = ?", false).Find(&devs)
	var size int = len(devs)
	devres := make([]*response.DevQueryInfo, size)
	for k := 0; k < size; k++ {
		devres[k].Seq = int64(k) + 1
		devres[k].DevID = devs[k].ID
		devres[k].State = GetState(devs[k].IsOnline)
		dal.Getdb().Model(&model.Driver{}).Select("id", "name", "mobile").Where("id = ?", devs[k].OwnerID).First(&driver)
		devres[k].DriverName = driver.Name
		devres[k].Mobile = driver.Mobile
		//TODO GET DRIVER LOCATION USING API
		devres[k].Location = driver.Location

		//TODO: GET SALES DATA
		//devres[k].SaleForToday=

	}
	c.JSON(200, response.DevResponse{
		response.Response{
			200,
			"查询全部设备数据",
		},
		devres,
	})
}

//从数据库获取相关信息
func GetActivated(c *gin.Context) {
	var devs []*model.Device
	var driver model.Driver
	dal.Getdb().Model(&model.Device{}).Where("is_activated = ?", true).Find(&devs)
	var size int = len(devs)
	devres := make([]*response.DevQueryInfo, size)
	for k := 0; k < size; k++ {
		devres[k].Seq = int64(k) + 1
		devres[k].DevID = devs[k].ID
		devres[k].State = GetState(devs[k].IsOnline)
		dal.Getdb().Model(&model.Driver{}).Select("id", "name", "mobile").Where("id = ?", devs[k].OwnerID).First(&driver)
		devres[k].DriverName = driver.Name
		devres[k].Mobile = driver.Mobile
		//TODO GET DRIVER LOCATION USING API
		devres[k].Location = driver.Location

		//TODO: GET SALES DATA
		//devres[k].SaleForToday=

	}
	c.JSON(200, response.DevResponse{
		response.Response{
			200,
			"查询全部设备数据",
		},
		devres,
	})
}
func GetNotActivated(c *gin.Context) {
	var devs []*model.Device
	var driver model.Driver
	dal.Getdb().Model(&model.Device{}).Where("is_activated = ?", false).Find(&devs)
	var size int = len(devs)
	devres := make([]*response.DevQueryInfo, size)
	for k := 0; k < size; k++ {
		devres[k].Seq = int64(k) + 1
		devres[k].DevID = devs[k].ID
		devres[k].State = GetState(devs[k].IsOnline)
		dal.Getdb().Model(&model.Driver{}).Select("id", "name", "mobile").Where("id = ?", devs[k].OwnerID).First(&driver)
		devres[k].DriverName = driver.Name
		devres[k].Mobile = driver.Mobile
		//TODO GET DRIVER LOCATION USING API
		devres[k].Location = driver.Location

		//TODO: GET SALES DATA
		//devres[k].SaleForToday=

	}
	c.JSON(200, response.DevResponse{
		response.Response{
			200,
			"查询全部设备数据",
		},
		devres,
	})
}
func AddDev(c *gin.Context) {
	var dev []*model.Device
	c.ShouldBind(&dev)
	if len(dev) == 0 {
		c.JSON(200, response.Response{
			400,
			"添加设备失败，请输入正确的设备信息",
		})
	} else {
		err := dal.Getdb().Model(&model.Device{}).Create(&dev[0])
		if err == nil {
			c.JSON(200, response.AddDevResponse{
				response.Response{200,
					"添加设备成功",
				},
				dev[0],
			})
		} else {
			c.JSON(200, response.Response{400, "添加设备失败"})
		}
	}
}
