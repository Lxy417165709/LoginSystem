package controls

import (
	"0_common/commonConst"
	"0_common/commonFunction"
	"0_common/commonStruct"
	"2_models/table"
	"3_transition"
	"checker"
	"github.com/astaxie/beego/logs"
	"net/http"
	"strings"
)

var checkerManager = checker.GetChecker()

func HandleErr(w http.ResponseWriter, Err *commonStruct.Error) {
	if Err != nil {
		if Err.ForDeveloper != nil && Err.ForUser != nil {
			logs.Error(Err.ForDeveloper)
			ResponseError(w, Err.ForUser)
			return
		}

		if Err.ForDeveloper != nil {
			logs.Error(Err.ForDeveloper)
			return
		}
		if Err.ForUser != nil {
			ResponseError(w, Err.ForUser)
			return
		}
	}
}

// 测试接口
func Test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

// 登录接口
func Login(w http.ResponseWriter, r *http.Request) {
	// 解析请求，获取内部数据
	loginData := commonStruct.LoginData{}
	if Err := ParseRequestData(r, &loginData); Err != nil {
		HandleErr(w, Err)
		return
	}
	// 数据加工
	loginData.Email = strings.ToLower(loginData.Email) // 统一小写

	// 校验
	if Err := checkerManager.Check(loginData); Err != nil {
		//for i:=3;i<6;i++{
		//	logs.SetLogFuncCallDepth(i)
		//	HandleErr(w, Err)
		//}
		HandleErr(w, Err)
		//HandleErr(w, Err)
		return
	}

	// 执行步骤 (可以单独分块)
	if err := transition.UpdateLastLoginTime(loginData.Email); err != nil {
		logs.Error(err)
		return
	}
	uid, err := transition.GetUid(loginData.Email)
	if err != nil {
		logs.Error(err)
		return
	}
	if err := SetUidToResponse(w, uid); err != nil {
		logs.Error(err)
		return
	}

	// 最终回复
	Response(w, &ResponseProto{
		Status: commonConst.LoginSuccessFlag,
		Msg:    "登录成功",
	})
}

// 注册接口
func Register(w http.ResponseWriter, r *http.Request) {
	// 从请求中获取注册数据
	registerData := commonStruct.RegisterData{}
	if Err := ParseRequestData(r, &registerData); Err != nil {
		HandleErr(w, Err)
		return
	}

	// 数据加工
	registerData.Email = strings.ToLower(registerData.Email) // 统一小写

	// 校验
	if Err := checkerManager.Check(registerData); Err != nil {
		HandleErr(w, Err)
		return
	}
	// 校验通过就删除验证码
	if Err := transition.DelRegisterVrc(registerData.Email); Err != nil {
		HandleErr(w, Err)
		return
	}
	// 产生新用户 (返回uid和错误)
	uid, Err := transition.GenerateNewUser(registerData.Email, registerData.Password)
	if Err != nil {
		HandleErr(w, Err)
		return
	}

	// 设置token
	if Err := SetUidToResponse(w, uid); Err != nil {
		HandleErr(w, Err)
		return
	}

	Response(w, &ResponseProto{
		Status: commonConst.RegisterSuccessFlag,
		Msg:    "注册成功",
	})
}

// 获取用户账号信息接口
func GetUai(w http.ResponseWriter, r *http.Request) {
	// 从请求中获取登录数据
	var uid int
	var uai *table.UserAccountInformation
	var Err *commonStruct.Error
	if uid, Err = GetUidFromRequest(r); Err != nil {
		HandleErr(w, Err)
		return
	}

	if uai, Err = transition.GetUai(uid); Err != nil {
		HandleErr(w, Err)
		return
	}

	Response(w, &ResponseProto{
		Status: commonConst.UaiGetSuccessFlag,
		Msg:    "用户账号信息获取成功",
		Data:   *uai,
	})
}

// 获取用户个人信息接口
func GetUpi(w http.ResponseWriter, r *http.Request) {
	// 从请求中获取登录数据
	var uid int
	var upi *table.UserPersonalInformation
	var Err *commonStruct.Error
	if uid, Err = GetUidFromRequest(r); Err != nil {
		HandleErr(w, Err)
		return
	}

	// 获取用户个人信息
	if upi, Err = transition.GetUpi(uid); Err != nil {
		HandleErr(w, Err)
		return
	}

	// 获取头像
	// 进行校验
	var base64Data string
	if base64Data, Err = transition.GetPhotoCheck(upi.UserPhotoUrl); Err != nil {
		HandleErr(w, Err)
		return
	}

	data := struct {
		table.UserPersonalInformation
		PhotoData string `json:"photoData"` // 头像数据
	}{
		*upi,
		base64Data,
	}

	Response(w, &ResponseProto{
		Status: commonConst.UpiGetSuccessFlag,
		Msg:    "用户个人信息获取成功",
		Data:   data,
	})
}

// 获取图片接口
func GetPhoto(w http.ResponseWriter, r *http.Request) {
	getPhotoData := commonStruct.GetPhotoData{}
	if Err := ParseRequestData(r, &getPhotoData); Err != nil {
		HandleErr(w, Err)
		return
	}

	// 进行校验
	var base64Data string
	var Err *commonStruct.Error
	if base64Data, Err = transition.GetPhoto(getPhotoData.PhotoName); Err != nil {
		HandleErr(w, Err)
		return
	}

	// 将base64Data封装入json中
	var rspDataBytes []byte
	if rspDataBytes, Err = FormatPhotoRspData(base64Data); Err != nil {
		HandleErr(w, Err)
		return
	}

	Response(w, &ResponseProto{
		Status: commonConst.GetPhotoSuccessFlag,
		Data:   string(rspDataBytes),
		Msg:    "图片获取成功",
	})
}

