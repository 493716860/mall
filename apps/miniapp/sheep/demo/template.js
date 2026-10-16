const style = {
  bgType: 'color',
  bgColor: '#ffffff',
  marginLeft: 10,
  marginRight: 10,
  marginTop: 8,
  marginBottom: 8,
  padding: 0,
};
const block = (id, property) => ({ id, property: { style, ...property } });
const navigationBar = {
  styleType: 'normal',
  showType: 'always',
  bgType: 'color',
  bgColor: '#ffffff',
  otherCells: [],
  mpCells: [],
};
const tabBar = {
  mode: 1,
  theme: 'orange',
  style: { color: '#87929c', activeColor: '#e29158', bgType: 'color', bgColor: '#ffffff' },
  items: [
    ['首页', 'index', 'home'],
    ['分类', 'category', 'category'],
    ['购物车', 'cart', 'cart'],
    ['我的', 'user', 'user'],
  ].map(([text, page, icon]) => ({
    text,
    url: `/pages/index/${page}`,
    iconUrl: `/static/demo/${icon}.png`,
    activeIconUrl: `/static/demo/${icon}-active.png`,
  })),
};
export const demoTemplate = {
  id: 1,
  name: '我的宠物商城',
  property: { tabBar },
  home: {
    page: { bgType: 'color', bgColor: '#f6f3ef' },
    navigationBar,
    components: [
      block('SearchBar', {
        placeholder: '搜索宠物食品、日常用品',
        borderRadius: 20,
        height: 40,
        textColor: '#999',
        hotKeywords: [],
      }),
      block('TitleBar', {
        title: '每日优选',
        description: '安心好物，陪伴它的每一天',
        titleColor: '#3e3935',
        descriptionColor: '#998e84',
        titleSize: 20,
        descriptionSize: 12,
        titleWeight: 700,
        textAlign: 'left',
        height: 68,
        marginLeft: 10,
      }),
      block('ProductCard', {
        layoutType: 'twoCol',
        spuIds: [101, 102, 103, 104, 105, 106, 107, 108],
        space: 8,
        borderRadiusTop: 12,
        borderRadiusBottom: 12,
        fields: {
          name: { show: true, color: '#3e3935' },
          introduction: { show: true, color: '#998e84' },
          price: { show: true, color: '#cf7543' },
          marketPrice: { show: false, color: '#999' },
          salesCount: { show: true, color: '#999' },
          stock: { show: false, color: '#999' },
        },
        badge: { show: false },
        btnBuy: { type: 'text', text: '选购', bgBeginColor: '#daa278', bgEndColor: '#c17b4d' },
      }),
    ],
  },
  user: {
    page: { bgType: 'color', bgColor: '#f6f3ef' },
    navigationBar,
    components: [
      block('UserCard', { space: 0 }),
      block('TitleBar', {
        title: '我的订单',
        titleColor: '#3e3935',
        titleSize: 18,
        titleWeight: 600,
        height: 50,
        marginLeft: 10,
      }),
      block('UserOrder', { space: 0 }),
      block('TitleBar', {
        title: '我的宠物商城 · 演示模式',
        description: '本页面为开源商城的独立预览',
        titleColor: '#998e84',
        descriptionColor: '#aaa',
        titleSize: 13,
        descriptionSize: 11,
        height: 70,
        textAlign: 'center',
      }),
    ],
  },
};
