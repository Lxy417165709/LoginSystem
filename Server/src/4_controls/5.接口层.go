package controls

import (
	"0_common/commonConst"
	"0_common/commonFunction"
	"4_controls/interaction/checker"
	"0_common/commonStruct"
	"2_models/table"
	"3_transition"
	"4_controls/interaction/cookie"
	"errors"
	"github.com/astaxie/beego/logs"
	"4_controls/interaction"
	"net/http"
	"strings"
)

var checkerManager = checker.GetChecker()
var cookieManager = cookie.GetCookieManager()
var interactionManger = interaction.GetInteractionManger()

func HandleErr(w http.ResponseWriter, Err *commonStruct.Error) {
	if Err != nil {
		if Err.ForDeveloper != nil && Err.ForUser != nil {
			logs.Error(Err.ForDeveloper)
			interactionManger.ResponseError(w, Err.ForUser)
			return
		}

		if Err.ForDeveloper != nil {
			logs.Error(Err.ForDeveloper)
			return
		}
		if Err.ForUser != nil {
			interactionManger.ResponseError(w, Err.ForUser)
			return
		}
	}
}



// 对表单进行校验，并返回响应结果
func CheckPart(w http.ResponseWriter,data interface{}) bool{

	if result := checkerManager.GetCheckResult(data); result != nil {
		if result.ForDeveloper != nil{
			logs.Error(result.ForDeveloper)
		}
		interactionManger.ResponseError(w, result.ForUser)
		return false
	}
	return true
}




// 测试接口
func Test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

// 登录接口
func Login(w http.ResponseWriter, r *http.Request) {
	// 解析请求，获取内部数据
	var loginData interaction.LoginData
	var err error
	if err = interactionManger.FormData(r, &loginData); err != nil {
		logs.Error(err)
		interactionManger.ResponseError(w,errors.New("登录表单获取失败"))
		return
	}

	// 数据加工
	loginData.Email = strings.ToLower(loginData.Email) // 统一小写

	// 校验
	if CheckPart(w,loginData)==false{
		return
	}

	// 执行步骤 (可以单独分块)
	if err = transition.UpdateLastLoginTime(loginData.Email); err != nil {
		logs.Error(err)
		interactionManger.ResponseError(w,errors.New("更新用户最近登录时间失败"))
		return
	}

	// 获取uid
	var uid int
	if uid, err = transition.GetUid(loginData.Email); err != nil {
		logs.Error(err)
		interactionManger.ResponseError(w,errors.New("获取 UserId 失败"))
		return
	}
	if err = cookieManager.SetUid(w, uid); err != nil {
		logs.Error(err)
		interactionManger.ResponseError(w,errors.New("token 设置失败"))
		return
	}

	// 最终回复
	interactionManger.Response(w, &interaction.ResponseProto{
		Status: commonConst.LoginSuccessFlag,
		Msg:    "登录成功",
	})
}

// 注册接口
func Register(w http.ResponseWriter, r *http.Request) {
	// 从请求中获取注册数据
	var registerData interaction.RegisterData
	var err error
	if err = interactionManger.FormData(r, &registerData); err != nil {
		logs.Error(err)
		interactionManger.ResponseError(w,errors.New("注册表单获取失败"))
		return
	}

	// 数据加工
	registerData.Email = strings.ToLower(registerData.Email) // 统一小写

	// 校验
	if CheckPart(w,registerData)==false{
		return
	}


	// 校验通过就删除验证码
	if err = transition.DelRegisterVrc(registerData.Email); err != nil {
		logs.Error(err)
		interactionManger.ResponseError(w,errors.New("注册验证码删除失败"))
		return
	}
	// 产生新用户 (返回uid和错误)
	var uid int

	if uid, err = transition.GenerateNewUser(registerData.Email, registerData.Password);err != nil {
		logs.Error(err)
		interactionManger.ResponseError(w,errors.New("用户创建失败"))
		return
	}

	// 设置token
	if err = cookieManager.SetUid(w, uid); err != nil {
		logs.Error(err)
		interactionManger.ResponseError(w,errors.New("token 设置失败"))
		return
	}

	interactionManger.Response(w, &interaction.ResponseProto{
		Status: commonConst.RegisterSuccessFlag,
		Msg:    "注册成功",
	})
}

