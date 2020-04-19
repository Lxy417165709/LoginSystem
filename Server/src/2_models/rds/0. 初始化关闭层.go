package rds

import (
	"1_env"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

type Redis struct {
	redis.Conn
}

func (r *Redis) Init() error {
	var err error
	if r.Conn, err = redis.Dial(env.Conf.Redis.Network, fmt.Sprintf("%s:%d", env.Conf.Redis.Host, env.Conf.Redis.Port)); err != nil {
		return err
	}
	return nil
}
func (r *Redis) Close() error {
	return r.Conn.Close()
}
