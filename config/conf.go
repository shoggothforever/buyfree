package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

const (
	Sqldsn        string = "dsn"
	Redisaddr     string = "addr"
	Redispassword string = "password"
	OK            int    = 200
	BAD           int    = 500
	FORBIDDEN     int    = 403
	REDIRECT      int    = 307
)
const (
	Role_0 string = "Factory"
	Role_1 string = "Driver"
	Role_2 string = "Passenger"
)

type MysqlParam struct {
	Host      string
	Port      int
	Database  string
	Username  string
	Password  string
	Charset   string
	ParseTime bool `toml:"parse_time"`
	Loc       string
}
type Config struct {
	DB MysqlParam `toml:"mysqldal"`
}

var Reader *viper.Viper

func Init() {
	Reader = viper.New()
	path, _ := os.Getwd()
	Reader.AddConfigPath(path + "./config")
	Reader.SetConfigName("config")
	Reader.SetConfigType("yaml")
	err := Reader.ReadInConfig() // 查找并读取配置文件
	if err != nil {              // 处理读取配置文件的错误
		logrus.Error("Read config file failed: %s \n", err)
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logrus.Info("no error in config file")
		} else {
			logrus.Error("found error in config file\n", ok)
		}
	}
}
