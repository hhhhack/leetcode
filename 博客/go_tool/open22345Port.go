/*
 *
 * 　　┏┓　　　┏┓+ +
 * 　┏┛┻━━━┛┻┓ + +
 * 　┃　　　　　　　┃
 * 　┃　　　━　　　┃ ++ + + +
 *  ████━████ ┃+
 * 　┃　　　　　　　┃ +
 * 　┃　　　┻　　　┃
 * 　┃　　　　　　　┃ + +
 * 　┗━┓　　　┏━┛
 * 　　　┃　　　┃
 * 　　　┃　　　┃ + + + +
 * 　　　┃　　　┃
 * 　　　┃　　　┃ +  神兽保佑
 * 　　　┃　　　┃    代码无bug
 * 　　　┃　　　┃　　+
 * 　　　┃　 　　┗━━━┓ + +
 * 　　　┃ 　　　　　　　┣┓
 * 　　　┃ 　　　　　　　┏┛
 * 　　　┗┓┓┏━┳┓┏┛ + + + +
 * 　　　　┃┫┫　┃┫┫
 * 　　　　┗┻┛　┗┻┛+ + + +
 *
 *
 * @Author: hhhhack
 * @Date: 2021-01-09 11:34:13
 * @LastEditTime: 2021-01-28 17:44:14
 * @LastEditors: hhhhack
 * @Description:
 * @FilePath: /code/leetcode/博客/go_tool/open22345Port.go
 * @
 */

// package main

// import (
// 	"bytes"
// 	"crypto/tls"
// 	"encoding/json"
// 	"flag"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"net/url"
// 	"strings"
// 	"sync"
// 	"time"
// 	"unsafe"

// 	"gopkg.in/ini.v1"
// )

// var (
// 	configFile *string = nil
// 	userName   *string = nil
// 	password   *string = nil
// 	ipAdderss  *string = nil
// 	operation  *string = nil
// 	httpport   *int    = nil
// )

// var waitqueue sync.WaitGroup

// func init() {
// 	configFile = flag.String("f", "", "congfig file path, you can use user@passwd:ip for every device")
// 	operation = flag.String("op", "open", "open or close ssh port")
// 	userName = flag.String("name", "admin", "consle user name")
// 	password = flag.String("pass", "admin", "consle user password")
// 	ipAdderss = flag.String("ip", "", "device ipadderss")
// 	httpport = flag.Int("port", 443, "service port")

// 	flag.Parse()
// 	log.SetFlags(log.Lshortfile)
// 	cfg, err := ini.Load(configFile)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	fmt.Println(cfg)
// }

// func sendRequest(config deviceInfo, ch chan<- int) int {
// 	loginUrl := "https://" + config.ipadderss + ":" + string(config.httpport) + "/api/ad/current-version/login"
// 	loginBody := `{"username":"` + config.userName + `","loginType":"account","password":"` + config.password + `"}`

// 	requestUrl := "https://" + config.ipadderss + ":" + string(config.httpport) + "/api/ad/current-version/sys/maintenance-mode"

// 	openBody := `{"console_service":{"lan_stat":"ENABLE","wan_stat":"DISABLE"},"troubleshooting_port":{"lan_stat":"ENABLE","wan_stat":"DISABLE","management_stat":"ENABLE","ssh_maintenance_port":22}}`
// 	closeBody := `{"console_service":{"lan_stat":"DISABLE","wan_stat":"DISABLE"},"troubleshooting_port":{"lan_stat":"DISABLE","wan_stat":"DISABLE","management_stat":"DISABLE","ssh_maintenance_port":22}}`

// 	loginRequest, err := http.NewRequest(http.MethodPost, loginUrl, bytes.NewReader([]byte(loginBody)))
// 	if err != nil {
// 		log.Printf("new a login request error ip is %v ", config.ipa)
// 		fmt.Println(err)
// 		return -1
// 	}
// 	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}} //如果需要测试自签名的证书 这里需要设置跳过证书检测 否则编译报错
// 	client := http.Client{Transport: tr, Timeout: time.Duration(10 * time.Second)}
// 	resp, err := client.Do(loginRequest)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	if resp.StatusCode >= 300 {
// 		fmt.Printf("login err")
// 		return -1
// 	}
// 	var session string
// 	cookies := strings.Split(resp.Header.Get("Set-Cookie"), ";")

// GETCOOKIE:
// 	for _, cookie := range cookies {
// 		if strings.HasPrefix(cookie, "session_id") {
// 			log.Printf("get cookie is %s ", cookie)
// 			tmpRune := []rune(cookie)
// 			for i, value := range tmpRune {
// 				if value == '=' {
// 					session = string(tmpRune[i+1:])
// 					log.Printf("get session is %s ", session)
// 					break GETCOOKIE
// 				}
// 			}
// 		}
// 	}
// 	respBytes, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return -1
// 	}
// 	respMap := make(map[string]interface{})
// 	err = json.Unmarshal(respBytes, &respMap)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	//byte数组直接转成string，优化内存
// 	params := url.Values{}
// 	params.Add("x_id", respMap["x_id"].(string))

// 	closeRequest, err := http.NewRequest(http.MethodPatch, requestUrl, bytes.NewReader([]byte(closeBody)))
// 	closeRequest.URL.RawQuery = params.Encode()
// 	if err != nil {
// 		log.Printf("new a request error")
// 		fmt.Println(err)
// 		return -1
// 	}
// 	closeRequest.AddCookie(&http.Cookie{
// 		Name:  "session_id_443",
// 		Value: session,
// 	})
// 	resp, err = client.Do(closeRequest)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	if operation != nil && *operation == "open" {

// 		openRequest, err := http.NewRequest(http.MethodPatch, requestUrl, bytes.NewReader([]byte(openBody)))
// 		openRequest.URL.RawQuery = params.Encode()
// 		if err != nil {
// 			log.Printf("new a request error")
// 			fmt.Println(err)
// 			return -1
// 		}
// 		openRequest.AddCookie(&http.Cookie{
// 			Name:  "session_id_443",
// 			Value: session,
// 		})
// 		resp, err = client.Do(openRequest)
// 		if err != nil {
// 			fmt.Println(err)
// 		}

// 		respBytes, err = ioutil.ReadAll(resp.Body)
// 		if err != nil {
// 			fmt.Println(err.Error())
// 			return -1
// 		}
// 		//byte数组直接转成string，优化内存
// 		str := (*string)(unsafe.Pointer(&respBytes))
// 		fmt.Println(*str)
// 	}
// 	waitqueue.Done()
// 	return 0
// }

// func main() {
// 	if (ipAdderss == nil || *ipAdderss == "") && (configFile == nil || *configFile == "") {
// 		fmt.Printf("argument error, please use -h for help")
// 		return
// 	}
// 	configs := make([]deviceInfo, 0)
// 	//configs = append(configs, {userName, password, ipAdderss})
// 	if configFile != nil {
// 		if parseConf.getConfig(configs) == -1 {
// 			log.Printf("read file fail")
// 			return
// 		}
// 	} else {
// 		configs = append(configs, deviceInfo{*userName, *password, *ipAdderss})
// 	}
// 	ch := make(chan int, len(configs))
// 	for _, conf := range configs {
// 		waitqueue.Add(1)
// 		go sendRequest(conf, ch)
// 	}
// 	waitqueue.Wait()

// }
