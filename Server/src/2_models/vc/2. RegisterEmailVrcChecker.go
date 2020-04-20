package vc

import (
	"0_common/commonStruct"
	"1_env"
	"fmt"
)

type RegisterEmailVrcChecker struct {
	*emailSender
	IEmailVrcStorage
	Prefix string
}


func NewRegisterEmailVrcChecker(storage IEmailVrcStorage) *RegisterEmailVrcChecker {
	email,password := env.Conf.EmailServer.User, env.Conf.EmailServer.Password
	return &RegisterEmailVrcChecker{
		newEmailSender(email, password, "qq"),
		storage,
		"register",
	}
}

func (re *RegisterEmailVrcChecker) SendVrc(receiver, vrc string, expiredTime int) error {
	email := &email{
		re.UserEmail,
		[]string{receiver},
		"Golang验证码邮件",
		fmt.Sprintf("<html><body><font size=\"6\">欢迎您注册 <font color=\"red\"><b>Kiss Me</b></font> 应用，你的邮件验证码是: <font color=\"blue\"><u><b>%s</b></u></font></font></body></html>", vrc),
		"html",
	}

	if err := re.send(email); err != nil {
		return err
	}
	// 存储操作
	return re.SetVrc(re.Prefix,receiver, vrc, expiredTime)
}

func (re *RegisterEmailVrcChecker) VrcIsRight(receiver, vrc string) *commonStruct.Error {
	var err error
	var rightVrc string
	// 获取正确的验证码
	if rightVrc, err = re.GetVrc(re.Prefix,receiver); err != nil {
		return commonStruct.NewError(
			fmt.Errorf("该邮箱验证码未发送或已失效"),
			err,
		)
	}
	if rightVrc != vrc {
		return commonStruct.NewError(
			fmt.Errorf("验证码：%s 错误",vrc),
			nil,
		)
	}
	return nil
}
