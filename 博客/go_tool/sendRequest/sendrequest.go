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
 * @Date: 2021-01-28 17:45:21
 * @LastEditTime: 2021-01-29 17:31:02
 * @LastEditors: hhhhack
 * @Description:
 * @FilePath: \go_tool\sendRequest\sendrequest.go
 * @
 */

package sendRequest

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"initConf"
	"io/ioutil"
	"log"
	"logger"
	"net/http"
	"net/url"
	"strings"
	"time"
	"unsafe"

	"go.uber.org/zap"
)

/**
 * @SendRequest
 * @description: 发送请求
 * @param {*initConf.DeviceInfo} config
 * @return {*}
 */
func SendRequest(config *initConf.DeviceInfo) int {
	if config == nil {
		return -1
	}
	defer initConf.Wg.Done()
	loginUrl := fmt.Sprintf("https://%s:%d/api/ad/current-version/login", config.Ipadderss, config.Httpport)
	logger.MyLogger.Info("init login url",
		zap.String("loginUrl", loginUrl),
	)
	bodyfmt := `{"username":"%s", "loginType":"account", "password":"%s"}`
	loginBody := fmt.Sprintf(bodyfmt, config.UserName, config.Password)
	logger.MyLogger.Info("init login body",
		zap.String("loginbody", loginBody),
	)
	requestUrl := fmt.Sprintf("https://%s:%d/api/ad/current-version/sys/maintenance-mode", config.Ipadderss, config.Httpport)

	openBody := `{"troubleshooting_port":{"lan_stat":"ENABLE","wan_stat":"DISABLE","management_stat":"ENABLE","ssh_maintenance_port":22}}`
	closeBody := `{"troubleshooting_port":{"lan_stat":"DISABLE","wan_stat":"DISABLE","management_stat":"DISABLE","ssh_maintenance_port":22}}`

	loginRequest, err := http.NewRequest(http.MethodPost, loginUrl, bytes.NewReader([]byte(loginBody)))
	if err != nil {
		log.Printf("new a login request error ip is %v ", config.Ipadderss)
		return -1
	}
	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}} //如果需要测试自签名的证书 这里需要设置跳过证书检测 否则编译报错
	client := http.Client{Transport: tr, Timeout: time.Duration(10 * time.Second)}
	var resp *http.Response
	go func() {
		resp, err = client.Do(loginRequest)
		if err != nil {
			fmt.Println(err)
		}
	}()

	if resp.StatusCode >= 300 {
		logger.MyLogger.Fatal("login err",
			zap.Int("Status_Code", resp.StatusCode),
		)
		return -1
	}
	var session string
	var cookie string
	cookies := strings.Split(resp.Header.Get("Set-Cookie"), ";")

GETCOOKIE:
	for _, cookie = range cookies {
		if strings.HasPrefix(cookie, "session_id") {
			log.Printf("get cookie is %s ", cookie)
			tmpRune := []rune(cookie)
			for i, value := range tmpRune {
				if value == '=' {
					session = string(tmpRune[i+1:])
					cookie = string(tmpRune[:i])
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

	logger.MyLogger.Info("x_id is ",
		zap.String("x_id", respMap["x_id"].(string)),
	)
	//byte数组直接转成string，优化内存
	params := url.Values{}
	params.Add("x_id", respMap["x_id"].(string))

	closeRequest, err := http.NewRequest(http.MethodPatch, requestUrl, bytes.NewReader([]byte(closeBody)))
	closeRequest.URL.RawQuery = params.Encode()
	if err != nil {
		log.Printf("new a request error")
		fmt.Println(err)
		return -1
	}
	closeRequest.AddCookie(&http.Cookie{
		Name:  cookie,
		Value: session,
	})
	logger.MyLogger.Info("session is  ",
		zap.String("session_name", cookie),
		zap.String("session_value", session),
	)
	resp, err = client.Do(closeRequest)
	if err != nil {
		fmt.Println(err)
	}
	if config.Option {

		openRequest, err := http.NewRequest(http.MethodPatch, requestUrl, bytes.NewReader([]byte(openBody)))
		openRequest.URL.RawQuery = params.Encode()
		if err != nil {
			log.Printf("new a request error")
			fmt.Println(err)
			return -1
		}
		openRequest.AddCookie(&http.Cookie{
			Name:  cookie,
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
	return 0
}
