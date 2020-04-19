package rds

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func (r *Redis) SetVrc(keyPrefix ,email,vrc string, expiredTime int) error {
	_, err := r.Do("SET", fmt.Sprintf("%s:%s", keyPrefix, email), vrc, "ex", expiredTime)
	return err
}
func (r *Redis) GetVrc(keyPrefix, email string) (string, error) {
	return redis.String(r.Do("get", fmt.Sprintf("%s:%s", keyPrefix, email)))
}
func (r *Redis) DelVrc(keyPrefix, email string) error {
	_, err := r.Do("del", fmt.Sprintf("%s:%s", keyPrefix, email))
	return err
}
