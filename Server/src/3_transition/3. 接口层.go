package transition

import (
	"0_common/commonFunction"
	"0_common/commonStruct"
	"2_models/table"
	"errors"
	"fmt"
)

func UpdateUpi(uid int, data commonStruct.UpiData) *commonStruct.Error {

	upi := table.UserPersonalInformation{}
	upi.UserId = uid
	upi.UserName = data.UserName
	upi.UserContactEmail = data.UserContactEmail
	upi.UserContactPhone = data.UserContactPhone
	upi.UserBirthday = data.UserBirthday
	upi.UserSex = data.UserSex
	// 校验还没写
	err := dataCenter.UpdateUpi(upi)
	if err != nil {
		return commonStruct.NewError(
			fmt.Errorf("服务器端发生错误：信息更新失败"),
			err,
		)
	}
	return nil
}

func UpdateLastLoginTime(email string) *commonStruct.Error {
	err := dataCenter.UpdateLastLoginTime(email)
	if err != nil {
		return commonStruct.NewError(
			fmt.Errorf("用户：%s，最近登录时间更新失败", email),
			err,
		)
	}
	return nil
}

func GetUai(uid int) (*table.UserAccountInformation, *commonStruct.Error) {

	uai, err := dataCenter.GetUaiByUid(uid)
	if err != nil {
		return uai, commonStruct.NewError(
			fmt.Errorf("用户账户信息获取失败"),
			err,
		)
	}

	// 返回用户信息
	return uai, nil
}

// 返回upi和error
func GetUpi(uid int) (*table.UserPersonalInformation, *commonStruct.Error) {
	upi, err := dataCenter.GetUpiByUid(uid)
	if err != nil {
		return upi, commonStruct.NewError(
			fmt.Errorf("用户个人信息获取失败"),
			err,
		)
	}
	return upi, nil
}

func GetUid(email string) (int, *commonStruct.Error) {
	uid, err := dataCenter.GetUid(email)
	if err != nil {
		return uid, commonStruct.NewError(
			fmt.Errorf("用户：%s, ID获取失败", email),
			err,
		)
	}
	return uid, nil
}

func GenerateNewUser(email, password string) (int, *commonStruct.Error) {
	uid, err := dataCenter.GenerateNewUser(email, password)
	if err != nil {
		return uid, commonStruct.NewError(
			fmt.Errorf("用户：%s, 新建失败", email),
			err,
		)
	}
	return uid, nil
}

// 有些冗余
func RegisterVrcIsRight(email, vrc string) *commonStruct.Error {
	return registerVrcManager.VrcIsRight(email, vrc)
}

// 图片信息校验
func UpdatePhoto(uid int, data commonStruct.UpdatePhotoData) *commonStruct.Error {

	// 存储图片
	var storeName string
	var err error
	if storeName, err = photoUploader.StorePhoto(data.PhotoBase64); err != nil {
		return commonStruct.NewError(
			errors.New("服务器错误，用户头像更新失败"),
			err,
		)
	}

	// 更新用户的photoUrl
	if err = dataCenter.UpdateUserPhotoUrl(uid, storeName); err != nil {
		return commonStruct.NewError(
			errors.New("服务器错误，用户头像更新失败"),
			err,
		)
	}
	return nil
}

// 获取图片
func GetPhotoCheck(name string) (string, *commonStruct.Error) {
	base64Str, err := photoUploader.GetTargetPhotoBase64ByName(name)
	if err != nil {
		return base64Str, commonStruct.NewError(
			errors.New("图片获取时发生错误"),
			err,
		)
	}
	return base64Str, nil
}

// 发送注册验证码
func SendRegisterVrc(email string, expiredTime int) *commonStruct.Error {
	vrc := commonFunction.CreatVrc()
	if err := registerVrcManager.SendVrc(email, vrc, expiredTime); err != nil {
		return commonStruct.NewError(
			errors.New("服务器在发送验证码时发生错误"),
			err,
		)
	}
	return nil
}





// 发送修改密码链接校验
//func SendChangePasswordLinkCheck(email string, expiredTime int) error {
//	// 发送验证码校验
//	ckus := []checkUnit{
//		NewCheckUnit(emailIsFitFormat, []interface{}{email}, true, fmt.Errorf("your email can't fit its format")),
//		NewCheckUnit(dataCenter.EmailIsExist, []interface{}{email}, true, fmt.Errorf("the email has not been registered")),
//	}
//	for _, cku := range ckus {
//		if err := cku.check(); err != nil {
//			return err
//		}
//	}
//
//	// 执行: 发送邮件
//	if err := SendVrc(changePasswordVrcChecker, email, expiredTime); err != nil {
//		return err
//	}
//	return nil
//
//}

// 修改密码链接访问校验
//func ChangePasswordLinkVisitCheck(email, vrc string) error {
//	// 发送验证码校验
//	ckus := []checkUnit{
//		NewCheckUnit(emailIsFitFormat, []interface{}{email}, true, fmt.Errorf("your email can't fit its format")),
//		NewCheckUnit(dataCenter.EmailIsExist, []interface{}{email}, true, fmt.Errorf("the email has not been registered")),
//		NewCheckUnit(vrcIsFitFormat, []interface{}{vrc}, true, fmt.Errorf("the vrc can't fit its format")),
//		NewCheckUnit(changePasswordVrcChecker.VrcIsRight, []interface{}{email, vrc, false}, true, fmt.Errorf("the link is wrong")),
//	}
//	for _, cku := range ckus {
//		if err := cku.check(); err != nil {
//			return err
//		}
//	}
//
//	return nil
//}

// 执行修改密码校验
//func ChangePasswordExecCheck(email, vrc, newPassword string) error {
//	// 这里的newPassword是明文
//
//	// 发送验证码校验
//	ckus := []checkUnit{
//		NewCheckUnit(emailIsFitFormat, []interface{}{email}, true, fmt.Errorf("your email can't fit its format")),
//		NewCheckUnit(dataCenter.EmailIsExist, []interface{}{email}, true, fmt.Errorf("the email has not been registered")),
//		NewCheckUnit(vrcIsFitFormat, []interface{}{vrc}, true, fmt.Errorf("the vrc can't fit its format")),
//		NewCheckUnit(changePasswordVrcChecker.VrcIsRight, []interface{}{ email, vrc, true}, true, fmt.Errorf("the link is wrong")),
//		NewCheckUnit(passwordIsFitFormat, []interface{}{newPassword}, true, fmt.Errorf("your password can't fit the format")),
//	}
//	for _, cku := range ckus {
//		if err := cku.check(); err != nil {
//			return err
//		}
//	}
//
//	// 执行修改密码操作
//	return dataCenter.UpdateUserPassword(email, newPassword)
//}
