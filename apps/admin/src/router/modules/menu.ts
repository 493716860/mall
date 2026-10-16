export const mallMenu = [
  {
    title: '商城概览',
    icon: 'ep:house',
    path: '/mall/home',
    name: 'MallHome'
  },
  {
    title: '商品中心',
    icon: 'ep:goods',
    items: [
      {
        path: '/mall/product/spu',
        name: 'ProductSpu',
        title: '商品管理'
      },
      {
        path: '/mall/product/category',
        name: 'ProductCategory',
        title: '商品分类'
      },
      {
        path: '/mall/product/brand',
        name: 'ProductBrand',
        title: '品牌管理'
      },
      {
        path: '/mall/product/property',
        name: 'ProductProperty',
        title: '商品规格'
      },
      {
        path: '/mall/product/comment',
        name: 'ProductComment',
        title: '商品评价'
      }
    ]
  },
  {
    title: '交易中心',
    icon: 'ep:shopping-cart',
    items: [
      {
        path: '/mall/trade/order',
        name: 'TradeOrder',
        title: '订单管理'
      },
      {
        path: '/mall/trade/after-sale',
        name: 'TradeAfterSale',
        title: '售后退款'
      },
      {
        path: '/mall/trade/delivery/express',
        name: 'DeliveryExpress',
        title: '物流公司'
      },
      {
        path: '/mall/trade/delivery/express-template',
        name: 'DeliveryExpressTemplate',
        title: '运费模板'
      },
      {
        path: '/mall/trade/delivery/pick-up-store',
        name: 'DeliveryPickUpStore',
        title: '自提门店'
      },
      {
        path: '/mall/trade/delivery/pick-up-order',
        name: 'DeliveryPickUpOrder',
        title: '自提核销'
      },
      {
        path: '/mall/trade/config',
        name: 'TradeConfig',
        title: '交易设置'
      }
    ]
  },
  {
    title: '营销中心',
    icon: 'ep:ticket',
    items: [
      {
        path: '/mall/promotion/coupon/template',
        name: 'PromotionCouponTemplate',
        title: '优惠券模板'
      },
      {
        path: '/mall/promotion/coupon',
        name: 'PromotionCoupon',
        title: '优惠券记录'
      },
      {
        path: '/mall/promotion/discount-activity',
        name: 'PromotionDiscountActivity',
        title: '限时折扣'
      },
      {
        path: '/mall/promotion/reward-activity',
        name: 'PromotionRewardActivity',
        title: '满减送活动'
      },
      {
        path: '/mall/promotion/seckill/activity',
        name: 'PromotionSeckillActivity',
        title: '秒杀活动'
      },
      {
        path: '/mall/promotion/seckill/config',
        name: 'PromotionSeckillConfig',
        title: '秒杀时段'
      },
      {
        path: '/mall/promotion/combination/activity',
        name: 'PromotionCombinationActivity',
        title: '拼团活动'
      },
      {
        path: '/mall/promotion/combination/record',
        name: 'PromotionCombinationRecord',
        title: '拼团记录'
      },
      {
        path: '/mall/promotion/bargain/activity',
        name: 'PromotionBargainActivity',
        title: '砍价活动'
      },
      {
        path: '/mall/promotion/bargain/record',
        name: 'PromotionBargainRecord',
        title: '砍价记录'
      },
      {
        path: '/mall/promotion/point/activity',
        name: 'PromotionPointActivity',
        title: '积分商城'
      },
      {
        path: '/mall/promotion/banner',
        name: 'PromotionBanner',
        title: '轮播广告'
      },
      {
        path: '/mall/promotion/article',
        name: 'PromotionArticle',
        title: '商城文章'
      },
      {
        path: '/mall/promotion/article/category',
        name: 'PromotionArticleCategory',
        title: '文章分类'
      },
      {
        path: '/mall/promotion/diy-template/diy-template',
        name: 'DiyTemplate',
        title: '店铺装修'
      },
      {
        path: '/mall/promotion/diy-template/diy-page',
        name: 'DiyPage',
        title: '装修页面'
      },
      {
        path: '/mall/promotion/kefu',
        name: 'PromotionKefu',
        title: '商城客服'
      }
    ]
  },
  {
    title: '会员中心',
    icon: 'ep:user',
    items: [
      {
        path: '/member/user',
        name: 'MemberUser',
        title: '会员管理'
      },
      {
        path: '/member/level',
        name: 'MemberLevel',
        title: '会员等级'
      },
      {
        path: '/member/group',
        name: 'MemberGroup',
        title: '会员分组'
      },
      {
        path: '/member/tag',
        name: 'MemberTag',
        title: '会员标签'
      },
      {
        path: '/member/point/record',
        name: 'MemberPointRecord',
        title: '积分记录'
      },
      {
        path: '/member/signin/config',
        name: 'MemberSigninConfig',
        title: '签到规则'
      },
      {
        path: '/member/signin/record',
        name: 'MemberSigninRecord',
        title: '签到记录'
      },
      {
        path: '/member/config',
        name: 'MemberConfig',
        title: '会员设置'
      }
    ]
  },
  {
    title: '分销中心',
    icon: 'ep:share',
    items: [
      {
        path: '/mall/trade/brokerage/user',
        name: 'TradeBrokerageUser',
        title: '分销用户'
      },
      {
        path: '/mall/trade/brokerage/record',
        name: 'TradeBrokerageRecord',
        title: '佣金记录'
      },
      {
        path: '/mall/trade/brokerage/withdraw',
        name: 'TradeBrokerageWithdraw',
        title: '提现审核'
      }
    ]
  },
  {
    title: '数据分析',
    icon: 'ep:data-analysis',
    items: [
      {
        path: '/mall/statistics/trade',
        name: 'StatisticsTrade',
        title: '交易分析'
      },
      {
        path: '/mall/statistics/member',
        name: 'StatisticsMember',
        title: '会员分析'
      },
      {
        path: '/mall/statistics/product',
        name: 'StatisticsProduct',
        title: '商品分析'
      }
    ]
  },
  {
    title: '支付与钱包',
    icon: 'ep:wallet',
    items: [
      {
        path: '/pay/order',
        name: 'PayOrder',
        title: '支付订单'
      },
      {
        path: '/pay/refund',
        name: 'PayRefund',
        title: '退款记录'
      }
    ]
  }
] as const
