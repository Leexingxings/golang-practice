package test

import (
	"fmt"
	"practice/pkg/auth"
	"testing"
)

func TestAuthToken(t *testing.T) {
	token, err := auth.BuildToken(98672341121234)
	if err != nil {
		t.Errorf("Build token failed: %s", err.Error())
	}
	fmt.Println("生成token为：", token)

	cliams, err := auth.ParseToken(token)
	if err != nil {
		t.Errorf("Parse token failed: %s", err.Error())
	}
	fmt.Printf("解析token为：%+v\n", cliams)
}
