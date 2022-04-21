package biz

import (
	"context"
	"fmt"
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

// PlanToRemind 推送计划消费
func (uc *ConsumerUsecase) PlanToRemind(_ context.Context, result []byte) error {
	// json unmarshal
	plan := NewDefaultPlan(0)
	if err := uc.parser.Unmarshal(result, plan); err != nil {
		return err
	}
	// send email todo config
	msg := fmt.Sprintf(uc.emailPayload.msg, plan.CronDesc, plan.Name)
	return uc.email.Send(uc.emailPayload.subject, msg, "804738464@qq.com")
}
