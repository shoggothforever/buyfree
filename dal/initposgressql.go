package dal

import (
	"buyfree/config"
	"buyfree/repo/model"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var dB *gorm.DB

type casbinModel struct {
	Enforcer *casbin.Enforcer
	Adapter  *gormadapter.Adapter
}

var casbinmodel casbinModel

func Getdb() *gorm.DB {
	return dB
}
func GetCasbinModel() *casbinModel {
	return &casbinmodel
}

var dsn string

func ReadPostgresSQLlinfo() {
	info := config.Reader.GetStringMapString("postgresql")
	dsn = info[config.Sqldsn]
}

func init() {
	ReadPostgresSQLlinfo()
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
		logrus.WithFields(logrus.Fields{"error": err}).Error("Open PostgresSQL")
		return
	}

	casbinmodel.Adapter, err = gormadapter.NewAdapterByDB(dB)
	if err != nil {
		logrus.WithFields(logrus.Fields{"error": err}).Error("load RABC model")
		return
	}
	casbinmodel.Enforcer, err = casbin.NewEnforcer("repo/model/design/rbac_model.conf", casbinmodel.Adapter)
	if err != nil {
		logrus.WithFields(logrus.Fields{"error": err}).Error("load RBAC model")
		return
	}
	err = casbinmodel.Adapter.SavePolicy(casbinmodel.Enforcer.GetModel())
	if err != nil {
		logrus.WithFields(logrus.Fields{"error": err}).Error("save policy")
		return
	}
	err = casbinmodel.Adapter.LoadPolicy(casbinmodel.Enforcer.GetModel())
	if err != nil {
		logrus.WithFields(logrus.Fields{"error": err}).Error("load policy")
		return
	}
	casbinmodel.Enforcer.LoadPolicy()
	//Create TABLES
	{
		dB.AutoMigrate(
			&model.FundInfo{},
			&model.Platform{},
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

}
