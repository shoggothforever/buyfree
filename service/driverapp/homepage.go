package driverapp

import "github.com/gin-gonic/gin"

type HomePageController struct {
	BaseDrController
}

// @Summary 车主端首页
// @Description 展示销售数据
// @Tags Driver
// @Accept json
// @Produce json
// @Success 200 {object} response.HomePageResponse
// @Failure 400 {object} response.Response
// @Router /dr/home [get]
func (h *HomePageController) GetStatic(c *gin.Context) {

}
