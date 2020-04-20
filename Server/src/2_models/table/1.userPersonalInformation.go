package table

import (
	"0_common/commonConst"
)

// 对应数据库的 tb_userPersonalInformation 表
type UserPersonalInformation struct {
	UserId           int    `json:"userId" isPrimaryKey:"true"`
	UserPhotoUrl     string `json:"userPhotoUrl"`
	UserName         string `json:"userName"`
	UserSex          int    `json:"userSex"`
	UserContactPhone string `json:"userContactPhone"`
	UserContactEmail string `json:"userContactEmail"`
	UserBirthday     int    `json:"userBirthday"`
	Reserved1        string `json:"reserved1"`
	Reserved2        string `json:"reserved2"`
}
func(upi *UserPersonalInformation)GetTableName() string{
	return commonConst.NameOfTableUpi
}

//func(upi *UserPersonalInformation)GetInsertSql()(sqlSentence string, valueSlice []interface{}, err error){
//	var tableName = upi.GetTableName()
//
//	uType := reflect.TypeOf(upi).Elem()
//	uValue := reflect.ValueOf(upi).Elem()
//	setPartSlice := make([]string, 0)
//	keyNameSlice := make([]string, 0)
//
//	for i := 0; i < uType.NumField(); i++ {
//		keyName := uType.Field(i).Tag.Get("json")
//
//		// userId == 0 时，则不把该字段加入sql语句
//		// 为 0 值时则不加入更新字段
//		uInterface := uValue.Field(i).Interface()
//		switch uInterface.(type) {
//		case int:
//			if uInterface == 0 {
//				continue
//			}
//		case string:
//			if uInterface == "" {
//				continue
//			}
//		}
//		keyNameSlice = append(keyNameSlice, keyName)
//		valueSlice = append(valueSlice, uInterface)
//		setPartSlice = append(setPartSlice, fmt.Sprintf("$%d", len(setPartSlice)+1))
//	}
//
//	sqlSentence = fmt.Sprintf("insert into %s(%s)values(%s)", tableName, strings.Join(keyNameSlice, ","), strings.Join(setPartSlice, ","))
//	return sqlSentence, valueSlice, nil
//}
//
//// 这个有点问题
//func(upi *UserPersonalInformation)GetDeleteSql(queryStr string, parameters ...interface{}) (sqlSentence string, err error) {
//	var tableName = upi.GetTableName()
//	sqlSentence = fmt.Sprintf("delete from %s %s", tableName, queryStr)
//	return sqlSentence, nil
//}
//
//func(upi *UserPersonalInformation)GetUpdateSql(queryStr string,parameters ...interface{}) (sqlSentence string, valueSlice []interface{}, err error){
//	var tableName = upi.GetTableName()
//
//	valueSlice = append(valueSlice, parameters...)
//
//	uType := reflect.TypeOf(upi).Elem()
//	uValue := reflect.ValueOf(upi).Elem()
//	//logs.Info(upi,uType)
//	//logs.Info(upi,uValue)
//	setPartSlice := make([]string, 0)
//	for i := 0; i < uType.NumField(); i++ {
//		keyName := uType.Field(i).Tag.Get("json")
//		uInterface := uValue.Field(i).Interface()
//		// 为0值时则不加入更新字段
//		switch uInterface.(type) {
//		case int:
//			if uInterface == 0 {
//				continue
//			}
//		case string:
//			if uInterface == "" {
//				continue
//			}
//		}
//		valueSlice = append(valueSlice, uInterface)
//		setPartSlice = append(setPartSlice, fmt.Sprintf("%s=$%d", keyName, len(valueSlice)))
//	}
//	sqlSentence = fmt.Sprintf("Update %s SET %s %s", tableName, strings.Join(setPartSlice, ","), queryStr)
//	//logs.Info(sqlSentence, valueSlice)
//	return sqlSentence, valueSlice, nil
//}
//
//func(upi *UserPersonalInformation)GetSelectSql(queryStr string,parameters ...interface{})(sqlSentence string, valueSlice []interface{}, err error){
//	var tableName = upi.GetTableName()
//
//	valueSlice = append(valueSlice, parameters...)
//	sqlSentence = fmt.Sprintf("select * from %s %s", tableName, queryStr)
//	return sqlSentence, valueSlice, nil
//}
//
//// 只能传入结构体指针！
//// 功能是 获取结构体的所有字段的指针
//func(upi *UserPersonalInformation) GetFieldAddr() []interface{}{
//	val := reflect.ValueOf(upi).Elem()
//	v := make([]interface{}, val.NumField())
//	for i := 0; i < val.NumField(); i++ {
//		valueField := val.Field(i)
//		v[i] = valueField.Addr().Interface()
//	}
//	return v
//}
//
