package commonStruct


// 登录需要两个字段 email,password
type LoginData struct {
	Email    string `json:"email" checkType:"email"`
	Password string `json:"password" checkType:"password"`
}


// 注册需要三个字段 email,password,vrc
type RegisterData struct {
	Email    string `json:"email" checkType:"email"`
	Password string `json:"password" checkType:"password"`
	Vrc      string `json:"vrc" checkType:"vrc"`
}

// 更新用户信息的数据结构
type UpiData struct {
	UserName         string `json:"userName" checkType:"email"`
	UserSex          int    `json:"userSex" checkType:"sex"`
	UserContactPhone string `json:"userContactPhone" checkType:"phone"`
	UserContactEmail string `json:"userContactEmail" checkType:"email"`
	UserBirthday     int    `json:"userBirthday" checkType:"birthday"`
}
// 修改用户头像数据结构
type UpdatePhotoData struct {
	PhotoBase64 string `json:"photoBase64" checkType:"photo"`
}

// 邮箱发送验证码数据结构
type EmailData struct {
	Email string `json:"email" checkType:"email"`
}

// 邮箱验证接口数据结构
type EmailVerificationData struct {
	Email string `json:"email"`
	Vrc   string `json:"vrc"`
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
