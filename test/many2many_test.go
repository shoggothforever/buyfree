package test

import (
	"buyfree/repo/model"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
)

func TestAD_Device(t *testing.T) {
	dsn := "host=localhost port=5432 user=bf dbname=bfdb password=bf123  sslmode=disable  TimeZone=Asia/Shanghai"
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		CreateBatchSize: 1000,
		//PrepareStmt:            true,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		logrus.WithFields(logrus.Fields{"error": err}).Error("Open PostgresSQL failed")
	} else {
		logrus.Info("Open postgresSQL successfully")
	}
	//DB.AutoMigrate(&model.User{}, &model.Driver{}, &model.Device{}, &model.Advertisement{})
	var id int64 = 1
	var ad model.Advertisement
	//var drivers []*model.Driver
	//var driver *model.Driver
	var devices []model.Device
	//var effinfo []*response.ADEfficientInfo
	DB.Model(&model.Advertisement{}).Where("id = ?", id).First(&ad)
	fmt.Println(ad)
	//TODO:好好写原生SQL语句
	DB.Model(&model.Device{}).Raw("select * from devices as d where d.id in "+
		"(select device_id from ad_devices where advertisement_id = ? )", id).Find(&devices)

	fmt.Println(devices)

}
