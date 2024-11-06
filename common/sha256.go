package common

import (
	"crypto/sha1"
	"encoding/hex"
)

// 使用SHA-256加密
func SHAEncoding(data string) string {
	// 创建 sha256 哈希对象
	hash := sha1.New()
	// 写入数据
	hash.Write([]byte(data))
	// 计算哈希值
	hashSum := hash.Sum(nil)
	return hex.EncodeToString(hashSum)
}
