package test

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestRegister(t *testing.T) {
	dsn := "host=localhost port=5432 user=bf dbname=bfdb password=bf123  sslmode=disable  TimeZone=Asia/Shanghai"
	DB, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{
		CreateBatchSize:        1000,
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	})
	//pt := model.Platform{}
	//pt.ID = 1
	//pt.PasswordSalt = "123"
	//pt.Password = "123"
	//pt.Role = 3
	//pt.Name = "dsm"
	//pt.Balance = 0
	//pt.Pic = "233"
	//pt.Level = 0
	//DB.Model(&model.Platform{}).Create(&pt)
	//l := model.LoginInfo{123, "123", "123", "123"}
	//DB.Model(&model.LoginInfo{}).Create(&l)
	//var si response.ScreenInfo
	//DB.Raw("select count(*) from advertisements").First(&si.DevNums)
	//t.Log(si.DevNums)

	//var partinfo []response.DevProductPartInfo
	//DB.Raw("select * from device_products").Find(&partinfo)
	//t.Log(partinfo)

	var start time.Time
	DB.Raw("select date_trunc('week',now())").First(&start)
	t.Log(start.In(time.Local))

}
