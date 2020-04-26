package cookie

import (
	"0_common/commonConst"
	"0_common/commonFunction"
	"4_controls/interaction/cookie/token"
	"fmt"
	"github.com/astaxie/beego/logs"
	"net/http"
)


type cookieManager struct {
	tokenManager *token.TokenManager
}
var ckm = &cookieManager{token.GetTokenManager()}
func GetCookieManager() *cookieManager{
	return  ckm
}

// 设置token(密文)
// token(密文) -> 响应
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
// 请求 -> token(明文)
func getTokenFromRequest(r *http.Request) (string, error) {
	cookie, err := &http.Cookie{}, fmt.Errorf("")
	if cookie, err = r.Cookie("token"); err != nil {
		return "", err
	}
	return commonFunction.Decode(cookie.Value, commonConst.AESKey)
}

// uid -> 响应
func (cm *cookieManager) SetUid(w http.ResponseWriter, userId int) error {
	var tokenString string
	var err error
	if tokenString, err = cm.tokenManager.GetToken(userId); err != nil {
		logs.Error(err)
		return err
	}
	if err = setTokenToResponse(w, tokenString); err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

// 请求 -> uid
func (cm *cookieManager) GetUid(r *http.Request) (int, error) {
	var tokenString string
	var err error
	if tokenString, err = getTokenFromRequest(r); err != nil {
		logs.Info(err)
		return 0, err
	}
	var uid int
	if uid, err = cm.tokenManager.ParseToken(tokenString); err != nil {
		logs.Info(err)
		return 0,err
	}

	return uid, nil
}
