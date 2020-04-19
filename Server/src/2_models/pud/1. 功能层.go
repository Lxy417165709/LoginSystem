package pud

import (
	"1_env"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

// 通过图片的二进制内容获取其base64编码
// 图片二进制 -> base64编码
func getPhotoBase64(photo []byte) string {
	return base64.StdEncoding.EncodeToString(photo)
}

// 获取指定路径下的图片二进制内容的BASE64编码
// 图片路径 -> base64编码
func getTargetPhotoBase64(photoPath string) (string, error) {
	file, err := os.Open(photoPath)
	if err != nil {
		return "", err
	}

	result := make([]byte, 0)
	for {
		buf := make([]byte, 1024)
		n, err := file.Read(buf)
		switch {
		case n < 0:
			return "", err
		case n == 0:
			return getPhotoBase64(result), nil
		case n > 0:
			result = append(result, buf...)
		}
	}
}

// 通过图片的二进制码得到其类型
func getImageFileSuffix(fileData []byte) string {
	tp := http.DetectContentType(fileData)
	arr := strings.Split(tp, "/")
	return arr[1]
}

// 创建文件名
func creatFileName(suffix string) string {
	return fmt.Sprintf("%d.%s", int(time.Now().UnixNano()), suffix)
}

// 存储图片
// 存储名字为 时间戳.图片类型，如: 159210324487.jpg
// 返回存储名字和错误
func (p *PhotoUploader) StorePhoto(photo []byte) (string, error) {

	storePhotoName := creatFileName(getImageFileSuffix(photo))
	return storePhotoName, ioutil.WriteFile(fmt.Sprintf("%s/%s", env.Conf.Server.PhotoPath, storePhotoName), photo, 0777) //buffer输出到jpg文件中（不做处理，直接写到文件）
}

// 通过文件名获取服务器图片文件夹中的图片BASE64
// 图片名 -> base64编码
func (p *PhotoUploader) GetTargetPhotoBase64ByName(name string) (string, error) {
	fmt.Println("图片路径",fmt.Sprintf("%s/%s", env.Conf.Server.PhotoPath, name))
	return getTargetPhotoBase64(fmt.Sprintf("%s/%s", env.Conf.Server.PhotoPath, name))
}

// 解析BASE64编码为图片二进制
// base64编码 -> 图片二进制
func (p *PhotoUploader) ParseBase64(photoBase64 string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(photoBase64)
}
