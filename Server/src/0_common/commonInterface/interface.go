package commonInterface

type ITable interface {
	GetTableName() string
}

type MainDb interface {
	MdbInit(host string,port int,user,password,dbname,sslmode string,maxIdleConns,maxOpenConns int) error
	MdbClose() error
	Insert(table ITable) error
	Delete(table ITable, queryStr string, parameters ...interface{}) error
	Update(table ITable, queryStr string, parameters ...interface{}) error
	Select(table ITable, queryStr string, parameters ...interface{}) ([]ITable, error)
}

type Cache interface {
	CacheInit(network string,host string ,port int) error
	CacheClose() error
	Set(key string, value []byte, expiredTime int) error
	Del(key string) error
	Get(key string) ([]byte, error)
}
