package test

import (
	"buyfree/repo/gen"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
)

func TestGen(t *testing.T) {
	dsn := "host=localhost port=5432 user=bf dbname=bfdb password=bf123  sslmode=disable  TimeZone=Asia/Shanghai"
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
	gen.SetDefault(DB)
	u, _ := gen.Platform.GetByID(238043646732013568)
	fmt.Println(u.ID)
	l, _ := gen.LoginInfo.GetByNameAndPsw(238043443408932864, "338951b7e7607b65262fb051e7804d91")
	fmt.Println(l.UserID)
}
