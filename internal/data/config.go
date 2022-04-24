package data

import (
	"github.com/nacos-group/nacos-sdk-go/vo"
	"plan-to-remind/internal/pkg/json"
)

type EmailConfig struct {
	Subject string   `json:"subject"`
	Msg     string   `json:"msg"`
	To      []string `json:"to"`
}

type Config struct {
	EmailConfig EmailConfig `json:"email_config"`
}

type ConfigRepo struct {
	data *Data
	json *json.Parser
}

func NewConfigRepo(data *Data, json *json.Parser) *ConfigRepo {
	return &ConfigRepo{data: data, json: json}
}

func (repo *ConfigRepo) getConfig(dataId string) (*Config, error) {
	cfgStr, err := repo.data.nacos.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  "DEFAULT_GROUP",
	})
	if err != nil {
		return nil, err
	}
	cfg := new(Config)
	if err := repo.json.Unmarshal([]byte(cfgStr), cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
