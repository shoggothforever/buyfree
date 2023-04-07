package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	Sqldsn        string = "dsn"
	Redisaddr     string = "addr"
	Redispassword string = "password"
	UploadPath    string = ""
	ACCESS_KEY    string = "******EA09VCy5EfN_*******************"
	SECRET_KEY    string = "******-yvwcYwImN6F*******************"
	BUCKET        string = "bucket"
	OK            int    = 200
	BAD           int    = 500
	FORBIDDEN     int    = 403
	REDIRECT      int    = 307
)
const (
	Role_1 string = "Driver"
	Role_2 string = "Factory"
	Role_3 string = "Passenger"
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

var (
	Reader    *viper.Viper
	QINIU_AK  string
	QINIU_SK  string
	QINIU_BK  string
	APPID     string
	APPSECRET string
	GRANTTYPE = "authorization_code"
)

func init() {
	Reader = viper.New()
	//path, _ := os.Getwd()
	path := "d:/desktop/pr/buyfree"
	//path := "/www/wwwroot/bf.shoggothy.xyz/buyfree"
	fmt.Println("config文件读取路径", path)
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
	info := Reader.GetStringMapString("qiniu")
	QINIU_AK = info["ak"]
	if QINIU_AK == "" {
		QINIU_AK = "krCw1c0mo4uyEbHNArbXQR6xpdz6QLamc99iAu_-"
	}
	QINIU_SK = info["sk"]
	if QINIU_SK == "" {
		QINIU_SK = "XHY438HM9qjh3c1uIOVmzdO-bjlLTSYUZzKEY7_4"
	}
	QINIU_BK = info["bk"]
	if QINIU_BK == "" {
		QINIU_BK = "bfcloud"
	}
	winfo := Reader.GetStringMapString("weixinapp")
	APPID = winfo["appid"]
	if APPID == "" {
		APPID = "wxd776834423fadf04"
	}
	APPSECRET = winfo["appsecret"]
	if APPSECRET == "" {
		APPSECRET = "00a3239022c4146ec4c3209792539c0b"
	}
}
