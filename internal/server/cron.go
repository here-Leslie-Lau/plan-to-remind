package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/robfig/cron/v3"
	"plan-to-remind/internal/service"
)

type CronServer struct {
	c   *cron.Cron
	log *log.Helper
	svc *service.TimerService
}

func NewCronServer(svc *service.TimerService, logger log.Logger) *CronServer {
	return &CronServer{c: cron.New(), log: log.NewHelper(logger), svc: svc}
}

func (c *CronServer) Start(context.Context) error {
	c.c.Start()
	c.log.Info("start cron_server success...")
	c.addFunc()
	return nil
}

func (c *CronServer) Stop(context.Context) error {
	c.c.Stop()
	c.log.Info("stop cron_server success...")
	return nil
}

type cronTab struct {
	// 定时任务名称
	Name string
	// cron expression
	Expression string
	// handle method
	F func()
}

func (c *CronServer) addFunc() {
	list := []*cronTab{
		{
			Name:       "扫描今日计划并推送",
			Expression: "1 0 * * *",
			F:          c.svc.UserPlanPush,
		},
		{
			Name:       "每日扫描，将不符合的计划失效",
			Expression: "0 0 * * *",
			F:          c.svc.PlanToInvalid,
		},
	}
	for _, tab := range list {
		_, err := c.c.AddFunc(tab.Expression, tab.F)
		if err != nil {
			panic(err)
		}
		c.log.Infow("cron_server addFunc success", "name:", tab.Name)
	}
}
