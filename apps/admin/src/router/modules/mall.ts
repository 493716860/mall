export const mallRoutes: AppRouteRecordRaw[] = [
  // 概览
  {
    path: '/mall/home',
    name: 'MallHome',
    component: () => import('@/views/mall/home/index.vue'),
    meta: { title: '商城概览', icon: 'ep:house' }
  },

  // 商品中心
  {
    path: '/mall/product/spu',
    name: 'ProductSpu',
    component: () => import('@/views/mall/product/spu/index.vue'),
    meta: { title: '商品管理', icon: 'ep:goods' }
  },
  {
    path: '/mall/product/category',
    name: 'ProductCategory',
    component: () => import('@/views/mall/product/category/index.vue'),
    meta: { title: '商品分类', icon: 'ep:goods' }
  },
  {
    path: '/mall/product/brand',
    name: 'ProductBrand',
    component: () => import('@/views/mall/product/brand/index.vue'),
    meta: { title: '品牌管理', icon: 'ep:goods' }
  },
  {
    path: '/mall/product/property',
    name: 'ProductProperty',
    component: () => import('@/views/mall/product/property/index.vue'),
    meta: { title: '商品规格', icon: 'ep:goods' }
  },
  {
    path: '/mall/product/comment',
    name: 'ProductComment',
    component: () => import('@/views/mall/product/comment/index.vue'),
    meta: { title: '商品评价', icon: 'ep:goods' }
  },

  // 交易中心
  {
    path: '/mall/trade/order',
    name: 'TradeOrder',
    component: () => import('@/views/mall/trade/order/index.vue'),
    meta: { title: '订单管理', icon: 'ep:shopping-cart' }
  },
  {
    path: '/mall/trade/after-sale',
    name: 'TradeAfterSale',
    component: () => import('@/views/mall/trade/afterSale/index.vue'),
    meta: { title: '售后退款', icon: 'ep:shopping-cart' }
  },
  {
    path: '/mall/trade/delivery/express',
    name: 'DeliveryExpress',
    component: () => import('@/views/mall/trade/delivery/express/index.vue'),
    meta: { title: '物流公司', icon: 'ep:shopping-cart' }
  },
  {
    path: '/mall/trade/delivery/express-template',
    name: 'DeliveryExpressTemplate',
    component: () => import('@/views/mall/trade/delivery/expressTemplate/index.vue'),
    meta: { title: '运费模板', icon: 'ep:shopping-cart' }
  },
  {
    path: '/mall/trade/delivery/pick-up-store',
    name: 'DeliveryPickUpStore',
    component: () => import('@/views/mall/trade/delivery/pickUpStore/index.vue'),
    meta: { title: '自提门店', icon: 'ep:shopping-cart' }
  },
  {
    path: '/mall/trade/delivery/pick-up-order',
    name: 'DeliveryPickUpOrder',
    component: () => import('@/views/mall/trade/delivery/pickUpOrder/index.vue'),
    meta: { title: '自提核销', icon: 'ep:shopping-cart' }
  },
  {
    path: '/mall/trade/config',
    name: 'TradeConfig',
    component: () => import('@/views/mall/trade/config/index.vue'),
    meta: { title: '交易设置', icon: 'ep:shopping-cart' }
  },

  // 营销中心
  {
    path: '/mall/promotion/coupon/template',
    name: 'PromotionCouponTemplate',
    component: () => import('@/views/mall/promotion/coupon/template/index.vue'),
    meta: { title: '优惠券模板', icon: 'ep:ticket' }
  },
  {
    path: '/mall/promotion/coupon',
    name: 'PromotionCoupon',
    component: () => import('@/views/mall/promotion/coupon/index.vue'),
    meta: { title: '优惠券记录', icon: 'ep:ticket' }
  },
  {
    path: '/mall/promotion/discount-activity',
    name: 'PromotionDiscountActivity',
    component: () => import('@/views/mall/promotion/discountActivity/index.vue'),
    meta: { title: '限时折扣', icon: 'ep:ticket' }
  },
  {
    path: '/mall/promotion/reward-activity',
    name: 'PromotionRewardActivity',
    component: () => import('@/views/mall/promotion/rewardActivity/index.vue'),
    meta: { title: '满减送活动', icon: 'ep:ticket' }
  },
  {
    path: '/mall/promotion/seckill/activity',
    name: 'PromotionSeckillActivity',
    component: () => import('@/views/mall/promotion/seckill/activity/index.vue'),
    meta: { title: '秒杀活动', icon: 'ep:ticket' }
  },
  {
    path: '/mall/promotion/seckill/config',
    name: 'PromotionSeckillConfig',
    component: () => import('@/views/mall/promotion/seckill/config/index.vue'),
    meta: { title: '秒杀时段', icon: 'ep:ticket' }
  },
  {
    path: '/mall/promotion/combination/activity',
    name: 'PromotionCombinationActivity',
    component: () => import('@/views/mall/promotion/combination/activity/index.vue'),
    meta: { title: '拼团活动', icon: 'ep:ticket' }
  },
  {
    path: '/mall/promotion/combination/record',
    name: 'PromotionCombinationRecord',
    component: () => import('@/views/mall/promotion/combination/record/index.vue'),
    meta: { title: '拼团记录', icon: 'ep:ticket' }
  },
  {
    path: '/mall/promotion/bargain/activity',
    name: 'PromotionBargainActivity',
    component: () => import('@/views/mall/promotion/bargain/activity/index.vue'),
    meta: { title: '砍价活动', icon: 'ep:ticket' }
  },
  {
    path: '/mall/promotion/bargain/record',
    name: 'PromotionBargainRecord',
    component: () => import('@/views/mall/promotion/bargain/record/index.vue'),
    meta: { title: '砍价记录', icon: 'ep:ticket' }
  },
  {
    path: '/mall/promotion/point/activity',
    name: 'PromotionPointActivity',
    component: () => import('@/views/mall/promotion/point/activity/index.vue'),
    meta: { title: '积分商城活动', icon: 'ep:ticket' }
  },
  {
    path: '/mall/promotion/banner',
    name: 'PromotionBanner',
    component: () => import('@/views/mall/promotion/banner/index.vue'),
    meta: { title: 'Banner 管理', icon: 'ep:ticket' }
  },
  {
    path: '/mall/promotion/article',
    name: 'PromotionArticle',
    component: () => import('@/views/mall/promotion/article/index.vue'),
    meta: { title: '文章管理', icon: 'ep:ticket' }
  },
  {
    path: '/mall/promotion/article/category',
    name: 'PromotionArticleCategory',
    component: () => import('@/views/mall/promotion/article/category/index.vue'),
    meta: { title: '文章分类', icon: 'ep:ticket' }
  },
  {
    path: '/mall/promotion/diy-template/diy-template',
    name: 'DiyTemplate',
    component: () => import('@/views/mall/promotion/diy/template/index.vue'),
    meta: { title: '装修模板', icon: 'ep:ticket' }
  },
  {
    path: '/mall/promotion/diy-template/diy-page',
    name: 'DiyPage',
    component: () => import('@/views/mall/promotion/diy/page/index.vue'),
    meta: { title: '装修页面', icon: 'ep:ticket' }
  },
  {
    path: '/mall/promotion/kefu',
    name: 'PromotionKefu',
    component: () => import('@/views/mall/promotion/kefu/index.vue'),
    meta: { title: '商城客服', icon: 'ep:ticket' }
  },

  // 会员中心
  {
    path: '/member/user',
    name: 'MemberUser',
    component: () => import('@/views/member/user/index.vue'),
    meta: { title: '会员管理', icon: 'ep:user' }
  },
  {
    path: '/member/level',
    name: 'MemberLevel',
    component: () => import('@/views/member/level/index.vue'),
    meta: { title: '会员等级', icon: 'ep:user' }
  },
  {
    path: '/member/group',
    name: 'MemberGroup',
    component: () => import('@/views/member/group/index.vue'),
    meta: { title: '会员分组', icon: 'ep:user' }
  },
  {
    path: '/member/tag',
    name: 'MemberTag',
    component: () => import('@/views/member/tag/index.vue'),
    meta: { title: '会员标签', icon: 'ep:user' }
  },
  {
    path: '/member/point/record',
    name: 'MemberPointRecord',
    component: () => import('@/views/member/point/record/index.vue'),
    meta: { title: '积分记录', icon: 'ep:user' }
  },
  {
    path: '/member/signin/config',
    name: 'MemberSigninConfig',
    component: () => import('@/views/member/signin/config/index.vue'),
    meta: { title: '签到规则', icon: 'ep:user' }
  },
  {
    path: '/member/signin/record',
    name: 'MemberSigninRecord',
    component: () => import('@/views/member/signin/record/index.vue'),
    meta: { title: '签到记录', icon: 'ep:user' }
  },
  {
    path: '/member/config',
    name: 'MemberConfig',
    component: () => import('@/views/member/config/index.vue'),
    meta: { title: '会员设置', icon: 'ep:user' }
  },

  // 分销中心
  {
    path: '/mall/trade/brokerage/user',
    name: 'TradeBrokerageUser',
    component: () => import('@/views/mall/trade/brokerage/user/index.vue'),
    meta: { title: '分销用户', icon: 'ep:share' }
  },
  {
    path: '/mall/trade/brokerage/record',
    name: 'TradeBrokerageRecord',
    component: () => import('@/views/mall/trade/brokerage/record/index.vue'),
    meta: { title: '佣金记录', icon: 'ep:share' }
  },
  {
    path: '/mall/trade/brokerage/withdraw',
    name: 'TradeBrokerageWithdraw',
    component: () => import('@/views/mall/trade/brokerage/withdraw/index.vue'),
    meta: { title: '提现审核', icon: 'ep:share' }
  },

  // 数据分析
  {
    path: '/mall/statistics/trade',
    name: 'StatisticsTrade',
    component: () => import('@/views/mall/statistics/trade/index.vue'),
    meta: { title: '交易分析', icon: 'ep:data-analysis' }
  },
  {
    path: '/mall/statistics/member',
    name: 'StatisticsMember',
    component: () => import('@/views/mall/statistics/member/index.vue'),
    meta: { title: '会员分析', icon: 'ep:data-analysis' }
  },
  {
    path: '/mall/statistics/product',
    name: 'StatisticsProduct',
    component: () => import('@/views/mall/statistics/product/index.vue'),
    meta: { title: '商品分析', icon: 'ep:data-analysis' }
  },

  // 支付与钱包
  {
    path: '/pay/order',
    name: 'PayOrder',
    component: () => import('@/views/pay/order/index.vue'),
    meta: { title: '支付订单', icon: 'ep:wallet' }
  },
  {
    path: '/pay/refund',
    name: 'PayRefund',
    component: () => import('@/views/pay/refund/index.vue'),
    meta: { title: '退款记录', icon: 'ep:wallet' }
  }
]
