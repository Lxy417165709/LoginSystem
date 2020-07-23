package dataCenter

import (
	"0_common/commonInterface"
	"github.com/astaxie/beego/logs"
)

// 数据中心，整合了缓存与主数据库
type DataCenter struct {
	cache  commonInterface.Cache
	mainDb commonInterface.MainDb
}

func NewDataCenter(cache commonInterface.Cache, mainDb commonInterface.MainDb) *DataCenter {
	return &DataCenter{
		cache,
		mainDb,
	}
}

func (dc *DataCenter) Init(mdbHost string, mdbPort int, user, password, dbname, sslmode string, maxIdleConns, maxOpenConns int, network, cachehost string, cacheport int) error {
	if err := dc.cache.CacheInit(network, cachehost, cacheport); err != nil {
		logs.Error(err)
		return err
	}
	if err := dc.mainDb.MdbInit(mdbHost, mdbPort, user, password, dbname, sslmode, maxIdleConns, maxOpenConns); err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

func (dc *DataCenter) Close() error {
	if err := dc.cache.CacheClose(); err != nil {
		logs.Error(err)
		return err
	}
	if err := dc.mainDb.MdbClose(); err != nil {
		logs.Error(err)
		return err
	}
	return nil
}
