<script lang="ts" setup>
import * as ProductSpuApi from '@/api/mall/product/spu'
import * as TradeStatisticsApi from '@/api/mall/statistics/trade'
import * as PayStatisticsApi from '@/api/mall/statistics/pay'
import { CardTitle } from '@/components/Card'
import { fenToYuan } from '@/utils'

/** 运营数据卡片 */
defineOptions({ name: 'OperationDataCard' })

const router = useRouter() // 路由

type OperationDataItem = {
  name: string
  value: number
  routerName: string
  path: string
  query?: Record<string, any>
  prefix?: string
  decimals?: number
}

/** 数据 */
const data = reactive<Record<string, OperationDataItem>>({
  orderUndelivered: {
    name: '待发货订单',
    value: 0,
    routerName: 'TradeOrder',
    path: '/mall/trade/order',
    query: { tabType: '1' }
  },
  orderAfterSaleApply: {
    name: '退款中订单',
    value: 0,
    routerName: 'TradeAfterSale',
    path: '/mall/trade/after-sale'
  },
  orderWaitePickUp: {
    name: '待核销订单',
    value: 0,
    routerName: 'DeliveryPickUpOrder',
    path: '/mall/trade/delivery/pick-up-order'
  },
  productAlertStock: {
    name: '库存预警',
    value: 0,
    routerName: 'ProductSpu',
    path: '/mall/product/spu',
    query: { tabType: '3' }
  },
  productForSale: {
    name: '上架商品',
    value: 0,
    routerName: 'ProductSpu',
    path: '/mall/product/spu',
    query: { tabType: '0' }
  },
  productInWarehouse: {
    name: '仓库商品',
    value: 0,
    routerName: 'ProductSpu',
    path: '/mall/product/spu',
    query: { tabType: '1' }
  },
  withdrawAuditing: {
    name: '提现待审核',
    value: 0,
    routerName: 'TradeBrokerageWithdraw',
    path: '/mall/trade/brokerage/withdraw'
  },
  rechargePrice: {
    name: '账户充值',
    value: 0.0,
    prefix: '￥',
    decimals: 2,
    routerName: 'PayWalletRecharge',
    path: '/pay/wallet/rechargePackage'
  }
})

/** 查询订单数据 */
const getOrderData = async () => {
  try {
    const orderCount = await TradeStatisticsApi.getOrderCount()
    if (orderCount) {
      if (orderCount.undelivered != null) {
        data.orderUndelivered.value = orderCount.undelivered
      }
      if (orderCount.afterSaleApply != null) {
        data.orderAfterSaleApply.value = orderCount.afterSaleApply
      }
      if (orderCount.pickUp != null) {
        data.orderWaitePickUp.value = orderCount.pickUp
      }
      if (orderCount.auditingWithdraw != null) {
        data.withdrawAuditing.value = orderCount.auditingWithdraw
      }
    }
  } catch (error) {
    console.warn('获取订单统计数据失败:', error)
  }
}

/** 查询商品数据 */
const getProductData = async () => {
  try {
    const productCount = await ProductSpuApi.getTabsCount()
    if (productCount) {
      data.productForSale.value = productCount['0'] || 0
      data.productInWarehouse.value = productCount['1'] || 0
      data.productAlertStock.value = productCount['3'] || 0
    }
  } catch (error) {
    console.warn('获取商品统计数据失败:', error)
  }
}

/** 查询钱包充值数据 */
const getWalletRechargeData = async () => {
  try {
    const paySummary = await PayStatisticsApi.getWalletRechargePrice()
    if (paySummary) {
      data.rechargePrice.value = Number(fenToYuan(paySummary.rechargePrice || 0))
    }
  } catch (error) {
    console.warn('获取钱包充值数据失败:', error)
  }
}

/**
 * 跳转到对应页面
 */
const handleClick = (item: OperationDataItem) => {
  if (item.path) {
    router.push({ path: item.path, query: item.query })
  } else if (item.routerName) {
    router.push({ name: item.routerName, query: item.query })
  }
}

/** 激活时 */
onActivated(() => {
  getOrderData()
  getProductData()
  getWalletRechargeData()
})

/** 初始化 **/
onMounted(() => {
  getOrderData()
  getProductData()
  getWalletRechargeData()
})
</script>

<template>
  <el-card shadow="never">
    <template #header>
      <CardTitle title="运营数据" />
    </template>
    <div class="flex flex-row flex-wrap items-center gap-8 p-4">
      <div
        v-for="item in data"
        :key="item.name"
        class="h-20 w-20% flex flex-col cursor-pointer items-center justify-center gap-2 hover:bg-gray-50 rounded transition-colors"
        @click="handleClick(item)"
      >
        <CountTo
          :decimals="item.decimals"
          :end-val="item.value"
          :prefix="item.prefix"
          class="text-3xl"
        />
        <span class="text-center">{{ item.name }}</span>
      </div>
    </div>
  </el-card>
</template>
