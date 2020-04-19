package models

type Cache interface {
	Set(key string, value []byte, expiredTime int) error
	Del(key string) error
	Get(key string) ([]byte, error)
}
