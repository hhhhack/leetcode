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
 * @Date: 2020-12-30 09:50:33
 * @LastEditTime: 2020-12-31 17:38:06
 * @LastEditors: hhhhack
 * @Description:
 * @FilePath: /code/leetcode/博客/go-gin/main.go
 * @
 */

package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"gopkg.in/ini.v1"
)

var ConfigFile string

func init() {
	Config := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	Config.StringVar(&ConfigFile, "c", "/etc/myapp.ini", "config file path")
	Config.Parse()
	if Config.Parsed() {
		cfg, err := ini.Load(ConfigFile)
		if err != nil {
			log.Printf("load config file fail %v ", err)
			panic("load config file fail")
		}
	}
	if cfg.Section("LOG") {
		os.OpenFile(cfg.Section("LOG").key("file"))
	}
}

func main() {
	server := http.Server{}
}
