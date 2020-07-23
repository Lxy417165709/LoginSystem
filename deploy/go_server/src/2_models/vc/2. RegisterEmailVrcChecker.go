package vc

import (
	"0_common/commonInterface"
	"0_common/commonTemplate"
	"fmt"
	"github.com/astaxie/beego/logs"
)

const (
	keyPattern = "email:%s:vrc"
)


type RegisterEmailVrcChecker struct {
	es      *emailSender
	storage commonInterface.Cache
}

func NewRegisterEmailVrcChecker(storage commonInterface.Cache) *RegisterEmailVrcChecker {
	return &RegisterEmailVrcChecker{
		nil,
		storage,
	}
}

func (re *RegisterEmailVrcChecker)Init(email,password string,smtpType string,network,host string,port int) error{
	if err := re.storage.CacheInit(network,host,port);err!=nil{
		logs.Error(err)
		return err
	}
	re.es = newEmailSender(email,password,smtpType)
	return nil
}

func (re *RegisterEmailVrcChecker)Close() error{
	if err := re.storage.CacheClose();err!=nil{
		logs.Error(err)
		return err
	}
	return nil
}



// 发送验证码
func (re *RegisterEmailVrcChecker) SendVrc(receiver, vrc string) error {
	email := &email{
		re.es.UserEmail,
		[]string{receiver},
		"Golang验证码邮件",
		fmt.Sprintf(commonTemplate.VrcEmail, vrc),
		"html",
	}
	if err := re.es.send(email); err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

// 获取验证码
func (re *RegisterEmailVrcChecker) GetVrc(email string) (string, error) {
	vrcBytes, err := re.storage.Get(formatKey(email))
	if err != nil {
		return "", err
	}
	return string(vrcBytes), nil
}

// 设置验证码
func (re *RegisterEmailVrcChecker) SetVrc(receiver, vrc string, expiredTime int) error{
	return re.storage.Set(formatKey(receiver),[]byte(vrc),expiredTime)
}

// 删除验证码
func (re *RegisterEmailVrcChecker) DelVrc(receiver string) error{
	return re.storage.Del(formatKey(receiver))
}

func formatKey(email string) string{
	return fmt.Sprintf(keyPattern, email)
}


