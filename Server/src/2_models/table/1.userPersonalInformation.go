package table

import (
	"0_common/commonConst"
)

// 对应数据库的 tb_userPersonalInformation 表
type UserPersonalInformation struct {
	UserId           int    `json:"userId" isPrimaryKey:"true"`
	UserPhotoUrl     string `json:"userPhotoUrl"`
	UserName         string `json:"userName"`
	UserSex          int    `json:"userSex"`
	UserContactPhone string `json:"userContactPhone"`
	UserContactEmail string `json:"userContactEmail"`
	UserBirthday     int    `json:"userBirthday"`
	Reserved1        string `json:"reserved1"`
	Reserved2        string `json:"reserved2"`
}
func(upi *UserPersonalInformation)GetTableName() string{
	return commonConst.NameOfTableUpi
}

// 和newUpi比较，如果不同，则修改upi
func (upi *UserPersonalInformation)Update(newUpi *UserPersonalInformation) {

	//if upi == nil{
	//	if newUpi==nil{
	//		return
	//	}
	//	upi = newUpi
	//	return
	//}
	if !isZero(newUpi.UserId){
		upi.UserId = newUpi.UserId
	}
	if !isZero(newUpi.UserPhotoUrl){
		upi.UserPhotoUrl = newUpi.UserPhotoUrl
	}
	if !isZero(newUpi.UserName){
		upi.UserName= newUpi.UserName
	}
	if !isZero(newUpi.UserSex){
		upi.UserSex = newUpi.UserSex
	}
	if !isZero(newUpi.UserContactPhone){
		upi.UserContactPhone = newUpi.UserContactPhone
	}
	if !isZero(newUpi.UserContactEmail){
		upi.UserContactEmail = newUpi.UserContactEmail
	}
	if !isZero(newUpi.UserBirthday){
		upi.UserBirthday = newUpi.UserBirthday
	}
	if !isZero(newUpi.Reserved1){
		upi.Reserved1 = newUpi.Reserved1
	}
	if !isZero(newUpi.Reserved2){
		upi.Reserved2 = newUpi.Reserved2
	}
}
