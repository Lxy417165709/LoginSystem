package controls

import (
	"0_common/commonConst"
	"2_models/table"
	"3_transition"
	"fmt"
	"github.com/astaxie/beego/logs"
	"net/http"
	"strings"
)

// 测试接口
func Test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

// 修改用户个人信息接口 (待写)
func UpdateUpi(w http.ResponseWriter, r *http.Request){
	upiData := table.UserPersonalInformation{}
	if err := ParseRequestData(r, &upiData); err != nil {
		logs.Error(err)
		ResponseError(w, err)
		return
	}
	var uid int
	var err error
	if uid, err = GetUidFromRequest(r); err != nil {
		logs.Error(err)
		ResponseError(w, err)
		return
	}

	upiData.UserId = uid
	fmt.Println(upiData)
	if err = transition.UpdateUpiCheck(upiData); err != nil {
		logs.Error(err)
		ResponseError(w, err)
		return
	}

	Response(w, &ResponseProto{
		Status: commonConst.UpiUpdateSuccessFlag,
		Msg:    "修改成功",
	})
}




// 注册接口
func Register(w http.ResponseWriter, r *http.Request) {
	// 从请求中获取注册数据
	registerData := RegisterData{}
	if err := ParseRequestData(r, &registerData); err != nil {
		logs.Error(err)
		ResponseError(w, err)
		return
	}
	registerData.Email = strings.ToLower(registerData.Email) // 统一小写

	var uid int
	var err error
	// 注册验证
	if uid, err = transition.RegisterCheck(registerData.Email, registerData.Password, registerData.Vrc); err != nil {
		logs.Error(err)
		ResponseError(w, err)
		return
	}

	// 设置token
	if err := SetUidToResponse(w, uid); err != nil {
		logs.Error(err)
		ResponseError(w, err)
		return
	}

	Response(w, &ResponseProto{
		Status: commonConst.RegisterSuccessFlag,
		Msg:    "注册成功",
	})
}

// 登录接口
func Login(w http.ResponseWriter, r *http.Request) {
	// 从请求中获取登录数据
	loginData := LoginData{}
	if err := ParseRequestData(r, &loginData); err != nil {
		logs.Error(err)
		ResponseError(w, err)
		return
	}
	loginData.Email = strings.ToLower(loginData.Email) // 统一小写

	// 登录校验
	var uid int
	var err error
	if uid, err = transition.LoginCheck(loginData.Email, loginData.Password); err != nil {
		logs.Error(err)
		ResponseError(w, err)
		return
	}

	// 设置token
	if err := SetUidToResponse(w, uid); err != nil {
		logs.Error(err)
		ResponseError(w, err)
		return
	}

	Response(w, &ResponseProto{
		Status: commonConst.LoginSuccessFlag,
		Msg:    "登录成功",
	})

}

// 获取用户个人信息
func GetUai(w http.ResponseWriter, r *http.Request) {
	// 从请求中获取登录数据
	var uid int
	var uai *table.UserAccountInformation
	var err error
	if uid, err = GetUidFromRequest(r); err != nil {
		logs.Error(err)
		ResponseError(w, err)
		return
	}

	if uai, err = transition.GetUaiCheck(uid); err != nil {
		logs.Error(err)
		ResponseError(w, err)
		return
	}

	Response(w, &ResponseProto{
		Status: commonConst.UaiGetSuccessFlag,
		Msg:    "用户账号信息获取成功",
		Data:   *uai,
	})
}

// 获取用户个人信息
func GetUpi(w http.ResponseWriter, r *http.Request) {
	// 从请求中获取登录数据
	var uid int
	var upi *table.UserPersonalInformation
	var err error
	if uid, err = GetUidFromRequest(r); err != nil {
		logs.Error(err)
		ResponseError(w, err)
		return
	}

	if upi, err = transition.GetUpiCheck(uid); err != nil {
		logs.Error(err)
		ResponseError(w, err)
		return
	}

	// 获取头像
	// 进行校验
	var base64Data string
	if base64Data, err = transition.GetPhotoCheck(upi.UserPhotoUrl); err != nil {
		logs.Error(err)
		ResponseError(w, err)
		return
	}


	data := struct {
		table.UserPersonalInformation
		PhotoData string                        `json:"photoData"`
	}{
		*upi,
		base64Data,
	}
	//fmt.Println(data)
	Response(w, &ResponseProto{
		Status: commonConst.UpiGetSuccessFlag,
		Msg:    "用户个人信息获取成功",
		Data:   data,
	})
}

