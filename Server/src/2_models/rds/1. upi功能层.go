package rds

////插入单个upi
////以 uid 为键，upi 为值
//// uid不能为0
//func (r *Redis) InsertUpi(upi table.UserPersonalInformation) error {
//	if upi.UserId == 0 {
//		return fmt.Errorf("the userid is zero")
//	}
//	uaiJson, err := json.Marshal(upi)
//	if err != nil {
//		return err
//	}
//
//	_, err = r.Do("Set", fmt.Sprintf("upi:uid:%d", upi.UserId), uaiJson)
//	return err
//}
//
//func (r *Redis) GetUpi(userId int) (upi table.UserPersonalInformation, err error) {
//	result, err := r.Do("Get", fmt.Sprintf("upi:uid:%d", userId))
//
//	// 这里表示出现错误 或 用户不存在
//	if err != nil || result == nil {
//		return upi, err
//	}
//
//	// 用户存在
//	upiJson, err := redis.Bytes(result, nil)
//	if err != nil {
//		return upi, err
//	}
//	err = json.Unmarshal(upiJson, &upi)
//	return upi, err
//}
