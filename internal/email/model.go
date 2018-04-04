import (
	"bytes"
	"html/template"
	"net/smtp"
)

type TemplateData struct {
	Name string
	URL  string
}

type Email struct {
	from    string
	to      []string
	subject string
	body    string
}

func NewEmail(to []string, subject, body string) *Email {
	return &Email{
		to:      to,
		subject: subject,
		body:    body,
	}
}

func (e *Email) SendEmail(auth smtp.Auth) (bool, error) {
	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + e.subject + "!\n"
	msg := []byte(subject + mime + "\n" + e.body)
	addr := "smtp.gmail.com:587"

	if err := smtp.SendMail(addr, auth, "roommates40plus@gmail.com", e.to, msg); err != nil {
		return false, err
	}
	return true, nil
}

func (e *Email) ParseTemplate(templateFileName string, data interface{}) error {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	e.body = buf.String()
	return nil
}