package dataCenter

import (
	"0_common/commonFunction"
	"2_models/table"
	"time"
)

// 获取操作
func (dbc DataCenter) GetUai(email string) (*table.UserAccountInformation, error) {
	// 暂时不用缓存
	//key := fmt.Sprintf("uai:email:%s", email)

	uais, err := dbc.mainDb.Select(&table.UserAccountInformation{}, "where UserEmail=$1", email)
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
	// 暂时不用缓存
	//key := fmt.Sprintf("uai:uid:%d", uid)

	uais, err := dbc.mainDb.Select(&table.UserAccountInformation{}, "where UserId=$1", uid)
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
	// 暂时不用缓存
	//key := fmt.Sprintf("upi:uid:%d", uid)

	upis, err := dbc.mainDb.Select(&table.UserPersonalInformation{}, "where UserId=$1", uid)
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
	// 暂时不用缓存
	// key := fmt.Sprintf("upi:uid:%d", upi.UserId)

	return dbc.mainDb.Update(&upi, "where UserId=$1", upi.UserId)
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
func (dbc DataCenter) UpdateLastLoginTime(email string) error {
	// 暂时不用缓存
	//key := fmt.Sprintf("uai:email:%s", email)

	return dbc.mainDb.Update(
		&table.UserAccountInformation{UserLastLoginTime: int(time.Now().Unix())},
		"where UserEmail=$1",
		email,
	)
}

// 更新upi
func (dbc DataCenter) UpdateUserPhotoUrl(userId int, newPhotoUrl string) error {
	return dbc.mainDb.Update(
		&table.UserPersonalInformation{UserPhotoUrl: newPhotoUrl},
		"where UserId=$1",
		userId,
	)
}

// 更新upi
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


	//key := fmt.Sprintf("uai:email:%s", email)
	return dbc.mainDb.Update(uai, "where userEmail=$1", email)
}

// 插入操作
// 返回uid和error
func (dbc DataCenter) GenerateNewUser(email string, password string) (int, error) {

	// 创建 uai 信息
	if err := dbc.mainDb.Insert(table.NewDefaultUai(0, email, password)); err != nil {
		return 0, err
	}

	// 获取uid
	uid, err := dbc.GetUid(email)
	if err != nil {
		return 0, err
	}

	// 创建 upi 信息
	if err := dbc.mainDb.Insert(table.NewDefaultUpi(uid, email)); err != nil {
		return 0, err
	}
	return uid, nil
}
