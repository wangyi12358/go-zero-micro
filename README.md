# go-zero-micro

## 1. 环境准备
1. Golang1.8及以上版本

2. etcd3.5版本
    * 用于做服务发现
    * [安装地址](https://etcd.io/docs/v3.5/install/)

3. Goctl
    * 用于自动生成rpc、api等代码
    * [安装地址](https://go-zero.dev/cn/docs/goctl/goctl)

4. Gen Tool
    * 用于生成数据库Model实体类
    * [安装地址](https://gorm.io/zh_CN/gen/gen_tool.html)

## 项目结构
```
  .
├── api // 对外访问api接口
├── Makefile
├── README.md
├── common  // 常用方法等
├── configs  // 不同环境的通用配置文件
│   └── config.dev.yaml
├── go.mod
├── go.sum
├── script  // 脚本文件
├── service  // 内部服务
│   ├── product  // 产品服务
│   │   ├── rpc // 
│   │   │   ├── etc // 当前服务配置文件
│   │   │   ├── internal // 业务逻辑代码
│   │   │   └── etc // 当前服务配置文件
│   │   └── model // 数据库操作层
│   └── user  // user服务
├── script  // 脚本文件
├── gen.yaml  // 数据库生成配置文件
└── script  // 脚本文件

```
