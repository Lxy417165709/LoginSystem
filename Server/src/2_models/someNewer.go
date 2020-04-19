package models

import (
	"0_common/commonConst"
	"2_models/table"
	"time"
)

func NewDefaultUai(uid int, email, password string) *table.UserAccountInformation {
	var saltPassword string
	var err error
	salt := CreatSalt()
	if saltPassword, err = SaltHash(password, salt); err != nil {
		return nil
	}
	// 插入账户信息
	uai := &table.UserAccountInformation{
		UserId:            uid, // 让数据库根据serial获取
		UserEmail:         email,
		UserLastLoginTime: int(time.Now().Unix()*1000),
		UserRegisterTime:  int(time.Now().Unix()*1000),
		UserPassword:      saltPassword,
		UserType:          commonConst.SmallUser,
		Salt:              salt,
		Reserved2:         " ",
	}
	return uai
}

func NewDefaultUpi(uid int, contactEmail string) *table.UserPersonalInformation {
	return &table.UserPersonalInformation{
		uid,
		commonConst.DefaultPhotoUrl,
		commonConst.DefaultUserName,
		commonConst.DefaultUserSex,
		commonConst.DefaultUserPhone,
		contactEmail,
		int(time.Now().Unix()*1000),
		" ",
		" ",
	}
}

