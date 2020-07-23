package checker

import (
	"0_common/commonConst"
	"encoding/base64"
	"fmt"
	"github.com/astaxie/beego/logs"
	"net/http"
	"regexp"
	"time"
)


func emailCheck(email string) (int, error) {
	var result bool
	var err error
	if result, err = regexp.MatchString(commonConst.EmailRegexp, email); err != nil {
		logs.Error(err)
		return ErrorFlag, err
	}
	if result == false {
		return EmailFormatErrorFlag, nil
	} else {
		return EmailFormatRightFlag, nil
	}
}


func passwordCheck(password string) (int, error) {
	var result bool
	var err error
	if result, err = regexp.MatchString(commonConst.PasswordRegexp, password); err != nil {
		logs.Error(err)
		return ErrorFlag, err
	}
	if result == false {
		return PasswordFormatErrorFlag, nil
	} else {
		return PasswordFormatRightFlag, nil
	}
}


func vrcCheck(vrc string) (int, error) {
	vrcRegexp := fmt.Sprintf("^[%s]{%d}$", commonConst.VrcPool, commonConst.VrcLength)

	var result bool
	var err error
	if result, err = regexp.MatchString(vrcRegexp, vrc); err != nil {
		logs.Error(err)
		return ErrorFlag, err
	}
	if result == false {
		return VrcFormatErrorFlag, nil
	} else {
		return VrcFormatRightFlag, nil
	}
}


func phoneCheck(phone string) (int, error) {
	var result bool
	var err error
	if result, err = regexp.MatchString(commonConst.PhoneRegexp, phone); err != nil {
		logs.Error(err)
		return ErrorFlag, err
	}
	if result == false {
		return PhoneFormatErrorFlag, nil
	} else {
		return PhoneFormatRightFlag, nil
	}
}


func sexCheck(sex int) (int, error) {
	if sex != commonConst.Man && sex != commonConst.Woman {
		return SexSelectErrorFlag, nil
	} else {
		return SexSelectRightFlag, nil
	}
}


func birthdayCheck(birthday int) (int, error) {
	if birthday <= 0 || birthday > int(time.Now().Unix()*commonConst.TimeRato) {
		return BirthdaySelectErrorFlag, nil
	} else {
		return BirthdaySelectRightFlag, nil
	}
}


func photoCheck(base64Data string, maxSize int) (int, error) {

	// 判断文件大小是否合法
	if len(base64Data) > maxSize {
		return PhotoTooLargeFlag, nil
	}

	// 转为二进制
	var bytes []byte
	var err error
	if bytes, err = base64.StdEncoding.DecodeString(base64Data); err != nil {
		return ErrorFlag, err
	}

	// 判断是否是图片类型
	var typeStr string
	var result bool
	typeStr = http.DetectContentType(bytes)

	if result, err = regexp.MatchString(commonConst.PhotoRegexp, typeStr); err != nil {
		return ErrorFlag, err
	}
	if result == false {
		return PhotoIsNotPhotoFileFlag, nil
	} else {
		return PhotoValidFlag, nil
	}
}


func userNameCheck(userName string) (int, error) {
	var chars []rune
	if chars = []rune(userName); len(chars) <= 0 || len(chars) > 10 {
		return UsernameFormatErrorFlag, nil
	} else {
		return UsernameFormatRightFlag, nil
	}
}
