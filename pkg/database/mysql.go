package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"sync"
)

var db *sqlx.DB
var once sync.Once

// InitDB 初始化DB
func InitDB() {
	once.Do(func() {
		dsn := "root:lijiaxing@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True"
		db = sqlx.MustConnect("mysql", dsn)

		db.SetMaxOpenConns(100)
		db.SetMaxIdleConns(50)
	})
}

// ShareBlog 获取blog库连接
func ShareBlog() *sqlx.DB {
	return db
}
