package token

import (
	"0_common/commonConst"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/dgrijalva/jwt-go"
	"time"
)



type UserClaims struct {
	jwt.StandardClaims
	UserId int
}

type TokenManager struct {
	tokenDuration time.Duration
	tokenKey      string
	uidKey 		  string	// 等于 UserClaims uid的字段名
}

var tokenManager = &TokenManager{time.Hour * 72, commonConst.TokenKey,"UserId"}

func GetTokenManager() *TokenManager {
	return tokenManager
}

// uid -> tokenString
func (tm *TokenManager) GetToken(uid int) (string, error) {
	// jwt 的 载荷部分
	var userClaims UserClaims
	userClaims = UserClaims{
		jwt.StandardClaims{
			ExpiresAt: int64(time.Now().Add(tm.tokenDuration).Unix()),
		},
		uid,
	}

	// jwt 的 头部、载荷部分
	var token *jwt.Token
	token = jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)

	// jwt 的 头部、载荷、签名部分，并通过base64编码，获得完整token的字符串形式
	var tokenStr string
	var err error
	if tokenStr, err = token.SignedString([]byte(tm.tokenKey)); err != nil {
		return "", err
	}
	return tokenStr, nil
}

// tokenString -> uid
func(tm *TokenManager) ParseToken(tokenString string) (int, error) {



	// tokenString -> token结构
	var token *jwt.Token
	var err error
	if token, err = jwt.Parse(tokenString, func(*jwt.Token) (interface{}, error) {
		return []byte(tm.tokenKey), nil
	}); err != nil {
		logs.Error(err)
		return 0, err
	}

	// token结构 -> mapClaims(内含userId)
	var mapClaims jwt.MapClaims
	var ok bool
	if mapClaims, ok = token.Claims.(jwt.MapClaims); !ok {
		return 0, fmt.Errorf("the assert for mapClaims is error")
	}
	if err = mapClaims.Valid(); err != nil {
		return 0, err
	}

	//  mapClaims(内含userId) -> userId
	if _, ok = mapClaims[tm.uidKey].(float64); !ok {
		return 0, fmt.Errorf("the assert for UserId is error")
	}
	return int(mapClaims[tm.uidKey].(float64)), nil
}
