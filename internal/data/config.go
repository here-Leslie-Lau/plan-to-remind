package data

import (
	"github.com/nacos-group/nacos-sdk-go/vo"
	"plan-to-remind/internal/biz"
	"plan-to-remind/internal/pkg/json"
)

type ConfigRepo struct {
	data *Data
	json *json.Parser
}

func NewConfigRepo(data *Data) {
	biz.RepoConfig = &ConfigRepo{data: data, json: json.NewParser()}
}

func (repo *ConfigRepo) getConfig(dataId, group string) (*biz.Config, error) {
	cfgStr, err := repo.data.nacos.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
	})
	if err != nil {
		return nil, err
	}
	cfg := new(biz.Config)
	if err := repo.json.Unmarshal([]byte(cfgStr), cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

// GetDefaultConfig get config in group DEFAULT_GROUP
func (repo *ConfigRepo) GetDefaultConfig(dataId string) (*biz.Config, error) {
	return repo.getConfig(dataId, "DEFAULT_GROUP")
}
