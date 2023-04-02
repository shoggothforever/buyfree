package utils

import (
	"buyfree/middleware"
	"buyfree/repo/model"
	"github.com/gin-gonic/gin"
)

func GetDriveInfo(c *gin.Context) (admin model.Driver, ok bool) {
	iadmin, ok := c.Get(middleware.DRADMIN)
	if ok != true {
		return model.Driver{}, false
	}
	admin, ok = iadmin.(model.Driver)
	if ok != true {
		return model.Driver{}, false
	}
	return admin, ok
}