// 修改用户个人信息接口
func UpdateUpi(w http.ResponseWriter, r *http.Request) {
	upiData := commonStruct.UpiData{}
	if Err := ParseRequestData(r, &upiData); Err != nil {
		HandleErr(w, Err)
		return
	}
	var uid int
	var Err *commonStruct.Error
	if uid, Err = GetUidFromRequest(r); Err != nil {
		HandleErr(w, Err)
		return
	}

	// 校验
	if Err = checkerManager.Check(upiData); Err != nil {
		HandleErr(w, Err)
		return
	}

	// 执行
	if Err = transition.UpdateUpi(uid, upiData.UserName, upiData.UserContactEmail, upiData.UserContactPhone, upiData.UserBirthday, upiData.UserSex); Err != nil {
		HandleErr(w, Err)
		return
	}

	Response(w, &ResponseProto{
		Status: commonConst.UpiUpdateSuccessFlag,
		Msg:    "修改成功",
	})
}

// 修改用户头像接口
func UpdatePhoto(w http.ResponseWriter, r *http.Request) {
	// 数据解析
	updatePhotoData := commonStruct.UpdatePhotoData{}
	var Err *commonStruct.Error
	if Err = ParseRequestData(r, &updatePhotoData); Err != nil {
		HandleErr(w, Err)
		return
	}
	var uid int
	if uid, Err = GetUidFromRequest(r); Err != nil {
		HandleErr(w, Err)
		return
	}

	// 校验
	if Err = checkerManager.Check(updatePhotoData); Err != nil {
		HandleErr(w, Err)
		return
	}

	// 执行
	if Err = transition.UpdatePhoto(uid, updatePhotoData); Err != nil {
		HandleErr(w, Err)
		return
	}

	Response(w, &ResponseProto{
		Status: commonConst.PhotoUploadSuccessFlag,
		Msg:    "头像更新成功",
	})
}

// 发送注册验证码接口
func SendRegisterVrc(w http.ResponseWriter, r *http.Request) {
	evd := commonStruct.EmailData{}
	if Err := ParseRequestData(r, &evd); Err != nil {
		HandleErr(w, Err)
		return
	}

	// 校验
	if Err := checkerManager.Check(evd); Err != nil {
		HandleErr(w, Err)
		return
	}

	// 执行发送操作
	// 发送校验
	vrc := commonFunction.CreatVrc()
	if Err := transition.SendRegisterVrc(evd.Email, vrc); Err != nil {
		HandleErr(w, Err)
		return
	}
	// 保存操作
	if Err := transition.SetRegisterVrc(evd.Email, vrc, commonConst.RegisterExpiredTime); Err != nil {
		HandleErr(w, Err)
		return
	}

	Response(w, &ResponseProto{
		Status: commonConst.EmailSendSuccessFlag,
		Msg:    "注册验证码发送成功",
	})
}

// 发送修改密码链接接口
//func SendChangePasswordLink(w http.ResponseWriter, r *http.Request) {
//	evd := EmailData{}
//	if err := ParseRequestData(r, &evd); err != nil {
//		logs.Error(err)
//		ResponseError(w, err)
//		return
//	}
//	// 发送校验
//	if err := transition.SendChangePasswordLinkCheck(evd.Email, commonConst.ChangePasswordExpiredTime); err != nil {
//		logs.Error(err)
//		ResponseError(w, err)
//		return
//	}
//
//	Response(w, &ResponseProto{
//		Status: commonConst.EmailSendSuccessFlag,
//		Msg:    "changePassword link send success",
//	})
//}

// 访问修改密码链接
// 包括GET方法、POST方法
// GET用于验证链接是否有效
// POST用于验证修改密码操作是否可执行
//func ChangePasswordLinkVisit(w http.ResponseWriter, r *http.Request) {
//	linkEmail, linkVrc := r.URL.Query().Get("email"), r.URL.Query().Get("vrc")
//
//	// 表示访问修改密码的链接
//	if r.Method == "GET" {
//		if err := transition.ChangePasswordLinkVisitCheck(linkEmail, linkVrc); err != nil {
//			logs.Error(err)
//			ResponseError(w, err)
//			return
//		}
//		Response(w, &ResponseProto{
//			Status: commonConst.LinkValidFlag,
//			Msg:    "the link is valid",
//		})
//		return
//	}
//
//	// 表示执行密码修改
//	if r.Method == "POST" {
//		cpd := ChangePasswordData{}
//		if err := ParseRequestData(r, &cpd); err != nil {
//			logs.Error(err)
//			ResponseError(w, err)
//			return
//		}
//		if err := transition.ChangePasswordExecCheck(linkEmail, linkVrc, cpd.NewPassword); err != nil {
//			logs.Error(err)
//			ResponseError(w, err)
//			return
//		}
//
//		Response(w, &ResponseProto{
//			Status: commonConst.PasswordChangeFlag,
//			Msg:    "your password has been changed",
//		})
//		return
//	}
//
//}
