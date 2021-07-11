package rds

import (
	"1_env"
	"github.com/astaxie/beego/logs"
)

func (r *Redis) Set(key string, value []byte, expiredTime int) error {
	var err error
	var args []interface{}
	if expiredTime == 0 {
		args = []interface{}{key, value}
	} else {
		args = []interface{}{key, value, "ex", expiredTime}
	}

	_, err = r.db.Do("SET", args...)
	// 重试
	if err != nil {
		if _, err = r.retry("SET", args...); err != nil {
			return err
		}
	}

	return err
}

func (r *Redis) Del(key string) error {
	_, err := r.db.Do("Del", key)
	// 重试
	if err != nil {
		if _, err = r.retry("Del", key); err != nil {
			return err
		}
	}
	return nil
}

func (r *Redis) Get(key string) ([]byte, error) {

	rpl, err := r.db.Do("get", key)

	// 重试操作
	if err != nil {
		logs.Error(err)
		if rpl, err = r.retry("get", key); err != nil {
			return []byte{}, err
		}
	}

	if rpl == nil {
		return []byte{}, nil
	}
	return rpl.([]byte), nil
}

func (r *Redis) retry(commandName string, args ...interface{}) (interface{}, error) {
	logs.Info("%s -> %s", commandName, "重试")
	var rpl interface{}
	var err error
	for i := 0; i < r.retryTimes; i++ {
		logs.Info("%s ->第 %d 次", "重试", i+1)
		// 这有点不严谨
		if err = r.CacheInit(env.Conf.Redis.Network,env.Conf.Redis.Host,env.Conf.Redis.Port); err != nil {
			continue
		}
		if rpl, err = r.db.Do(commandName, args...);err != nil {
			continue
		}
	}
	return rpl, err
}
