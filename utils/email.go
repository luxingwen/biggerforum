package utils

import (
	"net/smtp"
)

func SendEmail(content []byte) error {
	auth := smtp.PlainAuth("", "935232474@qq.com", "10086zzz@", "smtp.qq.com")
	var a []string = []string{"1147130352@qq.com"}
	err := smtp.SendMail("smtp.qq.com:465", auth, "935232474@qq.com", a, content)
	return err
}
