package utils

import "math/rand"

// 随机字符串
var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// RandStr 生成随机字串
func RandStr(strLen int) string {
	randBytes := make([]rune, strLen)
	for i := range randBytes {
		randBytes[i] = letters[rand.Intn(len(letters))]
	}
	return string(randBytes)
}
