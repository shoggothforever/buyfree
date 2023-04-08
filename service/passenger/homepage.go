package passenger

import (
	"buyfree/mrpc"
	"buyfree/service/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

type HomePageController struct {
	BasePaController
}

// @Summary 乘客端首页
// @Description 用户扫码打开小程序，
// @Tags Driver
// @Accept json
// @Produce json
// @Param id path int true "扫码获取的设备id"
// @Success 200 {object} response.PassengerHomeResponse
// @Failure 400 {object} response.Response
// @Router /home/{id} [get]
func (h *HomePageController) GetStatic(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	homereq := mrpc.NewHomeScanReq(id)
	mrpc.DriverService.ReqChan <- homereq
	<-homereq.DoneChan
	if !homereq.Res {
		h.Error(c, 400, "获取信息失败")
	} else {
		c.JSON(200, response.PassengerHomeResponse{
			response.Response{200, "获取信息成功"},
			homereq.ADUrls,
			homereq.DeviceProducts,
		})
	}
}
