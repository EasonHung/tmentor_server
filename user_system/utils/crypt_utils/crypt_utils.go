package crypt_utils

import (
	"crypto/md5"
	"fmt"
)

func EncryptedWithMD5(str string, key string) string {
	data := []byte(str)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}
