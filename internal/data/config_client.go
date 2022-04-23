package data

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"plan-to-remind/internal/conf"
)

func NewConfigClient(data *conf.Data) config_client.IConfigClient {
	sc := []constant.ServerConfig{{
		IpAddr: data.Nacos.IpAddr,
		Port:   data.Nacos.Port,
	}}
	cc := constant.ClientConfig{
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogLevel:            "debug",
	}
	client, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})
	if err != nil {
		panic(err)
	}
	return client
}
