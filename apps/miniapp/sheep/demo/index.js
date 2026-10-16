import {
  products as seedProducts,
  categories as seedCategories,
  members,
  orders,
  couponTemplates,
} from '../../../../packages/demo-data/index.js';
import { demoTemplate } from './template';
const clone = (value) => JSON.parse(JSON.stringify(value));
const localImage = (value) => value?.replace('/demo/', '/static/demo/').replace('.svg', '.png');
const products = clone(seedProducts).map((product) => ({
  ...product,
  picUrl: localImage(product.picUrl),
  sliderPicUrls: product.sliderPicUrls.map(localImage),
  description: '<p>演示商品。安心好物，陪伴它的每一天。</p>',
  skus: product.skus.map((sku) => ({ ...sku, picUrl: localImage(sku.picUrl) })),
}));
const categories = clone(seedCategories).map((row) => ({ ...row, picUrl: localImage(row.picUrl) }));
let cart = uni.getStorageSync('mall-demo-cart') || [];
const save = () => uni.setStorageSync('mall-demo-cart', cart);
export async function demoRequest(config) {
  const url = config.url.split('?')[0];
  const params = { ...config.params };
  const query = config.url.split('?')[1];
  if (query)
    query.split('&').forEach((pair) => {
      const [key, value] = pair.split('=');
      params[decodeURIComponent(key)] = decodeURIComponent(value || '');
    });
  const data = config.data || {};
  let result;
  if (url.startsWith('/promotion/diy-template/')) result = demoTemplate;
  else if (url === '/system/tenant/get-by-website') result = { id: 1 };
  else if (url === '/product/category/list') result = categories;
  else if (url === '/product/spu/list-by-ids')
    result = products.filter(
      (p) =>
        p.status === 1 && (!params.ids || String(params.ids).split(',').includes(String(p.id))),
    );
  else if (url === '/product/spu/page') {
    const rows = products.filter(
      (p) =>
        p.status === 1 &&
        (!params.keyword || p.name.includes(params.keyword)) &&
        (!params.categoryId ||
          p.categoryId === Number(params.categoryId) ||
          categories
            .filter((c) => c.parentId === Number(params.categoryId))
            .some((c) => c.id === p.categoryId)),
    );
    const page = Math.max(1, Number(params.pageNo) || 1),
      size = Math.max(1, Number(params.pageSize) || 10);
    result = { list: rows.slice((page - 1) * size, page * size), total: rows.length };
  } else if (url === '/product/spu/get-detail')
    result = products.find((p) => p.id === Number(params.id) && p.status === 1) || null;
  else if (url === '/trade/order/settlement-product')
    result = products
      .filter((p) =>
        String(params.spuIds || '')
          .split(',')
          .includes(String(p.id)),
      )
      .map((p) => ({ id: p.id, spuId: p.id, skus: p.skus }));
  else if (url === '/member/user/get')
    result = { ...members[0], avatar: '/static/demo/avatar.png' };
  else if (url === '/pay/wallet/get') result = { balance: 5000 };
  else if (url === '/trade/order/get-count')
    result = {
      allCount: 5,
      unpaidCount: 1,
      undeliveredCount: 2,
      deliveredCount: 1,
      uncommentedCount: 1,
      afterSaleCount: 0,
    };
  else if (url === '/promotion/coupon/get-unused-count') result = 0;
  else if (url === '/product/comment/get-summary')
    result = { allCount: 0, goodCount: 0, goodRate: 100 };
  else if (url === '/trade/cart/list') result = { validList: cart, invalidList: [] };
  else if (url === '/trade/cart/add') {
    const product = products.find((p) => p.skus.some((s) => s.id === data.skuId));
    if (!product) return { code: 400, msg: '商品不存在', data: null };
    const existing = cart.find((r) => r.sku.id === data.skuId);
    const count = Number(data.count) || 1;
    if (count < 1 || (existing?.count || 0) + count > product.stock)
      return { code: 400, msg: '数量超过库存', data: null };
    if (existing) existing.count += count;
    else cart.push({ id: Date.now(), count, selected: true, spu: product, sku: product.skus[0] });
    save();
    result = true;
  } else if (url === '/trade/cart/update-count') {
    const row = cart.find((r) => r.id === data.id);
    if (!row || Number(data.count) < 1 || Number(data.count) > row.sku.stock)
      return { code: 400, msg: '数量不合法', data: null };
    row.count = Number(data.count);
    save();
    result = true;
  } else if (url === '/trade/cart/update-selected') {
    cart.forEach((r) => {
      if (data.ids.includes(r.id)) r.selected = data.selected;
    });
    save();
    result = true;
  } else if (url === '/trade/cart/delete') {
    cart = cart.filter((r) => !String(params.ids).split(',').includes(String(r.id)));
    save();
    result = true;
  } else if (url === '/trade/order/page')
    result = {
      list: clone(orders).map((o) => ({
        ...o,
        items: o.items.map((i) => ({ ...i, picUrl: localImage(i.picUrl) })),
      })),
      total: orders.length,
    };
  else if (url === '/trade/order/get-detail')
    result = clone(orders.find((o) => o.id === Number(params.id)) || null);
  else if (url === '/member/address/get-default') result = null;
  else if ((config.method || 'GET').toUpperCase() !== 'GET') {
    uni.showToast({ title: '此操作需要正式后端，演示模式不产生交易', icon: 'none' });
    return { code: 501, msg: '需要正式后端', data: null };
  } else if (url.endsWith('/page')) result = { list: [], total: 0 };
  else if (url.includes('/list') || url.endsWith('-list')) result = [];
  else if (url.endsWith('count')) result = 0;
  else result = null;
  if (config.custom?.showSuccess)
    uni.showToast({ title: config.custom.successMsg || '演示操作已完成', icon: 'none' });
  return { code: 0, msg: '演示数据', data: clone(result ?? null) };
}
