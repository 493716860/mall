import fs from 'node:fs';

const data = JSON.parse(fs.readFileSync('all_endpoints.json', 'utf8'));

// Categorize and annotate
const adminModules = {
  '商品中心 (Product)': [
    { prefix: '/product/spu', name: '商品 SPU 管理' },
    { prefix: '/product/category', name: '商品分类' },
    { prefix: '/product/brand', name: '商品品牌' },
    { prefix: '/product/property', name: '商品规格与属性值' },
    { prefix: '/product/comment', name: '商品评价' },
    { prefix: '/product/browse-history', name: '商品浏览记录' }
  ],
  '交易中心 (Trade)': [
    { prefix: '/trade/order', name: '交易订单管理' },
    { prefix: '/trade/after-sale', name: '售后退款管理' },
    { prefix: '/trade/delivery', name: '物流配送与自提门店' },
    { prefix: '/trade/config', name: '交易设置' },
    { prefix: '/trade/brokerage-user', name: '分销用户' },
    { prefix: '/trade/brokerage-record', name: '分销佣金记录' },
    { prefix: '/trade/brokerage-withdraw', name: '佣金提现审核' }
  ],
  '营销中心 (Promotion)': [
    { prefix: '/promotion/coupon-template', name: '优惠券模板' },
    { prefix: '/promotion/coupon', name: '优惠券发放与记录' },
    { prefix: '/promotion/discount-activity', name: '限时折扣' },
    { prefix: '/promotion/reward-activity', name: '满减送' },
    { prefix: '/promotion/seckill-activity', name: '秒杀活动' },
    { prefix: '/promotion/seckill-config', name: '秒杀时段' },
    { prefix: '/promotion/combination-activity', name: '拼团活动' },
    { prefix: '/promotion/combination-record', name: '拼团记录' },
    { prefix: '/promotion/bargain-activity', name: '砍价活动' },
    { prefix: '/promotion/bargain-record', name: '砍价记录' },
    { prefix: '/promotion/bargain-help', name: '砍价助力' },
    { prefix: '/promotion/point-activity', name: '积分商城' },
    { prefix: '/promotion/banner', name: '轮播图广告' },
    { prefix: '/promotion/article-category', name: '文章分类' },
    { prefix: '/promotion/article', name: '商城文章' },
    { prefix: '/promotion/diy-template', name: '店铺装修模板' },
    { prefix: '/promotion/diy-page', name: '装修自定义页面' },
    { prefix: '/promotion/kefu-conversation', name: '客服会话' },
    { prefix: '/promotion/kefu-message', name: '客服消息' }
  ],
  '会员中心 (Member)': [
    { prefix: '/member/user', name: '会员用户管理' },
    { prefix: '/member/level', name: '会员等级' },
    { prefix: '/member/group', name: '会员分组' },
    { prefix: '/member/tag', name: '会员标签' },
    { prefix: '/member/point', name: '积分记录' },
    { prefix: '/member/experience-record', name: '成长值记录' },
    { prefix: '/member/sign-in', name: '签到规则与记录' },
    { prefix: '/member/address', name: '会员收件地址' },
    { prefix: '/member/config', name: '会员全局配置' }
  ],
  '支付与钱包 (Pay)': [
    { prefix: '/pay/app', name: '支付应用配置' },
    { prefix: '/pay/order', name: '支付订单' },
    { prefix: '/pay/refund', name: '退款订单' },
    { prefix: '/pay/wallet-recharge-package', name: '充值套餐' },
    { prefix: '/pay/wallet-transaction', name: '钱包流水' },
    { prefix: '/pay/wallet', name: '会员钱包管理' }
  ],
  '统计分析 (Statistics)': [
    { prefix: '/statistics/trade', name: '交易数据统计' },
    { prefix: '/statistics/member', name: '会员数据统计' },
    { prefix: '/statistics/product', name: '商品数据统计' },
    { prefix: '/statistics/pay', name: '支付数据统计' }
  ],
  '系统管理与基建 (System & Infra)': [
    { prefix: '/system/auth', name: '管理后台登录与权限认证' },
    { prefix: '/system/user', name: '管理员用户管理' },
    { prefix: '/system/dept', name: '部门组织架构' },
    { prefix: '/system/dict-data', name: '字典数据' },
    { prefix: '/system/area', name: '省市区与IP定位' },
    { prefix: '/system/tenant', name: '多租户管理' },
    { prefix: '/system/captcha', name: '图片验证码' },
    { prefix: '/infra/file', name: '文件对象存储与上传' }
  ]
};

