package commonInterface



type ITable interface {
	GetTableName() string
}

type MainDb interface {
	Insert(table ITable) error
	Delete(table ITable, queryStr string, parameters ...interface{}) error
	Update(table ITable, queryStr string, parameters ...interface{}) error
	Select(table ITable, queryStr string, parameters ...interface{}) ([]ITable, error)
}

type Cache interface {
	Set(key string, value []byte, expiredTime int) error
	Del(key string) error
	Get(key string) ([]byte, error)
}
