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


func stringfy(dt reflect.Type,dv reflect.Value,fieldIndex int) (string,*commonStruct.Error){
	value, ok := "", false
	if value, ok = dv.Field(fieldIndex).Interface().(string); !ok {
		return "",commonStruct.NewError(
			nil,
			fmt.Errorf("data结构，第 %d 个字段(名为：%s)，因类型问题不能进行 %s 校验", fieldIndex+1, dt.Field(fieldIndex).Tag.Get("checkType"),dt.Field(fieldIndex).Name),
		)
	}
	return value,nil
}
func intfy(dt reflect.Type,dv reflect.Value,fieldIndex int) (int,*commonStruct.Error){
	value, ok := 0, false
	if value, ok = dv.Field(fieldIndex).Interface().(int); !ok {
		return 0,commonStruct.NewError(
			nil,
			fmt.Errorf("data结构，第 %d 个字段(名为：%s)，因类型问题不能进行 %s 校验", fieldIndex+1, dt.Field(fieldIndex).Tag.Get("checkType"),dt.Field(fieldIndex).Name),
		)
	}
	return value,nil
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


	// 单字段校验
	for i := 0; i < dt.NumField(); i++ {
		checkTypeTag := dt.Field(i).Tag.Get("checkType")
		switch checkTypeTag {

		case "email":
			value,Err := stringfy(dt,dv,i)
			if Err != nil{
				return Err
			}
			if Err := emailCheck(value); Err != nil {
				return Err
			}

		case "password":
			value,Err := stringfy(dt,dv,i)
			if Err != nil{
				return Err
			}
			if err := passwordCheck(value); err != nil {
				return err
			}

		case "vrc":
			value,Err := stringfy(dt,dv,i)
			if Err != nil{
				return Err
			}
			if err := vrcCheck(value); err != nil {
				return err
			}

		case "birthday":
			value,Err := intfy(dt,dv,i)
			if Err != nil{
				return Err
			}
			if err := birthdayCheck(value); err != nil {
				return err
			}

		case "sex":
			value,Err := intfy(dt,dv,i)
			if Err != nil{
				return Err
			}
			if err := sexCheck(value); err != nil {
				return err
			}

		case "phone":
			value,Err := stringfy(dt,dv,i)
			if Err != nil{
				return Err
			}
			if err := phoneCheck(value); err != nil {
				return err
			}

		case "photo":
			value,Err := stringfy(dt,dv,i)
			if Err != nil{
				return Err
			}

			if err := photoCheck(value, 1*commonConst.MB); err != nil {
				return err
			}


		case "userName":
			value,Err := stringfy(dt,dv,i)
			if Err != nil{
				return Err
			}
			if Err := userNameCheck(value);Err!=nil{
				return Err
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
	case commonStruct.RegisterData:
		registerData := data.(commonStruct.RegisterData)
		if Err := EmailIsNotExist(registerData.Email); Err != nil {
			return Err
		}
		if Err := RegisterVrcIsRight(registerData.Email, registerData.Vrc); Err != nil {
			return Err
		}

	case commonStruct.UpiData:
	case commonStruct.UpdatePhotoData:
	case commonStruct.EmailData:
		emailData := data.(commonStruct.EmailData)
		if Err := EmailIsNotExist(emailData.Email); Err != nil {
			return Err
		}
	default:
		return commonStruct.NewError(
			nil,
			fmt.Errorf("不支持类型为 %s 的表单验证", dt.Name()),
		)
	}
	return nil
}
