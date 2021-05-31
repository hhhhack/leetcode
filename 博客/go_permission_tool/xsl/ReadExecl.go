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
 * @Date: 2021-05-27 17:45:59
 * @LastEditTime: 2021-05-27 19:41:26
 * @LastEditors: hhhhack
 * @Description:
 * @FilePath: /leetcode/博客/go_permission_tool/xsl/ReadExecl.go
 *
 */

package xsl

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func Readexecl(path string) {
	f, err := excelize.OpenFile(path)
	if err != nil {
		fmt.Print(err)
	}
	raws, err := f.GetRaws("sheet")
	for _, raw := range raws {
		fmt.Print(raw)
	}
}
