package checker

import (
	"0_common/commonConst"
	"4_controls/interaction"
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"reflect"
)

type checkResult struct {
	ForDeveloper error
	ForUser      error
}

var checkerInstance = &Checker{make(map[int]string)}

type Checker struct {
	msgMap map[int]string
}

func (c *Checker) init() {
	c.msgMap[ErrorFlag] = "服务器在校验表单时发生错误"
	c.msgMap[EmailNotExistFlag] = "邮箱不存在"
	c.msgMap[PasswordNotRightFlag] = "密码错误"
	c.msgMap[VrcEmptyFlag] = "验证码未发送或已过期"
	c.msgMap[VrcErrorFlag] = "验证码错误"
	c.msgMap[EmailFormatErrorFlag] = "邮箱格式错误"
	c.msgMap[PasswordFormatErrorFlag] = "密码格式错误"
	c.msgMap[VrcFormatErrorFlag] = "验证码格式错误"
	c.msgMap[PhoneFormatErrorFlag] = "手机格式错误"
	c.msgMap[SexSelectErrorFlag] = "性别选择错误"
	c.msgMap[BirthdaySelectErrorFlag] = "生日选择错误"
	c.msgMap[PhotoIsNotPhotoFileFlag] = "该文件不是图片文件"
	c.msgMap[PhotoTooLargeFlag] = "文件太大了"
	c.msgMap[UsernameFormatErrorFlag] = "用户名格式错误"
	c.msgMap[EmailExistErrorFlag] = "该邮箱已存在"
}

func GetChecker() *Checker {
	checkerInstance.init()
	return checkerInstance
}
func beString(dt reflect.Type, dv reflect.Value, fieldIndex int) (string, error) {
	value, ok := "", false
	if value, ok = dv.Field(fieldIndex).Interface().(string); !ok {
		return "", fmt.Errorf("data结构，第 %d 个字段(名为：%s)，因类型问题不能进行 %s 校验", fieldIndex+1, dt.Field(fieldIndex).Tag.Get("checkType"), dt.Field(fieldIndex).Name)
	}
	return value, nil
}
func beInt(dt reflect.Type, dv reflect.Value, fieldIndex int) (int, error) {
	value, ok := 0, false
	if value, ok = dv.Field(fieldIndex).Interface().(int); !ok {
		return 0, fmt.Errorf("data结构，第 %d 个字段(名为：%s)，因类型问题不能进行 %s 校验", fieldIndex+1, dt.Field(fieldIndex).Tag.Get("checkType"), dt.Field(fieldIndex).Name)
	}
	return value, nil
}

func check(data interface{}) (int, error) {
	dt := reflect.TypeOf(data)
	dv := reflect.ValueOf(data)
	if dt.Kind() != reflect.Struct {
		err := fmt.Errorf("data 不是一个结构体")
		logs.Error(err)
		return ErrorFlag, err
	}

	// 单字段校验
	for i := 0; i < dt.NumField(); i++ {
		checkTypeTag := dt.Field(i).Tag.Get("checkType")
		switch checkTypeTag {

		case "email":
			var value string
			var err error
			var flag int
			if value, err = beString(dt, dv, i); err != nil {
				logs.Error(err)
				return ErrorFlag, err
			}
			if flag, err = emailCheck(value); flag%2 == 0 {
				if err != nil {
					logs.Error(err)
				}
				return flag, err
			}

		case "password":
			var value string
			var err error
			var flag int
			if value, err = beString(dt, dv, i); err != nil {
				logs.Error(err)
				return ErrorFlag, err
			}
			if flag, err = passwordCheck(value); flag%2 == 0 {
				if err != nil {
					logs.Error(err)
				}
				return flag, err
			}

		case "vrc":
			var value string
			var err error
			var flag int
			if value, err = beString(dt, dv, i); err != nil {
				logs.Error(err)
				return ErrorFlag, err
			}
			if flag, err = vrcCheck(value); flag%2 == 0 {
				if err != nil {
					logs.Error(err)
				}
				return flag, err
			}

		case "birthday":
			var value int
			var err error
			var flag int
			if value, err = beInt(dt, dv, i); err != nil {
				logs.Error(err)
				return ErrorFlag, err
			}
			if flag, err = birthdayCheck(value); flag%2 == 0 {
				if err != nil {
					logs.Error(err)
				}
				return flag, err
			}

		case "sex":
			var value int
			var err error
			var flag int
			if value, err = beInt(dt, dv, i); err != nil {
				logs.Error(err)
				return ErrorFlag, err
			}
			if flag, err = sexCheck(value); flag%2 == 0 {
				if err != nil {
					logs.Error(err)
				}
				return flag, err
			}

		case "phone":
			var value string
			var err error
			var flag int
			if value, err = beString(dt, dv, i); err != nil {
				logs.Error(err)
				return ErrorFlag, err
			}
			if flag, err = phoneCheck(value); flag%2 == 0 {
				if err != nil {
					logs.Error(err)
				}
				return flag, err
			}

		case "photo":
			var value string
			var err error
			var flag int
			if value, err = beString(dt, dv, i); err != nil {
				logs.Error(err)
				return ErrorFlag, err
			}
			if flag, err = photoCheck(value, 1*commonConst.MB); flag%2 == 0 {
				if err != nil {
					logs.Error(err)
				}
				return flag, err
			}

		case "userName":
			var value string
			var err error
			var flag int
			if value, err = beString(dt, dv, i); err != nil {
				logs.Error(err)
				return ErrorFlag, err
			}
			if flag, err = userNameCheck(value); flag%2 == 0 {
				if err != nil {
					logs.Error(err)
				}
				return flag, err
			}
		default:
			err := fmt.Errorf("data结构，第 %d 个字段(名为：%s)，不存在校验类型为 %s 的函数", i+1, dt.Field(i).Name, checkTypeTag)
			logs.Error(err)
			return ErrorFlag, err
		}
	}

	// 表单校验
	switch data.(type) {

	case interaction.LoginData:
		loginData := data.(interaction.LoginData)
		if flag, err := EmailIsExist(loginData.Email); flag%2 == 0 {
			if err != nil {
				logs.Error(err)
			}
			return flag, err
		}
		if flag, err := PasswordIsRight(loginData.Email, loginData.Password); flag%2 == 0 {
			if err != nil {
				logs.Error(err)
			}
			return flag, err
		}
	case interaction.RegisterData:
		registerData := data.(interaction.RegisterData)
		if flag, err := EmailIsNotExist(registerData.Email); flag%2 == 0 {
			if err != nil {
				logs.Error(err)
			}
			return flag, err
		}
		if flag, err := RegisterVrcIsRight(registerData.Email, registerData.Vrc); flag%2 == 0 {
			if err != nil {
				logs.Error(err)
			}
			return flag, err
		}
	case interaction.UpiData:
	case interaction.UpdatePhotoData:
	case interaction.EmailData:
		emailData := data.(interaction.EmailData)
		if flag, err := EmailIsNotExist(emailData.Email); flag%2 == 0 {
			if err != nil {
				logs.Error(err)
			}
			return flag, err
		}
	default:
		err := fmt.Errorf("不支持类型为 %s 的表单验证", dt.Name())
		logs.Error(err)
		return ErrorFlag, err
	}
	return CheckPassFlag, nil
}

func (c *Checker) GetCheckResult(data interface{}) *checkResult {
	flag, err := check(data)
	if flag%2 == 0 {
		return &checkResult{
			err,
			errors.New(c.msgMap[flag]),
		}
	}
	return nil
}
