package driverapp

import (
	"buyfree/dal"
	"buyfree/mrpc"
	"buyfree/repo/model"
	"buyfree/service/response"
	"buyfree/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"strconv"
)

type DeviceController struct {
	BaseDrController
}

// @Summary 扫码激活设备
// @Description 扫码向服务端(平台)验签，验签成功，返回一个待激活设备号码
// @Tags Driver/Auth
// @Accept json
// @Produce json
// @Success 200 {object} response.ScanResponse
// @Failure 400 {object} response.Response
// @Router /dr/devices/scan [get]
func (d *DeviceController) Scan(c *gin.Context) {
	admin, ok := utils.GetDriveInfo(c)
	if !ok {
		d.Error(c, 400, "获取用户信息失败")
	}
	fmt.Println(admin.ID)
	var id int64
	var req *mrpc.ScanRequest = mrpc.NewScanRequest(admin.ID, &id)

	mrpc.PlatFormService.ReqChan <- req
	<-req.DoneChan
	if req.Res != true {
		d.Error(c, 400, "验签失败")
		return
	}
	if id != 0 {
		c.JSON(200, response.ScanResponse{response.Response{200, "扫码成功"}, id})
	} else {
		c.JSON(200, response.ScanResponse{response.Response{400, "请联系平台更新设备信息"}, 0})
	}
}

// @Summary 输入认证信息绑定设备
// @Description 向服务端(平台)验签，等待设备激活
// @Tags Driver/Auth
// @Accept json
// @Produce json
// @Param AuthInfo body response.DriverAuthInfo true "传入获得的设备ID,以及一些车主的相关信息"
// @Success 200 {object} response.BindDeviceResponse
// @Failure 400 {object} response.Response
// @Router /dr/devices/bind [post]
func (d *DeviceController) BindDevice(c *gin.Context) {
	var info response.DriverAuthInfo
	if err := c.ShouldBind(&info); err != nil {
		d.Error(c, 400, "获取提交信息失败")
		return
	}
	admin, ok := utils.GetDriveInfo(c)
	if !ok {
		d.Error(c, 400, "获取司机信息失败")
		return
	}
	var req = mrpc.NewDeviceAuthRequest(admin.ID, info.DeviceID, info.Name, info.Mobile)
	mrpc.PlatFormService.ReqChan <- req
	<-req.DoneChan
	if req.Res != true {
		d.Error(c, 400, "绑定用户信息失败")
		return
	} else {
		dal.Getdb().Model(&model.Driver{}).Where("id = ?", admin.ID).Update("is_auth", true)
		c.JSON(200, response.BindDeviceResponse{response.Response{200, "绑定用户信息成功"}, &info})
	}
}

// @Summary 获取二维码
// @Description 获取车主所有设备的二维码信息
// @Tags Driver/Auth
// @Accept json
// @Produce json
// @Success 200 {object} response.QRCodeResponse
// @Failure 400 {object} response.Response
// @Router /dr/devices/QR [get]
func (d *DeviceController) QR(c *gin.Context) {
	admin, ok := utils.GetDriveInfo(c)
	if !ok {
		d.Error(c, 400, "获取用户信息失败")
		return
	}
	var urlinfos []response.QRUrlInfo
	var ids []int64
	err := dal.Getdb().Raw("select id from devices where owner_id = ? ", admin.ID).Find(&ids).Error
	if err != nil {
		logrus.Info(err)
		d.Error(c, 400, "未获取到该设备的二维码信息")
		return
	}
	n := len(ids)
	urlinfos = make([]response.QRUrlInfo, n)
	rdb := dal.Getrdb()
	for i := 0; i < n; i++ {
		urlinfos[i].DeviceID = ids[i]
		iurl, err := rdb.Do(rdb.Context(), "get", "QR:"+strconv.FormatInt(ids[i], 10)).Result()
		if err != nil {
			continue
		}
		urlinfos[i].QRUrl = iurl.(string)
	}
	c.JSON(200, response.QRCodeResponse{response.Response{200, "扫码成功"}, urlinfos})

}
