package dal

import (
	"buyfree/config"
	"buyfree/repo/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"

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
	//dsn = "host=localhost port=5432 user=root dbname=root password=nyarlak  sslmode=disable  TimeZone=Asia/Shanghai"
}

func init() {
	ReadPostgresSQLlinfo()
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Silent, // 日志级别
			//IgnoreRecordNotFoundError: true,          // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful: false, // 禁用彩色打印
		},
	)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		CreateBatchSize:        1000,
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
		Logger:                 newLogger,
	})
	if err != nil {
		logrus.WithFields(logrus.Fields{"error": err}).Error("Open PostgresSQL failed")
	} else {
		logrus.Info("Open postgresSQL successfully")
	}
	//Create TABLES
	{
		DB.AutoMigrate(
			&model.BankCardInfo{},
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
			&model.FactoryProduct{},
			&model.Advertisement{},
			&model.Ad_Device{},
		)

	}
}
