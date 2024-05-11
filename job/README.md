### 参考
- 消息队列、延迟队列、定时任务本项目使用的是asynq ，基于redis开发的简单中间件
- 定时任务实现参考：[go-zero-looklook](https://github.com/Mikaelemmmm/go-zero-looklook/blob/main/README-cn.md)
- 链接：https://github.com/hibiken/asynq

### 注册任务脚本

- 新增任务：./internal/logic/xxx.go
- 注册任务：./internal/logic/routes.go


### 任务调度器

- **cronspec** 传递了一个字符串是："* * * * * *"

    - `* * * * *`
    - @every 3s
    - @every 20m
    - @every 1h30m
    - @every midnight

  ```
  # Minute, Hour, Dom, Month, Dow
  ┌───────────── min (0 - 59)
  │ ┌────────────── hour (0 - 23)
  │ │ ┌─────────────── day of month (1 - 31)
  │ │ │ ┌──────────────── month (1 - 12)
  │ │ │ │ ┌───────────────── day of week (0 - 6) (0 to 6 are Sunday to
  │ │ │ │ │                  Saturday)
  │ │ │ │ │
  │ │ │ │ │
  * * * * *
  ```
  参考 **github.com/robfig/cron/v3** cron.go

