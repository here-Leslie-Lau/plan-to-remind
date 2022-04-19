package email

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"plan-to-remind/internal/conf"
)

type IEmail interface {
	Send(subject, msg string, to ...string) error
}

type qqEmailClient struct {
	from   string
	dialer *gomail.Dialer
}

func NewQqEmailClient(data *conf.Data) IEmail {
	// vpcitbhtcbbabeac
	return &qqEmailClient{from: data.Email.Username, dialer: gomail.NewDialer(data.Email.Host, 465, data.Email.Username, data.Email.Password)}
}

func (cli *qqEmailClient) Send(subject, msg string, to ...string) error {
	m := gomail.NewMessage(gomail.SetEncoding(gomail.Base64))
	m.SetHeader("From", cli.from)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	body := fmt.Sprintf("<h2>%s</h2>", msg)
	m.SetBody("text/html", body)

	return cli.dialer.DialAndSend(m)
}
