package user

import (
	"crypto/md5"
	"fmt"
	"io"
	"staff-mall-center/pkg/setting"
)

func CryptoHandler(pwd *string) {

	h := md5.New()

	// 	第一次md5加密
	io.WriteString(h, setting.CryptoSetting.Seed1)

	pwmd5 := fmt.Sprintf("%x", h.Sum(nil))

	salt := setting.CryptoSetting.Seed2

	io.WriteString(h, salt)
	io.WriteString(h, *pwd)
	io.WriteString(h, pwmd5)

	*pwd = fmt.Sprintf("%x", h.Sum(nil))
}
