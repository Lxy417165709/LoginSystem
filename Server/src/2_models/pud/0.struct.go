package pud

import (
	"0_common/commonInterface"
	"github.com/astaxie/beego/logs"
)

type PhotoUploader struct {
	cache commonInterface.Cache
}

func NewPhotoUploader(cache commonInterface.Cache) *PhotoUploader {
	return &PhotoUploader{cache}
}

func (pu *PhotoUploader) Init(network, host string, port int) error {
	if err := pu.cache.CacheInit(network, host, port); err != nil {
		logs.Error(err)
		return err
	}
	return nil
}
func (pu *PhotoUploader) Close() error {
	if err := pu.cache.CacheClose(); err != nil {
		logs.Error(err)
		return err
	}
	return nil
}
