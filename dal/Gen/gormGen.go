package main

import (
	"buyfree/config"
	"buyfree/dal"
	"buyfree/repo/model"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

var dsn string

//需要注意如果传递的参数不止一个且其中包含了uuid，那么需要将uuid编程string
type Query interface {
	// SELECT * FROM @@table WHERE id=@id
	GetByID(id int64) (gen.T, error)

	// SELECT * FROM @@table WHERE name=@id
	GetByName(id int64) (gen.T, error)
}
type LoginQuery interface {
	// SELECT * FROM @@table WHERE user_id=@uid and password=@psw
	GetByNameAndPsw(uid int64, psw string) (gen.T, error)
}

type CartQuery interface {
	// SELECT * FROM @@table WHERE cart_id=@id
	GetByCardID(id int64) (gen.T, error)

	// sql(SELECT * FROM @@table where driver_id=(SELECT id from drivers where id =@driverid))
	GetAllCarts(driverid int64) ([]gen.T, error)
}
type ProductQuery interface {
	//sql(SELECT * FROM @@table where device_id=(SELECT id from devices where id=@deviceid))
	GetAllDeviceProduct(deviceid int64) ([]gen.T, error)
	//sql(SELECT * FROM @@table where cart_refer=(SELECT cart_id from driver_carts where cart_id=@cartrefer))
	DGetAllOrderProductReferCart(cartrefer int64) ([]gen.T, error)
	//sql(SELECT * FROM @@table where cart_refer=(SELECT cart_id from passenger_carts where cart_id=@cartrefer))
	PGetAllOrderProductReferCart(cartrefer int64) ([]gen.T, error)
	//sql(SELECT * FROM @@table where factory_id=(SELECT id from factories where id=@factoryid))
	GetAllOrderProductReferFactory(factoryid int64) ([]gen.T, error)
	//sql(SELECT * FROM @@table where order_refer=(SELECT order_id from driver_order_forms where order_id=@orderrefer))
	GetAllOrderProductReferDOrder(orderrefer string) ([]gen.T, error)
	//sql(SELECT * FROM @@table where order_refer=(SELECT order_id from passenger_order_forms where order_id=@orderrefer))
	GetAllOrderProductReferPOrder(orderrefer string) ([]gen.T, error)
	//sql(SELECT * FROM @@table where sku=@sku and factory_name=@fname)
	GetBySkuAndFName(sku, fname string) (gen.T, error)
}

type DeviceQuery interface {
	//sql(SELECT * FROM @@table where owner_id=(SELECT id from drivers where id=@id))
	GetAllDriverDevice(id int64) (gen.T, error)
	//sql(SELECT * FROM @@table where platform_id=(SELECT id from platforms where id=@id))
	GetAllPlatformDevice(id int64) (gen.T, error)

	//sql(SELECT * FROM @@table where is_online=@mode and platform_id=(SELECT id from platforms where id=@id))
	GetByOnlinePlatformDevice(id int64, mode bool) (gen.T, error)

	//sql(SELECT * FROM @@table where is_activated=@mode and platform_id=(SELECT id from platforms where id=@id))
	GetByActivatedPlatformDevice(id int64, mode bool) (gen.T, error)
	//sql(select * from devices as d where d.id in (select device_id from ad_devices where advertisement_id = @ad_id ))
	GetDeviceByAdvertiseID(ad_id int64) ([]gen.T, error)
}
type OrderFormQuery interface {
	//sql(SELECT * FROM @@table where factory_id =(SELECT id from factories where id=@id))
	FGetAllOrderForms(id int64) ([]gen.T, error)

	//sql(SELECT * FROM @@table where state=@mode and factory_id=(SELECT id from factories where id=@id))
	FGetByStateOrderForms(id int64, mode model.ORDERSTATE) ([]gen.T, error)

	//sql(SELECT * FROM @@table where driver_id =(SELECT id from drivers where id=@id))
	DGetAllOrderForms(id int64) ([]gen.T, error)

	//sql(SELECT * FROM @@table where state=@mode and driver_id =(SELECT id from drivers where id=@id))
	DGetByStateOrderForms(id int64, mode model.ORDERSTATE) ([]gen.T, error)

	//sql(SELECT * FROM @@table where passenger_id =(SELECT id from passengers where id=@id))
	PGetAllOrderForms(id int64) ([]gen.T, error)

	//sql(SELECT * FROM @@table where state=@mode and passenger_id =(SELECT id from passengers where id=@id))
	PGetByStateOrderForms(id int64, mode model.ORDERSTATE) ([]gen.T, error)
}

type AdvertisementQuery interface {
	//sql(select * from advertisements as a where a.id in(select advertisement_id from ad_devices where device_id = @dev_id))
	GetAdvertisementByDeviceID(dev_id int64) ([]gen.T, error)
	//select * from ad_devices where advertisement_id = @ad_id and device_id = @dev_id
	GetAdvertisementProfitAndPlayTimes(ad_id, dev_id int64) (gen.T, error)
}

func main() {
	var err error
	config.Init()
	dal.ReadPostgresSQLlinfo()
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		CreateBatchSize:        1000,
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		logrus.WithFields(logrus.Fields{"error": err}).Error("Open PostgresSQL failed")
	} else {
		logrus.Info("Open postgresSQL successfully")
	}
	//TODO:use gen to simplify CRUD
	g := gen.NewGenerator(gen.Config{
		OutPath: "D:\\desktop\\pr\\buyfree\\repo/gen",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})
	g.UseDB(DB)
	//g.ApplyBasic(&model.User{})
	g.ApplyInterface(func(Query) {}, &model.Passenger{},
		&model.Driver{}, &model.Factory{},
		&model.Platform{}, &model.DeviceProduct{}, &model.FactoryProduct{},
		&model.Device{}, &model.Advertisement{}, &model.BankCardInfo{})
	g.ApplyInterface(func(LoginQuery) {}, &model.LoginInfo{})
	g.ApplyInterface(func(CartQuery) {}, &model.PassengerCart{}, &model.DriverCart{})
	g.ApplyInterface(func(ProductQuery) {}, &model.OrderProduct{}, &model.DeviceProduct{}, &model.FactoryProduct{})
	g.ApplyInterface(func(DeviceQuery) {}, &model.Device{})
	g.ApplyInterface(func(OrderFormQuery) {}, &model.PassengerOrderForm{}, &model.DriverOrderForm{})
	g.ApplyInterface(func(AdvertisementQuery) {}, &model.Advertisement{})
	g.Execute()
}
