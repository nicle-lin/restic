package common

import (
	"github.com/astaxie/beego/orm"
	"os"
	"strings"
	"fmt"
	"crypto/md5"
)

//返回成功表示有错误
func CheckError(err error) bool {
	if err != nil && err != orm.ErrNoRows {
		return true
	}
	return false
}

//返回true就表示不存在
func CheckNoExist(err error) bool {
	if err != nil && err == orm.ErrNoRows {
		return true
	}
	return false
}

func IsPathExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil || os.IsNotExist(err) {
		return false
	}
	return true
}

func Md5sum(plaintext string) (ciphertext string) {
	h := md5.New()
	h.Write([]byte(plaintext))
	ciphertext = fmt.Sprintf("%x", h.Sum(nil))
	return
}

func Join(a []int64, sep string) string {
	return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), sep), "[]")
}

