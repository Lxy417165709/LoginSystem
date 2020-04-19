package transition

import (
	"0_common/commonConst"
	"2_models/table"
	"fmt"
)
func UpdateUpiCheck(upi table.UserPersonalInformation)  error{
	// 校验还没写

	return dataCenter.UpdateUpi(upi)
}
// 注册校验
func RegisterCheck(email, password, vrc string) (int,error) {
	// 注册校验
	ckus := []checkUnit{
		NewCheckUnit(emailIsFitFormat, []interface{}{email}, true, fmt.Errorf("your email can't fit its format")),
		NewCheckUnit(dataCenter.EmailIsExist, []interface{}{email}, false, fmt.Errorf("the email has been registered")),
		NewCheckUnit(passwordIsFitFormat, []interface{}{password}, true, fmt.Errorf("your password can't fit its format")),
		NewCheckUnit(vrcIsFitFormat, []interface{}{vrc}, true, fmt.Errorf("your vrc can't fit its format")),
		NewCheckUnit(registerVrcChecker.VrcIsRight, []interface{}{email, vrc,true}, true, fmt.Errorf("your vrc is wrong")),
	}
	for _, cku := range ckus {
		if err := cku.check(); err != nil {
			return 0,err
		}
	}

	// 产生新用户
	return dataCenter.GenerateNewUser(email, password)
}

// 登录校验
// 返回uid和error
func LoginCheck(email, password string) (int,error) {
	// 登录校验
	ckus := []checkUnit{
		NewCheckUnit(emailIsFitFormat, []interface{}{email}, true, fmt.Errorf("your email can't fit its format")),
		NewCheckUnit(dataCenter.EmailIsExist, []interface{}{email}, true, fmt.Errorf("the email is not exist")),
		NewCheckUnit(passwordIsFitFormat, []interface{}{password}, true, fmt.Errorf("your password can't fit its format")),
		NewCheckUnit(dataCenter.PasswordIsRight, []interface{}{email, password}, true, fmt.Errorf("your password is wrong")),
	}
	for _, cku := range ckus {
		if err := cku.check(); err != nil {
			return 0,err
		}
	}

	// 执行: 更新用户的登录时间
	if err := dataCenter.UpdateLastLoginTime(email); err != nil {
		return 0,err
	}

	// 返回用户uid
	return dataCenter.GetUid(email)

}


func GetUaiCheck(uid int) (*table.UserAccountInformation,error) {
	// 返回用户信息
	return dataCenter.GetUaiByUid(uid)
}

// 返回uid和error
func GetUpiCheck(uid int) (*table.UserPersonalInformation,error) {
	// 返回用户信息
	return dataCenter.GetUpiByUid(uid)
}



// 图片信息校验
func UpdatePhotoCheck(uid int, photoBase64 string) error {

	var photo []byte
	var err error

	if photo, err = parseBase64(photoBase64); err != nil {
		return err
	}

	// 检查uid的合法性、检测图片大小
	ckus := []checkUnit{
		NewCheckUnit(uidIsValid, []interface{}{uid}, true, fmt.Errorf("uid is not valid")),
		NewCheckUnit(photoSizeIsValid, []interface{}{photo, 1 * commonConst.MB}, true, fmt.Errorf("photo size is too large")),
		NewCheckUnit(fileIsPhoto, []interface{}{photo}, true, fmt.Errorf("the file is not a photo")),
	}
	for _, cku := range ckus {
		if err := cku.check(); err != nil {
			return err
		}
	}

	// 存储图片
	var storeName string
	if storeName, err = storePhoto(photo); err != nil {
		return err
	}

	// 更新用户的photoUrl
	if err := dataCenter.UpdateUserPhotoUrl(uid, storeName); err != nil {
		return err
	}
	return nil
}

// 获取图片
func GetPhotoCheck(name string) (string, error) {
	// 上面还要进行一些校验的，这里先测试
	return getTargetPhotoBase64ByName(name)
}

// 发送注册验证码校验
func SendRegisterVrcCheck(email string, expiredTime int) error {
	// 发送验证码校验
	ckus := []checkUnit{
		NewCheckUnit(emailIsFitFormat, []interface{}{email}, true, fmt.Errorf("your email can't fit its format")),
		NewCheckUnit(dataCenter.EmailIsExist, []interface{}{email}, false, fmt.Errorf("the email has been registered")),
	}
	for _, cku := range ckus {
		if err := cku.check(); err != nil {
			return err
		}
	}

	// 执行: 发送邮件
	if err := SendVrc(registerVrcChecker, email, expiredTime); err != nil {
		return err
	}
	return nil

}

// 发送修改密码链接校验
func SendChangePasswordLinkCheck(email string, expiredTime int) error {
	// 发送验证码校验
	ckus := []checkUnit{
		NewCheckUnit(emailIsFitFormat, []interface{}{email}, true, fmt.Errorf("your email can't fit its format")),
		NewCheckUnit(dataCenter.EmailIsExist, []interface{}{email}, true, fmt.Errorf("the email has not been registered")),
	}
	for _, cku := range ckus {
		if err := cku.check(); err != nil {
			return err
		}
	}

	// 执行: 发送邮件
	if err := SendVrc(changePasswordVrcChecker, email, expiredTime); err != nil {
		return err
	}
	return nil

}

// 修改密码链接访问校验
func ChangePasswordLinkVisitCheck(email, vrc string) error {
	// 发送验证码校验
	ckus := []checkUnit{
		NewCheckUnit(emailIsFitFormat, []interface{}{email}, true, fmt.Errorf("your email can't fit its format")),
		NewCheckUnit(dataCenter.EmailIsExist, []interface{}{email}, true, fmt.Errorf("the email has not been registered")),
		NewCheckUnit(vrcIsFitFormat, []interface{}{vrc}, true, fmt.Errorf("the vrc can't fit its format")),
		NewCheckUnit(changePasswordVrcChecker.VrcIsRight, []interface{}{email, vrc, false}, true, fmt.Errorf("the link is wrong")),
	}
	for _, cku := range ckus {
		if err := cku.check(); err != nil {
			return err
		}
	}

	return nil
}

// 执行修改密码校验
func ChangePasswordExecCheck(email, vrc, newPassword string) error {
	// 这里的newPassword是明文

	// 发送验证码校验
	ckus := []checkUnit{
		NewCheckUnit(emailIsFitFormat, []interface{}{email}, true, fmt.Errorf("your email can't fit its format")),
		NewCheckUnit(dataCenter.EmailIsExist, []interface{}{email}, true, fmt.Errorf("the email has not been registered")),
		NewCheckUnit(vrcIsFitFormat, []interface{}{vrc}, true, fmt.Errorf("the vrc can't fit its format")),
		NewCheckUnit(changePasswordVrcChecker.VrcIsRight, []interface{}{ email, vrc, true}, true, fmt.Errorf("the link is wrong")),
		NewCheckUnit(passwordIsFitFormat, []interface{}{newPassword}, true, fmt.Errorf("your password can't fit the format")),
	}
	for _, cku := range ckus {
		if err := cku.check(); err != nil {
			return err
		}
	}

	// 执行修改密码操作
	return dataCenter.UpdateUserPassword(email, newPassword)
}
