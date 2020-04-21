package transition

import (
	"0_common/commonStruct"
	"2_models/table"
	"errors"
	"fmt"
	"time"
)


//// 更新upi
//func (dbc DataCenter) UpdateUserPassword(email, newPassword string) error {
//	// 这里的 newPassword 是明文
//	var uai *table.UserAccountInformation
//	var err error
//	if uai, err = dbc.GetUai(email); err != nil {
//		return err
//	}
//	var newHashPassword string
//	if newHashPassword, err = commonFunction.SaltHash(newPassword, uai.Salt); err != nil {
//		return err
//	}
//	uai.UserPassword = newHashPassword
//
//	return dbc.mainDb.Update(uai, "where userEmail=$1", email)
//}
//



func UpdateUpi(uid int, uName, ucEmail, ucPhone string, uBirthday, uSex int) *commonStruct.Error {

	upi := &table.UserPersonalInformation{}
	upi.UserId = uid
	upi.UserName = uName
	upi.UserContactEmail = ucEmail
	upi.UserContactPhone = ucPhone
	upi.UserBirthday = uBirthday
	upi.UserSex = uSex
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
	uai := &table.UserAccountInformation{UserLastLoginTime: int(time.Now().Unix())}
	uai.UserEmail = email

	if err := dataCenter.UpdateUai(uai); err != nil {
		return commonStruct.NewError(
			fmt.Errorf("用户：%s，最近登录时间更新失败", email),
			err,
		)
	}
	return nil
}

func GetUai(uid int) (*table.UserAccountInformation, *commonStruct.Error) {

	uai, err := dataCenter.GetUai(uid)
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
	upi, err := dataCenter.GetUpi(uid)
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


	upi := &table.UserPersonalInformation{UserPhotoUrl: storeName}
	upi.UserId = uid
	// 更新用户的photoUrl
	if err = dataCenter.UpdateUpi(upi); err != nil {
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
func SendRegisterVrc(email string,vrc string) *commonStruct.Error {
	if err := registerVrcManager.SendVrc(email, vrc); err != nil {
		return commonStruct.NewError(
			errors.New("服务器在发送验证码时发生错误"),
			err,
		)
	}
	return nil
}

// 设置
func SetRegisterVrc(email string,vrc string,expiredTime int) *commonStruct.Error {
	if err := registerVrcManager.SetVrc(email,vrc,expiredTime);err!= nil {
		return commonStruct.NewError(
			errors.New("服务器在设置验证码时发生错误"),
			err,
		)
	}
	return  nil
}
func GetRegisterVrc(email string) (string, *commonStruct.Error) {
	vrc, err := registerVrcManager.GetVrc(email)
	if err != nil {
		return "", commonStruct.NewError(
			errors.New("服务器在获取验证码时发生错误"),
			err,
		)
	}
	return vrc, nil
}
func DelRegisterVrc(email string) *commonStruct.Error{
	if err := registerVrcManager.DelVrc(email);err!=nil{
		return commonStruct.NewError(
			errors.New("服务器在删除验证码时发生错误"),
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
