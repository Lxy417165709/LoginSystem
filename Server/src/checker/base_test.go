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



// 邮箱发送数据结构
type EmailData struct {
	Email string `json:"email"`
}

// 邮箱验证接口数据结构
type EmailVerificationData struct {
	Email string `json:"email"`
	Vrc   string `json:"vrc"`
}

// 修改用户头像数据结构
type UploadPhotoData struct {
	PhotoBase64 string `json:"photoBase64"`
}

// 获取图片接口数据结构
type GetPhotoData struct {
	PhotoName string `json:"photoName"`
}

// 图片响应数据结构
type GetPhotoRspData struct {
	PhotoBase64 string `json:"photoBase64"`
}

// 修改密码数据结构
type ChangePasswordData struct {
	NewPassword string `json:"newPassword"`
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
