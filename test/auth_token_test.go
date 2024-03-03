package test

import (
	"fmt"
	"practice/pkg/auth"
	"testing"
)

// TestAuthToken 测试JWT生成与解析
func TestAuthToken(t *testing.T) {
	payload := auth.CustomPayload{
		UserID: 1348327948732,
		Name:   "李星星",
	}

	token, err := auth.BuildJWT(payload, auth.NormalAccessTokenTTL)
	if err != nil {
		t.Errorf("Build token failed: %s", err.Error())
	}
	fmt.Println("生成token为：", token)

	claims, err := auth.ParseJWT(token)
	if err != nil {
		t.Errorf("Parse token failed: %s", err.Error())
	}
	fmt.Printf("解析token为：%+v\n", claims)
}
