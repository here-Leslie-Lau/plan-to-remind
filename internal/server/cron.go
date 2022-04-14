package server

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/robfig/cron/v3"
	"plan-to-remind/internal/conf"
)

type CronServer struct {
	c   *cron.Cron
	log *log.Helper
	// todo timer service
}

func NewCronServer(data *conf.Data, logger log.Logger) *CronServer {
	return &CronServer{c: cron.New(), log: log.NewHelper(logger)}
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
		&cronTab{
			Name:       "demo",
			Expression: "29 18 * * *",
			F: func() {
				fmt.Println("hello world!")
			},
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
