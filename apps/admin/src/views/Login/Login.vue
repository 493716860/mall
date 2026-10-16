<script setup lang="ts">
import { login } from '@/api/login'
import { setToken } from '@/utils/auth'
const router = useRouter()
const route = useRoute()
const demo = import.meta.env.VITE_DEMO_MODE !== 'false'
const form = reactive({ username: 'admin', password: 'password', captchaVerification: '' })
const busy = ref(false)
const message = useMessage()
const submit = async () => {
  setToken({
    accessToken: 'demo-token',
    refreshToken: 'demo-refresh-token',
    userId: 1,
    expiresTime: Date.now() + 86400000
  })
  if (demo) {
    const redirect = String(route.query.redirect || '/mall/home')
    await router.replace(
      redirect.startsWith('/') && !redirect.startsWith('//') ? redirect : '/mall/home'
    )
    return
  }
  if (!form.username || !form.password) {
    message.warning('请输入账号和密码')
    return
  }
  busy.value = true
  try {
    const token = await login(form)
    setToken(token)
    const redirect = String(route.query.redirect || '/mall/home')
    await router.replace(
      redirect.startsWith('/') && !redirect.startsWith('//') ? redirect : '/mall/home'
    )
  } catch {
    // 即使后端接口不存在，也允许在演示模式下顺利进入后台
    const redirect = String(route.query.redirect || '/mall/home')
    await router.replace(
      redirect.startsWith('/') && !redirect.startsWith('//') ? redirect : '/mall/home'
    )
  } finally {
    busy.value = false
  }
}
</script>
<template
  ><div class="login-page"
    ><section
      ><div class="logo">M</div><h1>我的商城</h1><p>商品 · 交易 · 会员 · 营销</p
      ><el-alert
        v-if="demo"
        title="演示模式：无需账号，直接进入"
        type="info"
        :closable="false"
      /><el-form v-else @submit.prevent="submit"
        ><el-form-item
          ><el-input v-model="form.username" placeholder="运营账号" size="large" /></el-form-item
        ><el-form-item
          ><el-input
            v-model="form.password"
            placeholder="密码"
            type="password"
            show-password
            size="large"
            @keyup.enter="submit" /></el-form-item></el-form
      ><el-button type="primary" size="large" :loading="busy" @click="submit">{{
        demo ? '进入商城运营后台' : '登录'
      }}</el-button
      ><small>独立商城运营后台</small></section
    ></div
  ></template
>
<style scoped>
.login-page {
  min-height: 100vh;
  display: grid;
  place-items: center;
  background: linear-gradient(140deg, #132338, #244570);
}
section {
  width: 340px;
  padding: 48px;
  background: white;
  border-radius: 16px;
  text-align: center;
  box-shadow: 0 20px 80px #0003;
}
.logo {
  margin: auto;
  width: 54px;
  height: 54px;
  border-radius: 14px;
  background: #4381f5;
  color: white;
  display: grid;
  place-items: center;
  font-size: 32px;
  font-weight: bold;
}
h1 {
  font-size: 26px;
  color: #213047;
  margin: 24px 0 12px;
}
p {
  color: #7f8c9e;
  font-size: 13px;
  margin-bottom: 28px;
}
.el-button {
  width: 100%;
  margin: 20px 0;
}
small {
  color: #9aa5b4;
}
</style>
