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
 * @Date: 2021-05-24 14:22:53
 * @LastEditTime: 2021-05-24 21:03:51
 * @LastEditors: hhhhack
 * @Description:
 * @FilePath: /leetcode/博客/go_permission_tool/myxml/ReadXml.go
 *
 */

package myxml

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

type Property struct {
	Name     string `xml:"name,attr"`
	DataType string `xml:"type,attr"`
	property string `xml:"property,chardata"`
}

type Message struct {
	Name     string     `xml:"name,attr"`
	property []Property `xml:"property"`
	message  []Message  `xml:"message"`
}

type XmlRet struct {
	XMLName xml.Name `xml:"top"`
	message []Message
	Lastid  Property `xml:"property"`
}

var buf bytes.Buffer

var logger *log.Logger

func init() {
	logFile, err := os.OpenFile("../error.log", os.O_RDONLY|os.O_WRONLY|os.O_CREATE, os.ModeSticky|os.ModeSetuid|os.ModeSetgid)
	if err != nil {
		fmt.Print("open log file err \n")
	}
	logger = log.New(logFile, "ReadXml", log.LstdFlags|log.Lshortfile)
}

func ReadConf(path string, xmlret *XmlRet) {
	if path == "" {
		logger.Panicf("path is empty or nil %v ", path)
	}
	ConfFile, err := os.Open(path)
	defer ConfFile.Close()
	if err != nil {
		logger.Fatalf("open Conf file %s fail , err is %v ", path, err)
	}
	fileinfo, err := os.Stat(path)
	if err != nil {
		logger.Fatalf("get Conf file info %s fail , err is %v ", path, err)
	}
	size := fileinfo.Size()
	fmt.Printf("file size is %d \n", size)
	xmlDecoder := xml.NewDecoder(ConfFile)
	xmlDecoder.Strict = false
	err = xmlDecoder.Decode(xmlret)
	if err != nil {
		fmt.Printf("decode err %v ", err)
		os.Exit(1)
	}
	fmt.Printf("permission size is %d last id is %v , xml name is %v \n", len(xmlret.message), xmlret.Lastid, xmlret.XMLName)
}
