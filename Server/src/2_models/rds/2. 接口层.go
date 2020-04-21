package rds

func (r *Redis) Set(key string, value []byte, expiredTime int) error {
	var err error
	if expiredTime==0{
		_,err = r.db.Do("SET", key, value)
	}else{
		_, err = r.db.Do("SET", key, value, "ex", expiredTime)
	}
	return err
}
func (r *Redis) Del(key string) error {
	_, err := r.db.Do("Del", key)
	return err
}
func (r *Redis) Get(key string) ([]byte, error) {
	rpl, err := r.db.Do("get", key)
	if err != nil {
		return nil, err
	}
	if rpl == nil {
		return []byte{}, err
	}
	return rpl.([]byte),nil

}
