package dataCenter

import (
	"2_models/table"
	"fmt"
	"github.com/astaxie/beego/logs"
)

const (
	uaiUidKeyFormat = "uai:uid:%d"
	upiUidKeyFormat = "upi:uid:%d"
	emailKeyFormat  = "email:%s"
	expireTime      = 120
)

// 插入操作
// 返回uid和error
// 这里有些BUG的，因为用户信息可能只创建了一半
// 主数据库 产生新用户
func (dbc DataCenter) GenerateNewUser(email string, password string) (int, error) {

	// 创建 uai 信息
	if err := dbc.mainDb.Insert(table.NewDefaultUai(0, email, password)); err != nil {
		logs.Error(err)
		return 0, err
	}

	// 获取uid
	uid, err := dbc.GetUid(email)
	if err != nil {
		logs.Error(err)
		return 0, err
	}
	// 创建 upi 信息
	if err := dbc.mainDb.Insert(table.NewDefaultUpi(uid, email)); err != nil {
		logs.Error(err)
		return 0, err
	}
	return uid, nil
}

// 主数据库+缓存 -> 通过 email | uid 获取 uai
func (dbc DataCenter) GetUai(identity interface{}) (*table.UserAccountInformation, error) {
	uai, err := dbc.cacheGetUai(identity)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	// 缓存未命中
	if uai == nil {
		// 查询主数据库
		if uai, err = dbc.mdbGetUai(identity); err != nil {
			logs.Error(err)
			return nil, err
		}
	}

	// 缓存更新
	if err := dbc.cacheSetUai(identity, *uai); err != nil {
		logs.Error(err)
		return nil, err
	}

	return uai, nil
}

// 主数据库+缓存 -> 通过 email | uid 获取 upi
func (dbc DataCenter) GetUpi(identity interface{}) (*table.UserPersonalInformation, error) {
	// 缓存
	upi, err := dbc.cacheGetUpi(identity)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	// 缓存未命中
	if upi == nil {
		// 查询主数据库
		if upi, err = dbc.mdbGetUpi(identity); err != nil {
			logs.Error(err)
			return nil, err
		}
	}

	// 缓存更新
	if err := dbc.cacheSetUpi(identity, *upi); err != nil {
		logs.Error(err)
		return nil, err
	}

	return upi, nil
}

// 主数据库+缓存 -> 更新upi
func (dbc DataCenter) UpdateUpi(upi *table.UserPersonalInformation) error {

	// 更新数据库
	if err := dbc.mainDb.Update(upi, "where UserId=$1", upi.UserId); err != nil {
		logs.Error(err)
		return err
	}

	// 删除缓存 (让下次命中，达到更新的目的 -> 这个方法可能有些慢)
	key := fmt.Sprintf(upiUidKeyFormat, upi.UserId)
	return dbc.cache.Del(key)
}

// 主数据库+缓存 -> 更新uai
func (dbc DataCenter) UpdateUai(uai *table.UserAccountInformation) error {

	uid, err := uai.UserId, error(nil)
	if uai.UserId == 0 {
		if uid, err = dbc.GetUid(uai.UserEmail);err != nil{
			logs.Error(err)
			return err
		}
	}

	// 更新数据库
	if err := dbc.mainDb.Update(uai, "where UserId=$1", uid); err != nil {
		logs.Error(err)
		return err
	}

	// 删除缓存 (让下次命中，达到更新的目的 -> 这个方法可能有些慢)
	key := fmt.Sprintf(uaiUidKeyFormat, uid)
	return dbc.cache.Del(key)

}
