package checker

import (
	"0_common/commonFunction"
	"0_common/commonStruct"
	"3_transition"
	"fmt"
)

// 邮箱是否存在 校验操作
func EmailIsExist(email string) *commonStruct.Error {
	// 其实可以通过email，直接获取uai
	// 这里走了个弯路，先获得uid
	uid, Err := transition.GetUid(email)
	if Err != nil {
		return Err
	}
	if uid == 0 {
		return commonStruct.NewError(
			fmt.Errorf("邮箱：%s 不存在", email),
			nil,
		)
	}
	return nil
}

// 邮箱是否不存在 校验操作
func EmailIsNotExist(email string) *commonStruct.Error {
	// 其实可以通过email，直接获取uai
	// 这里走了个弯路，先获得uid
	uid, Err := transition.GetUid(email)
	if Err != nil {
		return Err
	}
	if uid != 0 {
		return commonStruct.NewError(
			fmt.Errorf("邮箱：%s 已存在", email),
			nil,
		)
	}
	return nil
}

// 密码是否正确
func PasswordIsRight(email, password string) *commonStruct.Error {
	// 其实可以通过email，直接获取uai
	// 这里走了个弯路，先获得uid
	uid, Err := transition.GetUid(email)
	if Err != nil {
		return Err
	}
	uai, Err := transition.GetUai(uid)
	if Err != nil {
		return Err
	}

	// 判断是否正确
	var hashPassword string
	var err error
	if hashPassword, err = commonFunction.SaltHash(password, uai.Salt); err != nil {
		return commonStruct.NewError(
			nil,
			err,
		)
	}
	if uai.UserPassword != hashPassword {
		return commonStruct.NewError(
			fmt.Errorf("密码错误"),
			err,
		)
	}
	return nil
}

// 注册验证码是否正确
func RegisterVrcIsRight(email, vrc string) *commonStruct.Error {
	rightVrc, Err := transition.GetRegisterVrc(email)
	//fmt.Println("wwww",rightVrc,Err)
	if Err != nil {
		return Err
	}
	if rightVrc == "" {
		return commonStruct.NewError(
			fmt.Errorf("邮箱：%s，验证码未发送或已过期", email),
			nil,
		)
	}

	if rightVrc != vrc {
		return commonStruct.NewError(
			fmt.Errorf("验证码错误"),
			nil,
		)
	}
	return nil
}
