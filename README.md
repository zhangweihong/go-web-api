## 概述

> 使用gin + mysql + redis + gorm进行的项目配置的快捷web api接口

## 目录介绍

### logs
> 按日期进行分割保存日志的目录 有access 和 error 日志
> 具体相关配置在 middleware/logger.go文件中配置

### public 
> 静态公共资源目录可在 src/router/router.go 文件中配置

### upload
> 文件上传目录，在src/helper/helper.go中有获取方法 GetUploadsFilePath()

### src/config
>相关的配置，在config.go文件中

### src/config/controller
>操作service的控制器 用来判断传入参数合法性等 返回前端数据

### src/config/database
>相关数据库的配置

### src/config/helper
>一些工具方法和公共的函数

### src/middleware
>中间件的配置文件 比如日志中间件或者路由鉴权等

### src/model
>对应mysql数据库的model，ORM框架使用gorm插件

### src/schedule
>定时执行任务

### src/router
>对应的路由配置文件目录

### src/service
>主要的业务的逻辑操作，进行model相关交互和逻辑处理，供对应的controller来进行调用 

### src/setting
>项目需要提前初始化的设置

### src/status
>返回前端的状态码设置

### src/websocket
>还暂未进行先关扩展






