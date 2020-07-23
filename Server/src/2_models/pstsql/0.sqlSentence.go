package pstsql

import (
	"0_common/commonInterface"
	"fmt"
	"reflect"
	"strings"
)
// 只能传入结构体指针！

func getInsertSql(table commonInterface.ITable)(sqlSentence string, valueSlice []interface{}, err error){
	var tableName = table.GetTableName()

	uType := reflect.TypeOf(table).Elem()
	uValue := reflect.ValueOf(table).Elem()
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
	return sqlSentence, valueSlice, nil
}

func getDeleteSql(table commonInterface.ITable,queryStr string, parameters ...interface{}) (sqlSentence string, valueSlice []interface{},err error) {
	var tableName = table.GetTableName()
	valueSlice = append(valueSlice, parameters...)
	sqlSentence = fmt.Sprintf("delete from %s %s", tableName, queryStr)
	return sqlSentence, valueSlice,nil
}

func getUpdateSql(table commonInterface.ITable,queryStr string,parameters ...interface{}) (sqlSentence string, valueSlice []interface{}, err error){
	var tableName = table.GetTableName()

	valueSlice = append(valueSlice, parameters...)

	uType := reflect.TypeOf(table).Elem()
	uValue := reflect.ValueOf(table).Elem()
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

func getSelectSql(table commonInterface.ITable,queryStr string,parameters ...interface{})(sqlSentence string, valueSlice []interface{}, err error){
	var tableName = table.GetTableName()
	valueSlice = append(valueSlice, parameters...)
	sqlSentence = fmt.Sprintf("select * from %s %s", tableName, queryStr)
	return sqlSentence, valueSlice, nil
}

// 功能是 获取结构体的所有字段的指针
func getFieldAddr(table commonInterface.ITable) []interface{}{
	val := reflect.ValueOf(table).Elem()
	v := make([]interface{}, val.NumField())
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		v[i] = valueField.Addr().Interface()
	}
	return v
}

