package test

import (
	"fmt"
	"practice/pkg/database"
	"testing"
)

type user struct {
	Id   uint64
	Name string
}

// TestMysqlConn 测试MySQL连接
func TestMySQLConn(t *testing.T) {
	// 初始化数据库
	database.InitDB()

	var userInfo user
	sqlStr := "SELECT * FROM user"
	err := database.ShareBlog().Get(&userInfo, sqlStr)
	if err != nil {
		t.Errorf("Get user failed, %s", err.Error())
	}

	fmt.Printf("%+v\n", userInfo)
}
