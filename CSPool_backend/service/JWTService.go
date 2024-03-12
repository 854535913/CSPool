package service

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

const TokenExpireDuration = time.Hour * 24

var SecretKey = []byte("yingzuorui")

type CustomClaims struct {
	// 可根据需要自行添加字段
	UserID               int64  `json:"userid"`
	Username             string `json:"username"`
	jwt.RegisteredClaims        // 内嵌标准的声明
}

// GenerateTokenService 生成JWT
func GenerateTokenService(userid int64, username string) (string, error) {
	// 创建一个我们自己的声明
	claims := CustomClaims{
		userid,
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)),
			Issuer:    "yzr666.fun", // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(SecretKey)
}

// DecryptTokenService 解析JWT
func DecryptTokenService(tokenString string) (*CustomClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return SecretKey, nil
	})
	if err != nil {
		return nil, err
	}
	// 对token对象中的Claim进行类型断言
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
