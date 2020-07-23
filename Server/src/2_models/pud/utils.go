package pud

import (
	"encoding/base64"
	"fmt"
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
