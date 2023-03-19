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

	// sql(SELECT * FROM @@table where @driverid=(SELECT id from drivers where id =@driverid))
	GetAllCarts(driverid int64) ([]gen.T, error)
}
type ProductQuery interface {
	//sql(SELECT * FROM @@table where @deviceid=(SELECT id from devices where id=@deviceid))
	GetAllDeviceProduct(deviceid int64) ([]gen.T, error)
	//sql(SELECT * FROM @@table where @cartrefer=(SELECT cart_id from driver_carts where cart_id=@cartrefer))
	GetAllOrderProductReferDCart(cartrefer int64) ([]gen.T, error)
	//sql(SELECT * FROM @@table where @cartrefer=(SELECT cart_id from passenger_carts where cart_id=@cartrefer))
	GetAllOrderProductReferPCart(cartrefer int64) ([]gen.T, error)
	//sql(SELECT * FROM @@table where @factoryrefer=(SELECT id from factories where id=@factoryrefer))
	GetAllOrderProductReferFactory(factoryrefer int64) ([]gen.T, error)
	//sql(SELECT * FROM @@table where @orderrefer=(SELECT order_id from driver_order_forms where order_id=@orderrefer))
	GetAllOrderProductReferDOrder(orderrefer string) ([]gen.T, error)
	//sql(SELECT * FROM @@table where @orderrefer=(SELECT order_id from passenger_order_forms where order_id=@orderrefer))
	GetAllOrderProductReferPOrder(orderrefer string) ([]gen.T, error)
}

type DeviceQuery interface {
	//sql(SELECT * FROM @@table where @id=(SELECT id from drivers where id=@id))
	GetAllDriverDevice(id int64) (gen.T, error)
	//sql(SELECT * FROM @@table where @id=(SELECT id from platforms where id=@id))
	GetAllPlatformDevice(id int64) (gen.T, error)
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
	g.ApplyBasic(&model.User{})
	g.ApplyInterface(func(Query) {}, &model.Passenger{},
		&model.Driver{}, &model.Factory{},
		&model.Platform{}, &model.DeviceProduct{},
		&model.Device{}, &model.Advertisement{})
	g.ApplyInterface(func(LoginQuery) {}, &model.LoginInfo{})
	g.ApplyInterface(func(CartQuery) {}, &model.PassengerCart{}, &model.DriverCart{})
	g.ApplyInterface(func(query ProductQuery) {}, &model.OrderProduct{}, &model.DeviceProduct{})
	g.ApplyInterface(func(query DeviceQuery) {}, &model.Device{})
	g.Execute()
}
