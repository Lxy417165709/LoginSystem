package dataCenter

import (
	"0_common/commonInterface"
	"2_models/table"
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"strconv"
)

// 主数据库， email -> uid
func (dbc DataCenter) mdbGetUid(email string) (int, error) {
	uais, err := dbc.mainDb.Select(&table.UserAccountInformation{}, "where UserEmail=$1", email)
	if err != nil {
		return 0, err
	}
	if len(uais) == 0 {
		return 0, nil
	}
	return uais[0].(*table.UserAccountInformation).UserId, nil
}

// 缓存，设置 email->uid 映射
func (dbc DataCenter) setEmailToUid(email string, uid int) error {
	key := fmt.Sprintf(emailKeyFormat, email)
	if err := dbc.cache.Set(key, []byte(strconv.Itoa(uid)), 0); err != nil {
		return err
	}
	return nil
}

// 主数据库、缓存，email -> uid
func (dbc DataCenter) emailToUid(email string) (int, error) {

	// redis先查
	key := fmt.Sprintf(emailKeyFormat, email)
	bytes, err := dbc.cache.Get(key)
	if err != nil {
		return 0, err
	}
	// 表示redis中没有
	if len(bytes) == 0 {
		// 没有就去主数据库查
		uid, err := dbc.mdbGetUid(email)
		if err != nil {
			return 0, err
		}
		if uid == 0 {
			// 表示数据库也没有
			return 0, nil
		}

		// 缓存，设置 email->uid 映射
		if err := dbc.setEmailToUid(email, uid); err != nil {
			return 0, err
		}
		return uid, nil
	}
	uid, err := redis.Int(bytes, err)
	if err != nil {
		return 0, err
	}

	return uid, nil
}

// 主数据库，email,uid  -> uai
func (dbc DataCenter) mdbGetUai(value interface{}) (*table.UserAccountInformation, error) {
	uais, err := make([]commonInterface.ITable, 0), error(nil)
	uid, err := dbc.GetUid(value)
	if err != nil {
		return nil, err
	}
	uais, err = dbc.mainDb.Select(&table.UserAccountInformation{}, "where UserId=$1", uid)

	if err != nil || len(uais) == 0 {
		return nil, err
	}
	return uais[0].(*table.UserAccountInformation), nil
}

// 主数据库，email,uid  -> upi
func (dbc DataCenter) mdbGetUpi(value interface{}) (*table.UserPersonalInformation, error) {
	upis, err := make([]commonInterface.ITable, 0), error(nil)
	uid, err := dbc.GetUid(value)
	if err != nil {
		return nil, err
	}

	upis, err = dbc.mainDb.Select(&table.UserPersonalInformation{}, "where UserId=$1", uid)
	if err != nil || len(upis) == 0 {
		return nil, err
	}
	return upis[0].(*table.UserPersonalInformation), nil
}

// 缓存，email,uid -> upi
func (dbc DataCenter) cacheGetUpi(value interface{}) (*table.UserPersonalInformation, error) {

	upi := new(table.UserPersonalInformation)
	// 将uid,email 统一为 uid
	uid, err := dbc.GetUid(value)
	if err != nil {
		return nil, err
	}

	// 缓存
	key := fmt.Sprintf(upiUidKeyFormat, uid)
	upiBytes, err := dbc.cache.Get(key)
	if err != nil || len(upiBytes) == 0 {
		return nil, err
	}
	// 解析
	if err := json.Unmarshal(upiBytes, upi); err != nil {
		return nil, err
	}

	return upi, nil
}

// 缓存，通过email,uid 设置 upi
func (dbc DataCenter) cacheSetUpi(value interface{}, upi table.UserPersonalInformation) error {
	uid, err := dbc.GetUid(value)
	if err != nil {
		return err
	}

	upiBytes, err := make([]byte, 0), error(nil)
	if upiBytes, err = json.Marshal(&upi); err != nil {
		return err
	}

	key := fmt.Sprintf(upiUidKeyFormat, uid)
	if err = dbc.cache.Set(key, upiBytes, expireTime); err != nil {
		return err
	}
	return nil
}

// 主、缓存， email,uid -> uid
func (dbc DataCenter) GetUid(identify interface{}) (int, error) {
	uid, err := 0, error(nil)

	switch identify.(type) {
	case int:
		uid = identify.(int)
	case string:
		uid, err = dbc.emailToUid(identify.(string))
		if err != nil {
			return 0, err
		}
	default:
		return 0, fmt.Errorf("类型错误，无法获取uid")
	}
	return uid, nil
}
