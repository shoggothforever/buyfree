package dal

import (
	"buyfree/config"
	"buyfree/repo/model"
	"gorm.io/driver/postgres"

	//"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
	"gorm.io/gen"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Getdb() *gorm.DB {
	return DB
}

var dsn string

func ReadPostgresSQLlinfo() {
	info := config.Reader.GetStringMapString("postgresql")
	dsn = info[config.Sqldsn]
}

func InitPostgresSQL() {
	ReadPostgresSQLlinfo()
	//fmt.Println(dsn)
	var err error
	//DB, err = gorm.Open("postgres", dsn)
	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		logrus.WithFields(logrus.Fields{"error": err}).Error("Open PostgresSQL failed")
	} else {
		logrus.Info("Open postgresSQL successfully")
	}
	//Create PassengerEnd TABLES
	{
		DB.AutoMigrate(
			&model.LoginInfo{},
			&model.Passenger{},
			&model.PassengerCart{},
			&model.PassengerOrderForm{},
			&model.Factory{},
			&model.Platform{},
			&model.Driver{},
			&model.OrderProduct{},
			&model.DEVICE{},
			&model.DriverCart{},
			&model.DriverOrderForm{},
			&model.DeviceProduct{},
			&model.Advertisement{},
		)

	}
	////Create DriverEnd Tables
	//{
	//	DB.AutoMigrate(
	//
	//
	//
	//}
	//
	////Create FactoryEnd TABLES
	//{
	//	DB.AutoMigrate()
	//}
	////Create PlatFormEnd Tables
	//{
	//	DB.AutoMigrate(
	//
	//}

	g := gen.NewGenerator(gen.Config{
		OutPath: "../repo/gen",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})
	g.UseDB(DB)
}
