package vc

import (
	"0_common/commonInterface"
	"0_common/commonTemplate"
	"1_env"
	"fmt"
)

type RegisterEmailVrcChecker struct {
	es      *emailSender
	storage commonInterface.Cache
}

func NewRegisterEmailVrcChecker(storage commonInterface.Cache) *RegisterEmailVrcChecker {
	email, password := env.Conf.EmailServer.User, env.Conf.EmailServer.Password
	return &RegisterEmailVrcChecker{
		newEmailSender(email, password, "qq"),
		storage,
	}
}

// 发送，不存储
func (re *RegisterEmailVrcChecker) SendVrc(receiver, vrc string) error {
	email := &email{
		re.es.UserEmail,
		[]string{receiver},
		"Golang验证码邮件",
		fmt.Sprintf(commonTemplate.VrcEmail, vrc),
		"html",
	}

	if err := re.es.send(email); err != nil {
		return err
	}
	return nil
}

// 获取验证码
func (re *RegisterEmailVrcChecker) GetVrc(email string) (string, error) {
	key := fmt.Sprintf("email:%s:vrc", email)
	vrcBytes, err := re.storage.Get(key)
	if err != nil {
		return "", err
	}
	return string(vrcBytes), nil
}

// 设置验证码
func (re *RegisterEmailVrcChecker) SetVrc(receiver, vrc string, expiredTime int) error{
	key := fmt.Sprintf("email:%s:vrc", receiver)
	return re.storage.Set(key,[]byte(vrc),expiredTime)
}

// 删除验证码
func (re *RegisterEmailVrcChecker) DelVrc(receiver string) error{
	key := fmt.Sprintf("email:%s:vrc", receiver)
	return re.storage.Del(key)
}

