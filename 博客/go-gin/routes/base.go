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
 * @Date: 2021-01-21 17:46:28
 * @LastEditTime: 2021-02-01 11:41:33
 * @LastEditors: hhhhack
 * @Description:
 * @FilePath: /code/leetcode/博客/go-gin/routes/base.go
 * @
 */

package routes

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

const (
	GET   = 1 << iota
	POST  = 1 << iota
	PUT   = 1 << iota
	HEAD  = 1 << iota
	PATCH = 1 << iota
	DELEE = 1 << iota
)

/*******封装gin********/
type Myroute struct {
	*gin.Engine
}

//定义初始化接口
func New() *Myroute {
	return &Myroute{gin.New()}
}

/*******end********/
type Base struct {
	Method int
}

type Handle interface {
	get(c *gin.Context)
	post(c *gin.Context)
	put(c *gin.Context)
	head(c *gin.Context)
	patch(c *gin.Context)
	delete(c *gin.Context)
}

type ErrorMeta struct {
	ErrorCode int
	ErrorInfo string
}

const (
	SUCESS = iota
	ERROR_FILENAME_IS_NOT_SAFE
	LAST
)

var ErrorMsg = map[int]ErrorMeta{
	0: {200, "request ok "},
}

var MyEngine *gin.Engine

func RegisterError(errorcode int, msg string, httpcode int) int {
	if errorcode == LAST {
		log.Fatal("code is full, con't register error msg ")
	}
	if ErrorMsg[errorcode] != nil {
		log.Infof("this code %d had register ", errorcode)
		return -1
	}
	ErrorMsg[errorcode] = ErrorMeta{httpcode, msg}
	return 0
}

// func ReturnErr(err *ErrorMeta, parmars ...interface)string {
// 	return fmt.Sprintf(err.ErrorInfo, parmars)
// }

func (base *Base) get(c *gin.Context) {
	fmt.Printf(" this method is not allow")
}

func (base *Base) post(c *gin.Context) {
	fmt.Printf(" this method is not allow")
}

func (base *Base) put(c *gin.Context) {
	fmt.Printf(" this method is not allow")
}

func (base *Base) head(c *gin.Context) {
	fmt.Printf(" this method is not allow")
}

func (base *Base) delete(c *gin.Context) {
	fmt.Printf(" this method is not allow")
}

func (base *Base) patch(c *gin.Context) {
	fmt.Printf(" this method is not allow")
}
func (r *Myroute) Register(uri string, handle Handle) {
	r.GET(uri, handle.get)
	r.POST(uri, handle.post)
	r.PUT(uri, handle.put)
	r.HEAD(uri, handle.head)
}

func init(){
	MyEngine := routes.New()
	MyEngine.Use(gin.Logger())
	MyEngine.Use(gin.Recovery())
	gin.SetMode("debug")
}

// type Method interface {
// 	Init(uri)
// 	Get() int
// 	Put() int
// 	Post() int
// 	Delete() int
// }

// type Request struct{}

// type Response struct{}

// var route = gin.Default()

// type BaseRouteHandle struct {
// 	Context gin.Context
// 	Req Request
// 	Res Response
// }

// func Route(){
// 	baseroute := route.Group("/base"){
// 		baseroute.Get("/test", testHandle)
// 	}
// }

// func testHandle(c gin.Context*){

// }
