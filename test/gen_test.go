package test

import (
	"buyfree/repo/gen"
	"fmt"
	uuid "github.com/google/uuid"
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
	id, _ := uuid.Parse("a870e804-cf1e-3dc3-1190-5726a7d46039")
	u, _ := gen.Passenger.GetByUUID(id)
	fmt.Println(u.ID)
	l, _ := gen.LoginInfo.GetByUidAndPsw(u.ID, "123")
	fmt.Println(l.UserID)
}
