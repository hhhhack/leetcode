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
 * @Date: 2021-01-27 16:19:16
 * @LastEditTime: 2021-01-27 17:58:44
 * @LastEditors: hhhhack
 * @Description:
 * @FilePath: /code/leetcode/博客/go-gin/routes/user.go
 * @
 */

package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	*Base
}

func (user *User) get(c *gin.Context) {
	name := c.Param("name")
	fmt.Printf("name is %s \n", name)
	c.String(http.StatusOK, "Hello, %s ", name)
}

func (user *User) post(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")
	if name == "" {
		c.String(http.StatusBadRequest, "user name is empty")
	}
	if password == "" {
		c.String(http.StatusBadRequest, "user password is empty")
	}
}
