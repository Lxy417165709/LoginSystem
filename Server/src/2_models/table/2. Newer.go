package table

import (
	"0_common/commonConst"
	"0_common/commonFunction"
	"time"
)

func NewDefaultUai(uid int, email, password string) *UserAccountInformation {
	var saltPassword string
	var err error
	salt := commonFunction.CreatSalt()
	if saltPassword, err = commonFunction.SaltHash(password, salt); err != nil {
		return nil
	}
	// 插入账户信息
	uai := &UserAccountInformation{
		UserId:            uid, // 让数据库根据serial获取
		UserEmail:         email,
		UserLastLoginTime: int(time.Now().Unix()*commonConst.BirthDayRato),
		UserRegisterTime:  int(time.Now().Unix()*commonConst.BirthDayRato),
		UserPassword:      saltPassword,
		UserType:          commonConst.SmallUser,
		Salt:              salt,
		Reserved2:         " ",
	}
	return uai
}

func NewDefaultUpi(uid int, contactEmail string) *UserPersonalInformation {
	return &UserPersonalInformation{
		uid,
		commonConst.DefaultPhotoUrl,
		commonConst.DefaultUserName,
		commonConst.DefaultUserSex,
		commonConst.DefaultUserPhone,
		contactEmail,
		int(time.Now().Unix()*commonConst.BirthDayRato),
		" ",
		" ",
	}
}

