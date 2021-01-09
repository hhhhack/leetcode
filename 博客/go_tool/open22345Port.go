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
 * @LastEditTime: 2021-01-09 11:34:16
 * @LastEditors: hhhhack
 * @Description:
 * @FilePath: /code/leetcode/博客/go_tool/open22345Port.go
 * @
 */

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
 * @Date: 2020-09-10 17:29:25
 * @LastEditTime: 2021-01-08 16:31:13
 * @LastEditors: hhhhack
 * @Description:
 * @FilePath: \undefinedd:\project\go\open22345port.go
 * @
 */

package main

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
	"unsafe"

	"gopkg.in/ini.v1"
)

type deviceInfo struct {
	userName  string
	password  string
	ipadderss string
	httpport  int
}

var (
	configFile *string = nil
	userName   *string = nil
	password   *string = nil
	ipAdderss  *string = nil
	operation  *string = nil
	httpport   *int    = nil
)

var waitqueue sync.WaitGroup

func init() {
	configFile = flag.String("f", "", "congfig file path, you can use user@passwd:ip for every device")
	operation = flag.String("op", "open", "open or close ssh port")
	userName = flag.String("name", "admin", "consle user name")
	password = flag.String("pass", "admin", "consle user password")
	ipAdderss = flag.String("ip", "", "device ipadderss")
	httpport = flag.Int("port", 443, "service port")

	flag.Parse()
	log.SetFlags(log.Lshortfile)
	cfg, err := ini.Load(configFile)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(cfg)
}

//func doSend(req http.request)

func sendRequest(config deviceInfo, ch chan<- int) int {
	// defer func() {
	// 	ch <- 0
	// }()
	// var rsp io.Reader
	loginUrl := "https://" + config.ipadderss + ":" + string(config.httpport) + "/api/ad/current-version/login"
	loginBody := `{"username":"` + config.userName + `","loginType":"account","password":"` + config.password + `"}`

	requestUrl := "https://" + config.ipadderss + ":" + string(config.httpport) + "/api/ad/current-version/sys/maintenance-mode"

	openBody := `{"console_service":{"lan_stat":"ENABLE","wan_stat":"DISABLE"},"troubleshooting_port":{"lan_stat":"ENABLE","wan_stat":"DISABLE","management_stat":"ENABLE","ssh_maintenance_port":22}}`
	closeBody := `{"console_service":{"lan_stat":"DISABLE","wan_stat":"DISABLE"},"troubleshooting_port":{"lan_stat":"DISABLE","wan_stat":"DISABLE","management_stat":"DISABLE","ssh_maintenance_port":22}}`

	loginRequest, err := http.NewRequest(http.MethodPost, loginUrl, bytes.NewReader([]byte(loginBody)))
	if err != nil {
		log.Printf("new a login request error")
		fmt.Println(err)
		return -1
	}
	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}} //如果需要测试自签名的证书 这里需要设置跳过证书检测 否则编译报错
	client := http.Client{Transport: tr, Timeout: time.Duration(10 * time.Second)}
	resp, err := client.Do(loginRequest)
	if err != nil {
		fmt.Println(err)
	}
	if resp.StatusCode >= 300 {
		fmt.Printf("login err")
		return -1
	}
	var session string
	cookies := strings.Split(resp.Header.Get("Set-Cookie"), ";")

GETCOOKIE:
	for _, cookie := range cookies {
		if strings.HasPrefix(cookie, "session_id") {
			log.Printf("get cookie is %s ", cookie)
			tmpRune := []rune(cookie)
			for i, value := range tmpRune {
				if value == '=' {
					session = string(tmpRune[i+1:])
					log.Printf("get session is %s ", session)
					break GETCOOKIE
				}
			}
		}
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}
	respMap := make(map[string]interface{})
	err = json.Unmarshal(respBytes, &respMap)
	if err != nil {
		log.Println(err)
	}
	//byte数组直接转成string，优化内存
	params := url.Values{}
	params.Add("x_id", respMap["x_id"].(string))

	//requestUrl += ("?x_id=" + respMap["x_id"].(string))
	closeRequest, err := http.NewRequest(http.MethodPatch, requestUrl, bytes.NewReader([]byte(closeBody)))
	closeRequest.URL.RawQuery = params.Encode()
	if err != nil {
		log.Printf("new a request error")
		fmt.Println(err)
		return -1
	}
	closeRequest.AddCookie(&http.Cookie{
		Name:  "session_id_443",
		Value: session,
	})
	resp, err = client.Do(closeRequest)
	if err != nil {
		fmt.Println(err)
	}
	if operation != nil && *operation == "open" {

		openRequest, err := http.NewRequest(http.MethodPatch, requestUrl, bytes.NewReader([]byte(openBody)))
		openRequest.URL.RawQuery = params.Encode()
		if err != nil {
			log.Printf("new a request error")
			fmt.Println(err)
			return -1
		}

		//request.Header.Add("Authorization", "Basic "+base64.URLEncoding.EncodeToString([]byte(config.userName+":"+config.password)))
		openRequest.AddCookie(&http.Cookie{
			Name:  "session_id_443",
			Value: session,
		})
		resp, err = client.Do(openRequest)
		if err != nil {
			fmt.Println(err)
		}

		respBytes, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err.Error())
			return -1
		}
		//byte数组直接转成string，优化内存
		str := (*string)(unsafe.Pointer(&respBytes))
		fmt.Println(*str)
	}
	waitqueue.Done()
	return 0
}

func parseConf(conf string) deviceInfo {
	if len(conf) == 0 {
		log.Printf("conf is nil")
		return deviceInfo{"", "", ""}
	}
	if strings.Index(conf, ":") != -1 {
		confs := strings.Split(conf, ":")
		fmt.Println("conf is :", confs)
		if len(confs) != 3 {
			log.Printf("conf is illagal")
			return deviceInfo{"", "", ""}
		}
		return deviceInfo{strings.TrimSpace(confs[0]), strings.TrimSpace(confs[1]), strings.TrimSpace(confs[2])}
	}
	return deviceInfo{"admin", "admin", strings.TrimSpace(conf)}
}

func getConfig(configs []deviceInfo) int {
	f, err := os.Open(*configFile)
	if err != nil {
		fmt.Printf("open file fail")
		return -1
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n') //注意是字符
		//fmt.Printf("read line is %s \n", line)
		if err == io.EOF {
			if len(line) != 0 {
				configs = append(configs, parseConf(line))
			}
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return -1
		}
		configs = append(configs, parseConf(line))
	}
	return 0
}

func main() {
	if (ipAdderss == nil || *ipAdderss == "") && (configFile == nil || *configFile == "") {
		fmt.Printf("argument error, please use -h for help")
		return
	}
	configs := make([]deviceInfo, 0)
	//configs = append(configs, {userName, password, ipAdderss})
	if configFile != nil {
		if getConfig(configs) == -1 {
			log.Printf("read file fail")
			return
		}
	} else {
		configs = append(configs, deviceInfo{*userName, *password, *ipAdderss})
	}
	//context.BackGround
	ch := make(chan int, len(configs))
	for _, conf := range configs {
		waitqueue.Add(1)
		//fmt.Println(conf)
		go sendRequest(conf, ch)
	}
	waitqueue.Wait()
	for i := 0; i < len(configs); i++ {
		<-ch
	}
	//conf := deviceInfo{*userName, *password, *ipAdderss}
	//sendRequest(conf)
}
