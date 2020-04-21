package vc

import (
	"0_common/commonInterface"
	"0_common/commonTemplate"
	"1_env"
	"fmt"
)

const (
	keyPattern = "email:%s:vrc"
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


