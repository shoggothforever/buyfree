package utils

import (
	"buyfree/config"
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

const uploadexpiretime = 7200

var mac *qbox.Mac
var putPolicy storage.PutPolicy
var upToken string
var cfg storage.Config

func init() {
	mac = qbox.NewMac(config.QINIU_AK, config.QINIU_SK)
	//putPolicy = storage.PutPolicy{Scope: config.QINIU_BK, Expires: 7200}
	//upToken = putPolicy.UploadToken(mac)
	cfg.Region = &storage.ZoneHuanan
	cfg.UseCdnDomains = true
}

func UpLoad(filepath, key string) storage.PutRet {
	putPolicy = storage.PutPolicy{Scope: config.QINIU_BK,
		CallbackURL:      "bf.shoggothy.xyz/",
		CallbackBody:     `{key:$(key),hash:$(etcd),fsize:$(fsize),bucket:$(bucket),name:$(x:name),notify:$(x:notify)`,
		CallbackBodyType: "application/json",

		Expires: 7200}
	upToken = putPolicy.UploadToken(mac)
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name":   "lookup",
			"x:notify": "dida···",
		},
	}
	err := formUploader.PutFile(context.Background(), &ret, upToken, key, filepath, &putExtra)
	if err != nil {
		fmt.Println(err)
		return ret
	}
	fmt.Println(ret.Key, ret.Hash)
	return ret
}
