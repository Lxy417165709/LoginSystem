package dataCenter

import (
	"0_common/commonInterface"
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
