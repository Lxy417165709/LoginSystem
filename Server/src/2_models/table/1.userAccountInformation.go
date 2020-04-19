package table

import (
	"fmt"
	"reflect"
	"strings"
)
// 对应数据库的 tb_userAccountInformation 表
type UserAccountInformation struct {
	UserId            int    `json:"userId" isPrimaryKey:"true"`
	UserEmail         string `json:"userEmail"`
	UserPassword      string `json:"userPassword"`
	UserType          int    `json:"userType"`
	UserRegisterTime  int    `json:"userRegisterTime"`
	UserLastLoginTime int    `json:"userLastLoginTime"`
	Salt              string `json:"salt"`
	Reserved2         string `json:"reserved2"`
}

func(uai *UserAccountInformation)GetTableName() string{
	return "tb_userAccountInformation"
}

func(uai *UserAccountInformation)GetInsertSql()(sqlSentence string, valueSlice []interface{}, err error){
	var tableName = uai.GetTableName()

	uType := reflect.TypeOf(uai).Elem()
	uValue := reflect.ValueOf(uai).Elem()
	setPartSlice := make([]string, 0)
	keyNameSlice := make([]string, 0)

	for i := 0; i < uType.NumField(); i++ {
		keyName := uType.Field(i).Tag.Get("json")

		// userId == 0 时，则不把该字段加入sql语句
		// 为 0 值时则不加入更新字段
		uInterface := uValue.Field(i).Interface()
		switch uInterface.(type) {
		case int:
			if uInterface == 0 {
				continue
			}
		case string:
			if uInterface == "" {
				continue
			}
		}
		keyNameSlice = append(keyNameSlice, keyName)
		valueSlice = append(valueSlice, uInterface)
		setPartSlice = append(setPartSlice, fmt.Sprintf("$%d", len(setPartSlice)+1))
	}

	sqlSentence = fmt.Sprintf("insert into %s(%s)values(%s)", tableName, strings.Join(keyNameSlice, ","), strings.Join(setPartSlice, ","))
	//logs.Info(sqlSentence)
	return sqlSentence, valueSlice, nil
}

func(uai *UserAccountInformation)GetDeleteSql(queryStr string, parameters ...interface{}) (sqlSentence string, err error) {
	var tableName = uai.GetTableName()
	sqlSentence = fmt.Sprintf("delete from %s %s", tableName, queryStr)
	return sqlSentence, nil
}

func(uai *UserAccountInformation)GetUpdateSql(queryStr string,parameters ...interface{}) (sqlSentence string, valueSlice []interface{}, err error){
	var tableName = uai.GetTableName()

	valueSlice = append(valueSlice, parameters...)

	uType := reflect.TypeOf(uai).Elem()
	uValue := reflect.ValueOf(uai).Elem()
	setPartSlice := make([]string, 0)
	for i := 0; i < uType.NumField(); i++ {
		keyName := uType.Field(i).Tag.Get("json")
		uInterface := uValue.Field(i).Interface()
		// 为0值时则不加入更新字段
		switch uInterface.(type) {
		case int:
			if uInterface == 0 {
				continue
			}
		case string:
			if uInterface == "" {
				continue
			}
		}
		valueSlice = append(valueSlice, uInterface)
		setPartSlice = append(setPartSlice, fmt.Sprintf("%s=$%d", keyName, len(valueSlice)))
	}
	sqlSentence = fmt.Sprintf("Update %s SET %s %s", tableName, strings.Join(setPartSlice, ","), queryStr)
	return sqlSentence, valueSlice, nil
}

func(uai *UserAccountInformation)GetSelectSql(queryStr string,parameters ...interface{})(sqlSentence string, valueSlice []interface{}, err error){
	var tableName = uai.GetTableName()

	valueSlice = append(valueSlice, parameters...)
	sqlSentence = fmt.Sprintf("select * from %s %s", tableName, queryStr)
	return sqlSentence, valueSlice, nil
}



// 只能传入结构体指针！
// 功能是 获取结构体的所有字段的指针
func(uai *UserAccountInformation) GetFieldAddr() []interface{}{
	val := reflect.ValueOf(uai).Elem()
	v := make([]interface{}, val.NumField())
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		v[i] = valueField.Addr().Interface()
	}
	return v
}
