package utils

import (
	jwt "github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"
	"time"
	"toolbox/internal/errors"
	"toolbox/internal/logs"
)

// 定义过期时间
const TokenExpireDuration = time.Hour * 24

// 定义secret
var MySecret = []byte("7uORCD1Z7GGTsWvqfhqBhnZONBOvt4uq")

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenToken生成jwt
func GenToken(username string) (string, error) {
	c := MyClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "toolbox",
		},
	}
	//使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	//使用指定的secret签名并获得完成的编码后的字符串token
	tk, err := token.SignedString(MySecret)
	if err != nil {
		logs.Warn("sign jwt failed",
			zap.Error(err),
			zap.String("username", username))
		return "", errors.ErrGenerateTokenFailed
	}
	return tk, nil
}

// ParseToken解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	//解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		logs.Warn("parse token failed",
			zap.Error(err),
			zap.String("token", tokenString))
		return nil, errors.ErrParseTokenFailed
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.ErrInvalidToken
}
