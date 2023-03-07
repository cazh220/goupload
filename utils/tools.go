package utils

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"os"
	"strings"
)

/**
 * 字符串md5
 * str 待加密字符串
 * salt 盐
 * 返回值 md5加密串
 */
func Md5Crypt(str string, salt ...interface{}) (CryptStr string) {
	if l := len(salt); l > 0 {
		slice := make([]string, l+1)
		str = fmt.Sprintf(str+strings.Join(slice, "%v"), salt...)
	}
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

/**
 * base64 编码
 */
func Base64EnCode(str string) (CryptStr string) {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

/**
 * 判断目录或文件是否存在
 */
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
