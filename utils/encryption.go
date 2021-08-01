package utils

import (
	"crypto/sha256"
	"fmt"
	"io"
)

const slatLen = 8	// 随机盐的长度

// Encryption 加密密码
// @param password 原密码
// @return pw 加密后的密码
// @return salt 加密使用的盐
func Encryption(password string) (pw string, salt string) {
	salt = RandString(slatLen)
	return EncryptionBySalt(password, salt), salt
}

// EncryptionBySalt 使用指定盐加密密码
// @param password 原密码
// @param salt 加密的盐
// @return string 加密后的密码
func EncryptionBySalt(password, salt string) string {
	h := sha256.New()
	io.WriteString(h, password)
	io.WriteString(h, salt)
	return fmt.Sprintf("%x", h.Sum(nil))
}
