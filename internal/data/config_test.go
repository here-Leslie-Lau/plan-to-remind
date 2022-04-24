package data

import (
	"fmt"
	"plan-to-remind/internal/biz"
	"plan-to-remind/internal/conf"
	"testing"
)

func TestGetConfig(t *testing.T) {
	NewConfigRepo(&Data{nacos: NewConfigClient(&conf.Data{
		Nacos: &conf.Data_Nacos{
			Port:   8848,
			IpAddr: "127.0.0.1",
		},
	})})
	config, err := biz.RepoConfig.GetDefaultConfig("plan_to_remind")
	if err != nil {
		t.Fail()
		return
	}
	fmt.Println(config)
}
