# 我的商城

从芋道商城抽离的独立前端项目。保留原有商城业务页面，运营后台移除了其它业务系统；后续可通过接口适配对接 Go + Gin。

## 目录

- `apps/admin`：Vue 3 + Element Plus 运营后台。
- `apps/miniapp`：UniApp + Vue 3 客户端，支持微信小程序与 H5。
- `packages/demo-data`：两端使用的静态演示种子数据。
- `docs`：裁剪范围、接口契约和验收说明。
- `scripts`：演示接口回归检查。

Node.js 22.12+、pnpm 10+。两个应用各自拥有依赖锁文件，避免 UniApp 与后台框架版本互相影响。

```bash
pnpm install:apps
pnpm dev:admin
pnpm dev:miniapp
```

后台默认 `http://localhost:5173`；小程序 H5 默认 `http://localhost:3000/miniapp/`。默认开发启动使用**演示模式**，不连接上游演示服务器。

```bash
pnpm build:demo
pnpm typecheck
pnpm test:demo
pnpm build:weixin
```

## 正式后端联调

分别复制两端 `.env.example` 为 `.env.local`，填入自己的后端 URL。使用 `pnpm --dir apps/admin dev:api` 与 `pnpm --dir apps/miniapp dev:api`。正式构建 `pnpm build` 默认关闭演示数据。部署需提供 `/admin-api` 和 `/app-api` 或修改前缀，并配置后台 SPA fallback。

微信小程序使用自己的 AppID 填入 `apps/miniapp/manifest.json`，然后 `pnpm build:weixin`，将 `apps/miniapp/dist/build/mp-weixin` 导入微信开发者工具。正式发布还需配置合法域名、登录与支付接口；H5 预览不能验证微信登录和支付。

## 演示边界

后台商品、分类、品牌、订单、会员等页面使用虚构数据；部分商品 CRUD、订单发货与配置保存支持浏览器持久化。客户端支持商品浏览、分类、购物车。未覆盖的列表展示空数据，未实现的写入明确阻止；结算、支付、退款及其它实际业务必须接入后端。两端仅共享种子数据，**并未实现跨端数据同步**。

源码基于 MIT 开源项目，保留各应用的原始 LICENSE。请阅读 `docs/ARCHITECTURE.md` 和 `docs/ACCEPTANCE.md`。
