package helper

import (
	"bytes"
	"fmt"
	"net/smtp"
	"os"
	"path"
	"text/template"
)

var (
	// smtp server configuration.
	smtpHost = "smtp.gmail.com"
	smtpPort = "587"
)

func SendEmailForgotPassword(to []string, otpValue string) (error) {
	from := os.Getenv("EMAIL_ACCOUNT")
	auth := smtp.PlainAuth("", from, os.Getenv("EMAIL_PASSWORD"), smtpHost)

	prePath := path.Join("E://JavaScript Tutorial//atrobotics//server//internal//templates","resetPassword.html")
	tpl, err := template.ParseFiles(prePath)
	if err != nil {
		return err
	}

	var body bytes.Buffer
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	_ , err = body.Write([]byte(fmt.Sprint("Subject:  AT Robotics - Thông báo xác nhận quên mật khẩu", mimeHeaders)))

	if err != nil {
		return err
	}
	//email body
	bodyErr := tpl.Execute(&body, struct {
		Content string
	}{
		Content: otpValue,
	})
	if bodyErr != nil {
		return bodyErr
	}

	// Sending email.
	sendErr := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if sendErr != nil {
	  return sendErr
	}
	return nil
}