package email

import (
	"fmt"
	"plan-to-remind/internal/conf"
	"testing"
)

func TestSend(t *testing.T) {
	cli := NewQqEmailClient(&conf.Data{Email: &conf.Data_Email{
		Host:     "smtp.qq.com",
		Username: "XXX",
		Password: "XXX",
	}})
	if err := cli.Send("demo", "demo", "XXX"); err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}
}