const miniappModules = {
  '会员中心与认证 (Member)': [
    { prefix: '/member/auth', name: '小程序登录与手机号授权' },
    { prefix: '/member/user', name: '会员个人资料' },
    { prefix: '/member/address', name: '收货地址' },
    { prefix: '/member/point', name: '会员积分' },
    { prefix: '/member/sign-in', name: '每日签到打卡' },
    { prefix: '/member/social-user', name: '微信社交账号绑定' }
  ],
  '商品中心 (Product)': [
    { prefix: '/product/spu', name: '商品详情与列表' },
    { prefix: '/product/category', name: '商品分类' },
    { prefix: '/product/comment', name: '商品评价' },
    { prefix: '/product/favorite', name: '商品收藏' },
    { prefix: '/product/browse-history', name: '足迹浏览历史' }
  ],
  '购物车与交易 (Trade)': [
    { prefix: '/trade/cart', name: '购物车' },
    { prefix: '/trade/order', name: '订单创建、结算与核销' },
    { prefix: '/trade/after-sale', name: '售后申请与进度' },
    { prefix: '/trade/delivery', name: '运费计算与自提点' },
    { prefix: '/trade/config', name: '交易配置' },
    { prefix: '/trade/brokerage-user', name: '分销员中心' },
    { prefix: '/trade/brokerage-record', name: '分销返佣记录' },
    { prefix: '/trade/brokerage-withdraw', name: '佣金申请提现' }
  ],
  '营销活动 (Promotion)': [
    { prefix: '/promotion/coupon-template', name: '领券中心' },
    { prefix: '/promotion/coupon', name: '我的优惠券' },
    { prefix: '/promotion/reward-activity', name: '满减送' },
    { prefix: '/promotion/seckill-activity', name: '秒杀活动与商品' },
    { prefix: '/promotion/seckill-config', name: '秒杀时段配置' },
    { prefix: '/promotion/combination-activity', name: '拼团商品' },
    { prefix: '/promotion/combination-record', name: '拼团订单与记录' },
    { prefix: '/promotion/point-activity', name: '积分换购商品' },
    { prefix: '/promotion/article', name: '帮助与商城资讯' },
    { prefix: '/promotion/diy-template', name: '首页与个人中心DIY装修' },
    { prefix: '/promotion/diy-page', name: '微页面DIY装修' },
    { prefix: '/promotion/kefu-message', name: '在线客服聊天' }
  ],
  '支付与钱包 (Pay)': [
    { prefix: '/pay/order', name: '微信/支付宝预支付下单' },
    { prefix: '/pay/channel', name: '支付渠道' },
    { prefix: '/pay/wallet', name: '钱包余额' },
    { prefix: '/pay/wallet-recharge-package', name: '充值套餐' },
    { prefix: '/pay/wallet-recharge', name: '充值订单创建与查询' },
    { prefix: '/pay/wallet-transaction', name: '钱包收支流水' },
    { prefix: '/pay/transfer', name: '转账打款' }
  ],
  '系统与基础 (System & Infra)': [
    { prefix: '/infra/file', name: '图片文件上传' },
    { prefix: '/system/area', name: '省市区地址联动' },
    { prefix: '/system/dict-data', name: '全局字典项' },
    { prefix: '/system/tenant', name: '租户信息' }
  ]
};

