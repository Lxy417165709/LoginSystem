package pud

import (
	"encoding/base64"
	"fmt"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
)

const expiredTime = 60

// 存储图片
// 存储名字为 时间戳.图片类型，如: 159210324487.jpg
// base64编码 -> 图片名
func (p *PhotoUploader) StorePhoto(base64Data string) (string, error) {
	var bytes []byte
	var err error
	if bytes, err = base64.StdEncoding.DecodeString(base64Data); err != nil {
		logs.Error(err)
		return "", err
	}
	var storePhotoName = creatFileName(getImageFileSuffix(bytes))
	if err = ioutil.WriteFile(fmt.Sprintf("%s/%s", env.Conf.Server.PhotoPath, storePhotoName), bytes, 0777); err != nil {
		logs.Error(err)
		return "", err
	}
	// 缓存
	if err = p.cache.Set(storePhotoName, []byte(base64Data), expiredTime); err != nil {
		logs.Error(err)
		return "", err
	}
	return storePhotoName, nil
}

// 通过文件名获取服务器图片文件夹中的图片BASE64
// 图片名 -> base64编码
func (p *PhotoUploader) GetPhoto(name string) (string, error) {
	var bytes []byte
	var err error

	// 先在缓存中获取，此时获得的是base64编码
	if bytes, err = p.cache.Get(name); err != nil {
		return "", err
	}
	// 再到磁盘中获取
	var base64Str string
	if len(bytes) == 0 {
		if base64Str, err = getTargetPhotoBase64(fmt.Sprintf("%s/%s", env.Conf.Server.PhotoPath, name)); err != nil {
			return "", err
		}
		// 磁盘也没有
		if len(base64Str) == 0 {
			return "", nil
		}
	}else{
		base64Str = string(bytes)
	}


	// 缓存设置
	if err = p.cache.Set(name, []byte(base64Str), expiredTime); err != nil {
		logs.Error(err)
		return "", err
	}

	return base64Str, nil
}
