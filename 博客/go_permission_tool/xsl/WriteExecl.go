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
 * @Date: 2021-05-27 17:54:07
 * @LastEditTime: 2021-05-27 20:13:33
 * @LastEditors: hhhhack
 * @Description:
 * @FilePath: /leetcode/博客/go_permission_tool/xsl/WriteExecl.go
 *
 */

package xsl

import (
	"fmt"
	"myxml"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func WritExecl(xmlRet *myxml.XmlRet) {
	f := excelize.NewFile()
	for _, permission := range xmlRet.Permission {
		sheet := f.NewSheet(permission.Name)

		f.SetCellValue(permission.Name, "A1", permission.PerName[0].Data)
	}
	if err := f.SaveAs("sys_role_permission.xlsx"); err != nil {
		fmt.Print(err)
	}
}
