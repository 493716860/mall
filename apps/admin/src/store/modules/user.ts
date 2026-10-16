import { store } from '@/store'
import { defineStore } from 'pinia'
import { getAccessToken, removeToken } from '@/utils/auth'
import { CACHE_KEY, useCache, deleteUserCache } from '@/hooks/web/useCache'
import { getInfo, loginOut } from '@/api/login'

const { wsCache } = useCache()

interface UserVO {
  id: number
  avatar: string
  nickname: string
  deptId: number
}

interface UserInfoVO {
  // USER 缓存
  permissions: Set<string>
  roles: string[]
  isSetUser: boolean
  user: UserVO
}

export const useUserStore = defineStore('admin-user', {
  state: (): UserInfoVO => ({
    permissions: new Set<string>(),
    roles: [],
    isSetUser: false,
    user: {
      id: 0,
      avatar: '',
      nickname: '',
      deptId: 0
    }
  }),
  getters: {
    getPermissions(): Set<string> {
      return this.permissions
    },
    getRoles(): string[] {
      return this.roles
    },
    getIsSetUser(): boolean {
      return this.isSetUser
    },
    getUser(): UserVO {
      return this.user
    }
  },
  actions: {
    async setUserInfoAction() {
      if (!getAccessToken()) {
        this.resetState()
        return null
      }
      let userInfo = wsCache.get(CACHE_KEY.USER)
      if (!userInfo) {
        try {
          userInfo = await getInfo()
        } catch (error) {
          console.warn('获取用户信息失败:', error)
        }
      } else {
        // 特殊：在有缓存的情况下，进行加载。但是即使加载失败，也不影响后续的操作，保证可以进入系统
        try {
          const fresh = await getInfo()
          if (fresh) userInfo = fresh
        } catch (error) {}
      }
      if (!userInfo) {
        userInfo = {
          permissions: ['*:*:*'],
          roles: ['admin'],
          user: {
            id: 1,
            avatar: '',
            nickname: '管理员',
            deptId: 0
          },
          menus: []
        }
      }
      this.permissions = new Set(userInfo.permissions || ['*:*:*'])
      this.roles = userInfo.roles || ['admin']
      this.user = userInfo.user || {
        id: 1,
        avatar: '',
        nickname: '管理员',
        deptId: 0
      }
      this.isSetUser = true
      wsCache.set(CACHE_KEY.USER, userInfo)
      if (userInfo.menus) {
        wsCache.set(CACHE_KEY.ROLE_ROUTERS, userInfo.menus)
      }
      return userInfo
    },
    async setUserAvatarAction(avatar: string) {
      const userInfo = wsCache.get(CACHE_KEY.USER)
      this.user.avatar = avatar
      if (userInfo && userInfo.user) {
        userInfo.user.avatar = avatar
        wsCache.set(CACHE_KEY.USER, userInfo)
      }
    },
    async setUserNicknameAction(nickname: string) {
      const userInfo = wsCache.get(CACHE_KEY.USER)
      this.user.nickname = nickname
      if (userInfo && userInfo.user) {
        userInfo.user.nickname = nickname
        wsCache.set(CACHE_KEY.USER, userInfo)
      }
    },
    async loginOut() {
      await loginOut()
      removeToken()
      deleteUserCache() // 删除用户缓存
      this.resetState()
    },
    resetState() {
      this.permissions = new Set<string>()
      this.roles = []
      this.isSetUser = false
      this.user = {
        id: 0,
        avatar: '',
        nickname: '',
        deptId: 0
      }
    }
  }
})

export const useUserStoreWithOut = () => {
  return useUserStore(store)
}
