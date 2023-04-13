package config

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
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
	Reader *viper.Viper
)

type MConfigs struct {
	QINIU_AK         string `json:"QINIU_AK,omitempty"`
	QINIU_SK         string `json:"QINIU_SK,omitempty"`
	QINIU_BK         string `json:"QINIU_BK,omitempty"`
	APPID            string `json:"APPID,omitempty"`
	APPSECRET        string `json:"APPSECRET,omitempty"`
	GRANTTYPE        string `json:"GRANTTYPE,omitempty"`
	Mendpoint        string `json:"mendpoint,omitempty"`
	MAccessKeyID     string `json:"MAccessKeyID,omitempty"`
	MSecretAccessKey string `json:"MSecretAccessKey,omitempty"`
}

var Mcfg MConfigs
var D = flag.Bool("D", false, "默认为release，true为debug")

func init() {
	flag.Parse()

	gin.SetMode(gin.ReleaseMode)
	Mcfg.GRANTTYPE = "authorization_code"
	Reader = viper.New()
	//path, _ := os.Getwd()
	//path := "d:/desktop/pr/buyfree"
	path := "/www/wwwroot/bf.shoggothy.xyz/buyfree"
	fmt.Println("config文件读取路径", path)
	Reader.AddConfigPath(path + "/config")
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
	Mcfg.QINIU_AK = info["ak"]
	//if Mcfg.QINIU_AK == "" {
	//	Mcfg.QINIU_AK = "krCw1c0mo4uyEbHNArbXQR6xpdz6QLamc99iAu_-"
	//}
	Mcfg.QINIU_SK = info["sk"]
	//if Mcfg.QINIU_SK == "" {
	//	Mcfg.QINIU_SK = "XHY438HM9qjh3c1uIOVmzdO-bjlLTSYUZzKEY7_4"
	//}
	Mcfg.QINIU_BK = info["bk"]
	//if Mcfg.QINIU_BK == "" {
	//	Mcfg.QINIU_BK = "bfcloud"
	//}
	winfo := Reader.GetStringMapString("weixinapp")
	Mcfg.APPID = winfo["appid"]
	//if Mcfg.APPID == "" {
	//	Mcfg.APPID = "wxd776834423fadf04"
	//}
	Mcfg.APPSECRET = winfo["appsecret"]
	//if Mcfg.APPSECRET == "" {
	//	Mcfg.APPSECRET = "00a3239022c4146ec4c3209792539c0b"
	//}
	minioinfo := Reader.GetStringMapString("minio")
	Mcfg.Mendpoint = minioinfo["endpoint"]
	Mcfg.MAccessKeyID = minioinfo["accessKeyID"]
	Mcfg.MSecretAccessKey = minioinfo["secreatAccessKey"]
	//logger.Loger.Info(Mcfg)
}
