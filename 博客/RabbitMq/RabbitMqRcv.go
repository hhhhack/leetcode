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
 * @Date: 2021-01-26 11:38:03
 * @LastEditTime: 2021-01-26 12:04:34
 * @LastEditors: hhhhack
 * @Description:
 * @FilePath: /code/leetcode/博客/RabbitMq/RabbitMqRcv.go
 * @
 */

package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func ReportError(err error, msg string) {
	if err != nil {
		log.Fatalf("send err %s ", msg)
	}
}

func Receive() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	ReportError(err, "open connect fail")
	defer conn.Close()
	ch, err := conn.Channel()
	ReportError(err, "get channle fail")
	defer ch.Close()
	queue, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	ReportError(err, "get queue fail")

	msg, err := ch.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	ReportError(err, "get msg fail")
	forver := make(chan bool)
	go func() {
		for data := range msg {
			fmt.Printf("message is %s", data.Body)
		}
	}()
	<-forver
}

func main() {
	Receive()
}
