import { Layout } from '@/utils/routerHelper'
import { mallRoutes } from './mall'

const routes: AppRouteRecordRaw[] = [
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/views/Login/Login.vue'),
    meta: { hidden: true, title: '商城登录' }
  },
  {
    path: '/',
    component: Layout,
    name: 'MallRoot',
    redirect: '/mall/home',
    children: [
      {
        path: '/redirect/:path(.*)',
        name: 'Redirect',
        component: () => import('@/views/Redirect/Redirect.vue'),
        meta: { hidden: true }
      },
      ...mallRoutes,
      {
        path: '/mall/product/spu/add',
        component: () => import('@/views/mall/product/spu/form/index.vue'),
        name: 'ProductSpuAdd',
        meta: {
          noCache: false,
          hidden: true,
          canTo: true,
          icon: 'ep:edit',
          title: '商品添加',
          activeMenu: '/mall/product/spu'
        }
      },
      {
        path: '/mall/product/spu/edit/:id(\\d+)',
        component: () => import('@/views/mall/product/spu/form/index.vue'),
        name: 'ProductSpuEdit',
        meta: {
          noCache: true,
          hidden: true,
          canTo: true,
          icon: 'ep:edit',
          title: '商品编辑',
          activeMenu: '/mall/product/spu'
        }
      },
      {
        path: '/mall/product/spu/detail/:id(\\d+)',
        component: () => import('@/views/mall/product/spu/form/index.vue'),
        name: 'ProductSpuDetail',
        meta: {
          noCache: true,
          hidden: true,
          canTo: true,
          icon: 'ep:view',
          title: '商品详情',
          activeMenu: '/mall/product/spu'
        }
      },
      {
        path: '/mall/product/property/value/:propertyId(\\d+)',
        component: () => import('@/views/mall/product/property/value/index.vue'),
        name: 'ProductPropertyValue',
        meta: {
          noCache: true,
          hidden: true,
          canTo: true,
          icon: 'ep:view',
          title: '商品属性值',
          activeMenu: '/product/property'
        }
      },
      {
        path: '/mall/trade/order/detail/:id(\\d+)',
        component: () => import('@/views/mall/trade/order/detail/index.vue'),
        name: 'TradeOrderDetail',
        meta: { title: '订单详情', icon: 'ep:view', activeMenu: '/mall/trade/order' }
      },
      {
        path: '/mall/trade/after-sale/detail/:id(\\d+)',
        component: () => import('@/views/mall/trade/afterSale/detail/index.vue'),
        name: 'TradeAfterSaleDetail',
        meta: { title: '退款详情', icon: 'ep:view', activeMenu: '/mall/trade/after-sale' }
      },
      {
        path: '/member/user/detail/:id',
        name: 'MemberUserDetail',
        meta: {
          title: '会员详情',
          noCache: true,
          hidden: true
        },
        component: () => import('@/views/member/user/detail/index.vue')
      },
      {
        path: '/pay/cashier',
        name: 'PayCashier',
        meta: {
          title: '收银台',
          noCache: true,
          hidden: true
        },
        component: () => import('@/views/pay/cashier/index.vue')
      },
      {
        path: '/diy/template/decorate/:id',
        name: 'DiyTemplateDecorate',
        meta: {
          title: '模板装修',
          noCache: false,
          hidden: true,
          activeMenu: '/mall/promotion/diy-template/diy-template'
        },
        component: () => import('@/views/mall/promotion/diy/template/decorate.vue')
      },
      {
        path: '/diy/page/decorate/:id',
        name: 'DiyPageDecorate',
        meta: {
          title: '页面装修',
          noCache: false,
          hidden: true,
          activeMenu: '/mall/promotion/diy-template/diy-page'
        },
        component: () => import('@/views/mall/promotion/diy/page/decorate.vue')
      }
    ]
  },
  {
    path: '/:pathMatch(.*)*',
    name: 'NoFound',
    component: () => import('@/views/Error/404.vue'),
    meta: { hidden: true, title: '页面不存在' }
  }
]

export default routes
