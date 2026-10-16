import router from './router'
import { getAccessToken } from '@/utils/auth'
import { useDictStoreWithOut } from '@/store/modules/dict'
import { useUserStoreWithOut } from '@/store/modules/user'
import NProgress from 'nprogress'
import 'nprogress/nprogress.css'

NProgress.configure({
  showSpinner: false,
  speed: 280,
  minimum: 0.2
})

const demo = import.meta.env.VITE_DEMO_MODE !== 'false'

router.beforeEach(async (to) => {
  NProgress.start()
  if (!demo && !getAccessToken() && to.path !== '/login') {
    NProgress.done()
    return { path: '/login', query: { redirect: to.fullPath } }
  }
  if (to.path !== '/login') {
    const user = useUserStoreWithOut()
    if (!user.getIsSetUser) {
      if (demo) {
        user.$patch({
          permissions: new Set(['*:*:*']),
          roles: ['demo'],
          isSetUser: true,
          user: { id: 1, nickname: '演示管理员', avatar: '', deptId: 0 }
        })
      } else {
        try {
          await user.setUserInfoAction()
        } catch (error) {
          console.warn('获取用户信息异常:', error)
        }
      }
    }
    const dict = useDictStoreWithOut()
    if (!dict.getIsSetDict) {
      try {
        await dict.setDictMap()
      } catch (error) {
        console.warn('获取字典数据异常:', error)
      }
    }
  }
})

router.afterEach((to) => {
  document.title = `${to.meta.title || '商城'} · 我的商城`
  NProgress.done()
})

router.onError(() => {
  NProgress.done()
})
