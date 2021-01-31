package main

import (
	"crypto/tls"
	"encoding/base64"
	"flag"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	url2 "net/url"
	"time"
)
type Config struct {
	Email string
	Key string
}
var configs Config
func Readconfig(){
	bytes,err := ioutil.ReadFile("./unit.config")
	if err != nil{
		fmt.Print("配置文件读取错误")
		return
	}
	configs.Email = gjson.GetBytes(bytes,"Email").String()
	configs.Key = gjson.GetBytes(bytes,"Key").String()


}
func main(){
	var gjc string
	var size string
	flag.StringVar(&gjc,"g","","请输入你要输入的关键词")
	flag.StringVar(&size,"s","","请输入你要最大删除多少条")
	flag.Parse()
	Readconfig()
	urls := Ulrc(gjc,size)
	bodys := Fofarequest(urls)
	cas := Bodycl(bodys)
	for _,v := range cas{
		fmt.Print(v.Array()[0])
		fmt.Print("\n")
	}
}
func Bodycl(body string)[]gjson.Result{
	czx := gjson.Get(body,"results")
	return czx.Array()
}
func Ulrc(gjc,size string)string{
	url := "https://fofa.so/api/v1/search/all?email="+configs.Email+"&key="+configs.Key+"&qbase64="+url2.QueryEscape(base64.StdEncoding.EncodeToString([]byte(gjc)))+"&size="+size
	return url
}
func Fofarequest(urlc string)string{
	var body string
	c := colly.NewCollector()
	c.WithTransport(&http.Transport{
		MaxIdleConnsPerHost:   10,
		ResponseHeaderTimeout: time.Second * time.Duration(5),
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true}, //忽略https验证
	})
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_1) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/12.0.1 Safari/605.1.15")
		r.Headers.Set("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")

	})
	c.OnResponse(func(r *colly.Response) {
		body += string(r.Body)
	})
	c.Visit(urlc)
	return body
}