import { demoTemplate } from '../../../miniapp/sheep/demo/template.js'
import { ElMessage } from 'element-plus'
import {
  categories,
  brands,
  products,
  members,
  orders,
  couponTemplates,
  tradeConfig
} from '../../../../packages/demo-data/index.js'
const storageKey = 'mall-demo-admin-v2'
const seed: Record<string, any> = {
  '/product/spu': products,
  '/product/category': categories,
  '/product/brand': brands,
  '/product/property': [
    { id: 1, name: '规格', remark: '标准装', createTime: '2026-10-14 09:00:00' }
  ],
  '/product/property/value': [{ id: 1, propertyId: 1, name: '标准装', remark: '' }],
  '/member/user': members,
  '/member/level': [
    { id: 1, name: '普通会员', level: 1, experience: 0, discountPercent: 100, status: 0 }
  ],
  '/member/group': [{ id: 1, name: '新会员', remark: '新注册会员' }],
  '/member/tag': [{ id: 1, name: '宠物家庭' }],
  '/trade/order': orders,
  '/promotion/coupon-template': couponTemplates,
  '/trade/delivery/express': [{ id: 1, name: '顺丰速运', code: 'SF', sort: 1, status: 0 }],
  '/trade/delivery/express-template': [
    {
      id: 1,
      name: '标准快递',
      chargeMode: 1,
      sort: 1,
      templateCharge: [
        { areaIds: [310100], startCount: 1, startPrice: 600, extraCount: 1, extraPrice: 200 }
      ],
      templateFree: []
    }
  ],
  '/trade/config': tradeConfig,
  '/member/config': {
    id: 1,
    pointTradeDeductEnable: true,
    pointTradeDeductUnitPrice: 1,
    pointTradeDeductMaxPrice: 1000,
    pointTradeGivePoint: 1
  },
  '/promotion/diy-template': [
    {
      id: 1,
      name: '我的宠物商城',
      used: true,
      remark: '默认演示装修',
      previewPicUrls: [],
      createTime: '2026-10-14 09:00:00',
      property: '{}'
    }
  ]
}
const clone = <T>(value: T): T => JSON.parse(JSON.stringify(value))
let db = clone(seed)
try {
  const saved = localStorage.getItem(storageKey)
  if (saved) db = { ...db, ...JSON.parse(saved) }
} catch {
  /* 存储不可用时仍可浏览。 */
}
const persist = () => {
  try {
    localStorage.setItem(storageKey, JSON.stringify(db))
  } catch {
    ElMessage.warning('浏览器存储不可用，本次修改仅在当前页面保留')
  }
}
const dictionary: Record<string, Array<[number | string, string, string?]>> = {
  common_status: [
    [0, '启用', 'success'],
    [1, '停用', 'info']
  ],
  product_spu_status: [
    [1, '上架', 'success'],
    [0, '下架', 'info'],
    [-1, '回收站', 'danger']
  ],
  product_unit: [
    [1, '件'],
    [2, '袋'],
    [3, '盒']
  ],
  trade_order_status: [
    [0, '待付款', 'warning'],
    [10, '待发货', 'primary'],
    [20, '待收货', 'primary'],
    [30, '已完成', 'success'],
    [40, '已取消', 'info']
  ],
  trade_order_type: [
    [0, '普通订单'],
    [1, '秒杀'],
    [2, '拼团'],
    [3, '砍价']
  ],
  trade_order_terminal: [
    [10, '微信小程序'],
    [20, 'H5']
  ],
  terminal: [
    [10, '微信小程序'],
    [20, 'H5']
  ],
  trade_delivery_type: [
    [1, '快递'],
    [2, '自提']
  ],
  trade_order_after_sale_status: [
    [0, '无售后'],
    [10, '售后中', 'warning'],
    [20, '售后完成', 'success']
  ],
  trade_after_sale_status: [
    [10, '待审核', 'warning'],
    [20, '已同意', 'primary'],
    [30, '待退货'],
    [40, '已收货'],
    [50, '已退款', 'success'],
    [60, '已拒绝', 'danger']
  ],
  system_user_sex: [
    [1, '男'],
    [2, '女'],
    [0, '未知']
  ],
  user_type: [
    [1, '会员'],
    [2, '管理员']
  ],
  pay_channel_code: [
    ['wx_lite', '微信小程序'],
    ['alipay_wap', '支付宝']
  ],
  pay_order_status: [
    [0, '待支付'],
    [10, '已支付', 'success'],
    [20, '已退款'],
    [30, '已关闭']
  ],
  promotion_discount_type: [
    [1, '满减'],
    [2, '折扣']
  ],
  promotion_coupon_status: [
    [1, '未使用'],
    [2, '已使用'],
    [3, '已过期']
  ],
  promotion_coupon_product_scope: [
    [1, '全部商品'],
    [2, '指定商品'],
    [3, '指定分类']
  ],
  promotion_coupon_take_type: [
    [1, '直接领取'],
    [2, '指定发放']
  ],
  promotion_coupon_template_validity_type: [
    [1, '固定日期'],
    [2, '领取后有效']
  ],
  promotion_activity_status: [
    [0, '未开始'],
    [1, '进行中'],
    [2, '已结束']
  ],
  brokerage_enabled_condition: [
    [1, '全部会员'],
    [2, '指定会员']
  ],
  brokerage_bind_mode: [
    [1, '首次绑定'],
    [2, '每次绑定']
  ],
  brokerage_withdraw_type: [
    [1, '钱包余额'],
    [2, '银行卡']
  ]
}
const dictionaries = Object.entries(dictionary).flatMap(([dictType, rows]) =>
  rows.map(([value, label, colorType = '']) => ({
    dictType,
    value: String(value),
    label,
    colorType,
    cssClass: ''
  }))
)
const comparison = (value: any, reference: any) => ({ value, reference })
function filter(rows: any[], params: any) {
  return rows.filter((row) =>
    Object.entries(params).every(([key, value]) => {
      if (
        value === undefined ||
        value === null ||
        value === '' ||
        key === 'pageNo' ||
        key === 'pageSize' ||
        key === 'createTime' ||
        key === 'loginDate'
      )
        return true
      if (key === 'tabType')
        return Number(value) === 0
          ? row.status === 1
          : Number(value) === 1
          ? row.status === 0
          : Number(value) === 2
          ? row.stock === 0
          : Number(value) === 3
          ? row.stock < 10
          : Number(value) === 4
          ? row.status === -1
          : true
      if (!(key in row)) return true
      if (Array.isArray(value)) return value.length === 0 || value.includes(row[key])
      return typeof row[key] === 'string'
        ? row[key].includes(String(value))
        : row[key] === Number(value)
    })
  )
}
export async function demoRequest(option: any): Promise<any> {
  const parsed = new URL(option.url, 'https://demo.invalid')
  const url = parsed.pathname
  const params = { ...Object.fromEntries(parsed.searchParams), ...option.params }
  const method = String(option.method || 'GET').toUpperCase()
  const data = typeof option.data === 'string' ? JSON.parse(option.data) : option.data || {}
  if (url === '/infra/file/upload') {
    const file = data.file
    if (!(file instanceof File)) throw new Error('请选择文件')
    const value = await new Promise((resolve, reject) => {
      const reader = new FileReader()
      reader.onload = () => resolve(reader.result)
      reader.onerror = reject
      reader.readAsDataURL(file)
    })
    return { data: value }
  }
  const known = Object.keys(db)
    .sort((a, b) => b.length - a.length)
    .find((key) => url.startsWith(key + '/'))
  const action = url.split('/').pop() || ''
  if (url === '/system/dict-data/simple-list') return { data: clone(dictionaries) }
  if (url === '/system/area/tree')
    return {
      data: [
        {
          id: 310000,
          name: '上海市',
          children: [{ id: 310100, name: '上海市', children: [{ id: 310104, name: '徐汇区' }] }]
        }
      ]
    }
  if (url === '/system/area/get-by-ip') return { data: '上海市' }
  if (url === '/product/spu/get-count') {
    const rows = filter(db['/product/spu'], { ...params, tabType: undefined })
    return {
      data: {
        '0': rows.filter((p) => p.status === 1).length,
        '1': rows.filter((p) => p.status === 0).length,
        '2': rows.filter((p) => p.stock === 0).length,
        '3': rows.filter((p) => p.stock < 10).length,
        '4': rows.filter((p) => p.status === -1).length
      }
    }
  }
  if (url === '/trade/order/summary')
    return {
      data: {
        orderCount: 5,
        orderPayPrice: db['/trade/order'].reduce((n, o) => n + o.payPrice, 0),
        afterSaleCount: 0,
        afterSalePrice: 0
      }
    }
  if (url === '/trade/order/get-express-track-list')
    return { data: [{ time: '2026-10-14 10:00:00', content: '演示物流：快件已揽收' }] }
  if (url === '/statistics/trade/order-comparison')
    return {
      data: comparison(
        { orderPayCount: 24, orderPayPrice: 326800 },
        { orderPayCount: 19, orderPayPrice: 279500 }
      )
    }
  if (url === '/statistics/member/user-count-comparison')
    return {
      data: comparison(
        { visitUserCount: 326, registerUserCount: 12 },
        { visitUserCount: 289, registerUserCount: 8 }
      )
    }
  if (url === '/statistics/trade/order-count')
    return { data: { undelivered: 2, pickUp: 0, afterSaleApply: 0, auditingWithdraw: 0 } }
  if (url.includes('/order-count-trend'))
    return {
      data: Array.from({ length: 7 }, (_, i) =>
        comparison(
          {
            date: `2026-10-${String(i + 8).padStart(2, '0')}`,
            orderPayCount: 12 + i * 2,
            orderPayPrice: 160000 + i * 23000
          },
          {
            date: `2026-10-${String(i + 1).padStart(2, '0')}`,
            orderPayCount: 9 + i,
            orderPayPrice: 120000 + i * 19000
          }
        )
      )
    }
  if (url.includes('/register-count-list'))
    return {
      data: Array.from({ length: 7 }, (_, i) => ({ date: `2026-10-${i + 8}`, count: 5 + i }))
    }
  if (url.includes('/terminal-statistics-list'))
    return {
      data: [
        { terminal: 10, userCount: 236 },
        { terminal: 20, userCount: 90 }
      ]
    }
  if (url.includes('/sex-statistics-list'))
    return {
      data: [
        { sex: 1, userCount: 120 },
        { sex: 2, userCount: 206 }
      ]
    }
  if (url.includes('/area-statistics-list'))
    return {
      data: [
        {
          areaId: 310000,
          areaName: '上海市',
          userCount: 326,
          orderCreateUserCount: 46,
          orderPayUserCount: 24,
          orderPayPrice: 326800
        }
      ]
    }
  if (url === '/statistics/member/analyse')
    return {
      data: {
        visitUserCount: 326,
        orderUserCount: 46,
        payUserCount: 24,
        atv: 13617,
        comparison: comparison(
          { registerUserCount: 12, visitUserCount: 326, rechargeUserCount: 2 },
          { registerUserCount: 8, visitUserCount: 289, rechargeUserCount: 1 }
        )
      }
    }
  if (url === '/statistics/member/summary')
    return {
      data: { userCount: 326, rechargeUserCount: 2, rechargePrice: 20000, expensePrice: 326800 }
    }
  if (url.includes('/statistics/pay/')) return { data: { rechargePrice: 20000 } }
  if (url.startsWith('/statistics/') && action.endsWith('page'))
    return { data: { list: [], total: 0 } }
  if (url.startsWith('/statistics/')) {
    if (action === 'list' || action.endsWith('-list')) return { data: [] }
    return {
      data: comparison(
        {
          yesterdayOrderCount: 19,
          monthOrderCount: 148,
          yesterdayPayPrice: 279500,
          monthPayPrice: 1945800,
          turnoverPrice: 326800,
          orderPayPrice: 326800,
          rechargePrice: 20000,
          expensePrice: 326800,
          walletPayPrice: 0,
          brokerageSettlementPrice: 0,
          afterSaleRefundPrice: 0
        },
        {
          yesterdayOrderCount: 16,
          monthOrderCount: 126,
          yesterdayPayPrice: 246000,
          monthPayPrice: 1720000,
          turnoverPrice: 279500,
          orderPayPrice: 279500,
          rechargePrice: 10000,
          expensePrice: 279500,
          walletPayPrice: 0,
          brokerageSettlementPrice: 0,
          afterSaleRefundPrice: 0
        }
      )
    }
  }
  if (url === '/promotion/diy-template/get-property')
    return {
      data: {
        ...clone(db['/promotion/diy-template'][0]),
        property: JSON.stringify(demoTemplate.property),
        pages: [
          { id: 1, name: '首页', property: JSON.stringify(demoTemplate.home) },
          { id: 2, name: '我的', property: JSON.stringify(demoTemplate.user) }
        ]
      }
    }
  if (url === '/system/auth/login' || url === '/system/auth/sms-login') {
    return {
      data: {
        accessToken: 'demo-token',
        refreshToken: 'demo-refresh-token',
        userId: 1,
        expiresTime: Date.now() + 86400000
      }
    }
  }
  if (url === '/system/auth/get-permission-info') {
    return {
      data: {
        user: { id: 1, nickname: '演示管理员', avatar: '', deptId: 0 },
        roles: ['demo'],
        permissions: ['*:*:*'],
        menus: []
      }
    }
  }
  if (url === '/system/auth/refresh-token') {
    return {
      data: {
        accessToken: 'demo-token',
        refreshToken: 'demo-refresh-token',
        userId: 1,
        expiresTime: Date.now() + 86400000
      }
    }
  }
  if (url === '/system/auth/logout') return { data: true }
  if (url.startsWith('/system/tenant/')) return { data: { id: 1 } }
  if (url === '/member/user/get-statistics')
    return { data: { orderCount: 1, orderPayPrice: 12900, afterSaleCount: 0, afterSalePrice: 0 } }
  if (url === '/pay/wallet/get')
    return {
      data: {
        id: 1,
        userId: Number(params.userId) || 1,
        balance: 5000,
        totalExpense: 12900,
        totalRecharge: 20000
      }
    }
  if (known) {
    const records = db[known]
    if (method === 'GET') {
      if (!Array.isArray(records)) return { data: clone(records) }
      const rows = filter(records, params)
      if (action === 'page') {
        const pageNo = Math.max(1, Number(params.pageNo) || 1),
          size = Math.max(1, Number(params.pageSize) || 10)
        return {
          data: { list: clone(rows.slice((pageNo - 1) * size, pageNo * size)), total: rows.length }
        }
      }
      if (action.includes('list')) {
        let ids = params.ids || params.spuIds
        return {
          data: clone(
            ids ? rows.filter((r) => String(ids).split(',').includes(String(r.id))) : rows
          )
        }
      }
      if (action.startsWith('get'))
        return { data: clone(records.find((r) => r.id === Number(params.id)) || null) }
    }
    if (action === 'save' && !Array.isArray(records)) {
      db[known] = clone(data)
      persist()
      return { data: true }
    }
    if (action === 'create' && Array.isArray(records)) {
      const id = Math.max(0, ...records.map((r) => r.id)) + 1
      records.push({
        ...clone(data),
        id,
        createTime: new Date().toISOString(),
        price: data.skus?.[0]?.price ?? data.price,
        stock: data.skus?.reduce((n, s) => n + Number(s.stock || 0), 0) ?? data.stock,
        status: data.status ?? 0
      })
      persist()
      return { data: id }
    }
    if (action === 'delete' && Array.isArray(records)) {
      db[known] = records.filter((r) => r.id !== Number(params.id || data.id))
      persist()
      return { data: true }
    }
    if (
      [
        'update',
        'update-status',
        'update-remark',
        'delivery',
        'update-address',
        'update-price'
      ].includes(action) &&
      Array.isArray(records)
    ) {
      const row = records.find((r) => r.id === Number(data.id))
      if (!row) throw new Error('演示记录不存在')
      if (action === 'delivery' && row.status !== 10) throw new Error('仅待发货订单可执行发货')
      Object.assign(row, clone(data))
      if (action === 'delivery') {
        row.status = 20
        row.deliveryTime = new Date().toISOString()
        row.logs.push({ content: '演示管理员发货', createTime: row.deliveryTime, userType: 2 })
      }
      persist()
      return { data: true }
    }
  }
  // 未覆盖的业务展示为空，写操作明确阻止，避免模拟成功误导联调。
  if (method === 'GET') {
    if (action === 'page') return { data: { list: [], total: 0 } }
    if (action.includes('list') || action === 'tree') return { data: [] }
    if (action.endsWith('count')) return { data: 0 }
    return { data: null }
  }
  ElMessage.warning(`此操作需要正式后端接口：${url}`)
  throw new Error(`演示模式未实现 ${method} ${url}`)
}
