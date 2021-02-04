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
 * @Date: 2021-01-28 21:53:29
 * @LastEditTime: 2021-02-04 11:49:36
 * @LastEditors: hhhhack
 * @Description:
 * @FilePath: \go_tool\initConf\init.go
 * @
 */

package initConf

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

type DeviceInfo struct {
	UserName  string
	Password  string
	Ipadderss string
	Httpport  int
	Option    bool
}

var Device []DeviceInfo

var LogFilePath *string

var ConfPath *string

var Wg sync.WaitGroup

func InitConf() {
	LogFilePath = flag.String("L", "err.log", "日志文件")
	ConfPath = flag.String("C", "conf", "配置文件")
	flag.Parse()
	logFile, err := os.OpenFile(*LogFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open log file fail %v ", err)
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Lshortfile | log.Ltime)
	Device = make([]DeviceInfo, 0)
}

func parseConf(line string) DeviceInfo {
	var port int = 443
	var op bool = true
	var err error
	if len(line) == 0 {
		log.Printf("conf is nil")
		return DeviceInfo{"", "", "", 0, false}
	}
	if strings.Index(line, ":") != -1 {
		confs := strings.Split(line, ":")
		fmt.Printf("conf is :%v conf len is %d ", confs, len(confs))
		if len(confs) < 3 {
			log.Printf("conf is illagal")
			return DeviceInfo{"", "", "", 0, false}
		}
		if len(confs) > 3 {
			log.Printf("more than 3 field confs 4 is %v", confs[3])
			port, err = strconv.Atoi(confs[3])
			if err != nil {
				port = 443
			}
		}
		if len(confs) > 4 {
			log.Printf("more than 4 field confs 5 is %v", confs[4])
			if strings.TrimSpace(confs[4]) == "false" {
				op = false
			}
		}
		log.Printf("user is %s pass is %s host is %s port is % d, op is %v ", confs[0], confs[1], confs[2], port, op)
		return DeviceInfo{strings.TrimSpace(confs[0]), strings.TrimSpace(confs[1]), strings.TrimSpace(confs[2]), port, op}
	}
	return DeviceInfo{"admin", "admin", strings.TrimSpace(line), port, op}
}

func ReadConf() {
	if ConfPath == nil {
		log.Fatal("confpath is nil")
	}
	conf, err := os.Open(*ConfPath)
	if err != nil {
		log.Fatalf("read conf fail %v, file path is %s ", err, *ConfPath)
	}
	reader := bufio.NewReader(conf)
	Device = make([]DeviceInfo, 0)
	for {
		line, err := reader.ReadString('\n') //注意是字符
		if err == io.EOF {
			if len(line) != 0 {
				Device = append(Device, parseConf(line))
			}
			fmt.Println("file read finish")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		Device = append(Device, parseConf(line))
	}
}
