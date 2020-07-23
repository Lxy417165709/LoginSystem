package dataCenter

import (
	"0_common/commonInterface"
	"2_models/table"
	"github.com/astaxie/beego/logs"
)

const (
	uaiUidKeyFormat = "uai:uid:%d"
	upiUidKeyFormat = "upi:uid:%d"
	emailKeyFormat  = "email:%s"
	expireTime      = 120
)

// 第一个参数是标识，第二个参数是表实例
// 缓存、数据库 --- 通过 uid、email 获取 uai、upi
func (dbc DataCenter) get(identity interface{}, receiver commonInterface.ITable) (commonInterface.ITable, error) {
	var info commonInterface.ITable = nil
	var err error

	// 查缓存
	if info, err = dbc.cacheGet(identity, receiver); err != nil {
		logs.Error(err)
		return nil, err
	}

	// 缓存未命中，查数据库
	if info == nil {
		if info, err = dbc.mdbGet(identity, receiver); err != nil {
			logs.Error(err)
			return nil, err
		}
		// 数据库不存在，返回空
		if info == nil {
			return nil, nil
		}
	}

	// 缓存设置
	if err = dbc.cacheSet(identity, receiver); err != nil {
		logs.Error(err)
		return nil, err
	}
	return info, nil
}

// 缓存、数据库 --- 通过 uid、email 更新 uai、upi
func (dbc DataCenter) update(identity interface{}, result commonInterface.ITable) error {

	var err error
	// 更新主数据库
	if err = dbc.mdbUpdate(identity, result); err != nil {
		logs.Error(err)
		return err
	}

	// 更新缓存(让它失效)
	if err = dbc.cacheUpdate(identity, result); err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

// 插入操作
// 返回uid和error
// 这里有些BUG的，因为用户信息可能只创建了一半 --> 可以采用事务
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
	uai, err := dbc.get(identity, &table.UserAccountInformation{})
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	return uai.(*table.UserAccountInformation), nil
}

// 主数据库+缓存 -> 通过 email | uid 获取 upi
func (dbc DataCenter) GetUpi(identity interface{}) (*table.UserPersonalInformation, error) {
	upi, err := dbc.get(identity, &table.UserPersonalInformation{})
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	return upi.(*table.UserPersonalInformation), nil
}

// 主数据库+缓存 -> 通过 email | uid 更新upi
func (dbc DataCenter) UpdateUpi(identity interface{}, upi *table.UserPersonalInformation) error {
	return dbc.update(identity, upi)
}

// 主数据库+缓存 -> 通过 email | uid 更新uai
func (dbc DataCenter) UpdateUai(identity interface{}, uai *table.UserAccountInformation) error {
	return dbc.update(identity, uai)
}
