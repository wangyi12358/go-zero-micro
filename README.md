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

## 2. 项目规范

* 文件夹、文件名都采用下划线分割，并且首字母小写开头。
* .api、.proto的service、type、message 都以首字母大写开头。

## 3. 项目开发

* 修改`gen.yaml`配置，用于自动生成model
* `make gen_db_model` 会自动生成model到对应的服务下面，注意：需要选择服务。
* 每个服务下面会有 `gen_xxx_code` 命令，用于生成代码。

## 4. 项目结构

```
  .
├── api // 对外访问api接口
├── Makefile
├── README.md
├── common  // 常用方法等
├── core    // 常用中间件
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
