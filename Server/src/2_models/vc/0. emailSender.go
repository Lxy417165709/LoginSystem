package vc

import (
	"1_env"
	"fmt"
	"net/smtp"
)

const (
	QQStmpAddr = "smtp.qq.com"
	QQStmpPort = 25
	WyStmpAddr = "smtp.163.com"
	WyStmpPort = 25
)
// 邮件结构
type email struct {
	Sender    string
	Receivers []string
	Subject   string
	Body      string
	Type      string
}

// 邮箱客户端结构
type emailSender struct {
	UserEmail string // 邮箱
	Password  string // 授权码
	Addr      string
	Port      int
}

// 这里依赖了环境配置
var es = newEmailSender(
	env.Conf.EmailServer.User,
	env.Conf.EmailServer.Password,
	"qq",
)

// 单例
func GetEsInstance() *emailSender{
	return es
}

func newEmailSender(user string, password string, smtpType string) *emailSender {
	switch smtpType {
	case "qq":
		return &emailSender{
			user,
			password,
			QQStmpAddr,
			QQStmpPort,
		}
	case "wy":
		return &emailSender{
			user,
			password,
			WyStmpAddr,
			WyStmpPort,
		}
	default:
		return nil
	}
}

func (es *emailSender) send(email *email) error {
	if es == nil {
		return fmt.Errorf("the pointer of the client is nil")
	}
	auth := smtp.PlainAuth("", es.UserEmail, es.Password, es.Addr)

	var contentType string
	switch email.Type {
	case "html":
		contentType = fmt.Sprintf("text/%s; charset=UTF-8", email.Type)
	case "plain":
		contentType = fmt.Sprintf("text/%s; charset=UTF-8", email.Type)
	default:
		return fmt.Errorf(fmt.Sprintf("the type(%s) of the email can't be distinguished", email.Type))
	}

	// msg只是用来控制显示的,
	// To字段: 显示要发送给谁(可以不要)
	// From字段: 显示由谁发(可以不要)
	// Subject字段: 控制标题(可以不要)
	// Content-Type: 指定Body内容的类型，默认为plain。
	// Body: 邮件的内容
	msg := []byte(fmt.Sprintf("From: %s\r\nSubject: %s\r\nContent-Type: %s\r\n\r\n%s", email.Sender, email.Subject, contentType, email.Body))

	host := fmt.Sprintf("%s:%d", es.Addr, es.Port)
	return smtp.SendMail(host, auth, es.UserEmail, email.Receivers, msg)
}


