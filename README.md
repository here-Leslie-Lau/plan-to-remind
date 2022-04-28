# plan-to-remind

## Project introduction

just a tiny project to plan to remind

项目初衷，仅仅是因为本人经常忘记一些要计划做好的事情，所以打算写一个计划推送服务，来提醒自己。虽然可以直接用手机的备忘录，但是还是秉承着程序员想自己"造轮子"的心理，所以还是打算开搞😃

本项目采用bilili的`kratos`框架，再加上了自己对于DDD的理解，所以在它原本的示例项目中做了些改造。

## version management

- [x] v0.1: 时间表curd
- [x] v0.2: 计划表curd
- [x] v0.2.1: 技术债:biz里面的model改成充血模型
- [x] v0.2.2: 技术债:repo暴露方法规范问题
- [x] v0.3: 定时器基本框架搭建
- [x] v0.4: 延迟队列，计划推送生产者实现
- [x] v0.5: 延迟队列，计划推送消费者实现
- [x] v0.5.1: 本地打包docker镜像部署
- [x] v0.5.2: 远程服务器部署与基本环境搭建(db,mq...)
- [x] v0.5.3: 增加nacos读取配置封装
- [x] v0.6: nacos配置读取
- [x] v0.7: 计划完成接口
- v1.0: 每日计划完成度统计数据的生成
- v1.1: 每日计划完成度查询，修改，删除
- v1.2: 每日完成分析(数据分析，完成趋向等指标)
- v1.3: 时间解析器功能完善

## The technical architecture

todo

## Docker deployment

- plan(this project) port:8000(http)&9000(grpc)
- mysql port:3306
- pulsar port:6650
- nacos port:8848

## how to run this project

1. Get mysql, Pulsar, nacos ready
2. `git clone git@github.com:here-Leslie-Lau/plan-to-remind.git`
3. To initialize the database.run in terminal:`make init-local`
4. `make build && make run`
