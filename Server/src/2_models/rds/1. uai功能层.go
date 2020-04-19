package rds

//
////插入单个uai
////以 uid 和 email 为键，uai 为值
//// uid不能为0
//func (r *Redis) InsertUai(uai table.UserAccountInformation) error {
//	if uai.UserId == 0 {
//		return fmt.Errorf("the userid is zero")
//	}
//	uaiJson, err := json.Marshal(uai)
//	if err != nil {
//		return err
//	}
//
//	_, err = r.Do("Set", fmt.Sprintf("uai:uid:%d", uai.UserId), uaiJson)
//	if err != nil {
//		return err
//	}
//	_, err = r.Do("Set", fmt.Sprintf("uai:uemail:%s", uai.UserEmail), uaiJson)
//	return err
//}
//
//// 这个函数通过 account 或 userId 来获取用户信息 (根据identification的类型判断)
//func (r *Redis) GetUai(identification interface{}) (uai table.UserAccountInformation, err error) {
//	var result interface{}
//	switch identification.(type) {
//	case int:
//		id := identification.(int)
//		result, err = r.Do("Get", fmt.Sprintf("uai:uid:%d", id))
//
//	case string:
//		// 判断用户是否存在
//		email := identification.(string)
//		result, err = r.Do("Get", fmt.Sprintf("uai:uemail:%s", email))
//	default:
//		return uai, fmt.Errorf("the type of identification has not been registered")
//	}
//
//	// 这里表示出现错误 或 用户不存在
//	if err != nil || result == nil {
//		return uai, err
//	}
//
//	// 用户存在
//	uaiJson, err := redis.Bytes(result, nil)
//	if err != nil {
//		return uai, err
//	}
//	err = json.Unmarshal(uaiJson, &uai)
//	return uai, err
//}
//
//// uid == 0 表示用户不存在
//func (r *Redis) GetUid(email string) (uid int, err error) {
//	var uai table.UserAccountInformation
//	if uai, err = r.GetUai(email); err != nil {
//		return commonConst.ErrorUserId, err
//	}
//	return uai.UserId, nil
//}
//
