package platform

import (
	"buyfree/dal"
	"buyfree/repo/model"
	"buyfree/service/response"
	"buyfree/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DevadminController struct {
	BaseController
}

func GetOnlineState(state bool) string {
	if state == true {

		return "在线"
	} else {
		return "离线"
	}
}
func (d *DevadminController) GetdevBystate(c *gin.Context) {
	//mode =0 全部 1 在线 2离线 3已激活 4 未激活
	mode := c.Param("mode")
	var devs []*model.Device
	var driver model.Driver
	if mode == "1" {
		dal.Getdb().Model(&model.Device{}).Where("is_online = ?", true).Find(&devs)
	} else if mode == "2" {
		dal.Getdb().Model(&model.Device{}).Where("is_online = ?", false).Find(&devs)
	} else if mode == "3" {
		dal.Getdb().Model(&model.Device{}).Where("is_activated = ?", true).Find(&devs)
	} else if mode == "4" {
		dal.Getdb().Model(&model.Device{}).Where("is_activated = ?", false).Find(&devs)
	} else {
		dal.Getdb().Model(&model.Device{}).Find(&devs)
	}
	var size int = len(devs)
	devres := make([]response.DevQueryInfo, size)
	for k := 0; k < size; k++ {
		fmt.Println(devs[k])
		devres[k].Seq = int64(k) + 1
		devres[k].DevID = devs[k].ID
		devres[k].State = GetOnlineState(devs[k].IsOnline)
		err := dal.Getdb().Model(&model.Driver{}).Select("id", "name", "mobile").Where("id = ?", devs[k].OwnerID).First(&driver).Error
		if err == gorm.ErrRecordNotFound {
			continue
		}
		fmt.Println(driver)
		devres[k].DriverName = driver.Name
		devres[k].Mobile = driver.Mobile
		//TODO GET DRIVER LOCATION USING API
		//devre.Location = driver.Location
		//TODO: GET SALES DATA
		//devres[k].SaleForToday=
	}
	c.JSON(200, response.DevResponse{
		response.Response{
			200,
			"查询全部设备数据",
		},
		//TODO:序号交给前端
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
