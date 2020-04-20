package checker

import (
	"0_common/commonConst"
	"0_common/commonStruct"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"time"
)

// 邮箱格式校验
func emailCheck(email string) *commonStruct.Error {
	result, err := regexp.MatchString(commonConst.EmailRegexp, email)
	if err != nil {
		return commonStruct.NewError(
			nil,
			err,
		)
	}
	if result == false {
		return commonStruct.NewError(
			fmt.Errorf("邮箱: %s, 格式错误", email),
			nil,
		)
	}
	return nil
}

// 密码格式校验
func passwordCheck(password string) *commonStruct.Error {
	result, err := regexp.MatchString(commonConst.PasswordRegexp, password)
	if err != nil {
		return commonStruct.NewError(
			nil,
			err,
		)
	}
	if result == false {
		return commonStruct.NewError(
			fmt.Errorf("密码格式错误"),
			nil,
		)
	}
	return nil
}

// 验证码格式校验
func vrcCheck(vrc string) *commonStruct.Error {
	vrcRegexp := fmt.Sprintf("^[%s]{%d}$", commonConst.VrcPool, commonConst.VrcLength)
	result, err := regexp.MatchString(vrcRegexp, vrc)
	if err != nil {
		return commonStruct.NewError(
			nil,
			err,
		)
	}
	if result == false {
		return commonStruct.NewError(
			fmt.Errorf("验证码: %s, 格式错误", vrc),
			nil,
		)
	}
	return nil
}

// 手机格式校验
func phoneCheck(phone string) *commonStruct.Error {
	result, err := regexp.MatchString(commonConst.PhoneRegexp, phone)
	if err != nil {
		return commonStruct.NewError(
			nil,
			err,
		)
	}
	if result == false {
		return commonStruct.NewError(
			fmt.Errorf("手机格式错误"),
			nil,
		)
	}
	return nil
}

// 性别合法性校验
func sexCheck(sex int) *commonStruct.Error {
	if sex != commonConst.Man && sex != commonConst.Woman {
		return commonStruct.NewError(
			errors.New("性别信息设置有误"),
			nil,
		)
	}
	return nil
}

// 生日合法性校验
func birthdayCheck(birthday int) *commonStruct.Error {
	if birthday <= 0 || birthday > int(time.Now().Unix()*commonConst.BirthDayRato) {
		return commonStruct.NewError(
			errors.New("生日信息设置有误"),
			nil,
		)
	}
	return nil
}

// 图片合法性校验
func photoCheck(base64Data string, maxSize int) *commonStruct.Error {

	// 判断文件大小是否合法
	if len(base64Data) > maxSize {
		return commonStruct.NewError(
			fmt.Errorf("文件太大了"),
			nil,
		)
	}

	// 转为二进制
	binarys, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return commonStruct.NewError(
			nil,
			err,
		)
	}

	// 判断是否是图片类型
	typeStr := http.DetectContentType(binarys)
	result, err := regexp.MatchString(commonConst.PhotoRegexp, typeStr)
	if err != nil {
		return commonStruct.NewError(
			nil,
			err,
		)
	}
	if result == false {
		return commonStruct.NewError(
			fmt.Errorf("该文件的类型不是图片"),
			nil,
		)
	}

	return nil
}

// 用户名校验
func userNameCheck(userName string) *commonStruct.Error {
	chars := []rune(userName)
	if len(chars) <= 0 || len(chars) > 10 {
		return commonStruct.NewError(
			fmt.Errorf("用户名长度为 1-10 个字符"),
			nil,
		)
	}
	return nil
}
