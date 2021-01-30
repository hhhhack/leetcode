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
 * @LastEditTime: 2021-01-29 10:53:15
 * @LastEditors: hhhhack
 * @Description:
 * @FilePath: /code/leetcode/博客/go_tool/main.go
 * @
 */

package main

import (
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
	for _, device := range initConf.Device {
		initConf.Wg.Add(1)
		go sendRequest.SendRequest(&device)
	}
	initConf.Wg.Wait()
}
