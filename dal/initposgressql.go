package dal

import (
	"buyfree/config"
	"buyfree/repo/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
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
	DB, err = gorm.Open("postgres", dsn)
	if err != nil {
		logrus.WithFields(logrus.Fields{"error": err}).Error("Open PostgresSQL failed")
	} else {
		logrus.Info("Open postgresSQL successfully")
	}
	DB.AutoMigrate(&model.Passenger{}, &model.PassengerCart{}, &model.PassengerOrderForm{},
		&model.OrderProduct{})
	//&model.Passenger{}, &model.PassengerCart{}, &model.PassengerOrderForm{}, &model.Driver{}, model.Factory{}, model.Platform{}, model.VdMachine{}
}
