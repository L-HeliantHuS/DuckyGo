package util

import (
	"fmt"
	"net/smtp"
)

func SendEmail(username string, password string, host string, to []string, subject string, body string) {
	/*
		Params:
			username: 用户名
			password: 密码
			host: smtp服务器地址
			to: 填写多个接收者的邮箱
			subject: 主题
			body: 发送的主要信息
		Example:
			toUsers := []string{"...@qq.com", "....@qq.com"}
			SendEmail("...@qq.com", ".....", "smtp.qq.com", "testSubject", "The email is test")
	*/
	auth := smtp.PlainAuth("", username, password, host)
	content_type := "Content-Type: text/plain" + "; charset=UTF-8"
	isSuccess := make(chan string)
	for _, i := range to {
		// 异步同时发送多个邮件. goroutine!启动！
		go func(i string) {
			msg := []byte("To: " + i + "\r\n" +
				"From: " + username + "\r\n" +
				"Subject: " + subject + "\r\n" +
				content_type + "\r\n\r\n" +
				body)
			if err := smtp.SendMail(fmt.Sprintf("%s:25", host), auth, username, []string{i}, msg); err != nil {
				panic(err)
			}
			// 防止程序提前退出
			isSuccess <- i
		}(i)
	}

	for i := 0; i < len(to); i++ {
		temp := <-isSuccess
		fmt.Println(temp + "已经发送成功！")
	}
}
