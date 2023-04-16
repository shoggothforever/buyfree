package dal

import (
	"buyfree/config"
	"buyfree/repo/model"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var dB *gorm.DB

func Getdb() *gorm.DB {
	return dB
}

var dsn string

func ReadPostgresSQLlinfo() {
	info := config.Reader.GetStringMapString("postgresql")
	dsn = info[config.Sqldsn]
	if dsn == "" {
		dsn = "host=localhost port=5432 user=root dbname=root password=nyarlak  sslmode=disable  TimeZone=Asia/Shanghai"
	}
}

func init() {
	ReadPostgresSQLlinfo()
	//fmt.Println(dsn)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer 日志输出的目标，前缀和日志包含的内容
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Silent, // 日志级别
			//IgnoreRecordNotFoundError: true,          // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful: false, // 禁用彩色打印
		},
	)
	var err error
	dB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
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
		dB.AutoMigrate(
			&model.FundInfo{},
			&model.LoginInfo{},
			&model.Passenger{},
			&model.PassengerCart{},
			&model.PassengerOrderForm{},
			&model.Factory{},
			&model.Driver{},
			&model.CartProduct{},
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
	// Attach Cache Option
	//opt := gcache.DefaultOption{}
	//opt.Expires = 300
	//opt.Level = option.LevelSearch
	//opt.AsyncWrite = false
	//opt.PenetrationSafe = false
	//gcache.AttachDB(DB, &opt, &option.RedisOption{Addr: "localhost:6379"})

}
