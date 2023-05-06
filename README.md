# go-zero-micro

## 1. 环境准备

1. Golang1.8及以上版本

2. etcd3.5版本
    * 用于做服务发现、配置中心
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

### 生成model文件
* 修改`gen.yaml`配置，用于自动生成model
* `make gen_db_model` 会自动生成model到对应的服务下面，注意：需要选择服务。

### 生成服务文件
* 修改
* 每个服务下面会有 `gen_xxx_code` 命令，用于生成代码。

## 4. 项目结构

```
  .
├── api // 对外访问api接口
├── Makefile
├── README.md
├── common  // 常用方法等
├── core    // 常用中间件
├── config  // 会把一些通用配置放在这里
│   └── config.dev.yaml
├── go.mod
├── go.sum
├── script  // 脚本文件
│   ├── gen_db_model.go // 通过数据库field生成类型
│   └── push_etcd_config.go // 把通用配置写到etcd
├── service  // 内部服务
│   ├── product  // 产品服务
│   │   ├── rpc // 
│   │   │   ├── etc // 当前服务配置文件
│   │   │   ├── internal // 业务逻辑代码
│   │   │   └── etc // 当前服务配置文件
│   │   └── model // 数据库操作层
│   └── user  // user服务
├── sql  // sql文件
├── script  // 脚本文件
├── gen.yaml  // 数据库生成配置文件
└── script  // 脚本文件

```

## 5. commit type

| 类型       | 描述                           |
|----------|------------------------------|
| build    | 	编译相关的修改，例如发布版本、对项目构建或者依赖的改动 |
| chore    | 	其他修改, 比如改变构建流程、或者增加依赖库、工具等  |
| ci       | 	持续集成修改                      |
| docs     | 	文档修改                        |
| feat     | 	新特性、新功能                     |
| fix      | 	修改bug                       |
| perf     | 	优化相关，比如提升性能、体验              |
| refactor | 	代码重构                        |
| revert   | 	回滚到上一个版本                    |
| style    | 	代码格式修改, 注意不是 css 修改         |
| test     | 	测试用例修改                      |
