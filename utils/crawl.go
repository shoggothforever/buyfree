package utils

import (
	"encoding/base64"
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/go-resty/resty/v2"
	"golang.org/x/net/html"
	"net/http"
)

var apiurl = "https://cli.im/api/qrcode/code?text=%s&mhid=sRTBBAq/zMohMHYoLddTOqk"
var shortenurl = "https://www.mxnzp.com/api/shortlink/create"
var drinvurl = "http://bfp.shoggothy.xyz/home/%d"

func fetch(durl string) *html.Node {
	url := fmt.Sprintf(apiurl, durl)
	fmt.Println(url)
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

func shortenUrl(url string) string {
	r := resty.New()
	encodeurl := base64.StdEncoding.EncodeToString([]byte(url))
	var path = "url=%s&app_id=%s&app_secret=%s"
	var data shortres
	resp, _ := r.R().SetHeader("Content-Type", "application/json").
		SetQueryString(fmt.Sprintf(path, encodeurl, "urveqjgajrithmkd", "SEl2aHp5UTkySkFQWW5RdmJwNGRiZz09")).
		Get(shortenurl)
	Json.Unmarshal(resp.Body(), &data)
	return data.Data.ShortUrl
}

// 根据设备编号生成相应的图片url并且存储在redis数据库中
func GenerateSourceUrl(id int64) string {
	url := fmt.Sprintf(drinvurl, id)
	return shortenUrl(parseUrls(url))
}
