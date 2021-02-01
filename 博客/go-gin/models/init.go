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
 * @Date: 2021-01-21 11:28:44
 * @LastEditTime: 2021-01-30 11:26:26
 * @LastEditors: hhhhack
 * @Description:
 * @FilePath: /code/leetcode/博客/go-gin/models/init.go
 * @
 */

package init

import (
	"github.com/spf13/viper"
)

var AppMode int

type ServerConf struct {
	host string
	port int
}

type DbConf struct {
	host     string
	user     string
	password string
	dbname   string
}

var SvrCnf = &ServerConf{}

func ParseIni() {
	ini := viper.New()
	ini.SetDefault("", "localhost")
}

// import (
// 	"fmt"

// 	"github.com/jinzhu/gorm"
// 	_ "github.com/jinzhu/gorm/dialects/mysql"
// 	"go.uber.org/zap"
// )

// type DATABASE struct {
// 	username string
// 	host     string
// 	password string
// 	dbname   string
// }

// var database = &DATABASE{}

// func init() {
// 	log, err := zap.NewProduction()
// 	if err == nil {
// 		fmt.Printf("init log fail %v", err)
// 		return
// 	}

// }

// func DataBaseSetUp() {
// 	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
// 		database.username,
// 		database.password,
// 		database.host,
// 		database.dbname))
// 	if err == nil {
// 		fmt.Printf("open db fail %#v", err)
// 	}
// 	db.AutoMigrate(&DATABASE{})
// }

// func ParseConfig(){

// }
