## 1. 基本介绍

### 1.1 项目介绍

> Ginchat 是一个基于 [gin](https://gin-gonic.com) 开发的全栈前后端分离的在线聊天平台，目前还在构建服务端框架，充分参考学习了[gin-vue-admin](https://github.com/flipped-aurora/gin-vue-admin)。

## 2. 使用说明

```
- golang版本 >= v1.16
- IDE推荐：Visual Studio Code
```

### 2.1 server项目

使用 `Visual Studio Code` 等编辑工具，进入项目根目录，修改 `app.yaml` 中 `mysql` 相关配置。

```bash
# 克隆项目
git clone https://github.com/liao123-git/ginchat.git

# 使用 go mod 并安装go依赖包
go generate

# 直接运行 
go run .
```

### 2.3 swagger自动化API文档

#### 2.3.1 安装 swagger

````
go get -u github.com/swaggo/swag/cmd/swag
````

#### 2.3.2 生成API文档

```` shell
swag init
````

> 执行上面的命令后，server目录下会出现docs文件夹里的 `docs.go`, `swagger.json`, `swagger.yaml` 三个文件更新，启动go服务之后, 在浏览器输入 [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html) 即可查看swagger文档

## 3. 项目架构

### 3.1 目录结构

```shell
├── assest
├── config
├── controller
├── core
│   └── initialize
├── docs
├── global
├── middleware
├── model
│   ├── request
│   └── response
├── router
│   └── system
├── storage
│   ├── app
│   └── logs
├── test
├── utils
│   ├── request
|   └── response
```

| 文件夹        | 说明                   | 描述                        |
| -----------  | --------------------- | --------------------------- |
| `assest`     | 静态资源文件夹          | 负责存放静态文件               |
| `config`     | 配置包                 | app.yaml对应的配置结构体       |
| `controller` | 控制器                 | api层                       |
| `core`       | 核心文件                | 核心组件(gorm, viper, server)的初始化 |
| `--initialize` | 初始化                | router, gorm, validator, timer的初始化 |    
| `docs`       | swagger文档目录         | swagger文档目录              |
| `global`     | 全局对象                | 全局对象                     |
| `middleware` | 中间件层 | 用于存放 gin 中间件代码                       |
| `model`      | 模型层                  | 模型对应数据表                |
| `--request`  | 入参结构体              | 接收前端发送到后端的数据。       |
| `--response` | 出参结构体              | 返回给前端的数据结构体          |
| `router`     | 路由层                  | 路由层                       |
| `--system`   | 路由系统文件             | 封装 gin.Engine 资源路由方法   |
| `storage`    | 读写文件夹              | 存放有可能会读写的文件           |
| `--app`      | app文件夹               | 存放前端上传的文件              |
| `--logs`     | 日志文件夹              | 存放日志文件                   |
| `test`       | 测试文件                | 测试代码                      |
| `utils`      | 工具包                  | 工具函数封装                  |

## 4. 主要功能（更新中）

- 用户：注册，登陆。
- validator：自定义验证字段和错误信息。
  - 自定义验证字段：[model/request/user_request.go](https://github.com/liao123-git/ginchat/tree/main/model/request/user_request.go)
  - 自定义错误信息：[utils/request/request.go](https://github.com/liao123-git/ginchat/tree/main/utils/request/request.go)
- restful示例：可以参考资源路由模块中的示例API。
    - 后台文件参考: [router/system/resouce_router.go](https://github.com/liao123-git/ginchat/tree/main/router/system/resouce_router.go)
