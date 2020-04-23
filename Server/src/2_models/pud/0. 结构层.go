package pud

import (
	"0_common/commonInterface"
)

type PhotoUploader struct{
	cache commonInterface.Cache
}

func NewPhotoUploader(cache commonInterface.Cache) *PhotoUploader{
	return &PhotoUploader{cache}
}


