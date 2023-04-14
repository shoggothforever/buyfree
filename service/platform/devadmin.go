package platform

import (
	"buyfree/dal"
	"buyfree/logger"
	"buyfree/middleware"
	"buyfree/repo/model"
	"buyfree/service/response"
	"buyfree/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type DevadminController struct {
	BasePtController
}

func GetOnlineState(state bool) string {
	if state == true {

		return "在线"
	} else {
		return "离线"
	}
}

// TODO:swagger
// @Summary 获取设备信息
// @Description 传入字段名:mode;mode=0:获取全部设备信息,mode=1，2,3,4分别对应获取在线，离线,激活，未激活的设备信息
// @Tags Platform/Device
// @Accept json
// @Accept mpfd
// @Produce json
// @Success 200 {object} response.DevResponse "获取的设备信息"
// @Failuer 400 {object} response.Response "对应mode的失败信息“
// @Param mode path int true "1：在线，2：离线，3：激活，4：未激活"
// @Param page path int true "默认第一页，一页20条数据"
// @Router /pt/dev-admin/list/{mode}/{page} [get]
func (d *DevadminController) GetdevBystate(c *gin.Context) {
	//mode =0 全部 1 在线 2离线 3已激活 4 未激活
	spage := c.Param("page")
	page, _ := strconv.Atoi(spage)
	if page < 1 {
		page = 1
	}
	mode := c.Param("mode")
	_, ok := c.Get(middleware.PTADMIN)
	if ok != true {
		d.Error(c, 400, "获取用户信息失败")
		return
	}
	var devs []*model.Device
	var driver model.Driver
	var err error
	//if mode == "1" {
	//	err = dal.Getdb().Model(&model.Device{}).Where("is_online = ? and platform_id = ?", true, pid).Find(&devs).Error
	//	if err != nil {
	//		d.Error(c, 400, "获取在线设备信息失败")
	//	}
	//} else if mode == "2" {
	//	err = dal.Getdb().Model(&model.Device{}).Where("is_online = ? and platform_id = ?", false, pid).Find(&devs).Error
	//	if err != nil {
	//		d.Error(c, 400, "获取离线设备信息失败")
	//	}
	//} else if mode == "3" {
	//	err = dal.Getdb().Model(&model.Device{}).Where("is_activated = ? and platform_id = ?", true, pid).Find(&devs).Error
	//	if err != nil {
	//		d.Error(c, 400, "获取激活设备信息失败")
	//	}
	//} else if mode == "4" {
	//	err = dal.Getdb().Model(&model.Device{}).Where("is_activated = ? and platform_id = ?", false, pid).Find(&devs).Error
	//	if err != nil {
	//		d.Error(c, 400, "获取未激活设备信息失败")
	//	}
	//} else {
	//	err = dal.Getdb().Model(&model.Device{}).Where("platform_id = ? ", pid).Find(&devs).Error
	//	if err != nil {
	//		d.Error(c, 400, "获取设备信息失败")
	//	}
	//}
	if mode == "1" {
		err = dal.Getdb().Model(&model.Device{}).Where("is_online = ?", true).Offset((page - 1) * 20).Limit(20).Find(&devs).Error
		if err != nil {
			logger.Loger.Info(err)
			d.Error(c, 400, "获取在线设备信息失败")
			return
		}
	} else if mode == "2" {
		err = dal.Getdb().Model(&model.Device{}).Where("is_online = ?", false).Offset((page - 1) * 20).Limit(20).Find(&devs).Error
		if err != nil {
			logger.Loger.Info(err)
			d.Error(c, 400, "获取离线设备信息失败")
			return
		}
	} else if mode == "3" {
		err = dal.Getdb().Model(&model.Device{}).Where("is_activated = ?", true).Offset((page - 1) * 20).Limit(20).Find(&devs).Error
		if err != nil {
			logger.Loger.Info(err)
			d.Error(c, 400, "获取激活设备信息失败")
			return
		}
	} else if mode == "4" {
		err = dal.Getdb().Model(&model.Device{}).Where("is_activated = ?", false).Offset((page - 1) * 20).Limit(20).Find(&devs).Error
		if err != nil {
			logger.Loger.Info(err)
			d.Error(c, 400, "获取未激活设备信息失败")
			return
		}
	} else {
		err = dal.Getdb().Model(&model.Device{}).Offset((page - 1) * 20).Limit(20).Find(&devs).Error
		if err != nil {
			d.Error(c, 400, "获取设备信息失败")
			return
		}
	}
	var size = len(devs)
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
		//fmt.Println(driver)
		devres[k].DriverName = driver.Name
		devres[k].Mobile = driver.Mobile
		//TODO GET DRIVER LOCATION USING API
		devres[k].Location = driver.Address
		//TODO: GET SALES DATA
		rdb := dal.Getrdb()
		var sales string
		sales, _ = rdb.Get(c, utils.GetSalesKeyByMode(utils.Ranktype1, devres[k].DriverName, 0)).Result()
		if sales != "" {
			devres[k].SalesOfToday, err = strconv.ParseFloat(sales, 64)
			if err != nil {
				logger.Loger.Info(err)
				d.Error(c, 400, "获取用户当日营销额信息失败")
				return
			}
		} else {
			devres[k].SalesOfToday = 0
		}
	}
	c.JSON(200, response.DevResponse{
		response.Response{
			200,
			"查询全部设备数据",
		},
		devres,
	})

}

// TODO:swagger
// @Summary 添加设备信息
// @Description 按照Device的定义 传入json格式的数据,添加的设备默认为未激活，未上线状态
// @Tags	Platform/Device
// @Accept json
// @Accept mpfd
// @Produce json
// @Success 201 {object} response.AddDevResponse
// @Failure 400 {object} response.Response
// @Router /pt/dev-admin/devs [post]
func (d *DevadminController) AddDev(c *gin.Context) {
	var dev model.Device
	dev.IsActivated = false
	dev.IsOnline = false
	dev.Profit = 0
	var err error
	//err = c.ShouldBindJSON(&dev)
	dev.ID = utils.GetSnowFlake()
	if err != nil {
		fmt.Println(err)
		d.Error(c, 400, "添加设备失败")
		return
	}
	//fmt.Println(dev)
	iadmin, ok := c.Get(middleware.PTADMIN)
	if ok != true {
		d.Error(c, 400, "获取用户信息失败")
		return
	}
	admin := iadmin.(model.Platform)
	dev.PlatformID = admin.ID
	err = dal.Getdb().Model(&model.Device{}).Omit("owner_id", "activated_time").Create(&dev).Error
	if err == nil {
		var qrcode = utils.GenerateSourceUrl(dev.ID)
		rdb := dal.Getrdb()
		rdb.Do(rdb.Context(), "set", "QR:"+strconv.FormatInt(dev.ID, 10), qrcode)
		c.JSON(200, response.AddDevResponse{
			response.Response{201,
				"添加设备成功",
			}, qrcode,
			&dev,
		})
	} else {
		fmt.Println(err)
		d.Error(c, 400, "添加设备失败")
	}
}
