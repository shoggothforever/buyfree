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
	// SELECT * FROM @@table WHERE id=@uuid
	GetByID(id int64) (gen.T, error)
}
type LoginQuery interface {
	// SELECT * FROM @@table WHERE user_id=@@uid and password=@@psw
	GetByUidAndPsw(uid string, psw string) (gen.T, error)
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
	g.ApplyInterface(func(Query) {}, &model.Passenger{}, &model.Driver{}, &model.Factory{}, &model.Platform{}, &model.DeviceProduct{}, &model.DEVICE{})
	g.ApplyInterface(func(query LoginQuery) {}, &model.LoginInfo{})
	g.Execute()
}
