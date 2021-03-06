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
 * @Date: 2021-01-28 17:43:51
 * @LastEditTime: 2021-02-04 11:50:49
 * @LastEditors: hhhhack
 * @Description:
 * @FilePath: \go_tool\main.go
 * @
 */

package main

import (
	"fmt"
	"initConf"
	"logger"
	"sendRequest"
)

func init() {
	initConf.InitConf()
	logger.InitLog()
}

func main() {
	initConf.ReadConf()
	fmt.Println(initConf.Device)
	for _, device := range initConf.Device {
		initConf.Wg.Add(1)
		fmt.Println(device)
		go sendRequest.SendRequest(device)
	}
	initConf.Wg.Wait()
}
