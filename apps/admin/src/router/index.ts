import type { App } from 'vue'
import type { RouteRecordRaw } from 'vue-router'
import { createRouter, createWebHistory } from 'vue-router'
import remainingRouter from './modules/remaining'

// 创建路由实例
const router = createRouter({
  history: createWebHistory(import.meta.env.VITE_BASE_PATH), // createWebHashHistory URL带#，createWebHistory URL不带#
  strict: true,
  routes: remainingRouter as RouteRecordRaw[],
  scrollBehavior: () => {
    // 新开标签时、返回标签时，滚动条回到顶部，否则会保留上次标签的滚动位置。
    const scrollbarWrap = document.querySelector('.v-layout-content-scrollbar .el-scrollbar__wrap')
    if (scrollbarWrap) {
      scrollbarWrap.scrollTop = 0
    }
    return { left: 0, top: 0 }
  }
})

// 处理动态导入（如代码更新或热重载）模块加载失败时，优先 SPA 重试，避免跳出全页刷新
router.onError((error, to) => {
  if (
    error?.message?.includes('Failed to fetch dynamically imported module') ||
    error?.message?.includes('Importing a module script failed')
  ) {
    console.warn('动态模块加载中，正尝试无缝重试:', to.fullPath)
    router.replace(to.fullPath).catch(() => {
      window.location.assign(to.fullPath)
    })
  }
})

export const resetRouter = (): void => {
  const resetWhiteNameList = ['Redirect', 'RedirectRoot', 'Login', 'NoFound', 'Home']
  router.getRoutes().forEach((route) => {
    const { name } = route
    if (name && !resetWhiteNameList.includes(name as string)) {
      router.hasRoute(name) && router.removeRoute(name)
    }
  })
}

export const setupRouter = (app: App<Element>) => {
  app.use(router)
}

export default router
