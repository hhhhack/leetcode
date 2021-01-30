package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

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
 * @Date: 2021-01-12 21:10:02
 * @LastEditTime: 2021-01-29 10:03:19
 * @LastEditors: hhhhack
 * @Description:
 * @FilePath: /code/leetcode/博客/go_tool/parseConf.go
 * @
 */

package conf

import (
	"log"
	"strings"
)

type DeviceInfo struct {
	userName  string
	password  string
	ipadderss string
	httpport  int
}

func parseConf(conf string) DeviceInfo {
	if len(conf) == 0 {
		log.Printf("conf is nil")
		return DeviceInfo{"", "", ""}
	}
	if strings.Index(conf, ":") != -1 {
		confs := strings.Split(conf, ":")
		fmt.Println("conf is :", confs)
		if len(confs) != 3 {
			log.Printf("conf is illagal")
			return DeviceInfo{"", "", ""}
		}
		return DeviceInfo{strings.TrimSpace(confs[0]), strings.TrimSpace(confs[1]), strings.TrimSpace(confs[2])}
	}
	return DeviceInfo{"admin", "admin", strings.TrimSpace(conf)}
}

func getConfig(configs []DeviceInfo) int {
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