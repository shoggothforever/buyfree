package passenger

import "github.com/gin-gonic/gin"

type HomePageController struct {
	BasePaController
}

// @Summary 乘客端首页
// @Description 展示销售数据(日收入，日环比，周环比，本月收入，今日广告收入以及播放次数，两件热销商品)
// @Tags Driver
// @Accept json
// @Produce json
// @Success 200 {object} response.HomePageResponse
// @Failure 400 {object} response.Response
// @Router /dr/home [get]
func (h *HomePageController) GetStatic(c *gin.Context) {

}