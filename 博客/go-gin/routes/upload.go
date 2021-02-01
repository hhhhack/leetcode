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
 * @Date: 2021-01-30 11:32:55
 * @LastEditTime: 2021-02-01 11:49:53
 * @LastEditors: hhhhack
 * @Description:
 * @FilePath: /code/leetcode/博客/go-gin/routes/upload.go
 * @
 */

package routes

import (
	"fmt"
	"log"
	"models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Upload struct {
	*Base
}

func init() {
	RegisterError(ERROR_FILENAME_IS_NOT_SAFE, "file name %s has dangerous char", 400)
}

func (upload *Upload) post(c *gin.Context) {
	file, _ := c.FormFile("file")
	log.Println(file.Filename)
	if models.IsDangerousStr(file.Filename) {
		c.String(ErrorMsg[ERROR_FILENAME_IS_NOT_SAFE].ErrorCode,
			fmt.Sprintf(ErrorMsg[ERROR_FILENAME_IS_NOT_SAFE].ErrorInfo), file.Filename)
	}
	// Upload the file to specific dst.
	c.SaveUploadedFile(file, dst)

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
