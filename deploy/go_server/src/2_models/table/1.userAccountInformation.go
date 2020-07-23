package table

import (
	"0_common/commonConst"
	"fmt"
	"github.com/astaxie/beego/logs"
)

// 对应数据库的 tb_userAccountInformation 表
type UserAccountInformation struct {
	UserId            int    `json:"userId" isPrimaryKey:"true"`
	UserEmail         string `json:"userEmail"`
	UserPassword      string `json:"userPassword"`
	UserType          int    `json:"userType"`
	UserRegisterTime  int    `json:"userRegisterTime"`
	UserLastLoginTime int    `json:"userLastLoginTime"`
	Salt              string `json:"salt"`
	Reserved2         string `json:"reserved2"`
}

func (uai *UserAccountInformation) GetTableName() string {
	return commonConst.NameOfTableUai
}

func (uai *UserAccountInformation) Update(newUai *UserAccountInformation) {
	if !isZero(newUai.UserId) {
		uai.UserId = newUai.UserId
	}
	if !isZero(newUai.UserEmail) {
		uai.UserEmail = newUai.UserEmail
	}
	if !isZero(newUai.UserPassword) {
		uai.UserPassword = newUai.UserPassword
	}
	if !isZero(newUai.UserType) {
		uai.UserType = newUai.UserType
	}
	if !isZero(newUai.UserRegisterTime) {
		uai.UserRegisterTime = newUai.UserRegisterTime
	}
	if !isZero(newUai.UserLastLoginTime) {
		uai.UserLastLoginTime = newUai.UserLastLoginTime
	}
	if !isZero(newUai.Salt) {
		uai.Salt = newUai.Salt
	}
	if !isZero(newUai.Reserved2) {
		uai.Reserved2 = newUai.Reserved2
	}
}

func isZero(val interface{}) bool {
	switch val.(type) {
	case int:
		return val == 0
	case string:
		return val == ""
	case float64:
		return val == 0.0
	default:
		logs.Error(fmt.Sprintf("%T 该类型不支持判零", val))
		return false
	}
}
