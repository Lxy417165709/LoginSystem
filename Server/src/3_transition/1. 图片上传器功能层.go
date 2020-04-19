package transition

// 存储图片
func storePhoto(photo []byte) (string,error) {
	return photoUploader.StorePhoto(photo)
}

func parseBase64(photoBase64 string) ([]byte, error) {
	return photoUploader.ParseBase64(photoBase64)
}

// 通过名字获取图片的BASE64编码
// 名字 -> 图片base64编码
func getTargetPhotoBase64ByName(name string) (string, error) {
	return photoUploader.GetTargetPhotoBase64ByName(name)
}
