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
 * @LastEditTime: 2021-02-03 15:35:46
 * @LastEditors: hhhhack
 * @Description:
 * @FilePath: /code/leetcode/博客/go-gin/main.go
 * @
 */

package main

import (
	"fmt"
	"net/http"
	"routes"
	"time"

	"github.com/gin-gonic/gin"
)

/**
 * @description:
 * @param {*gin.Context} c
 * @return {*}
 */
func login(c *gin.Context) {
	fmt.Printf("%s \n", c.Param("user"))
}

func main() {

	routes.MyEngine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	routes.MyEngine.POST("/login", login)

	routes.MyEngine.Register("/usr/*name", &routes.User{})
	routes.MyEngine.Refister("/upload/*name" & routes.Upload{})
	srv := &http.Server{Addr: ":443",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	srv.ListenAndServe()

}
