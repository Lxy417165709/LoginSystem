package checker

import (
	"0_common/commonFunction"
	"2_models/table"
	"3_transition"
	"github.com/astaxie/beego/logs"
)

// 邮箱是否存在 校验操作
func EmailIsExist(email string) (int, error) {

	var uid int
	var err error
	if uid, err = transition.GetUid(email); err != nil {
		logs.Error(err)
		return ErrorFlag, err
	}
	if uid == 0 {
		return EmailNotExistFlag, nil
	} else {
		return EmailExistFlag, nil
	}
}

// 邮箱是否不存在 校验操作
func EmailIsNotExist(email string) (int, error) {
	var uid int
	var err error
	if uid, err = transition.GetUid(email); err != nil {
		logs.Error(err)
		return ErrorFlag, err
	}
	if uid == 0 {
		return EmailExistRightFlag, nil
	} else {
		return EmailExistErrorFlag, nil
	}
}

// 密码是否正确
func PasswordIsRight(email, password string) (int, error) {
	var uid int
	var err error
	if uid, err = transition.GetUid(email); err != nil {
		logs.Error(err)
		return ErrorFlag, err
	}

	// 这代码还能优化
	var uai *table.UserAccountInformation
	if uai, err = transition.GetUai(uid); err != nil {
		logs.Error(err)
		return ErrorFlag, err
	}

	// 判断是否正确
	var hashPassword string
	if hashPassword, err = commonFunction.SaltHash(password, uai.Salt); err != nil {
		logs.Error(err)
		return ErrorFlag, err
	}
	if uai.UserPassword != hashPassword {
		return PasswordNotRightFlag, nil
	} else {
		return PasswordRightFlag, nil
	}
}

// 注册验证码是否正确
func RegisterVrcIsRight(email, vrc string) (int, error) {

	var rightVrc string
	var err error
	if rightVrc, err = transition.GetRegisterVrc(email); err != nil {
		logs.Error(err)
		return ErrorFlag, err
	}

	if rightVrc == "" {
		return VrcEmptyFlag, nil
	}

	if rightVrc != vrc {
		return VrcErrorFlag, nil
	} else {
		return VrcRightFlag, nil
	}
}
