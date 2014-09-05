package main

import "fmt"
import (
	"os"
	"net/http"
	"log"
	"regexp"
	"strings"
//	"io/ioutil"
	"net/url"
	"compress/gzip"
	"io"
)

func Std_post (reqest http.Request) (string ){
	client := &http.Client{}


	reqest.Header.Set("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	reqest.Header.Set("Accept-Charset","GBK,utf-8;q=0.7,*;q=0.3")
	reqest.Header.Set("Accept-Encoding","gzip,deflate,sdch")
	reqest.Header.Set("Accept-Language","zh-CN,zh;q=0.8")
	reqest.Header.Set("Cache-Control","max-age=0")
	reqest.Header.Set("Connection","keep-alive")
	reqest.Header.Set("User-Agent","Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/37.0.2062.94 Safari/537.36")
	reqest.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	var body string
	response,_ := client.Do(&reqest)
	if response.StatusCode == 200 {
		response.Header.Get("Date")
		reader, _ := gzip.NewReader(response.Body)
		for {
			buf := make([]byte, 1024)
			n, err := reader.Read(buf)

			if err != nil && err != io.EOF {
				panic(err)
			}

			if n == 0 {
				break
			}
			body += string(buf)
		}

	}

	return body
}

func main() {
	Query_site := os.Args[1]
	Query_parameter := os.Args[2]
//	fmt.Printf("%s\n",Query_parameter)
//	fmt.Printf("%s\n",Query_site)


	Query_url1 := "http://ubtamator.github.io/acts/t.txt"
	resp1,err1 := http.Get(Query_url1)
	if err1 != nil{
		//handle error
		fmt.Println(err1)
		log.Fatal(err1)
	}
	Query_res1 := ""
	if resp1.StatusCode == http.StatusOK{
		buf := make([]byte, 1024)
		for {
			n,_ := resp1.Body.Read(buf)
			if 0 == n {break}
			Query_res1 += string(buf[:n])
		}
	}
	defer resp1.Body.Close()

	software_parameter := strings.Split(Query_res1,"||")
	if software_parameter[0] != "1" {
		fmt.Println("ERROR out of date")
		os.Exit(0)
	}


	Query_url := software_parameter[1] + Query_site
	resp,err := http.Get(Query_url)
	if err != nil{
		//handle error
		fmt.Println(err)
		log.Fatal(err)
	}
	Query_res := ""
	if resp.StatusCode == http.StatusOK{
		buf := make([]byte, 1024)
		for {
			n,_ := resp.Body.Read(buf)
			if 0 == n {break}
			Query_res += string(buf[:n])
		}
	}
	defer resp.Body.Close()

	Query_report := ""

	reg := regexp.MustCompile(`<li style=" width:130px">网站名称<br /><font>.+\r`)
	temp := reg.FindAllString(Query_res, -1)
	网站名称 := strings.Replace(temp[0],`<li style=" width:130px">网站名称<br /><font>`,``,-1)
	Query_report += "网站名称 ： "+网站名称 + "\n"

	reg = regexp.MustCompile(`<li style=" width:140px">网站首页网址<br /><font>.+\r`)
	temp = reg.FindAllString(Query_res, -1)
	网站首页网址 := strings.Replace(temp[0],`<li style=" width:140px">网站首页网址<br /><font>`,``,-1)
	Query_report += "网站首页网址 ： "+网站首页网址 + "\n"

	reg = regexp.MustCompile(`<li style=" width:230px">主办单位名称<br /><font>.+\r`)
	temp = reg.FindAllString(Query_res, -1)
	主办单位名称 := strings.Replace(temp[0],`<li style=" width:230px">主办单位名称<br /><font>`,``,-1)
	Query_report += "主办单位名称 ： "+主办单位名称 + "\n"

	reg = regexp.MustCompile(`<li>主办单位性质<br /><font>.+\r`)
	temp = reg.FindAllString(Query_res, -1)
	主办单位性质 := strings.Replace(temp[0],`<li>主办单位性质<br /><font>`,``,-1)
	Query_report += "主办单位性质 ： "+主办单位性质 + "\n"

	reg = regexp.MustCompile(`<li>网站备案/许可证号<br /><font><a style="color:green;cursor:hand".+\r`)
	temp = reg.FindAllString(Query_res, -1)
	sub_reg :=regexp.MustCompile(`\WICP\W.+$`)
	temp = sub_reg.FindAllString(temp[0],-1)
	网站备案许可证号 := temp[0]
	Query_report += "网站备案许可证号 ： "+网站备案许可证号 + "\n"

//	Std_post(http.NewRequest("POST","http://www.alexa.cn/inc.php",url.Values{"url":{"baidu.com"}}))




//	fmt.Println(http.PostForm("http://www.alexa.cn/inc.php",url.Values{"url":{"baidu.com"}}))

	v := url.Values{}
	v.Set("url",Query_site)
	request,_ := http.NewRequest("POST",software_parameter[2],strings.NewReader(v.Encode()))
	inc := Std_post(*request)

	v = url.Values{}
	v.Set("url", Query_site)
	v.Add("sig", "940a68c2b16d6288dc9988b083110c70")
	v.Add("keyt", "1409636251")
	request,_ = http.NewRequest("POST","http://www.alexa.cn/api0523.php",strings.NewReader(v.Encode()))
	api := Std_post(*request)

	fmt.Println(api)

	para := strings.Split(inc,"||")
	GooglePR := para[0]
	服务器IP := para[1]
	IP所在地 := para[2]
	服务器类型 := para[3]
	协议类型 := para[4]
	页面类型 := para[5]

	Query_report += "GooglePR ： "+GooglePR + "\n"
	Query_report += "服务器IP ： "+服务器IP + "\n"
	Query_report += "IP所在地 ： "+IP所在地 + "\n"
	Query_report += "服务器类型 ： "+服务器类型 + "\n"
	Query_report += "协议类型 ： "+协议类型 + "\n"
	Query_report += "页面类型 ： "+页面类型 + "\n"



	if Query_parameter == "all" {fmt.Println(Query_report)}
	if Query_parameter == "GooglePR" {fmt.Println(GooglePR)}
	if Query_parameter == "服务器IP" {fmt.Println(服务器IP)}
	if Query_parameter == "IP所在地" {fmt.Println(IP所在地)}
	if Query_parameter == "服务器类型" {fmt.Println(服务器类型)}
	if Query_parameter == "协议类型" {fmt.Println(协议类型)}
	if Query_parameter == "页面类型" {fmt.Println(页面类型)}
	if Query_parameter == "网站名称" {fmt.Println(网站名称)}
	if Query_parameter == "网站首页网址" {fmt.Println(网站首页网址)}
	if Query_parameter == "主办单位名称" {fmt.Println(主办单位名称)}
	if Query_parameter == "主办单位性质" {fmt.Println(主办单位性质)}
	if Query_parameter == "网站备案许可证号" {fmt.Println(网站备案许可证号)}




}
