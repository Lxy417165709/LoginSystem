package table

import (
	"0_common/commonConst"
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

func(uai *UserAccountInformation)GetTableName() string{
	return commonConst.NameOfTableUai
}

