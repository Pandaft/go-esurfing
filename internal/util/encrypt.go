package util

import (
	"crypto/md5"
	"fmt"
	"strings"
)

// CalMD5Hash 计算字符串的 MD5 哈希值
func CalMD5Hash(s string) string {
	hash := md5.New()
	hash.Write([]byte(s))
	bytes := hash.Sum(nil)
	str := fmt.Sprintf("%x", bytes)
	return strings.ToUpper(str)
}
