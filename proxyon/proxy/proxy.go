package proxy

import (
	"Proxyone/initone"
	"crypto/tls"
	"fmt"
	"github.com/elazarl/goproxy"
	"github.com/gocolly/colly"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type IPlb struct {
	ProxyIp string
	Timeone int
}
var SS IPlb
func init(){
	Xmapiinit()
}
func Xmjtproks(){
	proxy := goproxy.NewProxyHttpServer()
	proxy.Tr = &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse("http://"+ SS.ProxyIp)
		}}
	proxy.OnRequest().DoFunc(
		func(r *http.Request,ctx *goproxy.ProxyCtx)(*http.Request,*http.Response) {
			ts := time.Now().Add(time.Minute * -5)
			if int(ts.UnixNano()) > SS.Timeone {
				fmt.Println("准备更换IP")
				Xmapiinit()
			}
			return r,nil
		})
	fmt.Println("这次换的IP是："+SS.ProxyIp)
	proxy.ConnectDial = proxy.NewConnectDialToProxy("http://"+SS.ProxyIp)
	eerr := proxy.ConnectDial
	if eerr != nil{
		fmt.Println(eerr)
		fmt.Println("连接不上去了")
	}
	proxy.Verbose = true
	http.ListenAndServe(":8082", proxy)
}
func Xmapiinit(){
	var body string
	c := colly.NewCollector()
	c.WithTransport(&http.Transport{
		MaxIdleConnsPerHost:   10,
		ResponseHeaderTimeout: time.Second * time.Duration(5),
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true}, //忽略https验证
	})
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_1) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/12.0.1 Safari/605.1.15")
		r.Headers.Set("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
		r.Headers.Set("Accept-Encoding","gzip, deflate, br")
		r.Headers.Set("Accept-Language","zh-CN,zh;q=0.9")
		r.Headers.Set("Connection","keep-alive")
	})
	c.OnResponse(func(r *colly.Response) {
		boyds := string(r.Body)
		boyds = strings.Replace(boyds,"","",-1)
		boyds = strings.Replace(boyds,"\n","",-1)
		boyds = strings.Replace(boyds,"\r","",-1)
		body+=boyds
	})
	c.Visit(initone.InitConfig.Xmapi)
	SS.ProxyIp = body
	SS.Timeone = int(time.Now().UnixNano())
}