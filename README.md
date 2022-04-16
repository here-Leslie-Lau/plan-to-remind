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
- v0.4: 延迟队列，计划推送生产者实现
- v0.5: 延迟队列，计划推送消费者实现
- v0.6: 技术债:返回值新增状态码与描述信息
- v0.7: 时间解析器功能完善
- v1.0: 每日计划完成度统计数据的生成
- v1.1: 每日计划完成度查询，修改，删除
- v1.2: 每日完成分析(数据分析，完成趋向等指标)

## The technical architecture

todo