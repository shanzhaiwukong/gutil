package gutil

import (
	"fmt"

	gomail "gopkg.in/gomail.v2"
)

// Email 邮箱
type Email struct {
	host     string
	port     int
	userName string
	userPwd  string
}

// NewEmail 创建邮箱
func NewEmail(host, userName, userPwd string, port int) *Email {
	return &Email{
		host:     host,
		port:     port,
		userName: userName,
		userPwd:  userPwd,
	}
}

// Send 发送邮件
func (em *Email) Send(to, title, body, types string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("send email panic %v", r)
		}
	}()
	msg := gomail.NewMessage()
	msg.SetHeader("From", em.userName)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", title)
	msg.SetBody("text/"+types, body)
	d := gomail.NewDialer(em.host, em.port, em.userName, em.userPwd)
	return d.DialAndSend(msg)
}
