package dal

import (
	"buyfree/config"
	"buyfree/repo/model"
	"gorm.io/driver/postgres"
	//"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
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
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		CreateBatchSize:        1000,
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	})
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
			&model.Device{},
			&model.DriverCart{},
			&model.DriverOrderForm{},
			&model.DeviceProduct{},
			&model.Advertisement{},
		)

	}
}
