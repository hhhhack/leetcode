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
 * @LastEditTime: 2021-05-31 11:16:55
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
	var style excelize.Style
	style.Font = new(excelize.Font{"blod":true, "color":""})
	ret, err := f.NewStyle(style interface{})
	for _, permission := range xmlRet.Permission {
		sheet := f.NewSheet(permission.Name)
		f.SetColStyle(sheet, "B", ret) error

		f.SetCellBool(sheet, axis string, value bool) error
		f.SetCellValue()

		f.SetCellValue(permission.Name, "A1", permission.PerName[0].Data)
	}
	if err := f.SaveAs("sys_role_permission.xlsx"); err != nil {
		fmt.Print(err)
	}
}
