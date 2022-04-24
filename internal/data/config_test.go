package data

import (
	"fmt"
	"plan-to-remind/internal/conf"
	"plan-to-remind/internal/pkg/json"
	"testing"
)

func TestGetConfig(t *testing.T) {
	repo := NewConfigRepo(&Data{nacos: NewConfigClient(&conf.Data{
		Nacos: &conf.Data_Nacos{
			Port:   8848,
			IpAddr: "127.0.0.1",
		},
	})}, json.NewParser())
	config, err := repo.getConfig("plan_to_remind")
	if err != nil {
		t.Fail()
		return
	}
	fmt.Println(config)
}
