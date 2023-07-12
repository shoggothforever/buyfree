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
	"github.com/sirupsen/logrus"
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
	d.rwm.RLock()
	defer d.rwm.RUnlock()
	switch mode {
	case "1":
		err = dal.Getdb().Model(&model.Device{}).Where("is_online = ?", true).Offset((page - 1) * 20).Limit(20).Find(&devs).Error
		if err != nil {
			logger.Loger.Info(err)
			d.Error(c, 400, "获取在线设备信息失败")
			return
		}
	case "2":
		err = dal.Getdb().Model(&model.Device{}).Where("is_online = ?", false).Offset((page - 1) * 20).Limit(20).Find(&devs).Error
		if err != nil {
			logger.Loger.Info(err)
			d.Error(c, 400, "获取离线设备信息失败")
			return
		}
	case "3":
		err = dal.Getdb().Model(&model.Device{}).Where("is_activated = ?", true).Offset((page - 1) * 20).Limit(20).Find(&devs).Error
		if err != nil {
			logger.Loger.Info(err)
			d.Error(c, 400, "获取激活设备信息失败")
			return
		}
	case "4":
		err = dal.Getdb().Model(&model.Device{}).Where("is_activated = ?", false).Offset((page - 1) * 20).Limit(20).Find(&devs).Error
		if err != nil {
			logger.Loger.Info(err)
			d.Error(c, 400, "获取未激活设备信息失败")
			return
		}
	default:
		err = dal.Getdb().Model(&model.Device{}).Offset((page - 1) * 20).Limit(20).Find(&devs).Error
		if err != nil {
			d.Error(c, 400, "获取设备信息失败")
			return
		}
	}
	var size = len(devs)
	devres := make([]response.DevQueryInfo, size)
	for k := 0; k < size; k++ {
		devres[k].Seq = int64(k) + 1
		devres[k].DevID = devs[k].ID
		devres[k].State = GetOnlineState(devs[k].IsOnline)
		err := dal.Getdb().Raw("select * from drivers where id = ?", devs[k].OwnerID).First(&driver).Error
		if err == gorm.ErrRecordNotFound {
			logger.Loger.WithFields(logrus.Fields{"车主ID:": devs[k].OwnerID}).Info(err)
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
		rdb.Do(c, "set", "QR:"+strconv.FormatInt(dev.ID, 10), qrcode)
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

// TODO:swagger
// @Summary 展示设备详情信息
// @Description	输入设备的ID以查看对应设备的销量,绑定车主以及库存的信息
// @Tags	Platform/Device
// @Accept json
// @Accept mpfd
// @Produce json
// @Param id path int true "used to identify Device"
// @Success 200 {object} response.DevInfoResponse
// @Failure 400 {object} response.Response
// @Router /pt/dev-admin/infos/{id} [get]
func (d *DevadminController) LsInfo(c *gin.Context) {
	//TODO:分析数据的服务
	var err error
	dev := model.Device{}
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	//查找设备
	d.rwm.RLock()
	defer d.rwm.RUnlock()
	err = dal.Getdb().Model(&model.Device{}).Where("id=?", id).First(&dev).Error
	if err != nil {
		c.JSON(200, response.Response{
			400,
			"查找对应设备信息失败,请检查输入ID是否正确",
		})
		return
	}
	var driver model.Driver
	//获取设备表外键关联的司机表信息
	err = dal.Getdb().Model(&model.Driver{}).Where("id=?", dev.OwnerID).First(&driver).Error
	if err != gorm.ErrRecordNotFound && err != nil {
		c.JSON(200, response.Response{
			400,
			"查找对应设备绑定的车主信息失败,请检查设备是否合法",
		})
		return
	}
	var prinfos []response.DevProductPartInfo
	err = dal.Getdb().Raw("SELECT * FROM device_products where device_id=?", dev.ID).Find(&prinfos).Error
	if err != gorm.ErrRecordNotFound && err != nil {
		c.JSON(200, response.Response{
			400,
			"查找对应设备的商品信息失败",
		})
		return
	}
	rdb := dal.Getrdb()
	info, err := utils.GetSalesInfo(c, rdb, utils.Ranktype3, strconv.FormatInt(dev.ID, 10))
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
			salesinfo,
			response.DevInfo{
				dev.ID,
				dev.ActivatedTime,
				dev.UpdatedTime,
				driver.Address,
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

// @Summary 投放广告
// @Description 传入广告ID。激活的设备id
// @Tags	Platform/Advertisement
// @Accept json
// @Accept mpfd
// @Produce json
// @Param ad_ids body []int64 true "选中的广告ID"
// @Param dev_id path int true "设备ID"
// @Success 200 {object} response.ADLaunchResponse
// @Failure 400 {object} response.Response
// @Router /pt/dev-admin/launch/{dev_id} [post]
func (d *DevadminController) Launch(c *gin.Context) {
	var lrsponse response.ADLaunchResponse
	lrsponse.Response = response.Response{200, "广告成功上线"}
	var ad_ids []int64
	err := c.ShouldBind(&ad_ids)
	fmt.Println(ad_ids)
	if err != nil {
		d.Error(c, 400, "获取广告列表失败")
		return
	}
	dev_id := c.Param("dev_id")
	db := dal.Getdb()
	d.rwm.Lock()
	defer d.rwm.Unlock()
	err = db.Transaction(func(tx *gorm.DB) error {
		for _, ad_id := range ad_ids {
			var id int64
			db.Model(&model.Advertisement{}).Where("id = ?", ad_id).Update("ad_state", 1)
			ferr := db.Model(&model.Device{}).Select("owner_id").Where("id = ?", dev_id).First(&id).Error
			if ferr != nil {
				logger.Loger.Info(ferr)
				d.Error(c, 400, "获取用户信息失败")
				return ferr
			}
			var ids []int64
			ferr = db.Model(&model.Device{}).Select("id").Where("owner_id = ?", id).Find(&ids).Error
			if ferr != nil {
				logger.Loger.Info(ferr)
				d.Error(c, 400, "获取设备信息失败")
				return ferr
			}
			var ads model.Ad_Device
			ads.AdvertisementID = ad_id
			ads.PlayTimes = 0
			ads.Profit = 0
			pair := make([]response.PAIR_DEV_AD, len(ids))
			for k, v := range ids {
				ads.DeviceID = v
				pair[k].ADID = ad_id
				pair[k].DEVID = v
				db.Model(&model.Ad_Device{}).Create(&ads)
			}
			lrsponse.Pair = append(lrsponse.Pair, pair...)
		}
		return nil
	})
	if err == nil {
		c.JSON(200, lrsponse)
	} else {
		c.JSON(200, response.Response{400, "广告上线失败"})
	}

}
