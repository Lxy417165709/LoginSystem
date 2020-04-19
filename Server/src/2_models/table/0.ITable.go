package table

type ITable interface {
	GetTableName() string
	GetInsertSql() (sqlSentence string, valueSlice []interface{}, err error)
	GetDeleteSql(querySrt string, parameters ...interface{}) (sqlSentence string, err error)
	GetUpdateSql(querySrt string, parameters ...interface{}) (sqlSentence string, valueSlice []interface{}, err error)
	GetSelectSql(querySrt string, parameters ...interface{}) (sqlSentence string, valueSlice []interface{}, err error)
	GetFieldAddr() []interface{}
}
