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
 * @LastEditTime: 2021-05-27 17:33:19
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

var buf bytes.Buffer

var logger *log.Logger

func init() {
	logFile, err := os.OpenFile("../error.log", os.O_RDONLY|os.O_WRONLY|os.O_CREATE, os.ModeSticky|os.ModeSetuid|os.ModeSetgid)
	if err != nil {
		fmt.Print("open log file err \n")
	}
	logger = log.New(logFile, "ReadXml", log.LstdFlags|log.Lshortfile)
}

type Property struct {
	XMLName xml.Name `xml:"name,attr"`
	// Name     string   `xml:"name,attr"`
	DataType string `xml:"type,attr"`
	Data     string
}

type Message struct {
	XMLName xml.Name   `xml:"message"`
	Name    string     `xml:"name,attr"`
	PerName []Property `xml:"property"`
	Mess    []Message  `xml:"message"`
}

type XmlRet struct {
	XMLName    xml.Name  `xml:"top"`
	Permission []Message `xml:"message"`
	Lastid     Property  `xml:"property"`
}

func (property *Property) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	// property.XMLName = start.Name
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "name":
			property.XMLName = xml.Name{"", attr.Value}
		case "type":
			property.DataType = attr.Value
		default:
			property.DataType = "unknow"
		}
	}
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}
	property.Data = s
	// fmt.Println(s)

	return nil
}

func (property Property) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	fmt.Print(start)
	start.Attr = append(start.Attr, xml.Attr{xml.Name{"", "name"}, property.XMLName.Local})
	start.Attr = append(start.Attr, xml.Attr{xml.Name{"", "type"}, property.DataType})
	return e.EncodeElement(property.Data, start)
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
	// xmlret.message = make([]Message, 0)
	err = xmlDecoder.Decode(xmlret)
	if err != nil {
		fmt.Printf("decode err %v ", err)
		os.Exit(1)
	}
	fmt.Printf("permission size is %v last id is %v , xml name is %v \n", xmlret.Permission, xmlret.Lastid, xmlret.XMLName)
}
