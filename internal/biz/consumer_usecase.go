package biz

import (
	"context"
	"fmt"
	"plan-to-remind/internal/conf"
	"plan-to-remind/internal/pkg/email"
	"plan-to-remind/internal/pkg/json"
)

type ConsumerUsecase struct {
	parser  *json.Parser
	email   email.IEmail
	cfgRepo ConfigRepo
}

type EmailConfig struct {
	Subject string   `json:"subject"`
	Msg     string   `json:"msg"`
	To      []string `json:"to"`
}

type Config struct {
	EmailConfig EmailConfig `json:"email_config"`
}

type ConfigRepo interface {
	GetDefaultConfig(dataId string) (*Config, error)
}

var RepoConfig ConfigRepo

func NewConsumerUsecase(data *conf.Data, parser *json.Parser) *ConsumerUsecase {
	return &ConsumerUsecase{
		parser:  parser,
		email:   email.NewQqEmailClient(data),
		cfgRepo: RepoConfig,
	}
}

// PlanToRemind 推送计划消费
func (uc *ConsumerUsecase) PlanToRemind(_ context.Context, result []byte) error {
	// json unmarshal
	plan := new(Plan)
	if err := uc.parser.Unmarshal(result, plan); err != nil {
		return err
	}
	// get email payload from nacos config
	cfg, err := uc.cfgRepo.GetDefaultConfig("plan_to_remind")
	if err != nil {
		return err
	}
	// send email
	msg := fmt.Sprintf(cfg.EmailConfig.Msg, plan.Name, plan.CronDesc)
	return uc.email.Send(cfg.EmailConfig.Subject, msg, cfg.EmailConfig.To...)
}
