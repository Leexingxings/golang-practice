package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"practice/utils"
	"time"
)

const NormalAccessTokenTTL = time.Hour * 2
const FreshTokenTTL = time.Hour * 24 * 30

type CustomPayload struct {
	UserID uint64 `json:"uid,string"`
	Name   string `json:"name"`
}

type Payload struct {
	CustomPayload
	jwt.RegisteredClaims // 内嵌标准的声明
}

var secret []byte

func init() {
	secret = []byte("this-is-my-blog")
}

// BuildJWT 构建JWT
// customPayload 自定义字段
// ttl 过期时间
func BuildJWT(customPayload CustomPayload, ttl time.Duration) (string, error) {
	commonClaims := jwt.RegisteredClaims{
		ID:        utils.RandStr(10), // ID, 类似于盐值
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
	}
	claims := Payload{
		customPayload,
		commonClaims,
	}

	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(secret)
}

// ParseJWT 解析Token
// tokenStr JWT
func ParseJWT(tokenStr string) (*Payload, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Payload{}, func(token *jwt.Token) (i any, err error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}

	// 对token对象中的Claim进行类型断言
	if claims, ok := token.Claims.(*Payload); ok && token.Valid {
		// 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
