package controls

import (
	"0_common/commonConst"
	"0_common/commonFunction"
	"0_common/commonStruct"
	"fmt"
	"net/http"
)

// 设置token(密文)
func setTokenToResponse(w http.ResponseWriter, token string) error {
	encodeToken, err := commonFunction.Encode(token, commonConst.AESKey)
	if err != nil {
		return err
	}
	http.SetCookie(w, &http.Cookie{
		Name:  "token",
		Value: encodeToken,
	})
	return nil
}

// 获取token(明文)
func getTokenFromRequest(r *http.Request) (string, error) {
	cookie, err := &http.Cookie{}, fmt.Errorf("")
	if cookie, err = r.Cookie("token"); err != nil {
		return "", err
	}
	return commonFunction.Decode(cookie.Value, commonConst.AESKey)
}



func SetUidToResponse(w http.ResponseWriter, userId int) *commonStruct.Error {
	var tokenString string
	var err error
	if tokenString, err = CreatTokenString(userId); err != nil {
		return commonStruct.NewError(
			fmt.Errorf("登录状态保留失败"),
			err,
		)
	}
	if err := setTokenToResponse(w, tokenString); err != nil {
		return commonStruct.NewError(
			fmt.Errorf("登录状态保留失败"),
			err,
		)
	}
	return nil
}

func GetUidFromRequest(r *http.Request) (int,*commonStruct.Error) {
	var tokenString string
	var err error
	if tokenString, err = getTokenFromRequest(r); err != nil {
		return commonConst.ErrorUserId,commonStruct.NewError(
			fmt.Errorf("获取用户id失败"),
			err,
		)
	}
	var uid int
	if uid,err = GetUserIdFromTokenString(tokenString) ;err!=nil{
		return commonConst.ErrorUserId,commonStruct.NewError(
			fmt.Errorf("获取用户id失败"),
			err,
		)
	}

	return uid,nil
}
