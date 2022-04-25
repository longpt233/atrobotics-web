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

func SendEmailForgotPassword(to []string, newPassword string) (error) {
	from := os.Getenv("EMAIL_ACCOUNT")
	auth := smtp.PlainAuth("", from, os.Getenv("EMAIL_PASSWORD"), smtpHost)

	prePath := path.Join("E://JavaScript Tutorial//atrobotics//server//internal//templates","resetPassword.html")
	tpl, err := template.ParseFiles(prePath)
	if err != nil {
		return err
	}

	var body bytes.Buffer
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	_ , err = body.Write([]byte(fmt.Sprintf("Subject: Mail thông báo Mật khẩu mới của tài khoản "+to[0]+"\n%s\n\n", mimeHeaders)))

	if err != nil {
		return err
	}
	//email body
	bodyErr := tpl.Execute(&body, struct {
		BodyHeader string
		Content string
	}{
		BodyHeader: "Mật khẩu mới của bạn là:  ",
		Content: newPassword,
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