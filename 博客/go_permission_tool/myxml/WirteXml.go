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
 * @Date: 2021-05-27 16:19:28
 * @LastEditTime: 2021-05-27 17:37:28
 * @LastEditors: hhhhack
 * @Description:
 * @FilePath: /leetcode/博客/go_permission_tool/myxml/WirteXml.go
 *
 */

package myxml

import (
	"encoding/xml"
	"fmt"
)

func WriteConf(path string, xmlret *XmlRet) {
	// file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, os.ModeSticky|os.ModeSetuid|os.ModeSetgid)
	// if (err != nil){

	// }
	// var xmlEncoder = xml.NewEncoder(file)

	// err = xmlEncoder.Encode(xmlret)

	ret, err := xml.MarshalIndent(xmlret, "", "	")
	if err != nil {
		fmt.Printf("marsh err %c \n", err)
	}
	fmt.Print(string(ret))
}