//// 获取用户头像
//func GetSelfPhoto(w http.ResponseWriter,r *http.Request){
//
//}

// 上传图片
func UpdatePhoto(w http.ResponseWriter, r *http.Request) {
	uploadPhotoData := UploadPhotoData{}
	if err := ParseRequestData(r, &uploadPhotoData); err != nil {
		logs.Error(err)
		ResponseError(w, err)
		return
	}
	var uid int
	var err error
	if uid, err = GetUidFromRequest(r); err != nil {
		logs.Error(err)
		ResponseError(w, err)
		return
	}

	if err := transition.UpdatePhotoCheck(uid, uploadPhotoData.PhotoBase64); err != nil {
		logs.Error(err)
		ResponseError(w, err)
		return
	}
	Response(w, &ResponseProto{
		Status: commonConst.PhotoUploadSuccessFlag,
		Msg:    "upload photo success",
	})
}

// 获取图片
func GetPhoto(w http.ResponseWriter, r *http.Request) {
	getPhotoData := GetPhotoData{}
	if err := ParseRequestData(r, &getPhotoData); err != nil {
		logs.Error(err)
		ResponseError(w, err)
		return
	}

	// 进行校验
	var base64Data string
	var err error
	if base64Data, err = transition.GetPhotoCheck(getPhotoData.PhotoName); err != nil {
		logs.Error(err)
		ResponseError(w, err)
		return
	}

	// 将base64Data封装入json中
	var rspDataBytes []byte
	if rspDataBytes, err = FormatPhotoRspData(base64Data); err != nil {
		logs.Error(err)
		ResponseError(w, err)
		return
	}

	Response(w, &ResponseProto{
		Status: commonConst.GetPhotoSuccessFlag,
		Data:   string(rspDataBytes),
		Msg:    "get photo success",
	})
}

// 发送修改密码链接
func SendChangePasswordLink(w http.ResponseWriter, r *http.Request) {
	evd := EmailData{}
	if err := ParseRequestData(r, &evd); err != nil {
		logs.Error(err)
		ResponseError(w, err)
		return
	}
	// 发送校验
	if err := transition.SendChangePasswordLinkCheck(evd.Email, commonConst.ChangePasswordExpiredTime); err != nil {
		logs.Error(err)
		ResponseError(w, err)
		return
	}

	Response(w, &ResponseProto{
		Status: commonConst.EmailSendSuccessFlag,
		Msg:    "changePassword link send success",
	})
}

// 发送注册验证码
func SendRegisterVrc(w http.ResponseWriter, r *http.Request) {
	evd := EmailData{}
	if err := ParseRequestData(r, &evd); err != nil {
		logs.Error(err)
		ResponseError(w, err)
		return
	}
	// 发送校验
	if err := transition.SendRegisterVrcCheck(evd.Email, commonConst.RegisterExpiredTime); err != nil {
		logs.Error(err)
		ResponseError(w, err)
		return
	}

	Response(w, &ResponseProto{
		Status: commonConst.EmailSendSuccessFlag,
		Msg:    "register vrc send success",
	})
}

// 访问修改密码链接
// 包括GET方法、POST方法
// GET用于验证链接是否有效
// POST用于验证修改密码操作是否可执行
func ChangePasswordLinkVisit(w http.ResponseWriter, r *http.Request) {
	linkEmail, linkVrc := r.URL.Query().Get("email"), r.URL.Query().Get("vrc")

	// 表示访问修改密码的链接
	if r.Method == "GET" {
		if err := transition.ChangePasswordLinkVisitCheck(linkEmail, linkVrc); err != nil {
			logs.Error(err)
			ResponseError(w, err)
			return
		}
		Response(w, &ResponseProto{
			Status: commonConst.LinkValidFlag,
			Msg:    "the link is valid",
		})
		return
	}

	// 表示执行密码修改
	if r.Method == "POST" {
		cpd := ChangePasswordData{}
		if err := ParseRequestData(r, &cpd); err != nil {
			logs.Error(err)
			ResponseError(w, err)
			return
		}
		if err := transition.ChangePasswordExecCheck(linkEmail, linkVrc, cpd.NewPassword); err != nil {
			logs.Error(err)
			ResponseError(w, err)
			return
		}

		Response(w, &ResponseProto{
			Status: commonConst.PasswordChangeFlag,
			Msg:    "your password has been changed",
		})
		return
	}

}
