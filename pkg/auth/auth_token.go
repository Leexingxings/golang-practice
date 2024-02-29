package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const TokenExpireDuration = time.Hour * 2

type CustomClaims struct {
	Uid                  uint64 `json:"uid,string"`
	jwt.RegisteredClaims        // 内嵌标准的声明
}

var salt = []byte("this-is-my-blog")

func BuildToken(uid uint64) (string, error) {
	claims := CustomClaims{
		Uid: uid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)),
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(salt)
}

func ParseToken(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (i any, err error) {
		return salt, nil
	})
	if err != nil {
		return nil, err
	}

	// 对token对象中的Claim进行类型断言
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		// 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
