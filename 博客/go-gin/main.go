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
 * @LastEditTime: 2020-12-30 11:21:09
 * @LastEditors: hhhhack
 * @Description:
 * @FilePath: /code/leetcode/博客/go-gin/main.go
 * @
 */

package main

import (
	"flag"
	"os"
	"net/http"
)

struct Server{
	host string
	port int
}

var ConfigFile string
var ServerSet Server

func init() {
	ConfigFile := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	ConfigFile.StringVar(&ConfigFile, "c", "/etc/myapp.ini", "config file path")
	ConfigFile.StringVar(&Server.host, "h", "127.0.0.1", "server host")
	ConfigFile.UintVar(&Server.port, "p", 0, "server port")
	ConfigFile.Parse()
	if ConfigFile.Parsed() {
		cfg, err := ini.Load(configfile)
		if err != nil {
			fmt.Printf("load config file fail %v ", err)
		}
	}
}

func main() {
	server := http.Server{

	}
}
