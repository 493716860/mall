# 会员中心 (Member) 专项重构与 API 字典说明

本文档针对 **微信小程序端 (`apps/miniapp`)** 与 **WEB 运营后台 (`apps/admin`)** 的 **会员中心** 模块进行专项梳理。

---

## 📚 一、 会员中心 11 张核心数据表

| 数据表名 | 中文名称 | 说明与用途 |
|---|---|---|
| `member_user` | 会员用户主表 | 存储会员手机号、密码、昵称、头像、等级ID、分组ID、标签ID列表、积分、经验值等 |
| `member_level` | 会员等级配置表 | 定义 VIP1 ~ VIP5 等级名称、升级所需经验值、专属折扣百分比（如95折）、等级图标与背景 |
| `member_group` | 会员分组表 | 管理员对会员的分组（如高价值客户、沉睡客户等） |
| `member_tag` | 会员标签表 | 给会员打上的打标分类（如"喜好零食"、"复购达人"） |
| `member_address` | 会员收货地址表 | 用户收货人姓名、手机号、省市区 AreaID、详细地址、是否默认标记 |
| `member_point_record` | 会员积分变动记录表 | 记录积分加减明细、业务类型（签到、购物赠送、抵扣）、变动前后余额 |
| `member_experience_record`| 会员经验值变动记录表 | 记录经验值变动明细、触发事件（购物消费、签到） |
| `member_sign_in_config` | 会员签到配置表 | 配置连续签到 1-7 天获得的奖励积分与经验值 |
| `member_sign_in_record` | 会员签到历史表 | 记录用户每日打卡签到时间、第几天签到与获得的积分经验 |
| `member_brokerage_user` | 会员分销推广员表 | 分销员资格状态、绑定的上级推荐人 ID、累计佣金金额 |
| `member_brokerage_record` | 会员佣金明细表 | 佣金结算记录、订单来源、提现冻结状态与可提现金额 |

---

## 📡 二、 会员中心 35 个关键 API 接口

### 1. 小程序端 (`/app-api/member/*`)

* `GET /app-api/member/user/get` : 获取当前登录会员完整个人信息（含等级名称、积分余额、经验值、待付款/待收货订单数）
* `PUT /app-api/member/user/update-nickname` : 修改用户昵称
* `PUT /app-api/member/user/update-avatar` : 修改用户头像
* `POST /app-api/member/user/update-mobile` : 绑定/更换手机号（需微信授权 code 或短信验证码）
* `GET /app-api/member/level/list` : 获取所有 VIP 等级列表与升级条件
* `GET /app-api/member/address/list` : 获取当前用户的收货地址列表
* `GET /app-api/member/address/get?id={id}` : 获取单个收货地址详情
* `POST /app-api/member/address/create` : 新增收货地址
* `PUT /app-api/member/address/update` : 修改收货地址
* `DELETE /app-api/member/address/delete?id={id}` : 删除收货地址
* `GET /app-api/member/point/record/page` : 分页查询当前用户积分明细流水
* `GET /app-api/member/sign-in/config/list` : 获取签到奖励规则（7天轮次）
* `GET /app-api/member/sign-in/record/get-summary` : 获取个人签到汇总统计（连续签到天数、今日是否已签到）
* `POST /app-api/member/sign-in/record/create` : 执行每日打卡签到（原子增加积分与经验值，并检查触发 VIP 升级）

### 2. 运营后台端 (`/admin-api/member/*`)

* `GET /admin-api/member/user/page` : 管理员分页查询会员列表（支持按手机号、昵称、VIP等级、标签分组筛选）
* `GET /admin-api/member/user/get?id={id}` : 查看会员详情（含消费统计、订单列表、积分历史）
* `PUT /admin-api/member/user/update` : 修改会员资料或禁用账号
* `PUT /admin-api/member/user/update-level` : 手动调整会员 VIP 等级
* `PUT /admin-api/member/user/update-point` : 手动充值/扣减会员积分（记录管理员操作原因）
* `GET /admin-api/member/level/list` : 获取等级配置列表
* `POST /admin-api/member/level/create` : 新增 VIP 等级
* `PUT /admin-api/member/level/update` : 修改 VIP 等级配置与折扣
* `DELETE /admin-api/member/level/delete` : 删除 VIP 等级
* `GET /admin-api/member/group/list` : 获取会员分组列表
* `POST /admin-api/member/group/create` : 新增会员分组
* `GET /admin-api/member/tag/list` : 获取会员标签列表
* `POST /admin-api/member/tag/create` : 新增会员标签
