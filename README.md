# gint
## 介绍
gint是一个简单的代码生成工具，主要功能是帮助开发者快速生成gin项目代码，提高开发效率。
## 特性
**·** 基于gin框架，根据接口描述文件desc生成gin项目代码

**·** 路由自注册，生成的项目代码可直接运行，开发者只需关注业务逻辑

**·** 服务优雅退出

## 使用
### 安装
> go install github.com/TZY0101/gint  
> gint -v
### 快速开始
#### gint init
````text
# cd到项目文件夹
$ cd proj
# 执行指令生成demo项目
$ gitn init demo
````
指令执行完毕后会在proj目录下生成一个demo项目，项目目录如下
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
│       └── api       # API 处理层
│           └── demo       # demo模块
│               └── helloWorld.go       
│       ├── dto       # api入参、出参映射
│       └── logic     # 业务逻辑层
│           └── demo       # demo模块
│               └── helloWorld.go 
│       └── router    # 路由层
│           ├── demo_routes.go        
│           ├── router.go    
├── pkg
│   └── errorx        # 全局错误处理
│       ├── base.go   # 实现自定义错误
│       ├── code.go   # 错误码
│       ├── msg.go    # 错误信息
│   ├── ginx          # gin 扩展模块
```
gint init 命令使用的接口描述文件默认如下
```text
DTO DemoReq {
    Id int64
}

DTO DemoResp {
    Name string
}

APIS demo {
    get /demo helloWorld(DemoReq) DemoResp
}
```
#### gint gen
在生成的demo项目中修改接口描述文件后，可以用gint gen 命令来生成新的代码
```text
# cd到desc文件所在目录
$ cd demo/desc
# 执行指令生成新的接口
$ gint gen --desc demo.desc
```
### desc文件语法
- DTO
  - DTO定义与struct定义无差别，每个DTO都会生成一个struct与之对应
- APIS
```text
// 模块名为demo，在api和logic目录下有一个同名文件夹与之对应
APIS demo {
    // 请求方法为get
    // 路由为/demo
    // api和logic中对应的函数名为HelloWorld
    // 请求体为DemoReq，响应体为DemoResp
    get /demo helloWorld(DemoReq) DemoResp
}
```

## 计划
- 加入参数解析失败的自定义消息
- 根据表生成repo，提供基本curd
- 加入默认中间件：jwt认证、跨域处理
- desc文件加入对APIS模块的路由前缀定义、中间件使用