package rds

import (
	"1_env"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

type Redis struct {
	db redis.Conn
	retryTimes int
}

func (r *Redis) Init() error {
	var err error
	if r.db, err = redis.Dial(env.Conf.Redis.Network, fmt.Sprintf("%s:%d", env.Conf.Redis.Host, env.Conf.Redis.Port)); err != nil {
		return err
	}
	r.retryTimes = 3	// 默认为3
	return nil
}
func (r *Redis) Close() error {
	return r.db.Close()
}


