package encrypt

import (
	"crypto/md5"
	"fmt"
)

func Md5sum(plaintext string) (ciphertext string) {
	h := md5.New()
	h.Write([]byte(plaintext))
	ciphertext = fmt.Sprintf("%x", h.Sum(nil))
	return
}