<script setup lang="ts">
import { useAppStore } from '@/store/modules/app'
const appStore = useAppStore()
</script>
<template
  ><ConfigGlobal :size="appStore.getCurrentSize"><RouterView /></ConfigGlobal
></template>
<style>
html,
body,
#app {
  margin: 0;
  width: 100%;
  min-height: 100%;
  font-family: Inter, 'PingFang SC', 'Microsoft YaHei', sans-serif;
}
body {
  background: #f5f7fa;
}
</style>
