# Go + Gin 重构指南与架构模版

本文档为使用 **Go + Gin + GORM + Redis** 重构芋道 Mall 商城后端的架构指南。对应已实现的落地代码库路径为 **`apps/server-gin`**。

---

## 🏗️ 一、 工程项目分层架构

```text
apps/server-gin/
├── main.go               # 项目启动入口与路由挂载
├── config/               # 全局配置读取 (YAML / env)
├── global/               # 全局 DB / Redis 句柄句柄
├── middleware/           # JWT 认证、CORS 跨域、Recovery 日志中间件
├── dto/                  # 接口入参 DTO 与出参 Request/Response 定义
├── response/             # 100% 对齐 Java 的 CommonResult 统一返回体
├── model/                # GORM 数据库 Model 结构体定义 (全量 85 张表)
├── handler/              # Controller 路由处理层
├── service/              # 核心业务逻辑层 (事务控制、库存扣减、订单状态机)
└── router/               # 路由注册与分组 (app-api / admin-api)
```

---

## ⚡ 二、 快速启动与编译命令

```bash
cd apps/server-gin

# 1. 下载依赖
go mod tidy

# 2. 本地开发启动
go run main.go

# 3. 交叉编译 Linux 部署二进制文件
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server-gin main.go
```

---

## 🎯 三、 100% 对齐 Java 的统一返回体 (`response/result.go`)

在 Go 中使用以下结构体，确保小程序与后台前端无缝解析：

```go
package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommonResult struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, CommonResult{
		Code: 0,
		Data: data,
		Msg:  "success",
	})
}

func Error(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, CommonResult{
		Code: code,
		Data: nil,
		Msg:  msg,
	})
}
```
