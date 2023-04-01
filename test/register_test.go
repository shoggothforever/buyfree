package test

import (
	"buyfree/service/response"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
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

	//var start time.Time
	//DB.Raw("select date_trunc('week',now())").First(&start)
	//t.Log(start.In(time.Local))
	//var fa response.FactoryDetail
	//err := DB.Model(&model.Factory{}).Select("address", "description").Where("name=?", "penguin").First(&fa).Error
	//t.Log(fa, err)
	DB.Transaction(func(tx *gorm.DB) error {
		//var ids []int64
		//DB.Model(&model.Device{}).Select("id").Where("owner_id = ?", 1).Find(&ids)
		//var pname []string
		//DB.Model(&model.FactoryProduct{}).Select("name").Where("is_on_shelf = true and factory_name = ?", "cat").Find(&pname)
		//var m_inv []response.DriveInventory
		//DB.Raw("select name,sum(inventory) as inventory from device_products where device_id in ? and name in ? group by name", ids, pname).Find(&m_inv)
		var details []response.FactoryProductDetail
		//err := DB.Model(&model.FactoryProduct{}).Omit("m_inventory").Where("is_on_shelf = ? "+
		//	"and factory_name= ?", true, "cat").Find(&details).Error
		//if err != nil {
		//	return err
		//}
		err := DB.Raw("select fp.name,inventory,"+
			"dv.m_inventory,pic,type,monthly_sales,supply_price  "+
			"from factory_products as fp,(select dp.name,sum(dp.inventory)"+
			" as m_inventory from device_products dp where device_id in"+
			"(select id from devices where owner_id =?) and dp.name in "+
			"(select name from factory_products where factory_name=?)"+
			"group by (dp.name)) as dv where is_on_shelf=true "+
			"and factory_name=? and fp.name=dv.name", 242731811816345600, "cat", "cat").Find(&details).Error
		t.Log(details)
		return err
	})
	//select fp.name,inventory,dv.m_inventory,pic,type,monthly_sales,supply_price  from factory_products as fp,(select dp.name,sum(dp.inventory) as m_inventory from device_products dp where device_id in(select id from devices where owner_id =1) and dp.name in (select name from factory_products where factory_name='cat')group by (dp.name)) as dv where is_on_shelf=true and factory_name='cat' and fp.name=dv.name
}
