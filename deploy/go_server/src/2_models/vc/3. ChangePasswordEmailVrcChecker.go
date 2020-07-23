package vc
//
//import (
//	"1_env"
//	"fmt"
//)
//
//type ChangePasswordEmailVrcChecker struct {
//	*emailSender
//	IEmailVrcStorage
//	Prefix string
//}
//func NewChangePasswordEmailVrcChecker(storage IEmailVrcStorage) *ChangePasswordEmailVrcChecker {
//	email,password := env.Conf.EmailServer.User, env.Conf.EmailServer.Password
//	return &ChangePasswordEmailVrcChecker{
//		newEmailSender(email, password, "qq"),
//		storage,
//		"cgPassword",
//	}
//}
//func (re *ChangePasswordEmailVrcChecker) SendVrc(receiver, vrc string, expiredTime int) error {
//	link := fmt.Sprintf("http://127.0.0.1:8080/changePasswordLink/visit?email=%s&vrc=%s", receiver, vrc)
//	email := &email{
//		re.UserEmail,
//		[]string{receiver},
//		"Golang验证码邮件",
//		fmt.Sprintf("<html><body><font size=\"6\">请点击链接: <font color=\"blue\"><u><b>%s</b></u></font>进行密码修改~</font> </body></html>", link),
//		"html",
//	}
//
//	if err := re.send(email); err != nil {
//		return err
//	}
//	// 存储操作
//	return re.SetVrc(re.Prefix,receiver, vrc, expiredTime)
//}
//
//func (re *ChangePasswordEmailVrcChecker) VrcIsRight(receiver, vrc string, whenPassIsNeedToDelete bool) (bool, error) {
//	var err error
//	var rightVrc string
//	// 获取正确的验证码
//	if rightVrc, err = re.GetVrc(re.Prefix,receiver); err != nil {
//		return false, err
//	}
//
//	if rightVrc != vrc {
//		return false, nil
//	}
//
//	// 是否验证通过后删除
//	if whenPassIsNeedToDelete {
//		if err := re.DelVrc(re.Prefix,receiver); err != nil {
//			return false, err
//		}
//	}
//
//	return true, nil
//}
