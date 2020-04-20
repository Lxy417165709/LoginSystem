package checker

import (
	"0_common/commonConst"
	"0_common/commonStruct"
	"fmt"
	"reflect"
)

var checkerInstance = &Checker{}

type Checker struct {
}

func GetChecker() *Checker {
	return checkerInstance
}

func (c *Checker) Check(data interface{}) *commonStruct.Error {
	dt := reflect.TypeOf(data)
	dv := reflect.ValueOf(data)
	if dt.Kind() != reflect.Struct {
		return commonStruct.NewError(
			nil,
			fmt.Errorf("data 不是一个结构体"),
		)
	}
	// 格式校验
	for i := 0; i < dt.NumField(); i++ {
		checkTypeTag := dt.Field(i).Tag.Get("checkType")

		switch checkTypeTag {

		case "email":
			value, ok := "", false
			if value, ok = dv.Field(i).Interface().(string); !ok {
				return commonStruct.NewError(
					nil,
					fmt.Errorf("data结构，第 %d 个字段(名为：%s)，因类型问题不能进行邮箱校验", i+1, dt.Field(i).Name),
				)
			}
			if err := emailCheck(value); err != nil {
				return err
			}

		case "password":
			value, ok := "", false
			if value, ok = dv.Field(i).Interface().(string); !ok {
				return commonStruct.NewError(
					nil,
					fmt.Errorf("data结构，第 %d 个字段(名为：%s)，因类型问题不能进行密码校验", i+1, dt.Field(i).Name),
				)
			}
			if err := passwordCheck(value); err != nil {
				return err
			}

		case "vrc":
			value, ok := "", false
			if value, ok = dv.Field(i).Interface().(string); !ok {
				return commonStruct.NewError(
					nil,
					fmt.Errorf("data结构，第 %d 个字段(名为：%s)，因类型问题不能进行验证码校验", i+1, dt.Field(i).Name),
				)
			}
			if err := vrcCheck(value); err != nil {
				return err
			}

		case "birthday":
			value, ok := 0, false
			if value, ok = dv.Field(i).Interface().(int); !ok {
				return commonStruct.NewError(
					nil,
					fmt.Errorf("data结构，第 %d 个字段(名为：%s)，因类型问题不能进行验证码校验", i+1, dt.Field(i).Name),
				)
			}
			if err := birthdayCheck(value); err != nil {
				return err
			}

		case "sex":
			value, ok := 0, false
			if value, ok = dv.Field(i).Interface().(int); !ok {
				return commonStruct.NewError(
					nil,
					fmt.Errorf("data结构，第 %d 个字段(名为：%s)，因类型问题不能进行验证码校验", i+1, dt.Field(i).Name),
				)
			}
			if err := sexCheck(value); err != nil {
				return err
			}

		case "phone":
			value, ok := "", false
			if value, ok = dv.Field(i).Interface().(string); !ok {
				return commonStruct.NewError(
					nil,
					fmt.Errorf("data结构，第 %d 个字段(名为：%s)，因类型问题不能进行验证码校验", i+1, dt.Field(i).Name),
				)
			}
			if err := phoneCheck(value); err != nil {
				return err
			}

		case "photo":
			value, ok := "", false
			if value, ok = dv.Field(i).Interface().(string); !ok {
				return commonStruct.NewError(
					nil,
					fmt.Errorf("data结构，第 %d 个字段(名为：%s)，因类型问题不能进行验证码校验", i+1, dt.Field(i).Name),
				)
			}

			if err := photoCheck(value, 1*commonConst.MB); err != nil {
				return err
			}
		default:
			return commonStruct.NewError(
				nil,
				fmt.Errorf("data结构，第 %d 个字段(名为：%s)，不存在校验类型为 %s 的函数", i+1, dt.Field(i).Name, checkTypeTag),
			)
		}
	}

	// 表单校验
	switch data.(type) {

	case commonStruct.LoginData:
		loginData := data.(commonStruct.LoginData)
		if Err := EmailIsExist(loginData.Email); Err != nil {
			return Err
		}
		if Err := PasswordIsRight(loginData.Email, loginData.Password); Err != nil {
			return Err
		}
		return nil
	case commonStruct.RegisterData:
		registerData := data.(commonStruct.RegisterData)
		if Err := EmailIsNotExist(registerData.Email); Err != nil {
			return Err
		}
		if Err := VrcIsRight(registerData.Email, registerData.Vrc); Err != nil {
			return Err
		}
		return nil

	case commonStruct.UpiData:
		return nil
	case commonStruct.UpdatePhotoData:
		return nil
	case commonStruct.EmailData:
		emailData := data.(commonStruct.EmailData)
		if Err := EmailIsNotExist(emailData.Email); Err != nil {
			return Err
		}
		return nil
	default:
		return commonStruct.NewError(
			nil,
			fmt.Errorf("不支持类型为 %s 的表单验证", dt.Name()),
		)
	}

}
