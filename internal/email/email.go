package email

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"runtime"
	"seniors50plus/internal/models"
)

const (
	MIME = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
)

func catchPanic(err *error, functionName string) {
	if r := recover(); r != nil {
		fmt.Printf("%s : PANIC Defered : %v\n", functionName, r)

		// Capture the stack trace
		buf := make([]byte, 10000)
		runtime.Stack(buf, false)

		fmt.Printf("%s : Stack Trace : %s", functionName, string(buf))

		if err != nil {
			*err = fmt.Errorf("%v", r)
		}
	} else if err != nil && *err != nil {
		fmt.Printf("%s : ERROR : %v\n", functionName, *err)

		// Capture the stack trace
		buf := make([]byte, 10000)
		runtime.Stack(buf, false)

		fmt.Printf("%s : Stack Trace : %s", functionName, string(buf))
	}
}

func SendEmail(host string, port int, userName string, password string, to []string, subject string, templatePath string, info models.TemplateInfo) (err error) {
	defer catchPanic(&err, "SendEmail")

	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}
	buffer := new(bytes.Buffer)
	if err := t.Execute(buffer, info); err != nil {
		return err
	}
	body := "To: " + to[0] + "\r\nSubject: " + subject + "\r\n" + MIME + "\r\n" + buffer.String()
	auth := smtp.PlainAuth("", userName, password, host)

	err = smtp.SendMail(
		fmt.Sprintf("%s:%d", host, port),
		auth,
		userName,
		to,
		[]byte(body))

	return err
}
