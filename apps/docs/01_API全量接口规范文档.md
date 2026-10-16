# 芋道商城 (Mall) 全量 API 接口梳理规范文档

> 本文档 100% 完整覆盖了本项目运营后台 (`apps/admin`) 与 微信小程序端 (`apps/miniapp`) 所有被调用的真实 API 接口。
> - 运营后台 (`/admin-api/*`): 共 **305** 个接口
> - 微信小程序端 (`/app-api/*`): 共 **133** 个接口
> - 全量总计: **438** 个接口

---

## 壹、 WEB 运营后台 (`/admin-api`) 全量接口清单

### 1.商品中心 (Product)

#### 📌 商品 SPU 管理 (前缀: `/product/spu` - 10 个接口)

| HTTP 方法 | 后端完整路径 (`/admin-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/admin-api/product/spu/page` | `admin/src/api/mall/product/spu.ts` |
| `GET` | `/admin-api/product/spu/get-count` | `admin/src/api/mall/product/spu.ts` |
| `POST` | `/admin-api/product/spu/create` | `admin/src/api/mall/product/spu.ts` |
| `PUT` | `/admin-api/product/spu/update` | `admin/src/api/mall/product/spu.ts` |
| `PUT` | `/admin-api/product/spu/update-status` | `admin/src/api/mall/product/spu.ts` |
| `GET` | `/admin-api/product/spu/get-detail` | `admin/src/api/mall/product/spu.ts` |
| `GET` | `/admin-api/product/spu/list` | `admin/src/api/mall/product/spu.ts` |
| `DELETE` | `/admin-api/product/spu/delete` | `admin/src/api/mall/product/spu.ts` |
| `DOWNLOAD` | `/admin-api/product/spu/export-excel` | `admin/src/api/mall/product/spu.ts` |
| `GET` | `/admin-api/product/spu/list-all-simple` | `admin/src/api/mall/product/spu.ts` |

#### 📌 商品分类 (前缀: `/product/category` - 5 个接口)

| HTTP 方法 | 后端完整路径 (`/admin-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `POST` | `/admin-api/product/category/create` | `admin/src/api/mall/product/category.ts` |
| `PUT` | `/admin-api/product/category/update` | `admin/src/api/mall/product/category.ts` |
| `DELETE` | `/admin-api/product/category/delete` | `admin/src/api/mall/product/category.ts` |
| `GET` | `/admin-api/product/category/get` | `admin/src/api/mall/product/category.ts` |
| `GET` | `/admin-api/product/category/list` | `admin/src/api/mall/product/category.ts` |

#### 📌 商品品牌 (前缀: `/product/brand` - 6 个接口)

| HTTP 方法 | 后端完整路径 (`/admin-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `POST` | `/admin-api/product/brand/create` | `admin/src/api/mall/product/brand.ts` |
| `PUT` | `/admin-api/product/brand/update` | `admin/src/api/mall/product/brand.ts` |
| `DELETE` | `/admin-api/product/brand/delete` | `admin/src/api/mall/product/brand.ts` |
| `GET` | `/admin-api/product/brand/get` | `admin/src/api/mall/product/brand.ts` |
| `GET` | `/admin-api/product/brand/page` | `admin/src/api/mall/product/brand.ts` |
| `GET` | `/admin-api/product/brand/list-all-simple` | `admin/src/api/mall/product/brand.ts` |

#### 📌 商品规格与属性值 (前缀: `/product/property` - 12 个接口)

| HTTP 方法 | 后端完整路径 (`/admin-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `POST` | `/admin-api/product/property/create` | `admin/src/api/mall/product/property.ts` |
| `PUT` | `/admin-api/product/property/update` | `admin/src/api/mall/product/property.ts` |
| `DELETE` | `/admin-api/product/property/delete` | `admin/src/api/mall/product/property.ts` |
| `GET` | `/admin-api/product/property/get` | `admin/src/api/mall/product/property.ts` |
| `GET` | `/admin-api/product/property/page` | `admin/src/api/mall/product/property.ts` |
| `GET` | `/admin-api/product/property/simple-list` | `admin/src/api/mall/product/property.ts` |
| `GET` | `/admin-api/product/property/value/page` | `admin/src/api/mall/product/property.ts` |
| `GET` | `/admin-api/product/property/value/get` | `admin/src/api/mall/product/property.ts` |
| `POST` | `/admin-api/product/property/value/create` | `admin/src/api/mall/product/property.ts` |
| `PUT` | `/admin-api/product/property/value/update` | `admin/src/api/mall/product/property.ts` |
| `DELETE` | `/admin-api/product/property/value/delete` | `admin/src/api/mall/product/property.ts` |
| `GET` | `/admin-api/product/property/value/simple-list` | `admin/src/api/mall/product/property.ts` |

#### 📌 商品评价 (前缀: `/product/comment` - 5 个接口)

| HTTP 方法 | 后端完整路径 (`/admin-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/admin-api/product/comment/page` | `admin/src/api/mall/product/comment.ts` |
| `GET` | `/admin-api/product/comment/get` | `admin/src/api/mall/product/comment.ts` |
| `POST` | `/admin-api/product/comment/create` | `admin/src/api/mall/product/comment.ts` |
| `PUT` | `/admin-api/product/comment/update-visible` | `admin/src/api/mall/product/comment.ts` |
| `PUT` | `/admin-api/product/comment/reply` | `admin/src/api/mall/product/comment.ts` |

#### 📌 商品浏览记录 (前缀: `/product/browse-history` - 1 个接口)

| HTTP 方法 | 后端完整路径 (`/admin-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/admin-api/product/browse-history/page` | `admin/src/api/mall/product/history.ts` |

### 1.交易中心 (Trade)

#### 📌 交易订单管理 (前缀: `/trade/order` - 9 个接口)

| HTTP 方法 | 后端完整路径 (`/admin-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/admin-api/trade/order/summary` | `admin/src/api/mall/trade/order/index.ts` |
| `GET` | `/admin-api/trade/order/get-detail` | `admin/src/api/mall/trade/order/index.ts` |
| `PUT` | `/admin-api/trade/order/delivery` | `admin/src/api/mall/trade/order/index.ts` |
| `PUT` | `/admin-api/trade/order/update-remark` | `admin/src/api/mall/trade/order/index.ts` |
| `PUT` | `/admin-api/trade/order/update-price` | `admin/src/api/mall/trade/order/index.ts` |
| `PUT` | `/admin-api/trade/order/update-address` | `admin/src/api/mall/trade/order/index.ts` |
| `PUT` | `/admin-api/trade/order/pick-up-by-id` | `admin/src/api/mall/trade/order/index.ts` |
| `PUT` | `/admin-api/trade/order/pick-up-by-verify-code` | `admin/src/api/mall/trade/order/index.ts` |
| `GET` | `/admin-api/trade/order/get-by-pick-up-verify-code` | `admin/src/api/mall/trade/order/index.ts` |

#### 📌 售后退款管理 (前缀: `/trade/after-sale` - 6 个接口)

| HTTP 方法 | 后端完整路径 (`/admin-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/admin-api/trade/after-sale/get-detail` | `admin/src/api/mall/trade/afterSale/index.ts` |
| `PUT` | `/admin-api/trade/after-sale/agree` | `admin/src/api/mall/trade/afterSale/index.ts` |
| `PUT` | `/admin-api/trade/after-sale/disagree` | `admin/src/api/mall/trade/afterSale/index.ts` |
| `PUT` | `/admin-api/trade/after-sale/receive` | `admin/src/api/mall/trade/afterSale/index.ts` |
| `PUT` | `/admin-api/trade/after-sale/refuse` | `admin/src/api/mall/trade/afterSale/index.ts` |
| `PUT` | `/admin-api/trade/after-sale/refund` | `admin/src/api/mall/trade/afterSale/index.ts` |

#### 📌 物流配送与自提门店 (前缀: `/trade/delivery` - 17 个接口)

| HTTP 方法 | 后端完整路径 (`/admin-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/admin-api/trade/delivery/express/get` | `admin/src/api/mall/trade/delivery/express/index.ts` |
| `GET` | `/admin-api/trade/delivery/express/list-all-simple` | `admin/src/api/mall/trade/delivery/express/index.ts` |
| `POST` | `/admin-api/trade/delivery/express/create` | `admin/src/api/mall/trade/delivery/express/index.ts` |
| `PUT` | `/admin-api/trade/delivery/express/update` | `admin/src/api/mall/trade/delivery/express/index.ts` |
| `DELETE` | `/admin-api/trade/delivery/express/delete` | `admin/src/api/mall/trade/delivery/express/index.ts` |
| `DOWNLOAD` | `/admin-api/trade/delivery/express/export-excel` | `admin/src/api/mall/trade/delivery/express/index.ts` |
| `GET` | `/admin-api/trade/delivery/express-template/get` | `admin/src/api/mall/trade/delivery/expressTemplate/index.ts` |
| `GET` | `/admin-api/trade/delivery/express-template/list-all-simple` | `admin/src/api/mall/trade/delivery/expressTemplate/index.ts` |
| `POST` | `/admin-api/trade/delivery/express-template/create` | `admin/src/api/mall/trade/delivery/expressTemplate/index.ts` |
| `PUT` | `/admin-api/trade/delivery/express-template/update` | `admin/src/api/mall/trade/delivery/expressTemplate/index.ts` |
| `DELETE` | `/admin-api/trade/delivery/express-template/delete` | `admin/src/api/mall/trade/delivery/expressTemplate/index.ts` |
| `GET` | `/admin-api/trade/delivery/pick-up-store/get` | `admin/src/api/mall/trade/delivery/pickUpStore/index.ts` |
| `GET` | `/admin-api/trade/delivery/pick-up-store/simple-list` | `admin/src/api/mall/trade/delivery/pickUpStore/index.ts` |
| `POST` | `/admin-api/trade/delivery/pick-up-store/create` | `admin/src/api/mall/trade/delivery/pickUpStore/index.ts` |
| `PUT` | `/admin-api/trade/delivery/pick-up-store/update` | `admin/src/api/mall/trade/delivery/pickUpStore/index.ts` |
| `DELETE` | `/admin-api/trade/delivery/pick-up-store/delete` | `admin/src/api/mall/trade/delivery/pickUpStore/index.ts` |
| `POST` | `/admin-api/trade/delivery/pick-up-store/bind` | `admin/src/api/mall/trade/delivery/pickUpStore/index.ts` |

#### 📌 交易设置 (前缀: `/trade/config` - 2 个接口)

| HTTP 方法 | 后端完整路径 (`/admin-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/admin-api/trade/config/get` | `admin/src/api/mall/trade/config/index.ts` |
| `PUT` | `/admin-api/trade/config/save` | `admin/src/api/mall/trade/config/index.ts` |

#### 📌 分销用户 (前缀: `/trade/brokerage-user` - 5 个接口)

| HTTP 方法 | 后端完整路径 (`/admin-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `POST` | `/admin-api/trade/brokerage-user/create` | `admin/src/api/mall/trade/brokerage/user/index.ts` |
| `GET` | `/admin-api/trade/brokerage-user/get` | `admin/src/api/mall/trade/brokerage/user/index.ts` |
| `PUT` | `/admin-api/trade/brokerage-user/update-bind-user` | `admin/src/api/mall/trade/brokerage/user/index.ts` |
| `PUT` | `/admin-api/trade/brokerage-user/clear-bind-user` | `admin/src/api/mall/trade/brokerage/user/index.ts` |
| `PUT` | `/admin-api/trade/brokerage-user/update-brokerage-enable` | `admin/src/api/mall/trade/brokerage/user/index.ts` |

#### 📌 分销佣金记录 (前缀: `/trade/brokerage-record` - 2 个接口)

| HTTP 方法 | 后端完整路径 (`/admin-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/admin-api/trade/brokerage-record/page` | `admin/src/api/mall/trade/brokerage/record/index.ts` |
| `GET` | `/admin-api/trade/brokerage-record/get` | `admin/src/api/mall/trade/brokerage/record/index.ts` |

#### 📌 佣金提现审核 (前缀: `/trade/brokerage-withdraw` - 4 个接口)

| HTTP 方法 | 后端完整路径 (`/admin-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/admin-api/trade/brokerage-withdraw/page` | `admin/src/api/mall/trade/brokerage/withdraw/index.ts` |
| `GET` | `/admin-api/trade/brokerage-withdraw/get` | `admin/src/api/mall/trade/brokerage/withdraw/index.ts` |
| `PUT` | `/admin-api/trade/brokerage-withdraw/approve` | `admin/src/api/mall/trade/brokerage/withdraw/index.ts` |
| `PUT` | `/admin-api/trade/brokerage-withdraw/reject` | `admin/src/api/mall/trade/brokerage/withdraw/index.ts` |

### 1.营销中心 (Promotion)

#### 📌 优惠券模板 (前缀: `/promotion/coupon-template` - 8 个接口)

| HTTP 方法 | 后端完整路径 (`/admin-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `POST` | `/admin-api/promotion/coupon-template/create` | `admin/src/api/mall/promotion/coupon/couponTemplate.ts` |
| `PUT` | `/admin-api/promotion/coupon-template/update` | `admin/src/api/mall/promotion/coupon/couponTemplate.ts` |
| `PUT` | `/admin-api/promotion/coupon-template/update-status` | `admin/src/api/mall/promotion/coupon/couponTemplate.ts` |
| `DELETE` | `/admin-api/promotion/coupon-template/delete` | `admin/src/api/mall/promotion/coupon/couponTemplate.ts` |
| `GET` | `/admin-api/promotion/coupon-template/get` | `admin/src/api/mall/promotion/coupon/couponTemplate.ts` |
| `GET` | `/admin-api/promotion/coupon-template/page` | `admin/src/api/mall/promotion/coupon/couponTemplate.ts` |
| `GET` | `/admin-api/promotion/coupon-template/list` | `admin/src/api/mall/promotion/coupon/couponTemplate.ts` |
| `GET` | `/admin-api/promotion/coupon-template/export-excel` | `admin/src/api/mall/promotion/coupon/couponTemplate.ts` |

#### 📌 优惠券发放与记录 (前缀: `/promotion/coupon` - 11 个接口)

| HTTP 方法 | 后端完整路径 (`/admin-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `DELETE` | `/admin-api/promotion/coupon/delete` | `admin/src/api/mall/promotion/coupon/coupon.ts` |
| `GET` | `/admin-api/promotion/coupon/page` | `admin/src/api/mall/promotion/coupon/coupon.ts` |
| `POST` | `/admin-api/promotion/coupon/send` | `admin/src/api/mall/promotion/coupon/coupon.ts` |
| `POST` | `/admin-api/promotion/coupon-template/create` | `admin/src/api/mall/promotion/coupon/couponTemplate.ts` |
| `PUT` | `/admin-api/promotion/coupon-template/update` | `admin/src/api/mall/promotion/coupon/couponTemplate.ts` |
| `PUT` | `/admin-api/promotion/coupon-template/update-status` | `admin/src/api/mall/promotion/coupon/couponTemplate.ts` |
| `DELETE` | `/admin-api/promotion/coupon-template/delete` | `admin/src/api/mall/promotion/coupon/couponTemplate.ts` |
| `GET` | `/admin-api/promotion/coupon-template/get` | `admin/src/api/mall/promotion/coupon/couponTemplate.ts` |
| `GET` | `/admin-api/promotion/coupon-template/page` | `admin/src/api/mall/promotion/coupon/couponTemplate.ts` |
| `GET` | `/admin-api/promotion/coupon-template/list` | `admin/src/api/mall/promotion/coupon/couponTemplate.ts` |
| `GET` | `/admin-api/promotion/coupon-template/export-excel` | `admin/src/api/mall/promotion/coupon/couponTemplate.ts` |

#### 📌 限时折扣 (前缀: `/promotion/discount-activity` - 6 个接口)

| HTTP 方法 | 后端完整路径 (`/admin-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/admin-api/promotion/discount-activity/page` | `admin/src/api/mall/promotion/discount/discountActivity.ts` |
| `GET` | `/admin-api/promotion/discount-activity/get` | `admin/src/api/mall/promotion/discount/discountActivity.ts` |
| `POST` | `/admin-api/promotion/discount-activity/create` | `admin/src/api/mall/promotion/discount/discountActivity.ts` |
| `PUT` | `/admin-api/promotion/discount-activity/update` | `admin/src/api/mall/promotion/discount/discountActivity.ts` |
| `PUT` | `/admin-api/promotion/discount-activity/close` | `admin/src/api/mall/promotion/discount/discountActivity.ts` |
| `DELETE` | `/admin-api/promotion/discount-activity/delete` | `admin/src/api/mall/promotion/discount/discountActivity.ts` |

#### 📌 满减送 (前缀: `/promotion/reward-activity` - 6 个接口)

| HTTP 方法 | 后端完整路径 (`/admin-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `POST` | `/admin-api/promotion/reward-activity/create` | `admin/src/api/mall/promotion/reward/rewardActivity.ts` |
| `PUT` | `/admin-api/promotion/reward-activity/update` | `admin/src/api/mall/promotion/reward/rewardActivity.ts` |
| `GET` | `/admin-api/promotion/reward-activity/page` | `admin/src/api/mall/promotion/reward/rewardActivity.ts` |
| `GET` | `/admin-api/promotion/reward-activity/get` | `admin/src/api/mall/promotion/reward/rewardActivity.ts` |
| `DELETE` | `/admin-api/promotion/reward-activity/delete` | `admin/src/api/mall/promotion/reward/rewardActivity.ts` |
| `PUT` | `/admin-api/promotion/reward-activity/close` | `admin/src/api/mall/promotion/reward/rewardActivity.ts` |

#### 📌 秒杀活动 (前缀: `/promotion/seckill-activity` - 7 个接口)

| HTTP 方法 | 后端完整路径 (`/admin-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/admin-api/promotion/seckill-activity/page` | `admin/src/api/mall/promotion/seckill/seckillActivity.ts` |
| `GET` | `/admin-api/promotion/seckill-activity/list-by-ids` | `admin/src/api/mall/promotion/seckill/seckillActivity.ts` |
| `GET` | `/admin-api/promotion/seckill-activity/get` | `admin/src/api/mall/promotion/seckill/seckillActivity.ts` |
| `POST` | `/admin-api/promotion/seckill-activity/create` | `admin/src/api/mall/promotion/seckill/seckillActivity.ts` |
| `PUT` | `/admin-api/promotion/seckill-activity/update` | `admin/src/api/mall/promotion/seckill/seckillActivity.ts` |
| `PUT` | `/admin-api/promotion/seckill-activity/close` | `admin/src/api/mall/promotion/seckill/seckillActivity.ts` |
| `DELETE` | `/admin-api/promotion/seckill-activity/delete` | `admin/src/api/mall/promotion/seckill/seckillActivity.ts` |

#### 📌 秒杀时段 (前缀: `/promotion/seckill-config` - 7 个接口)

| HTTP 方法 | 后端完整路径 (`/admin-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/admin-api/promotion/seckill-config/page` | `admin/src/api/mall/promotion/seckill/seckillConfig.ts` |
| `GET` | `/admin-api/promotion/seckill-config/list` | `admin/src/api/mall/promotion/seckill/seckillConfig.ts` |
| `GET` | `/admin-api/promotion/seckill-config/get` | `admin/src/api/mall/promotion/seckill/seckillConfig.ts` |
| `POST` | `/admin-api/promotion/seckill-config/create` | `admin/src/api/mall/promotion/seckill/seckillConfig.ts` |
| `PUT` | `/admin-api/promotion/seckill-config/update` | `admin/src/api/mall/promotion/seckill/seckillConfig.ts` |
| `DELETE` | `/admin-api/promotion/seckill-config/delete` | `admin/src/api/mall/promotion/seckill/seckillConfig.ts` |
| `PUT` | `/admin-api/promotion/seckill-config/update-status` | `admin/src/api/mall/promotion/seckill/seckillConfig.ts` |

#### 📌 拼团活动 (前缀: `/promotion/combination-activity` - 7 个接口)

| HTTP 方法 | 后端完整路径 (`/admin-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/admin-api/promotion/combination-activity/page` | `admin/src/api/mall/promotion/combination/combinationActivity.ts` |
| `GET` | `/admin-api/promotion/combination-activity/get` | `admin/src/api/mall/promotion/combination/combinationActivity.ts` |
| `GET` | `/admin-api/promotion/combination-activity/list-by-ids` | `admin/src/api/mall/promotion/combination/combinationActivity.ts` |
| `POST` | `/admin-api/promotion/combination-activity/create` | `admin/src/api/mall/promotion/combination/combinationActivity.ts` |
| `PUT` | `/admin-api/promotion/combination-activity/update` | `admin/src/api/mall/promotion/combination/combinationActivity.ts` |
| `PUT` | `/admin-api/promotion/combination-activity/close` | `admin/src/api/mall/promotion/combination/combinationActivity.ts` |
| `DELETE` | `/admin-api/promotion/combination-activity/delete` | `admin/src/api/mall/promotion/combination/combinationActivity.ts` |

#### 📌 拼团记录 (前缀: `/promotion/combination-record` - 2 个接口)

| HTTP 方法 | 后端完整路径 (`/admin-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/admin-api/promotion/combination-record/page` | `admin/src/api/mall/promotion/combination/combinationRecord.ts` |
| `GET` | `/admin-api/promotion/combination-record/get-summary` | `admin/src/api/mall/promotion/combination/combinationRecord.ts` |

#### 📌 砍价活动 (前缀: `/promotion/bargain-activity` - 6 个接口)

| HTTP 方法 | 后端完整路径 (`/admin-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/admin-api/promotion/bargain-activity/page` | `admin/src/api/mall/promotion/bargain/bargainActivity.ts` |
| `GET` | `/admin-api/promotion/bargain-activity/get` | `admin/src/api/mall/promotion/bargain/bargainActivity.ts` |
| `POST` | `/admin-api/promotion/bargain-activity/create` | `admin/src/api/mall/promotion/bargain/bargainActivity.ts` |
| `PUT` | `/admin-api/promotion/bargain-activity/update` | `admin/src/api/mall/promotion/bargain/bargainActivity.ts` |
| `PUT` | `/admin-api/promotion/bargain-activity/close` | `admin/src/api/mall/promotion/bargain/bargainActivity.ts` |
| `DELETE` | `/admin-api/promotion/bargain-activity/delete` | `admin/src/api/mall/promotion/bargain/bargainActivity.ts` |

#### 📌 砍价记录 (前缀: `/promotion/bargain-record` - 1 个接口)

| HTTP 方法 | 后端完整路径 (`/admin-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/admin-api/promotion/bargain-record/page` | `admin/src/api/mall/promotion/bargain/bargainRecord.ts` |

#### 📌 砍价助力 (前缀: `/promotion/bargain-help` - 1 个接口)

| HTTP 方法 | 后端完整路径 (`/admin-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/admin-api/promotion/bargain-help/page` | `admin/src/api/mall/promotion/bargain/bargainHelp.ts` |

#### 📌 积分商城 (前缀: `/promotion/point-activity` - 7 个接口)

| HTTP 方法 | 后端完整路径 (`/admin-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/admin-api/promotion/point-activity/page` | `admin/src/api/mall/promotion/point/index.ts` |
| `GET` | `/admin-api/promotion/point-activity/get` | `admin/src/api/mall/promotion/point/index.ts` |
| `GET` | `/admin-api/promotion/point-activity/list-by-ids` | `admin/src/api/mall/promotion/point/index.ts` |
| `POST` | `/admin-api/promotion/point-activity/create` | `admin/src/api/mall/promotion/point/index.ts` |
| `PUT` | `/admin-api/promotion/point-activity/update` | `admin/src/api/mall/promotion/point/index.ts` |
| `DELETE` | `/admin-api/promotion/point-activity/delete` | `admin/src/api/mall/promotion/point/index.ts` |
| `PUT` | `/admin-api/promotion/point-activity/close` | `admin/src/api/mall/promotion/point/index.ts` |

#### 📌 轮播图广告 (前缀: `/promotion/banner` - 5 个接口)

| HTTP 方法 | 后端完整路径 (`/admin-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/admin-api/promotion/banner/page` | `admin/src/api/mall/promotion/banner/index.ts` |
| `GET` | `/admin-api/promotion/banner/get` | `admin/src/api/mall/promotion/banner/index.ts` |
| `POST` | `/admin-api/promotion/banner/create` | `admin/src/api/mall/promotion/banner/index.ts` |
| `PUT` | `/admin-api/promotion/banner/update` | `admin/src/api/mall/promotion/banner/index.ts` |
| `DELETE` | `/admin-api/promotion/banner/delete` | `admin/src/api/mall/promotion/banner/index.ts` |

#### 📌 文章分类 (前缀: `/promotion/article-category` - 6 个接口)

| HTTP 方法 | 后端完整路径 (`/admin-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/admin-api/promotion/article-category/page` | `admin/src/api/mall/promotion/articleCategory/index.ts` |
| `GET` | `/admin-api/promotion/article-category/list-all-simple` | `admin/src/api/mall/promotion/articleCategory/index.ts` |
| `GET` | `/admin-api/promotion/article-category/get` | `admin/src/api/mall/promotion/articleCategory/index.ts` |
| `POST` | `/admin-api/promotion/article-category/create` | `admin/src/api/mall/promotion/articleCategory/index.ts` |
| `PUT` | `/admin-api/promotion/article-category/update` | `admin/src/api/mall/promotion/articleCategory/index.ts` |
| `DELETE` | `/admin-api/promotion/article-category/delete` | `admin/src/api/mall/promotion/articleCategory/index.ts` |

#### 📌 商城文章 (前缀: `/promotion/article` - 11 个接口)

| HTTP 方法 | 后端完整路径 (`/admin-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/admin-api/promotion/article/page` | `admin/src/api/mall/promotion/article/index.ts` |
| `GET` | `/admin-api/promotion/article/get` | `admin/src/api/mall/promotion/article/index.ts` |
| `POST` | `/admin-api/promotion/article/create` | `admin/src/api/mall/promotion/article/index.ts` |
| `PUT` | `/admin-api/promotion/article/update` | `admin/src/api/mall/promotion/article/index.ts` |
| `DELETE` | `/admin-api/promotion/article/delete` | `admin/src/api/mall/promotion/article/index.ts` |
| `GET` | `/admin-api/promotion/article-category/page` | `admin/src/api/mall/promotion/articleCategory/index.ts` |
| `GET` | `/admin-api/promotion/article-category/list-all-simple` | `admin/src/api/mall/promotion/articleCategory/index.ts` |
| `GET` | `/admin-api/promotion/article-category/get` | `admin/src/api/mall/promotion/articleCategory/index.ts` |
| `POST` | `/admin-api/promotion/article-category/create` | `admin/src/api/mall/promotion/articleCategory/index.ts` |
| `PUT` | `/admin-api/promotion/article-category/update` | `admin/src/api/mall/promotion/articleCategory/index.ts` |
| `DELETE` | `/admin-api/promotion/article-category/delete` | `admin/src/api/mall/promotion/articleCategory/index.ts` |

#### 📌 店铺装修模板 (前缀: `/promotion/diy-template` - 8 个接口)

| HTTP 方法 | 后端完整路径 (`/admin-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/admin-api/promotion/diy-template/page` | `admin/src/api/mall/promotion/diy/template.ts` |
| `GET` | `/admin-api/promotion/diy-template/get` | `admin/src/api/mall/promotion/diy/template.ts` |
| `POST` | `/admin-api/promotion/diy-template/create` | `admin/src/api/mall/promotion/diy/template.ts` |
| `PUT` | `/admin-api/promotion/diy-template/update` | `admin/src/api/mall/promotion/diy/template.ts` |
| `DELETE` | `/admin-api/promotion/diy-template/delete` | `admin/src/api/mall/promotion/diy/template.ts` |
| `PUT` | `/admin-api/promotion/diy-template/use` | `admin/src/api/mall/promotion/diy/template.ts` |
| `GET` | `/admin-api/promotion/diy-template/get-property` | `admin/src/api/mall/promotion/diy/template.ts` |
| `PUT` | `/admin-api/promotion/diy-template/update-property` | `admin/src/api/mall/promotion/diy/template.ts` |

#### 📌 装修自定义页面 (前缀: `/promotion/diy-page` - 7 个接口)

| HTTP 方法 | 后端完整路径 (`/admin-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/admin-api/promotion/diy-page/page` | `admin/src/api/mall/promotion/diy/page.ts` |
| `GET` | `/admin-api/promotion/diy-page/get` | `admin/src/api/mall/promotion/diy/page.ts` |
| `POST` | `/admin-api/promotion/diy-page/create` | `admin/src/api/mall/promotion/diy/page.ts` |
| `PUT` | `/admin-api/promotion/diy-page/update` | `admin/src/api/mall/promotion/diy/page.ts` |
| `DELETE` | `/admin-api/promotion/diy-page/delete` | `admin/src/api/mall/promotion/diy/page.ts` |
| `GET` | `/admin-api/promotion/diy-page/get-property` | `admin/src/api/mall/promotion/diy/page.ts` |
| `PUT` | `/admin-api/promotion/diy-page/update-property` | `admin/src/api/mall/promotion/diy/page.ts` |

#### 📌 客服会话 (前缀: `/promotion/kefu-conversation` - 4 个接口)

| HTTP 方法 | 后端完整路径 (`/admin-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/admin-api/promotion/kefu-conversation/list` | `admin/src/api/mall/promotion/kefu/conversation/index.ts` |
| `GET` | `/admin-api/promotion/kefu-conversation/get` | `admin/src/api/mall/promotion/kefu/conversation/index.ts` |
| `PUT` | `/admin-api/promotion/kefu-conversation/update-conversation-pinned` | `admin/src/api/mall/promotion/kefu/conversation/index.ts` |
| `DELETE` | `/admin-api/promotion/kefu-conversation/delete` | `admin/src/api/mall/promotion/kefu/conversation/index.ts` |

#### 📌 客服消息 (前缀: `/promotion/kefu-message` - 3 个接口)

| HTTP 方法 | 后端完整路径 (`/admin-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `POST` | `/admin-api/promotion/kefu-message/send` | `admin/src/api/mall/promotion/kefu/message/index.ts` |
| `PUT` | `/admin-api/promotion/kefu-message/update-read-status` | `admin/src/api/mall/promotion/kefu/message/index.ts` |
| `GET` | `/admin-api/promotion/kefu-message/list` | `admin/src/api/mall/promotion/kefu/message/index.ts` |

### 1.会员中心 (Member)

#### 📌 会员用户管理 (前缀: `/member/user` - 4 个接口)

| HTTP 方法 | 后端完整路径 (`/admin-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/admin-api/member/user/get` | `admin/src/api/member/user/index.ts` |
| `PUT` | `/admin-api/member/user/update` | `admin/src/api/member/user/index.ts` |
| `PUT` | `/admin-api/member/user/update-level` | `admin/src/api/member/user/index.ts` |
| `PUT` | `/admin-api/member/user/update-point` | `admin/src/api/member/user/index.ts` |

#### 📌 会员等级 (前缀: `/member/level` - 6 个接口)

| HTTP 方法 | 后端完整路径 (`/admin-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/admin-api/member/level/list` | `admin/src/api/member/level/index.ts` |
| `GET` | `/admin-api/member/level/get` | `admin/src/api/member/level/index.ts` |
| `GET` | `/admin-api/member/level/list-all-simple` | `admin/src/api/member/level/index.ts` |
| `POST` | `/admin-api/member/level/create` | `admin/src/api/member/level/index.ts` |
| `PUT` | `/admin-api/member/level/update` | `admin/src/api/member/level/index.ts` |
| `DELETE` | `/admin-api/member/level/delete` | `admin/src/api/member/level/index.ts` |

#### 📌 会员分组 (前缀: `/member/group` - 6 个接口)

| HTTP 方法 | 后端完整路径 (`/admin-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/admin-api/member/group/page` | `admin/src/api/member/group/index.ts` |
| `GET` | `/admin-api/member/group/get` | `admin/src/api/member/group/index.ts` |
| `POST` | `/admin-api/member/group/create` | `admin/src/api/member/group/index.ts` |
| `GET` | `/admin-api/member/group/list-all-simple` | `admin/src/api/member/group/index.ts` |
| `PUT` | `/admin-api/member/group/update` | `admin/src/api/member/group/index.ts` |
| `DELETE` | `/admin-api/member/group/delete` | `admin/src/api/member/group/index.ts` |

#### 📌 会员标签 (前缀: `/member/tag` - 6 个接口)

| HTTP 方法 | 后端完整路径 (`/admin-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/admin-api/member/tag/page` | `admin/src/api/member/tag/index.ts` |
| `GET` | `/admin-api/member/tag/get` | `admin/src/api/member/tag/index.ts` |
| `GET` | `/admin-api/member/tag/list-all-simple` | `admin/src/api/member/tag/index.ts` |
| `POST` | `/admin-api/member/tag/create` | `admin/src/api/member/tag/index.ts` |
| `PUT` | `/admin-api/member/tag/update` | `admin/src/api/member/tag/index.ts` |
| `DELETE` | `/admin-api/member/tag/delete` | `admin/src/api/member/tag/index.ts` |

#### 📌 积分记录 (前缀: `/member/point` - 1 个接口)

| HTTP 方法 | 后端完整路径 (`/admin-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/admin-api/member/point/record/page` | `admin/src/api/member/point/record/index.ts` |

#### 📌 成长值记录 (前缀: `/member/experience-record` - 1 个接口)

| HTTP 方法 | 后端完整路径 (`/admin-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/admin-api/member/experience-record/get` | `admin/src/api/member/experience-record/index.ts` |

#### 📌 签到规则与记录 (前缀: `/member/sign-in` - 6 个接口)

| HTTP 方法 | 后端完整路径 (`/admin-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/admin-api/member/sign-in/config/list` | `admin/src/api/member/signin/config/index.ts` |
| `GET` | `/admin-api/member/sign-in/config/get` | `admin/src/api/member/signin/config/index.ts` |
| `POST` | `/admin-api/member/sign-in/config/create` | `admin/src/api/member/signin/config/index.ts` |
| `PUT` | `/admin-api/member/sign-in/config/update` | `admin/src/api/member/signin/config/index.ts` |
| `DELETE` | `/admin-api/member/sign-in/config/delete` | `admin/src/api/member/signin/config/index.ts` |
| `GET` | `/admin-api/member/sign-in/record/page` | `admin/src/api/member/signin/record/index.ts` |

#### 📌 会员收件地址 (前缀: `/member/address` - 1 个接口)

| HTTP 方法 | 后端完整路径 (`/admin-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/admin-api/member/address/list` | `admin/src/api/member/address/index.ts` |

#### 📌 会员全局配置 (前缀: `/member/config` - 2 个接口)

| HTTP 方法 | 后端完整路径 (`/admin-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/admin-api/member/config/get` | `admin/src/api/member/config/index.ts` |
| `PUT` | `/admin-api/member/config/save` | `admin/src/api/member/config/index.ts` |

### 1.支付与钱包 (Pay)

#### 📌 支付应用配置 (前缀: `/pay/app` - 7 个接口)

| HTTP 方法 | 后端完整路径 (`/admin-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/admin-api/pay/app/page` | `admin/src/api/pay/app/index.ts` |
| `GET` | `/admin-api/pay/app/get` | `admin/src/api/pay/app/index.ts` |
| `POST` | `/admin-api/pay/app/create` | `admin/src/api/pay/app/index.ts` |
| `PUT` | `/admin-api/pay/app/update` | `admin/src/api/pay/app/index.ts` |
| `PUT` | `/admin-api/pay/app/update-status` | `admin/src/api/pay/app/index.ts` |
| `DELETE` | `/admin-api/pay/app/delete` | `admin/src/api/pay/app/index.ts` |
| `GET` | `/admin-api/pay/app/list` | `admin/src/api/pay/app/index.ts` |

#### 📌 支付订单 (前缀: `/pay/order` - 4 个接口)

| HTTP 方法 | 后端完整路径 (`/admin-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/admin-api/pay/order/get` | `admin/src/api/pay/order/index.ts` |
| `GET` | `/admin-api/pay/order/get-detail` | `admin/src/api/pay/order/index.ts` |
| `POST` | `/admin-api/pay/order/submit` | `admin/src/api/pay/order/index.ts` |
| `DOWNLOAD` | `/admin-api/pay/order/export-excel` | `admin/src/api/pay/order/index.ts` |

#### 📌 退款订单 (前缀: `/pay/refund` - 5 个接口)

| HTTP 方法 | 后端完整路径 (`/admin-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/admin-api/pay/refund/get` | `admin/src/api/pay/refund/index.ts` |
| `POST` | `/admin-api/pay/refund/create` | `admin/src/api/pay/refund/index.ts` |
| `PUT` | `/admin-api/pay/refund/update` | `admin/src/api/pay/refund/index.ts` |
| `DELETE` | `/admin-api/pay/refund/delete` | `admin/src/api/pay/refund/index.ts` |
| `DOWNLOAD` | `/admin-api/pay/refund/export-excel` | `admin/src/api/pay/refund/index.ts` |

#### 📌 充值套餐 (前缀: `/pay/wallet-recharge-package` - 5 个接口)

| HTTP 方法 | 后端完整路径 (`/admin-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/admin-api/pay/wallet-recharge-package/page` | `admin/src/api/pay/wallet/rechargePackage/index.ts` |
| `GET` | `/admin-api/pay/wallet-recharge-package/get` | `admin/src/api/pay/wallet/rechargePackage/index.ts` |
| `POST` | `/admin-api/pay/wallet-recharge-package/create` | `admin/src/api/pay/wallet/rechargePackage/index.ts` |
| `PUT` | `/admin-api/pay/wallet-recharge-package/update` | `admin/src/api/pay/wallet/rechargePackage/index.ts` |
| `DELETE` | `/admin-api/pay/wallet-recharge-package/delete` | `admin/src/api/pay/wallet/rechargePackage/index.ts` |

#### 📌 钱包流水 (前缀: `/pay/wallet-transaction` - 1 个接口)

| HTTP 方法 | 后端完整路径 (`/admin-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/admin-api/pay/wallet-transaction/page` | `admin/src/api/pay/wallet/transaction/index.ts` |

#### 📌 会员钱包管理 (前缀: `/pay/wallet` - 9 个接口)

| HTTP 方法 | 后端完整路径 (`/admin-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/admin-api/pay/wallet/get` | `admin/src/api/pay/wallet/balance/index.ts` |
| `GET` | `/admin-api/pay/wallet/page` | `admin/src/api/pay/wallet/balance/index.ts` |
| `PUT` | `/admin-api/pay/wallet/update-balance` | `admin/src/api/pay/wallet/balance/index.ts` |
| `GET` | `/admin-api/pay/wallet-recharge-package/page` | `admin/src/api/pay/wallet/rechargePackage/index.ts` |
| `GET` | `/admin-api/pay/wallet-recharge-package/get` | `admin/src/api/pay/wallet/rechargePackage/index.ts` |
| `POST` | `/admin-api/pay/wallet-recharge-package/create` | `admin/src/api/pay/wallet/rechargePackage/index.ts` |
| `PUT` | `/admin-api/pay/wallet-recharge-package/update` | `admin/src/api/pay/wallet/rechargePackage/index.ts` |
| `DELETE` | `/admin-api/pay/wallet-recharge-package/delete` | `admin/src/api/pay/wallet/rechargePackage/index.ts` |
| `GET` | `/admin-api/pay/wallet-transaction/page` | `admin/src/api/pay/wallet/transaction/index.ts` |

### 1.统计分析 (Statistics)

#### 📌 交易数据统计 (前缀: `/statistics/trade` - 3 个接口)

| HTTP 方法 | 后端完整路径 (`/admin-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/admin-api/statistics/trade/list` | `admin/src/api/mall/statistics/trade.ts` |
| `DOWNLOAD` | `/admin-api/statistics/trade/export-excel` | `admin/src/api/mall/statistics/trade.ts` |
| `GET` | `/admin-api/statistics/trade/order-count` | `admin/src/api/mall/statistics/trade.ts` |

#### 📌 会员数据统计 (前缀: `/statistics/member` - 6 个接口)

| HTTP 方法 | 后端完整路径 (`/admin-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/admin-api/statistics/member/summary` | `admin/src/api/mall/statistics/member.ts` |
| `GET` | `/admin-api/statistics/member/analyse` | `admin/src/api/mall/statistics/member.ts` |
| `GET` | `/admin-api/statistics/member/area-statistics-list` | `admin/src/api/mall/statistics/member.ts` |
| `GET` | `/admin-api/statistics/member/sex-statistics-list` | `admin/src/api/mall/statistics/member.ts` |
| `GET` | `/admin-api/statistics/member/terminal-statistics-list` | `admin/src/api/mall/statistics/member.ts` |
| `GET` | `/admin-api/statistics/member/register-count-list` | `admin/src/api/mall/statistics/member.ts` |

#### 📌 商品数据统计 (前缀: `/statistics/product` - 3 个接口)

| HTTP 方法 | 后端完整路径 (`/admin-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/admin-api/statistics/product/list` | `admin/src/api/mall/statistics/product.ts` |
| `DOWNLOAD` | `/admin-api/statistics/product/export-excel` | `admin/src/api/mall/statistics/product.ts` |
| `GET` | `/admin-api/statistics/product/rank-page` | `admin/src/api/mall/statistics/product.ts` |

#### 📌 支付数据统计 (前缀: `/statistics/pay` - 1 个接口)

| HTTP 方法 | 后端完整路径 (`/admin-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/admin-api/statistics/pay/summary` | `admin/src/api/mall/statistics/pay.ts` |

### 1.系统管理与基建 (System & Infra)

#### 📌 管理后台登录与权限认证 (前缀: `/system/auth` - 9 个接口)

| HTTP 方法 | 后端完整路径 (`/admin-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `POST` | `/admin-api/system/auth/login` | `admin/src/api/login/index.ts` |
| `POST` | `/admin-api/system/auth/register` | `admin/src/api/login/index.ts` |
| `POST` | `/admin-api/system/auth/logout` | `admin/src/api/login/index.ts` |
| `GET` | `/admin-api/system/auth/get-permission-info` | `admin/src/api/login/index.ts` |
| `POST` | `/admin-api/system/auth/send-sms-code` | `admin/src/api/login/index.ts` |
| `POST` | `/admin-api/system/auth/sms-login` | `admin/src/api/login/index.ts` |
| `POST` | `/admin-api/system/auth/social-login` | `admin/src/api/login/index.ts` |
| `GET` | `/admin-api/system/auth/social-auth-redirect` | `admin/src/api/login/index.ts` |
| `POST` | `/admin-api/system/auth/reset-password` | `admin/src/api/login/index.ts` |

#### 📌 管理员用户管理 (前缀: `/system/user` - 14 个接口)

| HTTP 方法 | 后端完整路径 (`/admin-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/admin-api/system/user/simple-list` | `admin/src/api/system/user/index.ts` |
| `GET` | `/admin-api/system/user/get-simple` | `admin/src/api/system/user/index.ts` |
| `GET` | `/admin-api/system/user/list-by-nickname` | `admin/src/api/system/user/index.ts` |
| `GET` | `/admin-api/system/user/page` | `admin/src/api/system/user/index.ts` |
| `GET` | `/admin-api/system/user/list` | `admin/src/api/system/user/index.ts` |
| `GET` | `/admin-api/system/user/get` | `admin/src/api/system/user/index.ts` |
| `POST` | `/admin-api/system/user/create` | `admin/src/api/system/user/index.ts` |
| `PUT` | `/admin-api/system/user/update` | `admin/src/api/system/user/index.ts` |
| `DELETE` | `/admin-api/system/user/delete` | `admin/src/api/system/user/index.ts` |
| `DELETE` | `/admin-api/system/user/delete-list` | `admin/src/api/system/user/index.ts` |
| `DOWNLOAD` | `/admin-api/system/user/export-excel` | `admin/src/api/system/user/index.ts` |
| `DOWNLOAD` | `/admin-api/system/user/get-import-template` | `admin/src/api/system/user/index.ts` |
| `PUT` | `/admin-api/system/user/update-password` | `admin/src/api/system/user/index.ts` |
| `PUT` | `/admin-api/system/user/update-status` | `admin/src/api/system/user/index.ts` |

#### 📌 部门组织架构 (前缀: `/system/dept` - 7 个接口)

| HTTP 方法 | 后端完整路径 (`/admin-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/admin-api/system/dept/simple-list` | `admin/src/api/system/dept/index.ts` |
| `GET` | `/admin-api/system/dept/list` | `admin/src/api/system/dept/index.ts` |
| `GET` | `/admin-api/system/dept/get` | `admin/src/api/system/dept/index.ts` |
| `POST` | `/admin-api/system/dept/create` | `admin/src/api/system/dept/index.ts` |
| `PUT` | `/admin-api/system/dept/update` | `admin/src/api/system/dept/index.ts` |
| `DELETE` | `/admin-api/system/dept/delete` | `admin/src/api/system/dept/index.ts` |
| `DELETE` | `/admin-api/system/dept/delete-list` | `admin/src/api/system/dept/index.ts` |

#### 📌 字典数据 (前缀: `/system/dict-data` - 9 个接口)

| HTTP 方法 | 后端完整路径 (`/admin-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/admin-api/system/dict-data/simple-list` | `admin/src/api/system/dict/dict.data.ts` |
| `GET` | `/admin-api/system/dict-data/page` | `admin/src/api/system/dict/dict.data.ts` |
| `GET` | `/admin-api/system/dict-data/get` | `admin/src/api/system/dict/dict.data.ts` |
| `GET` | `/admin-api/system/dict-data/type` | `admin/src/api/system/dict/dict.data.ts` |
| `POST` | `/admin-api/system/dict-data/create` | `admin/src/api/system/dict/dict.data.ts` |
| `PUT` | `/admin-api/system/dict-data/update` | `admin/src/api/system/dict/dict.data.ts` |
| `DELETE` | `/admin-api/system/dict-data/delete` | `admin/src/api/system/dict/dict.data.ts` |
| `DELETE` | `/admin-api/system/dict-data/delete-list` | `admin/src/api/system/dict/dict.data.ts` |
| `DOWNLOAD` | `/admin-api/system/dict-data/export-excel` | `admin/src/api/system/dict/dict.data.ts` |

#### 📌 省市区与IP定位 (前缀: `/system/area` - 2 个接口)

| HTTP 方法 | 后端完整路径 (`/admin-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/admin-api/system/area/tree` | `admin/src/api/system/area/index.ts` |
| `GET` | `/admin-api/system/area/get-by-ip` | `admin/src/api/system/area/index.ts` |

#### 📌 多租户管理 (前缀: `/system/tenant` - 2 个接口)

| HTTP 方法 | 后端完整路径 (`/admin-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/admin-api/system/tenant/get-id-by-name` | `admin/src/api/login/index.ts` |
| `GET` | `/admin-api/system/tenant/get-by-website` | `admin/src/api/login/index.ts` |

#### 📌 文件对象存储与上传 (前缀: `/infra/file` - 6 个接口)

| HTTP 方法 | 后端完整路径 (`/admin-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/admin-api/infra/file/page` | `admin/src/api/infra/file/index.ts` |
| `DELETE` | `/admin-api/infra/file/delete` | `admin/src/api/infra/file/index.ts` |
| `DELETE` | `/admin-api/infra/file/delete-list` | `admin/src/api/infra/file/index.ts` |
| `GET` | `/admin-api/infra/file/presigned-url` | `admin/src/api/infra/file/index.ts` |
| `POST` | `/admin-api/infra/file/create` | `admin/src/api/infra/file/index.ts` |
| `UPLOAD` | `/admin-api/infra/file/upload` | `admin/src/api/infra/file/index.ts` |

---

## 贰、 微信小程序端 (`/app-api`) 全量接口清单

### 2.会员中心与认证 (Member)

#### 📱 小程序登录与手机号授权 (前缀: `/member/auth` - 9 个接口)

| HTTP 方法 | 后端完整路径 (`/app-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `POST` | `/app-api/member/auth/login` | `miniapp/sheep/api/member/auth.js` |
| `POST` | `/app-api/member/auth/sms-login` | `miniapp/sheep/api/member/auth.js` |
| `POST` | `/app-api/member/auth/send-sms-code` | `miniapp/sheep/api/member/auth.js` |
| `POST` | `/app-api/member/auth/logout` | `miniapp/sheep/api/member/auth.js` |
| `POST` | `/app-api/member/auth/refresh-token` | `miniapp/sheep/api/member/auth.js` |
| `GET` | `/app-api/member/auth/social-auth-redirect` | `miniapp/sheep/api/member/auth.js` |
| `POST` | `/app-api/member/auth/social-login` | `miniapp/sheep/api/member/auth.js` |
| `POST` | `/app-api/member/auth/weixin-mini-app-login` | `miniapp/sheep/api/member/auth.js` |
| `POST` | `/app-api/member/auth/create-weixin-jsapi-signature` | `miniapp/sheep/api/member/auth.js` |

#### 📱 会员个人资料 (前缀: `/member/user` - 6 个接口)

| HTTP 方法 | 后端完整路径 (`/app-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/app-api/member/user/get` | `miniapp/sheep/api/member/user.js` |
| `PUT` | `/app-api/member/user/update` | `miniapp/sheep/api/member/user.js` |
| `PUT` | `/app-api/member/user/update-mobile` | `miniapp/sheep/api/member/user.js` |
| `PUT` | `/app-api/member/user/update-mobile-by-weixin` | `miniapp/sheep/api/member/user.js` |
| `PUT` | `/app-api/member/user/update-password` | `miniapp/sheep/api/member/user.js` |
| `PUT` | `/app-api/member/user/reset-password` | `miniapp/sheep/api/member/user.js` |

#### 📱 收货地址 (前缀: `/member/address` - 5 个接口)

| HTTP 方法 | 后端完整路径 (`/app-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/app-api/member/address/list` | `miniapp/sheep/api/member/address.js` |
| `POST` | `/app-api/member/address/create` | `miniapp/sheep/api/member/address.js` |
| `PUT` | `/app-api/member/address/update` | `miniapp/sheep/api/member/address.js` |
| `GET` | `/app-api/member/address/get` | `miniapp/sheep/api/member/address.js` |
| `DELETE` | `/app-api/member/address/delete` | `miniapp/sheep/api/member/address.js` |

#### 📱 会员积分 (前缀: `/member/point` - 1 个接口)

| HTTP 方法 | 后端完整路径 (`/app-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/app-api/member/point/record/page` | `miniapp/sheep/api/member/point.js` |

#### 📱 每日签到打卡 (前缀: `/member/sign-in` - 4 个接口)

| HTTP 方法 | 后端完整路径 (`/app-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/app-api/member/sign-in/config/list` | `miniapp/sheep/api/member/signin.js` |
| `GET` | `/app-api/member/sign-in/record/get-summary` | `miniapp/sheep/api/member/signin.js` |
| `POST` | `/app-api/member/sign-in/record/create` | `miniapp/sheep/api/member/signin.js` |
| `GET` | `/app-api/member/sign-in/record/page` | `miniapp/sheep/api/member/signin.js` |

#### 📱 微信社交账号绑定 (前缀: `/member/social-user` - 5 个接口)

| HTTP 方法 | 后端完整路径 (`/app-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/app-api/member/social-user/get` | `miniapp/sheep/api/member/social.js` |
| `POST` | `/app-api/member/social-user/bind` | `miniapp/sheep/api/member/social.js` |
| `DELETE` | `/app-api/member/social-user/unbind` | `miniapp/sheep/api/member/social.js` |
| `GET` | `/app-api/member/social-user/get-subscribe-template-list` | `miniapp/sheep/api/member/social.js` |
| `POST` | `/app-api/member/social-user/wxa-qrcode` | `miniapp/sheep/api/member/social.js` |

### 2.商品中心 (Product)

#### 📱 商品详情与列表 (前缀: `/product/spu` - 3 个接口)

| HTTP 方法 | 后端完整路径 (`/app-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/app-api/product/spu/list-by-ids` | `miniapp/sheep/api/product/spu.js` |
| `GET` | `/app-api/product/spu/page` | `miniapp/sheep/api/product/spu.js` |
| `GET` | `/app-api/product/spu/get-detail` | `miniapp/sheep/api/product/spu.js` |

#### 📱 商品分类 (前缀: `/product/category` - 2 个接口)

| HTTP 方法 | 后端完整路径 (`/app-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/app-api/product/category/list` | `miniapp/sheep/api/product/category.js` |
| `GET` | `/app-api/product/category/list-by-ids` | `miniapp/sheep/api/product/category.js` |

#### 📱 商品评价 (前缀: `/product/comment` - 1 个接口)

| HTTP 方法 | 后端完整路径 (`/app-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/app-api/product/comment/page` | `miniapp/sheep/api/product/comment.js` |

#### 📱 商品收藏 (前缀: `/product/favorite` - 4 个接口)

| HTTP 方法 | 后端完整路径 (`/app-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/app-api/product/favorite/page` | `miniapp/sheep/api/product/favorite.js` |
| `GET` | `/app-api/product/favorite/exits` | `miniapp/sheep/api/product/favorite.js` |
| `POST` | `/app-api/product/favorite/create` | `miniapp/sheep/api/product/favorite.js` |
| `DELETE` | `/app-api/product/favorite/delete` | `miniapp/sheep/api/product/favorite.js` |

#### 📱 足迹浏览历史 (前缀: `/product/browse-history` - 3 个接口)

| HTTP 方法 | 后端完整路径 (`/app-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `DELETE` | `/app-api/product/browse-history/delete` | `miniapp/sheep/api/product/history.js` |
| `DELETE` | `/app-api/product/browse-history/clean` | `miniapp/sheep/api/product/history.js` |
| `GET` | `/app-api/product/browse-history/page` | `miniapp/sheep/api/product/history.js` |

### 2.购物车与交易 (Trade)

#### 📱 购物车 (前缀: `/trade/cart` - 5 个接口)

| HTTP 方法 | 后端完整路径 (`/app-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `POST` | `/app-api/trade/cart/add` | `miniapp/sheep/api/trade/cart.js` |
| `PUT` | `/app-api/trade/cart/update-count` | `miniapp/sheep/api/trade/cart.js` |
| `PUT` | `/app-api/trade/cart/update-selected` | `miniapp/sheep/api/trade/cart.js` |
| `DELETE` | `/app-api/trade/cart/delete` | `miniapp/sheep/api/trade/cart.js` |
| `GET` | `/app-api/trade/cart/list` | `miniapp/sheep/api/trade/cart.js` |

#### 📱 订单创建、结算与核销 (前缀: `/trade/order` - 11 个接口)

| HTTP 方法 | 后端完整路径 (`/app-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/app-api/trade/order/settlement-product` | `miniapp/sheep/api/product/spu.js` |
| `GET` | `/app-api/trade/order/settlement` | `miniapp/sheep/api/trade/order.js` |
| `POST` | `/app-api/trade/order/create` | `miniapp/sheep/api/trade/order.js` |
| `GET` | `/app-api/trade/order/get-detail` | `miniapp/sheep/api/trade/order.js` |
| `GET` | `/app-api/trade/order/page` | `miniapp/sheep/api/trade/order.js` |
| `PUT` | `/app-api/trade/order/receive` | `miniapp/sheep/api/trade/order.js` |
| `DELETE` | `/app-api/trade/order/cancel` | `miniapp/sheep/api/trade/order.js` |
| `DELETE` | `/app-api/trade/order/delete` | `miniapp/sheep/api/trade/order.js` |
| `GET` | `/app-api/trade/order/get-express-track-list` | `miniapp/sheep/api/trade/order.js` |
| `GET` | `/app-api/trade/order/get-count` | `miniapp/sheep/api/trade/order.js` |
| `POST` | `/app-api/trade/order/item/create-comment` | `miniapp/sheep/api/trade/order.js` |

#### 📱 售后申请与进度 (前缀: `/trade/after-sale` - 6 个接口)

| HTTP 方法 | 后端完整路径 (`/app-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/app-api/trade/after-sale/page` | `miniapp/sheep/api/trade/afterSale.js` |
| `POST` | `/app-api/trade/after-sale/create` | `miniapp/sheep/api/trade/afterSale.js` |
| `GET` | `/app-api/trade/after-sale/get` | `miniapp/sheep/api/trade/afterSale.js` |
| `DELETE` | `/app-api/trade/after-sale/cancel` | `miniapp/sheep/api/trade/afterSale.js` |
| `GET` | `/app-api/trade/after-sale-log/list` | `miniapp/sheep/api/trade/afterSale.js` |
| `PUT` | `/app-api/trade/after-sale/delivery` | `miniapp/sheep/api/trade/afterSale.js` |

#### 📱 运费计算与自提点 (前缀: `/trade/delivery` - 3 个接口)

| HTTP 方法 | 后端完整路径 (`/app-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/app-api/trade/delivery/express/list` | `miniapp/sheep/api/trade/delivery.js` |
| `GET` | `/app-api/trade/delivery/pick-up-store/list` | `miniapp/sheep/api/trade/delivery.js` |
| `GET` | `/app-api/trade/delivery/pick-up-store/get` | `miniapp/sheep/api/trade/delivery.js` |

#### 📱 交易配置 (前缀: `/trade/config` - 1 个接口)

| HTTP 方法 | 后端完整路径 (`/app-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/app-api/trade/config/get` | `miniapp/sheep/api/trade/config.js` |

#### 📱 分销员中心 (前缀: `/trade/brokerage-user` - 7 个接口)

| HTTP 方法 | 后端完整路径 (`/app-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `PUT` | `/app-api/trade/brokerage-user/bind` | `miniapp/sheep/api/trade/brokerage.js` |
| `GET` | `/app-api/trade/brokerage-user/get` | `miniapp/sheep/api/trade/brokerage.js` |
| `GET` | `/app-api/trade/brokerage-user/get-summary` | `miniapp/sheep/api/trade/brokerage.js` |
| `GET` | `/app-api/trade/brokerage-user/get-rank-by-price` | `miniapp/sheep/api/trade/brokerage.js` |
| `GET` | `/app-api/trade/brokerage-user/rank-page-by-price` | `miniapp/sheep/api/trade/brokerage.js` |
| `GET` | `/app-api/trade/brokerage-user/rank-page-by-user-count` | `miniapp/sheep/api/trade/brokerage.js` |
| `GET` | `/app-api/trade/brokerage-user/child-summary-page` | `miniapp/sheep/api/trade/brokerage.js` |

#### 📱 分销返佣记录 (前缀: `/trade/brokerage-record` - 2 个接口)

| HTTP 方法 | 后端完整路径 (`/app-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/app-api/trade/brokerage-record/page` | `miniapp/sheep/api/trade/brokerage.js` |
| `GET` | `/app-api/trade/brokerage-record/get-product-brokerage-price` | `miniapp/sheep/api/trade/brokerage.js` |

#### 📱 佣金申请提现 (前缀: `/trade/brokerage-withdraw` - 3 个接口)

| HTTP 方法 | 后端完整路径 (`/app-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `POST` | `/app-api/trade/brokerage-withdraw/create` | `miniapp/sheep/api/trade/brokerage.js` |
| `GET` | `/app-api/trade/brokerage-withdraw/page` | `miniapp/sheep/api/trade/brokerage.js` |
| `GET` | `/app-api/trade/brokerage-withdraw/get` | `miniapp/sheep/api/trade/brokerage.js` |

### 2.营销活动 (Promotion)

#### 📱 领券中心 (前缀: `/promotion/coupon-template` - 4 个接口)

| HTTP 方法 | 后端完整路径 (`/app-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/app-api/promotion/coupon-template/list-by-ids` | `miniapp/sheep/api/promotion/coupon.js` |
| `GET` | `/app-api/promotion/coupon-template/list` | `miniapp/sheep/api/promotion/coupon.js` |
| `GET` | `/app-api/promotion/coupon-template/page` | `miniapp/sheep/api/promotion/coupon.js` |
| `GET` | `/app-api/promotion/coupon-template/get` | `miniapp/sheep/api/promotion/coupon.js` |

#### 📱 我的优惠券 (前缀: `/promotion/coupon` - 8 个接口)

| HTTP 方法 | 后端完整路径 (`/app-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/app-api/promotion/coupon-template/list-by-ids` | `miniapp/sheep/api/promotion/coupon.js` |
| `GET` | `/app-api/promotion/coupon-template/list` | `miniapp/sheep/api/promotion/coupon.js` |
| `GET` | `/app-api/promotion/coupon-template/page` | `miniapp/sheep/api/promotion/coupon.js` |
| `GET` | `/app-api/promotion/coupon-template/get` | `miniapp/sheep/api/promotion/coupon.js` |
| `GET` | `/app-api/promotion/coupon/page` | `miniapp/sheep/api/promotion/coupon.js` |
| `POST` | `/app-api/promotion/coupon/take` | `miniapp/sheep/api/promotion/coupon.js` |
| `GET` | `/app-api/promotion/coupon/get` | `miniapp/sheep/api/promotion/coupon.js` |
| `GET` | `/app-api/promotion/coupon/get-unused-count` | `miniapp/sheep/api/promotion/coupon.js` |

#### 📱 满减送 (前缀: `/promotion/reward-activity` - 1 个接口)

| HTTP 方法 | 后端完整路径 (`/app-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/app-api/promotion/reward-activity/get` | `miniapp/sheep/api/promotion/rewardActivity.js` |

#### 📱 秒杀活动与商品 (前缀: `/promotion/seckill-activity` - 1 个接口)

| HTTP 方法 | 后端完整路径 (`/app-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/app-api/promotion/seckill-activity/list-by-ids` | `miniapp/sheep/api/promotion/seckill.js` |

#### 📱 拼团商品 (前缀: `/promotion/combination-activity` - 3 个接口)

| HTTP 方法 | 后端完整路径 (`/app-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/app-api/promotion/combination-activity/page` | `miniapp/sheep/api/promotion/combination.js` |
| `GET` | `/app-api/promotion/combination-activity/get-detail` | `miniapp/sheep/api/promotion/combination.js` |
| `GET` | `/app-api/promotion/combination-activity/list-by-ids` | `miniapp/sheep/api/promotion/combination.js` |

#### 📱 拼团订单与记录 (前缀: `/promotion/combination-record` - 4 个接口)

| HTTP 方法 | 后端完整路径 (`/app-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/app-api/promotion/combination-record/get-head-list` | `miniapp/sheep/api/promotion/combination.js` |
| `GET` | `/app-api/promotion/combination-record/page` | `miniapp/sheep/api/promotion/combination.js` |
| `GET` | `/app-api/promotion/combination-record/get-detail` | `miniapp/sheep/api/promotion/combination.js` |
| `GET` | `/app-api/promotion/combination-record/get-summary` | `miniapp/sheep/api/promotion/combination.js` |

#### 📱 积分换购商品 (前缀: `/promotion/point-activity` - 1 个接口)

| HTTP 方法 | 后端完整路径 (`/app-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/app-api/promotion/point-activity/list-by-ids` | `miniapp/sheep/api/promotion/point.js` |

#### 📱 帮助与商城资讯 (前缀: `/promotion/article` - 1 个接口)

| HTTP 方法 | 后端完整路径 (`/app-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/app-api/promotion/article/get` | `miniapp/sheep/api/promotion/article.js` |

#### 📱 首页与个人中心DIY装修 (前缀: `/promotion/diy-template` - 2 个接口)

| HTTP 方法 | 后端完整路径 (`/app-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/app-api/promotion/diy-template/used` | `miniapp/sheep/api/promotion/diy.js` |
| `GET` | `/app-api/promotion/diy-template/get` | `miniapp/sheep/api/promotion/diy.js` |

#### 📱 微页面DIY装修 (前缀: `/promotion/diy-page` - 1 个接口)

| HTTP 方法 | 后端完整路径 (`/app-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/app-api/promotion/diy-page/get` | `miniapp/sheep/api/promotion/diy.js` |

#### 📱 在线客服聊天 (前缀: `/promotion/kefu-message` - 2 个接口)

| HTTP 方法 | 后端完整路径 (`/app-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `POST` | `/app-api/promotion/kefu-message/send` | `miniapp/sheep/api/promotion/kefu.js` |
| `GET` | `/app-api/promotion/kefu-message/list` | `miniapp/sheep/api/promotion/kefu.js` |

### 2.支付与钱包 (Pay)

#### 📱 微信/支付宝预支付下单 (前缀: `/pay/order` - 2 个接口)

| HTTP 方法 | 后端完整路径 (`/app-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/app-api/pay/order/get` | `miniapp/sheep/api/pay/order.js` |
| `POST` | `/app-api/pay/order/submit` | `miniapp/sheep/api/pay/order.js` |

#### 📱 支付渠道 (前缀: `/pay/channel` - 1 个接口)

| HTTP 方法 | 后端完整路径 (`/app-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/app-api/pay/channel/get-enable-code-list` | `miniapp/sheep/api/pay/channel.js` |

#### 📱 钱包余额 (前缀: `/pay/wallet` - 6 个接口)

| HTTP 方法 | 后端完整路径 (`/app-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/app-api/pay/wallet/get` | `miniapp/sheep/api/pay/wallet.js` |
| `GET` | `/app-api/pay/wallet-transaction/page` | `miniapp/sheep/api/pay/wallet.js` |
| `GET` | `/app-api/pay/wallet-transaction/get-summary` | `miniapp/sheep/api/pay/wallet.js` |
| `GET` | `/app-api/pay/wallet-recharge-package/list` | `miniapp/sheep/api/pay/wallet.js` |
| `POST` | `/app-api/pay/wallet-recharge/create` | `miniapp/sheep/api/pay/wallet.js` |
| `GET` | `/app-api/pay/wallet-recharge/page` | `miniapp/sheep/api/pay/wallet.js` |

#### 📱 充值套餐 (前缀: `/pay/wallet-recharge-package` - 1 个接口)

| HTTP 方法 | 后端完整路径 (`/app-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/app-api/pay/wallet-recharge-package/list` | `miniapp/sheep/api/pay/wallet.js` |

#### 📱 充值订单创建与查询 (前缀: `/pay/wallet-recharge` - 3 个接口)

| HTTP 方法 | 后端完整路径 (`/app-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/app-api/pay/wallet-recharge-package/list` | `miniapp/sheep/api/pay/wallet.js` |
| `POST` | `/app-api/pay/wallet-recharge/create` | `miniapp/sheep/api/pay/wallet.js` |
| `GET` | `/app-api/pay/wallet-recharge/page` | `miniapp/sheep/api/pay/wallet.js` |

#### 📱 钱包收支流水 (前缀: `/pay/wallet-transaction` - 2 个接口)

| HTTP 方法 | 后端完整路径 (`/app-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/app-api/pay/wallet-transaction/page` | `miniapp/sheep/api/pay/wallet.js` |
| `GET` | `/app-api/pay/wallet-transaction/get-summary` | `miniapp/sheep/api/pay/wallet.js` |

#### 📱 转账打款 (前缀: `/pay/transfer` - 1 个接口)

| HTTP 方法 | 后端完整路径 (`/app-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/app-api/pay/transfer/sync` | `miniapp/sheep/api/pay/transfer.js` |

### 2.系统与基础 (System & Infra)

#### 📱 图片文件上传 (前缀: `/infra/file` - 2 个接口)

| HTTP 方法 | 后端完整路径 (`/app-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/app-api/infra/file/presigned-url` | `miniapp/sheep/api/infra/file.js` |
| `POST` | `/app-api/infra/file/create` | `miniapp/sheep/api/infra/file.js` |

#### 📱 省市区地址联动 (前缀: `/system/area` - 1 个接口)

| HTTP 方法 | 后端完整路径 (`/app-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/app-api/system/area/tree` | `miniapp/sheep/api/system/area.js` |

#### 📱 全局字典项 (前缀: `/system/dict-data` - 1 个接口)

| HTTP 方法 | 后端完整路径 (`/app-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/app-api/system/dict-data/type` | `miniapp/sheep/api/system/dict.js` |

#### 📱 租户信息 (前缀: `/system/tenant` - 1 个接口)

| HTTP 方法 | 后端完整路径 (`/app-api` + 路径) | 前端源码调用位置 |
|---|---|---|
| `GET` | `/app-api/system/tenant/get-by-website` | `miniapp/sheep/api/infra/tenant.js` |

---

## 叁、 核心功能能力专项清单 (WebSocket、微信生态、OSS上传、验证码与风控、Excel导出、地理位置)

> 本节专门梳理了前端在标准 REST 数据 CRUD 之外调用的底层能力接口，确保新后端迁移时各类扩展功能不遗漏、不降级。

### 3.1 WebSocket 实时客服长连接通讯

| 功能说明 | 协议 / 路径 | 前端调用源码位置 | 消息数据协议 |
|---|---|---|---|
| 小程序在线客服实时通讯 | `wss://${HOST}/infra/ws` 或 `SHOPRO_WEBSOCKET_PATH` | `apps/miniapp/pages/chat/util/useWebSocket.js` | JSON `{ type: "message", data: { content, contentType } }` |
| 管理后台客服工作台实时长连接 | `wss://${HOST}/infra/ws` | `apps/admin/src/store/modules/mall/kefu.ts` | 双向心跳 ping/pong，实时派发新会话与文本/图片/商品卡片消息 |

### 3.2 文件对象存储与多媒体上传能力 (OSS / File Upload)

| 功能说明 | HTTP 方法 | 后端完整路径 | 前端调用源码位置 |
|---|---|---|---|
| 管理后台商品图、SKU图、活动海报上传 | `POST (multipart/form-data)` | `/admin-api/infra/file/upload` | `apps/admin/src/components/UploadFile/src/useUpload.ts` |
| 小程序买家评价晒图、售后退款凭据、用户头像上传 | `POST (multipart/form-data)` | `/app-api/infra/file/upload` | `apps/miniapp/sheep/components/s-uploader/choose-and-upload-file.js` |
| 管理后台文件列表分页与存储记录查询 | `GET` | `/admin-api/infra/file/page` | `apps/admin/src/api/infra/file/index.ts` |
| 管理后台文件物理/逻辑删除 | `DELETE` | `/admin-api/infra/file/delete` | `apps/admin/src/api/infra/file/index.ts` |

### 3.3 微信生态专有能力 (WeChat MiniApp Capabilities)

| 功能说明 | HTTP 方法 | 后端完整路径 | 前端调用源码位置 |
|---|---|---|---|
| 微信小程序一键授权登录 (code2Session) | `POST` | `/app-api/member/auth/weixin-mini-app-login` | `apps/miniapp/sheep/api/member/auth.js` |
| 微信手机号授权一键绑定 | `POST` | `/app-api/member/auth/create-weixin-mini-app-mobile` | `apps/miniapp/sheep/api/member/auth.js` |
| 微信社交账号绑定 (OpenID / UnionID) | `POST` | `/app-api/member/social-user/bind-weixin-mini-app` | `apps/miniapp/sheep/api/member/social.js` |
| 微信支付预下单拉起收银台 (`wx_lite`) | `POST` | `/app-api/pay/order/submit` | `apps/miniapp/sheep/platform/pay.js` |
| 小程序到店自提二维码核销能力 | `GET` | `/app-api/trade/order/pick-up-verify` | `apps/miniapp/pages/order/pickUpVerify.vue` |
| 微信官方小程序直播间列表同步 | `GET` | `/app-api/app/mplive/get-room-list` | `apps/miniapp/sheep/api/migration/app.js` |
| 微信官方小程序直播回放视频列表 | `GET` | `/app-api/app/mplive/get-replay-list` | `apps/miniapp/sheep/api/migration/app.js` |

### 3.4 安全验证与风控验证码能力 (Captcha & SMS)

| 功能说明 | HTTP 方法 | 后端完整路径 | 前端调用源码位置 |
|---|---|---|---|
| 滑块/点选验证码图片与校验Token获取 | `POST` | `/admin-api/system/captcha/get` | `apps/admin/src/api/login/index.ts` |
| 验证码滑动轨迹与坐标二次核验 | `POST` | `/admin-api/system/captcha/check` | `apps/admin/src/api/login/index.ts` |
| 管理员登录短信验证码发送 | `POST` | `/admin-api/system/auth/send-sms-code` | `apps/admin/src/api/login/index.ts` |
| 管理员短信验证码登录 | `POST` | `/admin-api/system/auth/sms-login` | `apps/admin/src/api/login/index.ts` |
| 买家端短信验证码发送 | `POST` | `/app-api/member/auth/send-sms-code` | `apps/miniapp/sheep/api/member/auth.js` |
| 买家短信无密码快捷登录 | `POST` | `/app-api/member/auth/sms-login` | `apps/miniapp/sheep/api/member/auth.js` |
| 短信验证重置买家登录密码 | `POST` | `/app-api/member/auth/reset-password` | `apps/miniapp/sheep/api/member/auth.js` |

### 3.5 Excel 报表二进制下载导出能力 (Excel Export)

| 报表名称 | HTTP 方法 | 后端完整路径 | 前端响应处理 |
|---|---|---|---|
| 商品 SPU 基础信息与销售导出 | `GET` | `/admin-api/product/spu/export-excel` | 二进制流 `responseType: blob` 自动触发浏览器文件下载 |
| 交易中心订单与买家明细导出 | `GET` | `/admin-api/trade/order/export-excel` | 二进制流 `responseType: blob` |
| 售后退款单与原因导出 | `GET` | `/admin-api/trade/after-sale/export-excel` | 二进制流 `responseType: blob` |
| 会员用户与积分成长值导出 | `GET` | `/admin-api/member/user/export-excel` | 二进制流 `responseType: blob` |
| 支付交易流水订单导出 | `GET` | `/admin-api/pay/order/export-excel` | 二进制流 `responseType: blob` |
| 支付退款流水记录导出 | `GET` | `/admin-api/pay/refund/export-excel` | 二进制流 `responseType: blob` |
| 分销佣金提现申请明细导出 | `GET` | `/admin-api/trade/brokerage-withdraw/export-excel` | 二进制流 `responseType: blob` |

### 3.6 地理位置联动与 IP 逆地址解析 (Location)

| 功能说明 | HTTP 方法 | 后端完整路径 | 用途 |
|---|---|---|---|
| 全国省市区完整三级联动树 | `GET` | `/admin-api/system/area/tree` & `/app-api/system/area/tree` | 用于运费计费区域选择、门店地址、买家收货地址选择器 |
| 客户端访问 IP 逆解析城市名 | `GET` | `/admin-api/system/area/get-by-ip` | 登录日志记录、自动定位买家所在城市 |

### 3.7 苹果生态账号验证能力 (Apple Sign-In)

| 功能说明 | HTTP 方法 | 后端完整路径 | 前端源码调用位置 |
|---|---|---|---|
| 苹果身份令牌 identityToken 验签与绑定 | `POST` | `/app-api/third/apple/verify` | `apps/miniapp/sheep/api/migration/third.js` |

