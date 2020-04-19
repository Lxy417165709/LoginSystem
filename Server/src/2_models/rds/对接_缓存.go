package rds

func (r *Redis) Set(key string, value []byte, expiredTime int) error {
	_, err := r.Do("SET", key, value, "ex", expiredTime)
	return err
}
func (r *Redis) Del(key string) error {
	_, err := r.Do("Del", key)
	return err
}
func (r *Redis) Get(key string) ([]byte, error) {
	rpl, err := r.Do("get", key)
	if err != nil {
		return nil, err
	}
	if rpl == nil {
		return []byte{}, err
	}
	return rpl.([]byte),nil

}
