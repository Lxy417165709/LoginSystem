package models

import (
	"0_common/commonFunction"
	"2_models/table"
	"fmt"
	"time"
)

// 获取操作
func (dbc DataCenter) GetUai(email string) (*table.UserAccountInformation, error) {
	key := fmt.Sprintf("uai:email:%s", email)
	uais, err := dbc.Select(key, &table.UserAccountInformation{}, "where UserEmail=$1", email)
	if err != nil {
		return nil, err
	}
	if len(uais) == 0 {
		return nil, nil
	}
	uai := uais[0].(*table.UserAccountInformation)
	return uai, nil
}

// 获取操作
func (dbc DataCenter) GetUaiByUid(uid int) (*table.UserAccountInformation, error) {
	key := fmt.Sprintf("uai:uid:%d", uid)
	uais, err := dbc.Select(key, &table.UserAccountInformation{}, "where UserId=$1", uid)
	if err != nil {
		return nil, err
	}
	if len(uais) == 0 {
		return nil, nil
	}
	uai := uais[0].(*table.UserAccountInformation)
	return uai, nil
}

// 获取操作
func (dbc DataCenter) GetUpiByUid(uid int) (*table.UserPersonalInformation, error) {
	key := fmt.Sprintf("upi:uid:%d", uid)
	upis, err := dbc.Select(key, &table.UserPersonalInformation{}, "where UserId=$1", uid)
	if err != nil {
		return nil, err
	}
	if len(upis) == 0 {
		return nil, nil
	}
	upi := upis[0].(*table.UserPersonalInformation)
	return upi, nil
}

// 更新upi
func (dbc DataCenter) UpdateUpi(upi table.UserPersonalInformation) error {
	key := fmt.Sprintf("upi:uid:%d", upi.UserId)
	return dbc.Update(key, &upi, "where UserId=$1", upi.UserId)
}

// 更新upi
func (dbc DataCenter) GetUid(email string) (int, error) {
	uai, err := dbc.GetUai(email)
	if err != nil {
		return 0, err
	}
	if uai == nil {
		return 0, nil
	}
	return uai.UserId, nil
}

// 更新操作
// 更新upi
func (dbc DataCenter) UpdateLastLoginTime(email string) error {
	key := fmt.Sprintf("uai:email:%s", email)
	return dbc.Update(
		key,
		&table.UserAccountInformation{UserLastLoginTime: int(time.Now().Unix())},
		"where UserEmail=$1",
		email,
	)
}

// 更新upi (系统错误)
func (dbc DataCenter) UpdateUserPhotoUrl(userId int, newPhotoUrl string) error {
	return dbc.Update(
		fmt.Sprintf("upi:uid:%d", userId),
		&table.UserPersonalInformation{UserPhotoUrl: newPhotoUrl},
		"where UserId=$1",
		userId,
	)
}

// 更新upi (系统错误)
func (dbc DataCenter) UpdateUserPassword(email, newPassword string) error {
	// 这里的 newPassword 是明文
	var uai *table.UserAccountInformation
	var err error
	if uai, err = dbc.GetUai(email); err != nil {
		return err
	}
	var newHashPassword string
	if newHashPassword, err = commonFunction.SaltHash(newPassword, uai.Salt); err != nil {
		return err
	}
	uai.UserPassword = newHashPassword
	key := fmt.Sprintf("uai:email:%s", email)
	return dbc.Update(key, uai, "where userEmail=$1", email)
}

// 插入操作  (系统错误)
// 返回uid和error
func (dbc DataCenter) GenerateNewUser(email string, password string) (int, error) {

	// 这里表示不使用缓存
	if err := dbc.Insert("unUse", NewDefaultUai(0, email, password), 1); err != nil {
		return 0, err
	}

	// 获取uid
	uid, err := dbc.GetUid(email)
	if err != nil {
		return 0, err
	}
	upiKey := fmt.Sprintf("upi:uid:%d", uid)

	// 插入userPersonalInformation
	if err := dbc.Insert(upiKey, NewDefaultUpi(uid, email), 60); err != nil {
		return 0, err
	}
	return uid, nil
}


