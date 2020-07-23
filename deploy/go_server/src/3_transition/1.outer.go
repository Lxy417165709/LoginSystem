package transition

import (
	"0_common/commonConst"
	"2_models/table"
	"github.com/astaxie/beego/logs"
	"time"
)





func (fpi *FPI)UpdateUpi(uid int, uName, ucEmail, ucPhone string, uBirthday, uSex int) error {

	upi := &table.UserPersonalInformation{}
	upi.UserId = uid
	upi.UserName = uName
	upi.UserContactEmail = ucEmail
	upi.UserContactPhone = ucPhone
	upi.UserBirthday = uBirthday
	upi.UserSex = uSex

	if err := fpi.dataCenter.UpdateUpi(uid,upi);err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

func (fpi *FPI)UpdateLastLoginTime(email string) error {
	uai := &table.UserAccountInformation{UserLastLoginTime: int(time.Now().Unix())*commonConst.TimeRato}
	if err := fpi.dataCenter.UpdateUai(email,uai); err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

func (fpi *FPI)GetUai(uid int) (*table.UserAccountInformation, error) {
	uai, err := fpi.dataCenter.GetUai(uid)
	if err != nil {
		logs.Error(err)
		return nil,err
	}

	// 返回用户信息
	return uai, nil
}

// 返回upi和error
func (fpi *FPI)GetUpi(uid int) (*table.UserPersonalInformation, error) {
	upi, err := fpi.dataCenter.GetUpi(uid)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	return upi, nil
}

func (fpi *FPI)GetUid(email string) (int, error) {
	uid, err := fpi.dataCenter.GetUid(email)
	if err != nil {
		logs.Error(err)
		return 0,err
	}
	return uid, nil
}

func (fpi *FPI)GenerateNewUser(email, password string) (int, error) {
	uid, err := fpi.dataCenter.GenerateNewUser(email, password)
	if err != nil {
		logs.Error(err)
		return 0,err
	}
	return uid, nil
}

// 图片信息校验
func (fpi *FPI) UpdatePhoto(uid int, photoBase64 string) error {

	// 存储图片
	var storeName string
	var err error
	if storeName, err = fpi.photoUploader.StorePhoto(photoBase64); err != nil {
		logs.Error(err)
		return err
	}

	// 更新用户的photoUrl
	upi := &table.UserPersonalInformation{UserPhotoUrl: storeName}
	if err = fpi.dataCenter.UpdateUpi(uid,upi); err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

// 获取图片
func (fpi *FPI)GetPhoto(photoName string) (string, error) {
	base64Str, err := fpi.photoUploader.GetPhoto(photoName)
	if err != nil {
		logs.Error(err)
		return base64Str,err
	}
	return base64Str, nil
}

// 发送注册验证码
func (fpi *FPI)SendRegisterVrc(email string,vrc string) error {
	if err := fpi.registerVrcManager.SendVrc(email, vrc); err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

// 设置
func (fpi *FPI)SetRegisterVrc(email string,vrc string,expiredTime int) error {
	if err := fpi.registerVrcManager.SetVrc(email,vrc,expiredTime);err!= nil {
		logs.Error(err)
		return err
	}
	return  nil
}
func (fpi *FPI)GetRegisterVrc(email string) (string, error) {
	vrc, err := fpi.registerVrcManager.GetVrc(email)
	if err != nil {
		logs.Error(err)
		return "", err
	}
	return vrc, nil
}
func (fpi *FPI)DelRegisterVrc(email string) error{
	if err := fpi.registerVrcManager.DelVrc(email);err!=nil{
		logs.Error(err)
		return err
	}
	return nil
}
