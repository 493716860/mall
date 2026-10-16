<!-- 首页，支持店铺装修 -->
<template>
  <view v-if="template">
    <s-layout
      title="首页"
      navbar="custom"
      tabbar="/pages/index/index"
      :bgStyle="template.page"
      :navbarStyle="template.navigationBar"
      onShareAppMessage
    >
      <view v-if="isDemo" class="demo-hero"><view class="hero-eyebrow">MY PET STORE</view><view class="hero-title">好好陪伴<br/>每一个毛孩子</view><view class="hero-subtitle">精选宠物食品与生活好物</view><view class="hero-badge">独立演示 · 无真实交易</view></view>
      <s-block
        v-for="(item, index) in template.components"
        :key="index"
        :styles="item.property.style"
      >
        <s-block-item :type="item.id" :data="item.property" :styles="item.property.style" />
      </s-block>
    </s-layout>
  </view>
</template>

<script setup>
  import { computed } from 'vue';
  import { onLoad, onShow, onPageScroll, onPullDownRefresh } from '@dcloudio/uni-app';
  import sheep from '@/sheep';
  import $share from '@/sheep/platform/share';
  // 隐藏原生tabBar
  uni.hideTabBar({
    fail: () => {},
  });

  const isDemo = import.meta.env.SHOPRO_DEMO_MODE === 'true';
  const template = computed(() => sheep.$store('app').template?.home);
  // 在此处拦截改变一下首页轮播图 此处先写死后期复活 放到启动函数里
  // (async function() {
  // console.log('原代码首页定制化数据',template)
  // let {
  // 	data
  // } = await index2Api.decorate();
  // console.log('首页导航配置化过高无法兼容',JSON.parse(data[1].value))
  // 改变首页底部数据 但是没有通过数组id获取商品数据接口
  // let {
  // 	data: datas
  // } = await index2Api.spids();
  // template.value.data[9].data.goodsIds = datas.list.map(item => item.id);
  // template.value.data[0].data.list = JSON.parse(data[0].value).map(item => {
  // 	return {
  // 		src: item.picUrl,
  // 		url: item.url,
  // 		title: item.name,
  // 		type: "image"
  // 	}
  // })
  // }())

  onLoad((options) => {
    // #ifdef MP
    // 小程序识别二维码
    if (options.scene) {
      const sceneParams = decodeURIComponent(options.scene).split('=');
      console.log('sceneParams=>', sceneParams);
      options[sceneParams[0]] = sceneParams[1];
    }
    // #endif

    // 预览模板
    if (options.templateId) {
      sheep.$store('app').init(options.templateId);
    }

    // 解析分享信息
    if (options.spm) {
      $share.decryptSpm(options.spm);
    }

    // 进入指定页面(完整页面路径)
    if (options.page) {
      sheep.$router.go(decodeURIComponent(options.page));
    }
  });

  onShow(async() => {
    // #ifdef APP-PLUS
    // ios首次授权网络，需要重新加载一次应用初始化
    // 可能需要考虑上uni.onNetworkStatusChange，uni.offNetworkStatusChange组合拳以及主动主动唤起权限申请
    // 一开始放app.vue，感觉负载太重，搬到这里来了。
    // 如果你的首页不是这个页面，需要把代码搬过去。
    if (sheep.$platform.os === 'ios') {
      if (await sheep.$platform.checkNetwork()) {
        await sheep.$store('app').init();
      }
    }
    // #endif
  });

  // 下拉刷新
  onPullDownRefresh(() => {
    sheep.$store('app').init();
    setTimeout(function () {
      uni.stopPullDownRefresh();
    }, 800);
  });

  onPageScroll(() => {});
</script>

<style scoped>.demo-hero{margin:12px 16px 0;padding:28px 24px;background:linear-gradient(125deg,#e8dbc9,#f5ece1);border-radius:16px;color:#554638;position:relative;overflow:hidden}.hero-eyebrow{letter-spacing:4px;font-size:10px;color:#947e65}.hero-title{font-size:28px;font-weight:700;line-height:1.5;margin:12px 0}.hero-subtitle{font-size:12px;color:#9b826b}.hero-badge{display:inline-block;background:#fff8;color:#9f8266;border-radius:20px;padding:6px 10px;font-size:10px;margin-top:18px}</style>
