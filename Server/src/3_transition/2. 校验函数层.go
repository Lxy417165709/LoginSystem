package transition

import (
	"0_common/commonConst"
	"fmt"
	"net/http"
	"regexp"
)

// 格式校验
func emailIsFitFormat(email string) (bool, error) {
	return regexp.MatchString(commonConst.EmailRegexp, email)
}

func passwordIsFitFormat(password string) (bool, error) {
	return regexp.MatchString(commonConst.PasswordRegexp, password)
}

func vrcIsFitFormat(vrcValue string) (bool, error) {
	vrcRegexp := fmt.Sprintf("^[%s]{%d}$", commonConst.VrcPool, commonConst.VrcLength)
	return regexp.MatchString(vrcRegexp, vrcValue)
}

func photoBase64IsFitFormat(photoBase64 string) (bool, error) {

	return regexp.MatchString(commonConst.Base64Regexp, photoBase64)
}

// 合法性校验
func uidIsValid(uid int) (bool, error) {
	return uid != commonConst.ErrorUserId, nil
}

func photoSizeIsValid(photo []byte, size int) (bool, error) {
	// 图片大小要小于1M
	return photoUploader.PhotoSizeIsValid(photo, size)
}

func fileIsPhoto(fileData []byte) (bool, error) {
	typeStr := http.DetectContentType(fileData)
	return regexp.MatchString(`^image/.+`, typeStr)
}

