package rds

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

type Redis struct {
	db redis.Conn
	retryTimes int
}
func NewRedis() *Redis{
	return &Redis{retryTimes:3}
}
func (r *Redis) CacheInit(network string,host string,port int) error {
	var err error
	if r.db, err = redis.Dial(network, fmt.Sprintf("%s:%d",host,port)); err != nil {
		return err
	}
	return nil
}
func (r *Redis) CacheClose() error {
	return r.db.Close()
}


