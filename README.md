# 本项目对K8s二次开发, 基于golang 1.12

## 运行环境：go 1.12 + iris + mysql

## 运行步骤

### 配置：
> 在simple-k8-config上填写您的配置。本机新建mydb数据库，在你使用的数据库中运行_doc目录下的init.sql

### 构建：

> 1. 设置项目使用go mod的方式启动： set GO111MODULE=on

> 2. 终端代理开启(需要可以访问google页面), 默认是1087端口： export http_proxy='127.0.0.1:1087' 且  export https_proxy='127.0.0.1:1087'。

```text

macOS环境直接 vim ~/.zshrc

添加：

export http_proxy='127.0.0.1:1087'

export https_proxy='127.0.0.1:1087'

让配置生效：

source ~/.zshrc

```

> 3. goland的idea，最好设置代理。goland -> preferences -> 勾选Enable Go Modules(vgo) integration设置为https://goproxy.io

> 4. 执行go mod加载相关依赖： go mod tidy

> 5. 第四步骤没问题的情况下构建该项目：go build -x 根据自己需要部署的环境，打对应环境的二进制包

> 6. 在对应的环境中，后台挂起的方式执行本项目： nohup ./simple-k8 -c simple-k8-config.yml &