// 获取用户账号信息接口
func GetUai(w http.ResponseWriter, r *http.Request) {
	// 从请求中获取登录数据
	var uid int
	var uai *table.UserAccountInformation
	var err error
	if uid, err = cookieManager.GetUid(r); err != nil {
		logs.Error(err)
		interactionManger.ResponseError(w,errors.New("token 解析失败"))
		return
	}

	if uai, err = transition.GetUai(uid); err != nil {
		logs.Error(err)
		interactionManger.ResponseError(w,errors.New("用户账号信息获取失败"))
		return
	}

	interactionManger.Response(w, &interaction.ResponseProto{
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
	var err error
	if uid, err = cookieManager.GetUid(r); err != nil {
		logs.Error(err)
		interactionManger.ResponseError(w,errors.New("token 解析失败"))
		return
	}

	if upi, err = transition.GetUpi(uid); err != nil {
		logs.Error(err)
		interactionManger.ResponseError(w,errors.New("用户个人信息获取失败"))
		return
	}

	// 获取头像
	// 进行校验
	var base64Data string
	if base64Data, err = transition.GetPhoto(upi.UserPhotoUrl); err != nil {
		logs.Error(err)
		interactionManger.ResponseError(w,errors.New("用户头像获取失败"))
		return
	}

	data := struct {
		table.UserPersonalInformation
		PhotoData string `json:"photoData"` // 头像数据
	}{
		*upi,
		base64Data,
	}

	interactionManger.Response(w, &interaction.ResponseProto{
		Status: commonConst.UpiGetSuccessFlag,
		Msg:    "用户个人信息获取成功",
		Data:   data,
	})
}

// 获取图片接口
func GetPhoto(w http.ResponseWriter, r *http.Request) {
	var getPhotoData interaction.GetPhotoData
	var err error
	if err := interactionManger.FormData(r, &getPhotoData); err != nil {
		interactionManger.ResponseError(w,err)
		return
	}

	// 进行校验
	var base64Data string
	if base64Data, err = transition.GetPhoto(getPhotoData.PhotoName); err != nil {
		logs.Error(err)
		interactionManger.ResponseError(w,errors.New("图片获取失败"))
		return
	}

	// 将base64Data封装入json中
	var rspDataBytes []byte
	if rspDataBytes, err = interactionManger.GetPhotoRspData(base64Data); err != nil {
		logs.Error(err)
		interactionManger.ResponseError(w,errors.New("生成图片响应结构失败"))
		return
	}

	interactionManger.Response(w, &interaction.ResponseProto{
		Status: commonConst.GetPhotoSuccessFlag,
		Data:   string(rspDataBytes),
		Msg:    "图片获取成功",
	})
}

// 修改用户个人信息接口
func UpdateUpi(w http.ResponseWriter, r *http.Request) {
	var upiData interaction.UpiData
	var err error
	if err = interactionManger.FormData(r, &upiData); err != nil {
		logs.Error(err)
		interactionManger.ResponseError(w,errors.New("用户个人信息表单获取失败"))
		return
	}

	var uid int
	if uid, err = cookieManager.GetUid(r); err != nil {
		logs.Error(err)
		interactionManger.ResponseError(w,errors.New("UserId 获取失败"))
		return
	}

	// 校验
	if CheckPart(w,upiData)==false{
		return
	}

	// 执行
	if err = transition.UpdateUpi(uid, upiData.UserName, upiData.UserContactEmail, upiData.UserContactPhone, upiData.UserBirthday, upiData.UserSex); err != nil {
		logs.Error(err)
		interactionManger.ResponseError(w,errors.New("用户个人信息更新失败"))
		return
	}

	interactionManger.Response(w, &interaction.ResponseProto{
		Status: commonConst.UpiUpdateSuccessFlag,
		Msg:    "更新成功",
	})
}

// 修改用户头像接口
func UpdatePhoto(w http.ResponseWriter, r *http.Request) {
	// 数据解析
	var updatePhotoData interaction.UpdatePhotoData
	var err error
	if err = interactionManger.FormData(r, &updatePhotoData); err != nil {
		logs.Error(err)
		interactionManger.ResponseError(w,errors.New("图片数据表单获取失败"))
		return
	}

	var uid int
	if uid, err = cookieManager.GetUid(r); err != nil {
		logs.Error(err)
		interactionManger.ResponseError(w,errors.New("UserId 获取失败"))
		return
	}

	// 校验
	if CheckPart(w,updatePhotoData)==false{
		return
	}

	// 执行
	if err = transition.UpdatePhoto(uid, updatePhotoData.PhotoBase64); err != nil {
		logs.Error(err)
		interactionManger.ResponseError(w,errors.New("头像更新失败"))
		return
	}

	interactionManger.Response(w, &interaction.ResponseProto{
		Status: commonConst.PhotoUploadSuccessFlag,
		Msg:    "头像更新成功",
	})
}

// 发送注册验证码接口
func SendRegisterVrc(w http.ResponseWriter, r *http.Request) {
	var evd interaction.EmailData
	var err error
	if err = interactionManger.FormData(r, &evd); err != nil {
		logs.Error(err)
		interactionManger.ResponseError(w,errors.New("邮箱表单获取失败"))
		return
	}
	// 处理
	evd.Email = strings.ToLower(evd.Email)

	// 校验
	if CheckPart(w,evd)==false{
		return
	}

	// 执行发送操作
	// 发送校验
	var vrc = commonFunction.CreatVrc()
	if err := transition.SendRegisterVrc(evd.Email, vrc); err != nil {
		logs.Error(err)
		interactionManger.ResponseError(w,errors.New("发送注册验证码失败"))
		return
	}
	// 保存操作
	if err = transition.SetRegisterVrc(evd.Email, vrc, commonConst.RegisterExpiredTime); err != nil {
		logs.Error(err)
		interactionManger.ResponseError(w,errors.New("保存注册验证码失败"))
		return
	}

	interactionManger.Response(w, &interaction.ResponseProto{
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
//}
