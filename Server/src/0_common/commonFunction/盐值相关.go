package commonFunction

//// 创建盐值
//func CreatSalt() string {
//	rand.Seed(time.Now().Unix())
//	salt := make([]byte, commonConst.SaltLength)
//	for i := 0; i < commonConst.SaltLength; i++ {
//		salt[i] = commonConst.SaltPool[rand.Intn(len(commonConst.SaltPool))]
//	}
//	return string(salt)
//}
//
//// 对 password 进行加盐哈希，其中盐值为 salt
//// 返回十六进制哈希值(string)
//func SaltHash(password string, salt string) (string, error) {
//	// 加盐
//	firstLayPassword := password + salt
//
//	// 开始哈希
//	h := sha1.New()
//	if _, err := h.Write([]byte(firstLayPassword)); err != nil {
//		return "", err
//	}
//
//	return fmt.Sprintf("%x", h.Sum(nil)), nil
//}
