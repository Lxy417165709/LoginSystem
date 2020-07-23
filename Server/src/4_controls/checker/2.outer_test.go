package checker

import (
	"fmt"
	"testing"
)
// 注册需要三个字段 email,password,vrc
type RegisterData struct {
	Email    string `json:"email" checkType:"email"`
	Password string `json:"password" checkType:"password"`
	Vrc      string `json:"vrc"  checkType:"vrc"`
}


func TestChecker_Check(t *testing.T) {
	ch := GetChecker()
	testSlice := []RegisterData{
		{
			"417165709@qq.com",
			"123456789",
			"123456",
		},
		{
			"417165709.com",
			"123789",
			"123456110",
		},
		{
			"417165707279@qq.com",
			"123456789",
			"1234256",
		},
		{
			"417165709@qq.com",
			"",
			"16",
		},
	}

	for i:=0;i<len(testSlice);i++{
		fmt.Println(i,ch.Check(testSlice[i]))
	}
}
