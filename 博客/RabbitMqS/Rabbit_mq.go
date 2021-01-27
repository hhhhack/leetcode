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
 * @Date: 2021-01-25 12:57:36
 * @LastEditTime: 2021-01-26 11:38:10
 * @LastEditors: hhhhack
 * @Description:
 * @FilePath: /code/leetcode/博客/Rabbit_mq.go
 * @
 */

package main

import (
	"log"

	"github.com/streadway/amqp"
)

func ReportError(err error, msg string) {
	if err != nil {
		log.Fatalf("send err %s ", msg)
	}
}

func Send1(){
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	ReportError(err, "open connect fail")
	defer conn.close()
	ch, err := conn.Channel()
	ReportError(err, "get channel fail")
	defer ch.close()
	queue, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	ReportError(err, "create queue fail")
	body := "hello World"
	err = ch.Publish(
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			Contentype:"text/plain",
			Body: []byte(body)
		}
	)

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			DeliverMode: amqp.Persistent,
			Contentype:"text/plain",
			Body: []byte(body)
		}
	)
	ReportError(err, "Publish msg fail")
}

func Send2(){

}

func main() {
	Send1()
}
