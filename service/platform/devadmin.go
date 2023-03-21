package platform

import (
	"buyfree/dal"
	"buyfree/repo/model"
	"buyfree/service/response"
	"buyfree/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

type DevadminController struct {
	BaseController
}

func GetState(state bool) string {
	if state == true {

		return "在线"
	} else {
		return "离线"
	}
}
func (d *DevadminController) GetAlldev(c *gin.Context) {
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

func (d *DevadminController) GetOndev(c *gin.Context) {
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

func (d *DevadminController) GetOffdev(c *gin.Context) {
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
func (d *DevadminController) GetActivated(c *gin.Context) {
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
func (d *DevadminController) GetNotActivated(c *gin.Context) {
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
func (d *DevadminController) AddDev(c *gin.Context) {
	var dev model.Device
	var err error
	err = c.ShouldBindJSON(&dev)
	fmt.Println(utils.GetSnowFlake())
	dev.ID = utils.GetSnowFlake()
	if err != nil {
		c.JSON(200, response.Response{
			400,
			"添加设备失败，请输入正确的设备信息",
		})
		return
	}
	fmt.Println(dev)
	err = dal.Getdb().Model(&model.Device{}).Create(&dev).Error
	if err == nil {
		c.JSON(200, response.AddDevResponse{
			response.Response{200,
				"添加设备成功",
			},
			&dev,
		})
	} else {
		c.JSON(200, response.Response{400, "添加设备失败"})
	}
}
