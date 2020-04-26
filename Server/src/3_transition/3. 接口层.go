package transition

import (
	"0_common/commonConst"
	"2_models/table"
	"github.com/astaxie/beego/logs"
	"time"
)


//// 更新用户密码
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




func UpdateUpi(uid int, uName, ucEmail, ucPhone string, uBirthday, uSex int) error {

	upi := &table.UserPersonalInformation{}
	upi.UserId = uid
	upi.UserName = uName
	upi.UserContactEmail = ucEmail
	upi.UserContactPhone = ucPhone
	upi.UserBirthday = uBirthday
	upi.UserSex = uSex
	// 校验还没写
	err := dataCenter.UpdateUpi(uid,upi)
	if err != nil {
		return err
	}
	return nil
}

func UpdateLastLoginTime(email string) error {
	uai := &table.UserAccountInformation{UserLastLoginTime: int(time.Now().Unix())*commonConst.TimeRato}

	if err := dataCenter.UpdateUai(email,uai); err != nil {
		return err
	}
	return nil
}

func GetUai(uid int) (*table.UserAccountInformation, error) {
	uai, err := dataCenter.GetUai(uid)
	if err != nil {
		return nil,err
	}

	// 返回用户信息
	return uai, nil
}

// 返回upi和error
func GetUpi(uid int) (*table.UserPersonalInformation, error) {
	upi, err := dataCenter.GetUpi(uid)
	if err != nil {
		return nil, err
	}
	return upi, nil
}

func GetUid(email string) (int, error) {
	uid, err := dataCenter.GetUid(email)
	if err != nil {
		logs.Error(err)
		return 0,err
	}
	return uid, nil
}

func GenerateNewUser(email, password string) (int, error) {
	uid, err := dataCenter.GenerateNewUser(email, password)
	if err != nil {
		logs.Error(err)
		return 0,err
	}
	return uid, nil
}

// 图片信息校验
func UpdatePhoto(uid int, photoBase64 string) error {

	// 存储图片
	var storeName string
	var err error
	if storeName, err = photoUploader.StorePhoto(photoBase64); err != nil {
		logs.Error(err)
		return err
	}
	upi := &table.UserPersonalInformation{UserPhotoUrl: storeName}

	// 更新用户的photoUrl
	if err = dataCenter.UpdateUpi(uid,upi); err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

// 获取图片
func GetPhoto(photoName string) (string, error) {
	base64Str, err := photoUploader.GetPhoto(photoName)
	if err != nil {
		logs.Error(err)
		return base64Str,err
	}
	return base64Str, nil
}

// 发送注册验证码
func SendRegisterVrc(email string,vrc string) error {
	if err := registerVrcManager.SendVrc(email, vrc); err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

// 设置
func SetRegisterVrc(email string,vrc string,expiredTime int) error {
	if err := registerVrcManager.SetVrc(email,vrc,expiredTime);err!= nil {
		logs.Error(err)
		return err
	}
	return  nil
}
func GetRegisterVrc(email string) (string, error) {
	vrc, err := registerVrcManager.GetVrc(email)
	if err != nil {
		logs.Error(err)
		return "", err
	}
	return vrc, nil
}
func DelRegisterVrc(email string) error{
	if err := registerVrcManager.DelVrc(email);err!=nil{
		logs.Error(err)
		return err
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
