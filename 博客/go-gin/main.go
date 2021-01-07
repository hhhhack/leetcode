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
 * @LastEditTime: 2021-01-04 14:35:41
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

	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
)

var ConfigFile string
var HttpPort uint

func init() {
	Config := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	Config.StringVar(&ConfigFile, "c", "/etc/myapp.ini", "config file path")
	Config.Parse(os.Args[0])
	if Config.Parsed() {
		cfg, err := ini.Load(ConfigFile)
		if err != nil {
			log.Printf("load config file fail %v ", err)
			panic("load config file fail")
		}
	}
	if cfg.Section("LOG") {
		os.OpenFile(cfg.Section("log").key("logpath").String())
	}
	HttpPort = cfg.Section("server").Key("port").MustInt(65535)
}

func main() {

	server := http.Server{
		Addr:    "0.0.0.0:" + string(HttpPort),
		Handler: gin.New(),
	}
	server.Hander.User(gin.Logger())
	apiv1 := server.Hander.Group("/api/vi")
	apiv1.Get("/users", v1.getuser)
	//apiv1.GET("/ping", )
	server.ListenAndServe()
}

type userinfo struct {
	username string `form:"user" json:"user"`
}

func getuser(gc *gin.Context) {
	var user userinfo
	err := gc.ShouleBindJSON(&user)
	if err != nil {
		return
	}
	gc.JSON(http.StatusOK, gin.H{"username": user.username})
}
