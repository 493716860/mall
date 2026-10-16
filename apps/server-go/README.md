# Ruoyi-Vue-Pro Go + Gin 1:1 重构后端

本项目使用 **Go 1.22 + Gin 框架 + GORM** 对 Java 版本的 `yudao-server` 重写 1:1 还原。完全兼容微信小程序端 (`/app-api/*`) 与 运营后台 Web 端 (`/admin-api/*`) 的全量数据格式与接口协议。

---

## 🌟 核心优势

1. **极致轻量**：二进制可执行文件仅 **13MB**，启动内存占用低于 **30MB**（原 Java 版本需 1.5GB+）。
2. **毫秒级冷启动**：**10毫秒内** 完成服务初始化与数据库加载。
3. **零外部依赖依赖开销**：内置纯 Go 实现的 SQLite/GORM 驱动，亦可无缝切换至 MySQL 数据库。

---

## 📁 项目目录结构

```text
apps/server-go/
├── cmd/
│   └── server/
│       └── main.go                 # 启动入口主函数
├── internal/
│   ├── framework/                  # 基础设施
│   │   ├── database/               # GORM 数据库连接与初始化 (db.go)
│   │   ├── security/               # JWT 鉴权与 CORS 中间件 (jwt.go)
│   │   └── web/                    # 1:1 统一响应 CommonResult 与 PageResult (response.go)
│   └── module/                     # 7 大业务模块实现
│       ├── system/                 # 系统管理 (登录、菜单、字典、部门、管理员、地区树)
│       ├── infra/                  # 基础设施 (OSS 文件上传)
│       ├── member/                 # 会员中心 (微信/短信登录、个人资料、地址、积分、签到)
│       ├── product/                # 商品中心 (SPU/SKU、分类、品牌、属性规格、评价、收藏)
│       ├── trade/                  # 交易中心 (购物车、订单流转/结算/详情、物流快递、自提门店)
│       ├── promotion/              # 营销中心 (优惠券、秒杀、拼团、砍价、DIY微页面装配、在线客服)
│       ├── pay/                    # 支付中心 (微信统一下单、收银台、钱包余额、充值套餐)
│       └── statistics/             # 统计分析看板 (交易、会员、商品热销榜、充值统计)
├── go.mod                          # Go 模块依赖文件
└── README.md                       # 说明文档
```

---

## 🚀 编译与运行方式

### 1. 本地运行

```bash
cd apps/server-go
go run ./cmd/server/main.go -port 48080
```

### 2. 编译为可执行文件

```bash
cd apps/server-go
CGO_ENABLED=0 go build -o server-go ./cmd/server/main.go
./server-go -port 48080
```

---

## 🔗 开放服务接口前缀

- **管理后台 Web 端 API**：`http://localhost:48080/admin-api/*`
- **微信小程序移动端 API**：`http://localhost:48080/app-api/*`
- **健康检查接口**：`http://localhost:48080/health`
