package biz

import (
	"context"
	"fmt"
	"plan-to-remind/internal/conf"
	"plan-to-remind/internal/pkg/email"
	"plan-to-remind/internal/pkg/json"
)

type emailPayload struct {
	subject, msg string
}

type ConsumerUsecase struct {
	parser       *json.Parser
	email        email.IEmail
	emailPayload *emailPayload
}

func NewConsumerUsecase(data *conf.Data, parser *json.Parser) *ConsumerUsecase {
	return &ConsumerUsecase{
		parser: parser,
		email:  email.NewQqEmailClient(data),
		emailPayload: &emailPayload{
			// todo config
			subject: "来自计划提醒",
			msg:     "你的%s计划日程已到，计划时间%s",
		},
	}
}

// PlanToRemind 推送计划消费
func (uc *ConsumerUsecase) PlanToRemind(_ context.Context, result []byte) error {
	// json unmarshal
	plan := new(Plan)
	if err := uc.parser.Unmarshal(result, plan); err != nil {
		return err
	}
	// send email
	msg := fmt.Sprintf(uc.emailPayload.msg, plan.Name, plan.CronDesc)
	return uc.email.Send(uc.emailPayload.subject, msg, "804738464@qq.com")
}
