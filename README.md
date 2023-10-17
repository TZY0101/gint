# gint
## 介绍
gint是一个简单的代码生成工具，主要功能是帮助开发者快速生成gin项目代码，提高开发效率。
## 特性
**·** 基于gin框架，根据接口描述文件desc生成gin项目代码

**·** 路由自注册，生成的项目代码可直接运行，开发者只需关注业务逻辑

**·** 服务优雅退出

## 使用
- git install github.com/TZY0101/gint
- 验证：gint -v
- 初始化项目：gint init demo
- 根据desc文件生成代码：
  - cd 到desc文件所在目录
  - gint gen --desc demo.desc

### 目录结构
```text
├── cmd
│   └── main.go       # 入口文件
├── config            # 配置
│   ├── config.go        
│   ├── config.yaml       
├── desc              # 接口描述
│   ├── demo.desc       
├── internal
│   └── app
│       ├── api       # API 处理层
│       ├── dto       # api入参、出参映射
│       ├── logic     # 业务逻辑层
│       ├── router    # 路由层
├── pkg
│   └── errorx        # 全局错误处理
│       ├── base.go   # 实现自定义错误
│       ├── code.go   # 错误码
│       ├── msg.go    # 错误信息
│   ├── ginx          # gin 扩展模块
```

## 计划
- 加入参数解析失败的自定义消息
- 根据表生成repo，提供基本curd