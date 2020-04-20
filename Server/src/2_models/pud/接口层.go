package pud

import (
	"1_env"
	"encoding/base64"
	"fmt"
	"io/ioutil"
)



// 存储图片
// 存储名字为 时间戳.图片类型，如: 159210324487.jpg
// base64编码 -> 图片名
func (p *PhotoUploader) StorePhoto(base64Data string) (string, error) {
	binarys := make([]byte,0)
	err := error(nil)
	if binarys,err = base64.StdEncoding.DecodeString(base64Data);err!=nil{
		return "",err
	}
	storePhotoName := creatFileName(getImageFileSuffix(binarys))
	return storePhotoName, ioutil.WriteFile(fmt.Sprintf("%s/%s", env.Conf.Server.PhotoPath, storePhotoName), binarys, 0777) //buffer输出到jpg文件中（不做处理，直接写到文件）
}


// 通过文件名获取服务器图片文件夹中的图片BASE64
// 图片名 -> base64编码
func (p *PhotoUploader) GetTargetPhotoBase64ByName(name string) (string, error) {
	return getTargetPhotoBase64(fmt.Sprintf("%s/%s", env.Conf.Server.PhotoPath, name))
}

