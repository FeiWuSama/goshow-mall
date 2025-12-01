package md5

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5 计算字符串的MD5值
// 参数: str - 要加密的字符串
// 返回: 32位小写MD5字符串
func MD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// MD5WithSalt 使用盐值进行MD5加密
// 参数: str - 要加密的字符串, salt - 盐值
// 返回: 加盐后的MD5字符串
func MD5WithSalt(str, salt string) string {
	// 可以根据需要调整盐值的使用方式
	// 方式1: str + salt
	// 方式2: salt + str + salt
	// 这里使用方式1
	return MD5(str + salt)
}

// MD5Verify 验证字符串的MD5值
// 参数: str - 原始字符串, md5Str - 要验证的MD5字符串
// 返回: 是否匹配
func MD5Verify(str, md5Str string) bool {
	return MD5(str) == md5Str
}

// MD5VerifyWithSalt 验证加盐后的MD5值
// 参数: str - 原始字符串, salt - 盐值, md5Str - 要验证的MD5字符串
// 返回: 是否匹配
func MD5VerifyWithSalt(str, salt, md5Str string) bool {
	return MD5WithSalt(str, salt) == md5Str
}
