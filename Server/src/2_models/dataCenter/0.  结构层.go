package dataCenter

import (
	"0_common/commonInterface"
)

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
