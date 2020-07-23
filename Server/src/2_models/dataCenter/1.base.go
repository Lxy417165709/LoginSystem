package dataCenter

import (
	"0_common/commonInterface"
	"2_models/table"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/garyburd/redigo/redis"
	"strconv"
)


// 主、缓存， email,uid -> uid
func (dbc DataCenter) GetUid(identify interface{}) (int, error) {
	uid, err := 0, error(nil)
	if err != nil{
		logs.Error(err)
		return 0,err
	}
	switch identify.(type) {
	case int:
		uid = identify.(int)
	case string:
		uid, err = dbc.emailToUid(identify.(string))
		if err != nil {
			logs.Error(err)
			return 0, err
		}
	default:
		err := fmt.Errorf("类型错误，无法获取uid")
		logs.Error(err)
		return 0, err
	}
	return uid, nil
}


// 主数据库， email -> uid
func (dbc DataCenter) mdbGetUid(email string) (int, error) {
	uais, err := dbc.mainDb.Select(&table.UserAccountInformation{}, "where UserEmail=$1", email)
	if err != nil {
		logs.Error(err)
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
		logs.Error(err)
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
		logs.Error(err)
		return 0, err
	}
	// 表示redis中没有
	if len(bytes) == 0 {
		// 没有就去主数据库查
		uid, err := dbc.mdbGetUid(email)
		if err != nil {
			logs.Error(err)
			return 0, err
		}
		if uid == 0 {
			// 表示数据库也没有
			return 0, nil
		}

		// 缓存，设置 email->uid 映射
		if err := dbc.setEmailToUid(email, uid); err != nil {
			logs.Error(err)
			return 0, err
		}
		return uid, nil
	}
	uid, err := redis.Int(bytes, err)
	if err != nil {
		logs.Error(err)
		return 0, err
	}

	return uid, nil
}

// 主存，email | uid -> uai | upi
func (dbc DataCenter) mdbGet(identity interface{}, receiver commonInterface.ITable) (commonInterface.ITable, error) {
	infos, err := make([]commonInterface.ITable, 0), error(nil)
	uid, err := dbc.GetUid(identity)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	if infos, err = dbc.mainDb.Select(receiver, "where UserId=$1", uid); err != nil {
		logs.Error(err)
		return nil, err
	}
	if len(infos) == 0 {
		return nil, nil
	}
	return infos[0], nil
}

// 主存 email | uid -> 更新 uai | upi
func (dbc DataCenter) mdbUpdate(identify interface{}, receiver commonInterface.ITable) error {
	uid, err := dbc.GetUid(identify)
	if err != nil {
		logs.Error(err)
		return err
	}
	// 更新数据库
	if err := dbc.mainDb.Update(receiver, "where UserId=$1", uid); err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

// 缓存 email | uid -> 让 uai | upi 失效
func (dbc DataCenter) cacheUpdate(identify interface{}, receiver commonInterface.ITable) error {
	// 缓存的更新方式就是让它失效 (最方便的解法)
	var key string
	var err error
	if key,err = dbc.getKey(identify,receiver);err!=nil{
		logs.Error(err)
		return err
	}
	if err = dbc.cache.Del(key); err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

// 缓存，email | uid -> uai | upi
func (dbc DataCenter) cacheGet(identity interface{}, receiver commonInterface.ITable) (commonInterface.ITable, error) {
	var key string
	var err error
	if key, err = dbc.getKey(identity, receiver); err != nil {
		logs.Error(err)
		return nil, err
	}

	// 缓存
	var bytes []byte
	if bytes, err = dbc.cache.Get(key); err != nil {
		logs.Error(err)
		return nil, err
	}
	if len(bytes) == 0 {
		return nil, nil
	}
	if err = json.Unmarshal(bytes, receiver); err != nil {
		logs.Error(err)
		return nil, err
	}
	return receiver, nil
}

// 缓存，通过 email | uid 设置 uai | upi
func (dbc DataCenter) cacheSet(identity interface{}, receiver commonInterface.ITable) error {
	var key string
	var err error
	if key, err = dbc.getKey(identity, receiver); err != nil {
		logs.Error(err)
		return err
	}
	var bytes []byte
	if bytes, err = json.Marshal(receiver); err != nil {
		logs.Error(err)
		return err
	}
	if err = dbc.cache.Set(key, bytes, expireTime); err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

// 通过表名 email | id -> key
func (dbc DataCenter) getKey(identity interface{}, receiver commonInterface.ITable) (string, error) {
	uid, err := dbc.GetUid(identity)
	if err != nil {
		logs.Error(err)
		return "", err
	}
	var key string
	switch receiver.(type) {
	case *table.UserAccountInformation:
		key = fmt.Sprintf(uaiUidKeyFormat, uid)
	case *table.UserPersonalInformation:
		key = fmt.Sprintf(upiUidKeyFormat, uid)
	default:
		return "", fmt.Errorf("接收者类型错误，不存在keyformat")
	}
	return key, nil
}
