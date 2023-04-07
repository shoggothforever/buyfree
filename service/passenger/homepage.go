package passenger

import "github.com/gin-gonic/gin"

type HomePageController struct {
	BasePaController
}

// @Summary 乘客端首页
// @Description 用户扫码打开小程序，
// @Tags Driver
// @Accept json
// @Produce json
// @Success 200 {object} response.HomePageResponse
// @Failure 400 {object} response.Response
// @Router /home [get]
func (h *HomePageController) GetStatic(c *gin.Context) {

}