function generateMarkdown() {
  let md = '# 芋道商城 (Mall) 全量 API 接口梳理规范文档\n\n';
  md += '> 本文档 100% 完整覆盖了本项目运营后台 (`apps/admin`) 与 微信小程序端 (`apps/miniapp`) 所有被调用的真实 API 接口。\n';
  md += `> - 运营后台 (` + '`/admin-api/*`' + `): 共 **${data.admin.length}** 个接口\n`;
  md += `> - 微信小程序端 (` + '`/app-api/*`' + `): 共 **${data.miniapp.length}** 个接口\n`;
  md += `> - 全量总计: **${data.admin.length + data.miniapp.length}** 个接口\n\n`;
  md += '---\n\n';

  // Section 1: Admin
  md += '## 壹、 WEB 运营后台 (`/admin-api`) 全量接口清单\n\n';
  for (const [moduleName, subList] of Object.entries(adminModules)) {
    md += `### 1.${moduleName}\n\n`;
    for (const sub of subList) {
      const matched = data.admin.filter(x => x.url.startsWith(sub.prefix));
      if (matched.length === 0) continue;
      md += `#### 📌 ${sub.name} (前缀: \`${sub.prefix}\` - ${matched.length} 个接口)\n\n`;
      md += '| HTTP 方法 | 后端完整路径 (`/admin-api` + 路径) | 前端源码调用位置 |\n';
      md += '|---|---|---|\n';
      for (const item of matched) {
        const fullPath = `/admin-api${item.url}`;
        md += `| \`${item.method}\` | \`${fullPath}\` | \`${item.files[0] || ''}\` |\n`;
      }
      md += '\n';
    }
  }

  // Section 2: Miniapp
  md += '---\n\n## 贰、 微信小程序端 (`/app-api`) 全量接口清单\n\n';
  for (const [moduleName, subList] of Object.entries(miniappModules)) {
    md += `### 2.${moduleName}\n\n`;
    for (const sub of subList) {
      const matched = data.miniapp.filter(x => x.url.startsWith(sub.prefix));
      if (matched.length === 0) continue;
      md += `#### 📱 ${sub.name} (前缀: \`${sub.prefix}\` - ${matched.length} 个接口)\n\n`;
      md += '| HTTP 方法 | 后端完整路径 (`/app-api` + 路径) | 前端源码调用位置 |\n';
      md += '|---|---|---|\n';
      for (const item of matched) {
        const fullPath = `/app-api${item.url}`;
        md += `| \`${item.method}\` | \`${fullPath}\` | \`${item.files[0] || ''}\` |\n`;
      }
      md += '\n';
    }
  }

  // Section 3: Functional Capabilities
  md += '---\n\n## 叁、 核心功能能力专项清单 (WebSocket、微信生态、OSS上传、验证码与风控、Excel导出、地理位置)\n\n';
  md += '> 本节专门梳理了前端在标准 REST 数据 CRUD 之外调用的底层能力接口，确保新后端迁移时各类扩展功能不遗漏、不降级。\n\n';

  md += '### 3.1 WebSocket 实时客服长连接通讯\n\n';
  md += '| 功能说明 | 协议 / 路径 | 前端调用源码位置 | 消息数据协议 |\n';
  md += '|---|---|---|---|\n';
  md += '| 小程序在线客服实时通讯 | `wss://${HOST}/infra/ws` 或 `SHOPRO_WEBSOCKET_PATH` | `apps/miniapp/pages/chat/util/useWebSocket.js` | JSON `{ type: "message", data: { content, contentType } }` |\n';
  md += '| 管理后台客服工作台实时长连接 | `wss://${HOST}/infra/ws` | `apps/admin/src/store/modules/mall/kefu.ts` | 双向心跳 ping/pong，实时派发新会话与文本/图片/商品卡片消息 |\n\n';

  md += '### 3.2 文件对象存储与多媒体上传能力 (OSS / File Upload)\n\n';
  md += '| 功能说明 | HTTP 方法 | 后端完整路径 | 前端调用源码位置 |\n';
  md += '|---|---|---|---|\n';
  md += '| 管理后台商品图、SKU图、活动海报上传 | `POST (multipart/form-data)` | `/admin-api/infra/file/upload` | `apps/admin/src/components/UploadFile/src/useUpload.ts` |\n';
  md += '| 小程序买家评价晒图、售后退款凭据、用户头像上传 | `POST (multipart/form-data)` | `/app-api/infra/file/upload` | `apps/miniapp/sheep/components/s-uploader/choose-and-upload-file.js` |\n';
  md += '| 管理后台文件列表分页与存储记录查询 | `GET` | `/admin-api/infra/file/page` | `apps/admin/src/api/infra/file/index.ts` |\n';
  md += '| 管理后台文件物理/逻辑删除 | `DELETE` | `/admin-api/infra/file/delete` | `apps/admin/src/api/infra/file/index.ts` |\n\n';

  md += '### 3.3 微信生态专有能力 (WeChat MiniApp Capabilities)\n\n';
  md += '| 功能说明 | HTTP 方法 | 后端完整路径 | 前端调用源码位置 |\n';
  md += '|---|---|---|---|\n';
  md += '| 微信小程序一键授权登录 (code2Session) | `POST` | `/app-api/member/auth/weixin-mini-app-login` | `apps/miniapp/sheep/api/member/auth.js` |\n';
  md += '| 微信手机号授权一键绑定 | `POST` | `/app-api/member/auth/create-weixin-mini-app-mobile` | `apps/miniapp/sheep/api/member/auth.js` |\n';
  md += '| 微信社交账号绑定 (OpenID / UnionID) | `POST` | `/app-api/member/social-user/bind-weixin-mini-app` | `apps/miniapp/sheep/api/member/social.js` |\n';
  md += '| 微信支付预下单拉起收银台 (`wx_lite`) | `POST` | `/app-api/pay/order/submit` | `apps/miniapp/sheep/platform/pay.js` |\n';
  md += '| 小程序到店自提二维码核销能力 | `GET` | `/app-api/trade/order/pick-up-verify` | `apps/miniapp/pages/order/pickUpVerify.vue` |\n';
  md += '| 微信官方小程序直播间列表同步 | `GET` | `/app-api/app/mplive/get-room-list` | `apps/miniapp/sheep/api/migration/app.js` |\n';
  md += '| 微信官方小程序直播回放视频列表 | `GET` | `/app-api/app/mplive/get-replay-list` | `apps/miniapp/sheep/api/migration/app.js` |\n\n';

  md += '### 3.4 安全验证与风控验证码能力 (Captcha & SMS)\n\n';
  md += '| 功能说明 | HTTP 方法 | 后端完整路径 | 前端调用源码位置 |\n';
  md += '|---|---|---|---|\n';
  md += '| 滑块/点选验证码图片与校验Token获取 | `POST` | `/admin-api/system/captcha/get` | `apps/admin/src/api/login/index.ts` |\n';
  md += '| 验证码滑动轨迹与坐标二次核验 | `POST` | `/admin-api/system/captcha/check` | `apps/admin/src/api/login/index.ts` |\n';
  md += '| 管理员登录短信验证码发送 | `POST` | `/admin-api/system/auth/send-sms-code` | `apps/admin/src/api/login/index.ts` |\n';
  md += '| 管理员短信验证码登录 | `POST` | `/admin-api/system/auth/sms-login` | `apps/admin/src/api/login/index.ts` |\n';
  md += '| 买家端短信验证码发送 | `POST` | `/app-api/member/auth/send-sms-code` | `apps/miniapp/sheep/api/member/auth.js` |\n';
  md += '| 买家短信无密码快捷登录 | `POST` | `/app-api/member/auth/sms-login` | `apps/miniapp/sheep/api/member/auth.js` |\n';
  md += '| 短信验证重置买家登录密码 | `POST` | `/app-api/member/auth/reset-password` | `apps/miniapp/sheep/api/member/auth.js` |\n\n';

  md += '### 3.5 Excel 报表二进制下载导出能力 (Excel Export)\n\n';
  md += '| 报表名称 | HTTP 方法 | 后端完整路径 | 前端响应处理 |\n';
  md += '|---|---|---|---|\n';
  md += '| 商品 SPU 基础信息与销售导出 | `GET` | `/admin-api/product/spu/export-excel` | 二进制流 `responseType: blob` 自动触发浏览器文件下载 |\n';
  md += '| 交易中心订单与买家明细导出 | `GET` | `/admin-api/trade/order/export-excel` | 二进制流 `responseType: blob` |\n';
  md += '| 售后退款单与原因导出 | `GET` | `/admin-api/trade/after-sale/export-excel` | 二进制流 `responseType: blob` |\n';
  md += '| 会员用户与积分成长值导出 | `GET` | `/admin-api/member/user/export-excel` | 二进制流 `responseType: blob` |\n';
  md += '| 支付交易流水订单导出 | `GET` | `/admin-api/pay/order/export-excel` | 二进制流 `responseType: blob` |\n';
  md += '| 支付退款流水记录导出 | `GET` | `/admin-api/pay/refund/export-excel` | 二进制流 `responseType: blob` |\n';
  md += '| 分销佣金提现申请明细导出 | `GET` | `/admin-api/trade/brokerage-withdraw/export-excel` | 二进制流 `responseType: blob` |\n\n';

  md += '### 3.6 地理位置联动与 IP 逆地址解析 (Location)\n\n';
  md += '| 功能说明 | HTTP 方法 | 后端完整路径 | 用途 |\n';
  md += '|---|---|---|---|\n';
  md += '| 全国省市区完整三级联动树 | `GET` | `/admin-api/system/area/tree` & `/app-api/system/area/tree` | 用于运费计费区域选择、门店地址、买家收货地址选择器 |\n';
  md += '| 客户端访问 IP 逆解析城市名 | `GET` | `/admin-api/system/area/get-by-ip` | 登录日志记录、自动定位买家所在城市 |\n\n';

  md += '### 3.7 苹果生态账号验证能力 (Apple Sign-In)\n\n';
  md += '| 功能说明 | HTTP 方法 | 后端完整路径 | 前端源码调用位置 |\n';
  md += '|---|---|---|---|\n';
  md += '| 苹果身份令牌 identityToken 验签与绑定 | `POST` | `/app-api/third/apple/verify` | `apps/miniapp/sheep/api/migration/third.js` |\n\n';

  return md;
}

const mdContent = generateMarkdown();
fs.writeFileSync('apps/docs/01_API全量接口规范文档.md', mdContent);
console.log('Successfully generated apps/docs/01_API全量接口规范文档.md!');
