package controls

import (
	"0_common/commonConst"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type UserClaims struct {
	jwt.StandardClaims
	UserId int
}

// 以 UserClaims 为载荷创建 tokenString ，UserClaims内含userId
// userId -> TokenString
func CreatTokenString(userId int) (string, error) {
	// jwt 的 载荷部分
	userClaims := UserClaims{
		jwt.StandardClaims{
			ExpiresAt: int64(time.Now().Add(time.Hour * 72).Unix()),
		},
		userId,
	}

	// jwt 的 头部、载荷部分
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)

	// jwt 的 头部、载荷、签名部分，并通过base64编码，获得完整token的字符串形式
	return token.SignedString([]byte(commonConst.TokenKey))
}

// 从tokenString中获取userId
// tokenString -> userId
func GetUserIdFromTokenString(tokenString string) (int, error) {

	var mapClaims jwt.MapClaims
	var token *jwt.Token
	var err error
	// tokenString -> token结构
	if token, err = jwt.Parse(tokenString, func(*jwt.Token) (interface{}, error) {
		return []byte(commonConst.TokenKey), nil
	}); err != nil {
		return commonConst.ErrorUserId, err
	}
	// token结构 -> mapClaims(内含userId)
	var ok bool
	if mapClaims, ok = token.Claims.(jwt.MapClaims); !ok {
		return commonConst.ErrorUserId, fmt.Errorf("the assert for mapClaims is error")
	}
	if err = mapClaims.Valid(); err != nil {
		return commonConst.ErrorUserId, err
	}

	//  mapClaims(内含userId) -> userId
	if _, ok = mapClaims["UserId"].(float64); !ok {
		return commonConst.ErrorUserId, fmt.Errorf("the assert for UserId is error")
	}
	return int(mapClaims["UserId"].(float64)), nil
}


