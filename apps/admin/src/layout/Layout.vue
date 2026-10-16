<script setup lang="ts">
import { mallMenu } from '@/router/modules/menu'
import { useUserStore } from '@/store/modules/user'
import { removeToken } from '@/utils/auth'
import { deleteUserCache } from '@/hooks/web/useCache'

const route = useRoute()
const router = useRouter()
const user = useUserStore()
const collapsed = ref(false)
const demo = import.meta.env.VITE_DEMO_MODE === 'true'

const activeMenu = computed(() => String(route.meta.activeMenu || route.path))

// 默认展开所有多级菜单分组
const defaultOpeneds = computed(() =>
  mallMenu
    .map((item, i) => (item.items ? String(i) : null))
    .filter((v): v is string => v !== null)
)

const logout = () => {
  removeToken()
  deleteUserCache()
  user.resetState()
  router.push('/login')
}
</script>

<template>
  <div class="mall-layout" :class="{ collapsed }">
    <aside class="mall-sidebar">
      <RouterLink class="mall-brand" to="/mall/home">
        <span class="brand-mark">M</span>
        <div v-if="!collapsed"><strong>我的商城</strong><small>运营工作台</small></div>
      </RouterLink>
      <el-scrollbar class="nav-scroll">
        <el-menu
          router
          :default-active="activeMenu"
          :default-openeds="defaultOpeneds"
          :unique-opened="false"
          :collapse="collapsed"
          background-color="#142338"
          text-color="#b9c5d7"
          active-text-color="#ffffff"
          :collapse-transition="true"
        >
          <template v-for="(group, index) in mallMenu" :key="group.title">
            <!-- 顶层单级菜单（无二级子菜单，如：商城概览） -->
            <el-menu-item v-if="!group.items" :index="group.path">
              <Icon :icon="group.icon" class="nav-icon" />
              <template #title>
                <span>{{ group.title }}</span>
              </template>
            </el-menu-item>

            <!-- 包含二级子菜单的菜单分组 -->
            <el-sub-menu v-else :index="String(index)">
              <template #title>
                <Icon :icon="group.icon" class="nav-icon" />
                <span>{{ group.title }}</span>
              </template>
              <el-menu-item
                v-for="item in group.items"
                :key="item.path"
                :index="item.path"
              >
                {{ item.title }}
              </el-menu-item>
            </el-sub-menu>
          </template>
        </el-menu>
      </el-scrollbar>
      <div class="sidebar-footer" v-if="!collapsed">
        <span class="status-dot"></span>{{ demo ? '演示环境 · 无真实交易' : '商城运营环境' }}
      </div>
    </aside>

    <section class="mall-main">
      <header class="mall-header">
        <div class="header-left">
          <el-button text @click="collapsed = !collapsed" aria-label="切换菜单">
            <Icon icon="ep:fold" />
          </el-button>
          <el-breadcrumb separator="/">
            <el-breadcrumb-item :to="'/mall/home'">商城</el-breadcrumb-item>
            <el-breadcrumb-item>{{ route.meta.title }}</el-breadcrumb-item>
          </el-breadcrumb>
        </div>
        <div class="header-right">
          <el-tag :type="demo ? 'warning' : 'success'" effect="light">
            {{ demo ? '演示数据' : '正式接口' }}
          </el-tag>
          <el-dropdown @command="logout">
            <span class="admin-user">
              <el-avatar :size="30">管</el-avatar>
              {{ user.user.nickname || '管理员' }}
              <Icon icon="ep:arrow-down" />
            </span>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item command="logout">退出登录</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </div>
      </header>

      <div v-if="demo" class="demo-banner">
        <Icon icon="ep:info-filled" />
        当前为独立演示模式。操作只保存于当前浏览器，不会调用上游服务器或产生真实支付。
        <el-button text size="small" @click="router.push('/mall/product/spu')">开始管理商品 →</el-button>
      </div>

      <main class="mall-content">
        <RouterView v-slot="{ Component }">
          <component :is="Component" :key="route.path" />
        </RouterView>
      </main>
    </section>
  </div>
</template>

<style scoped>
.mall-layout {
  --sidebar-width: 232px;
  display: flex;
  min-height: 100vh;
  color: #22344c;
}
.mall-layout.collapsed {
  --sidebar-width: 64px;
}
.mall-sidebar {
  width: var(--sidebar-width);
  background: #142338;
  position: fixed;
  inset: 0 auto 0 0;
  z-index: 20;
  display: flex;
  flex-direction: column;
  transition: width 0.25s cubic-bezier(0.4, 0, 0.2, 1);
}
.mall-brand {
  height: 84px;
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 0 18px;
  color: #fff;
  text-decoration: none;
  white-space: nowrap;
}
.brand-mark {
  background: #4381f5;
  border-radius: 10px;
  font-size: 24px;
  font-weight: 750;
  width: 36px;
  height: 36px;
  display: grid;
  place-items: center;
  flex-shrink: 0;
}
.mall-brand strong {
  font-size: 19px;
  letter-spacing: 2px;
}
.mall-brand small {
  display: block;
  color: #8d9db6;
  margin-top: 6px;
  font-size: 11px;
  letter-spacing: 3px;
}
.nav-scroll {
  flex: 1;
  min-height: 0;
}
.el-menu {
  border: 0;
}
.nav-icon {
  margin-right: 10px;
}
.sidebar-footer {
  padding: 18px;
  font-size: 11px;
  color: #95a7c0;
  white-space: nowrap;
}
.status-dot {
  display: inline-block;
  width: 6px;
  height: 6px;
  border-radius: 50%;
  background: #52cba0;
  margin-right: 8px;
}
.mall-main {
  margin-left: var(--sidebar-width);
  width: calc(100% - var(--sidebar-width));
  min-width: 0;
  position: relative;
  transition: margin-left 0.25s cubic-bezier(0.4, 0, 0.2, 1), width 0.25s cubic-bezier(0.4, 0, 0.2, 1);
}

.mall-header {
  height: 64px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  background: #fff;
  padding: 0 24px;
  border-bottom: 1px solid #e9edf3;
  position: sticky;
  top: 0;
  z-index: 10;
}
.header-left,
.header-right,
.admin-user {
  display: flex;
  align-items: center;
  gap: 16px;
}
.header-right a {
  font-size: 13px;
  color: #4972aa;
  text-decoration: none;
}
.admin-user {
  cursor: pointer;
  font-size: 13px;
  gap: 8px;
}
.demo-banner {
  background: #eef4ff;
  border-bottom: 1px solid #dae6fa;
  padding: 9px 24px;
  font-size: 12px;
  color: #45638e;
  display: flex;
  align-items: center;
  gap: 8px;
}
.demo-banner .el-button {
  margin-left: auto;
}
.mall-content {
  padding: 22px;
  max-width: 1800px;
  margin: auto;
  min-height: calc(100vh - 120px);
}
.mall-content :deep(.el-card) {
  border-radius: 8px;
}
.el-menu :deep(.el-menu-item.is-active) {
  background: #2c4e7b;
}
.collapsed .mall-brand {
  padding: 0 14px;
}

@media (max-width: 850px) {
  .mall-layout {
    --sidebar-width: 64px;
  }
  .mall-brand div,
  .sidebar-footer {
    display: none;
  }
  .mall-brand {
    padding: 0 14px;
  }
  .mall-header {
    padding: 0 10px;
  }
  .header-right a,
  .admin-user {
    display: none;
  }
  .demo-banner {
    padding: 8px 12px;
    line-height: 1.8;
  }
  .demo-banner .el-button {
    display: none;
  }
  .mall-content {
    padding: 12px;
  }
}
</style>
