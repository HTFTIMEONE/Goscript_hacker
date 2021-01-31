package httpd

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)
func Redfile(filename string)(listone []string){
	var filelist []string
	fi,err := os.Open(filename)
	if err != nil {
		fmt.Printf("读取文件错误")
		return
	}
	br := bufio.NewReader(fi)
	for {
		a,_,c := br.ReadLine()
		if c == io.EOF{
			break
		}
		filelist = append(filelist,string(a))
	}
	return filelist
}
func Reqs(urls,host string)(bool,string){
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	reqest, err := http.NewRequest("GET", urls, nil)
	reqest.Host=host
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_1) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/12.0.1 Safari/605.1.15")
	reqest.Header.Add("Accept-Language","zh-CN,zh;q=0.9")
	reqest.Header.Add("Connection","keep-alive")

	if err != nil {
		fmt.Errorf("error")
		return false,"no"
	}
	body, errs := client.Do(reqest)
	if errs != nil {
		return false,"no"
	}
	defer body.Body.Close()
	htmltext,err :=ioutil.ReadAll(body.Body)
	if (body.StatusCode == 500 || body.StatusCode == 502 || body.StatusCode == 302){
		return false,"0"
	}else {
		title := Qtitle(string(htmltext))
		return true,title
	}
}
func Qtitle(body string) string  {
	titles := ""
	buffer := bytes.NewBufferString(body)
	bodysa := io.Reader(buffer)
	doc, err := goquery.NewDocumentFromReader(bodysa)
	if err != nil {
		fmt.Println(err)
		return "取标题错误"
	}
	doc.Find("html").Each(func(i int, s *goquery.Selection) {
		title := s.Find("title").Text()
		if title == "" {
			return
		}
		titles += title+"\n"
	})
	return titles
}