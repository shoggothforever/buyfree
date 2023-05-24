package utils

import (
	"buyfree/dal"
	"context"
	"fmt"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"math"
	"net/http"
	"time"
)

var apiurl = "https://cli.im/api/qrcode/code?text=%s&mhid=sRTBBAq/zMohMHYoLddTOqk"

// var shortenurl = "https://www.mxnzp.com/api/shortlink/create"
var shortenurl = "http://slink.shoggothy.xyz/url/create"
var drinvurl = "http://bfp.shoggothy.xyz/home/%d"
var drscanurl = "http://bfd.shoggothy.xyz/dr/devices/scan"
var shortKey = "shortenUrl:%s"
var local = "localhost:9003/"
var remote = "https://bf.shoggothy.xyz/"

func fetch(durl string) *html.Node {
	url := fmt.Sprintf(apiurl, durl)
	//fmt.Println(url)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36 Edg/111.0.1661.62")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("调用api失败")
	}
	if resp.StatusCode != 200 {
		fmt.Println("Http status code:", resp.StatusCode)
	}
	defer resp.Body.Close()
	doc, err := htmlquery.Parse(resp.Body)
	return doc
}
func parseUrls(durl string) string {
	doc := fetch(durl)
	nodes := htmlquery.Find(doc, `/html/body/div[2]/div`)
	if len(nodes) != 0 {
		url := htmlquery.FindOne(nodes[0], "./img/@src")
		return "https:" + htmlquery.InnerText(url)
	}
	return ""
}

type urldata struct {
	ShortUrl string `json:"shortUrl,omitempty"`
	Url      string `json:"url,omitempty"`
}

type shortres struct {
	Code int64   `json:"code,omitempty"`
	Msg  string  `json:"msg,omitempty"`
	Data urldata `json:"data"`
}

//type urlinfo struct {
//	Id         int    `form:"id" json:"id"`
//	UserId     int    `form:"user_id" json:"user_id"`
//	Origin     string `form:"origin" json:"origin"`
//	ShortUrl   string `form:"short" json:"short"`
//	Comment    string `form:"comment" json:"comment"`
//	StartTime  time.Time
//	ExpireTime time.Time
//}
//type shortres struct {
//	Code int64   `json:"code,omitempty"`
//	Msg  string  `json:"msg,omitempty"`
//	Data urlinfo `json:"urlinfo"`
//}

func GenShort(shortUrl string) string {
	ans := ""
	if shortUrl == "" {
		//未给定短链接, 通过当前时间的纳秒数生成新的短链接存储
		temp := time.Now().UnixNano() % int64(math.Pow(62, 6))
		ans := ""
		for {
			if temp == 0 {
				break
			}
			now := temp % 62
			if now >= 0 && now <= 25 { //generate A-Z
				ans = ans + string(65+now)
			} else if now >= 26 && now <= 51 { //generate a-z
				ans = ans + string(71+now)
			} else if now >= 52 && now <= 61 { //generate 0-9
				ans = ans + string(now-4)
			}
			temp /= 62
		}
		//return local + "shorten/" + ans
		return remote + "shorten/" + ans
	} else {
		for _, s := range shortUrl {
			if (s <= 57 && s >= 48) || (s <= 90 && s >= 65) || (s <= 122 && s >= 97) {
				ans = ans + string(s)
			}
		}
		//return local + "shorten/" + ans
		return remote + "shorten/" + ans
	}
}

//func shortenUrl(url string) string {
//	r := resty.New()
//	encodeurl := base64.StdEncoding.EncodeToString([]byte(url))
//	var path = "url=%s&app_id=%s&app_secret=%s"
//	var data shortres
//	resp, _ := r.R().SetHeader("Content-Type", "application/json").
//		SetQueryString(fmt.Sprintf(path, encodeurl, "urveqjgajrithmkd", "SEl2aHp5UTkySkFQWW5RdmJwNGRiZz09")).
//		Get(shortenurl)
//	Json.Unmarshal(resp.Body(), &data)
//	return data.Data.ShortUrl
//}

//func shortenUrl(url string) string {
//	r := resty.New()
//	var data shortres
//	resp, _ := r.R().SetHeader("Content-Type", "application/json").
//		SetFormData(map[string]string{"origin": url, "commen": "buyfree device query", "short": ""}).
//		Get(shortenurl)
//	Json.Unmarshal(resp.Body(), &data)
//	return data.Data.ShortUrl
//}

func GetShortenKey(short string) string {
	return fmt.Sprintf(shortKey, short)
}

// 根据设备编号生成相应的图片url并且存储在redis数据库中
func GenerateSourceUrl(id int64) string {
	url := fmt.Sprintf(drinvurl, id)
	origin := parseUrls(url)
	short := GenShort("")
	key := GetShortenKey(short)
	rdb := dal.Getrdb()
	rdb.Do(context.TODO(), "set", key, origin)
	return short
}

//	func GenerateSourceUrl1(id int64) string {
//		url := fmt.Sprintf(drinvurl, id)
//		//return parseUrls(url)
//		return shortenUrl(parseUrls(url))
//	}
func GenerateScanUrl() string {
	url := fmt.Sprintf(drscanurl)
	//return shortenUrl(parseUrls(url))
	return parseUrls(url)
}
