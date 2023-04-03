package driverapp

import (
	"buyfree/service/response"
	"buyfree/transport"
	"buyfree/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

type DeviceController struct {
	BaseDrController
}

// @Summary 扫码激活设备
// @Description 扫码向服务端(平台)验签，验签成功，返回待激活设备号码
// @Tags Driver/Auth
// @Accept json
// @Produce json
// @Success 200 {object} response.ScanResponse
// @Failure 400 {onject} response.Response
// @Router /dr/devices/scan [get]
func (d *DeviceController) Scan(c *gin.Context) {
	admin, ok := utils.GetDriveInfo(c)
	if !ok {
		d.Error(c, 400, "获取用户信息失败")
	}
	fmt.Println(admin.ID)
	var id int64
	var req *transport.ScanRequest = transport.NewScanRequest(admin.ID, &id)

	transport.PlatFormService.ReqChan <- req
	res := <-req.ReplyChan
	if res != true {
		d.Error(c, 400, "验签失败")
		return
	}
	c.JSON(200, response.ScanResponse{response.Response{200, "扫码成功"}, id})

}

// @Summary 输入认证信息绑定设备
// @Description 向服务端(平台)验签，等待设备激活
// @Tags Driver/Auth
// @Accept json
// @Produce json
// @Param AuthInfo body response.DriverAuthInfo true "传入获得的设备ID,以及一些车主的相关信息"
// @Success 200 {object} response.BindDeviceResponse
// @Failure 400 {onject} response.Response
// @Router /dr/devices/bind [post]
func (d *DeviceController) BindDevice(c *gin.Context) {
	var info response.DriverAuthInfo
	if err := c.ShouldBind(&info); err != nil {
		d.Error(c, 400, "获取提交信息失败")
		return
	}
	var req *transport.DeviceAuthRequest = transport.NewDeviceAuthRequest(info.DriverID, info.DeviceID, info.Name, info.Mobile)
	transport.PlatFormService.ReqChan <- req
	res := <-req.ReplyChan
	if res != true {
		d.Error(c, 400, "绑定用户信息失败")
		return
	} else {
		c.JSON(200, response.BindDeviceResponse{response.Response{200, "绑定用户信息成功"}, &info})
	}
}
