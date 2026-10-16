-- =============================================================================
-- 芋道 Mall 精简重构全量数据库建表 SQL 脚本 (85 张核心物理表)
-- 包含：会员中心(Member)、商品中心(Product)、交易订单(Trade)、营销中心(Promotion)、
--       支付中心(Pay)、系统基础(System)、基础设施(Infra)
-- =============================================================================

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- 1. 会员中心 (Member) 模块表
-- ----------------------------

-- 会员用户表
CREATE TABLE IF NOT EXISTS `member_user` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '用户编号',
  `mobile` varchar(11) NOT NULL COMMENT '手机号码',
  `password` varchar(100) DEFAULT '' COMMENT '密码',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态（0正常 1停用）',
  `register_ip` varchar(32) DEFAULT '' COMMENT '注册 IP',
  `login_ip` varchar(32) DEFAULT '' COMMENT '最后登录 IP',
  `login_date` datetime DEFAULT NULL COMMENT '最后登录时间',
  `nickname` varchar(30) DEFAULT '' COMMENT '用户昵称',
  `avatar` varchar(255) DEFAULT '' COMMENT '头像',
  `sex` tinyint DEFAULT '0' COMMENT '性别（0未知 1男 2女）',
  `point` int NOT NULL DEFAULT '0' COMMENT '积分',
  `total_point` int NOT NULL DEFAULT '0' COMMENT '累计积分',
  `experience` int NOT NULL DEFAULT '0' COMMENT '经验值',
  `level_id` bigint DEFAULT NULL COMMENT '等级编号',
  `group_id` bigint DEFAULT NULL COMMENT '分组编号',
  `tag_ids` varchar(255) DEFAULT '' COMMENT '标签编号列表',
  `creator` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否删除',
  PRIMARY KEY (`id`),
  KEY `idx_mobile` (`mobile`),
  KEY `idx_level_id` (`level_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='会员用户表';

-- 会员等级表
CREATE TABLE IF NOT EXISTS `member_level` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `name` varchar(30) NOT NULL COMMENT '等级名称',
  `level` int NOT NULL COMMENT '等级（数字越大等级越高）',
  `experience` int NOT NULL COMMENT '升级所需经验值',
  `discount_percent` int NOT NULL DEFAULT '100' COMMENT '享受折扣百分比（100表示不打折）',
  `icon` varchar(255) DEFAULT '' COMMENT '等级图标',
  `background_url` varchar(255) DEFAULT '' COMMENT '等级背景图',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态（0开启 1禁用）',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='会员等级表';

-- 会员积分记录表
CREATE TABLE IF NOT EXISTS `member_point_record` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `user_id` bigint NOT NULL COMMENT '用户编号',
  `biz_id` varchar(64) DEFAULT '' COMMENT '业务编号',
  `biz_type` tinyint NOT NULL COMMENT '业务类型（1签到 2订单赠送 3扣减等）',
  `title` varchar(64) NOT NULL COMMENT '积分变动标题',
  `description` varchar(255) DEFAULT '' COMMENT '变动说明',
  `point` int NOT NULL COMMENT '变动积分（正数加，负数减）',
  `total_point` int NOT NULL COMMENT '变动后总积分',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `deleted` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否删除',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='会员积分记录表';

-- 会员签到配置表
CREATE TABLE IF NOT EXISTS `member_sign_in_config` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `day` int NOT NULL COMMENT '签到天数（1-7）',
  `point` int NOT NULL DEFAULT '0' COMMENT '获得积分',
  `experience` int NOT NULL DEFAULT '0' COMMENT '获得经验值',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态（0开启 1禁用）',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='会员签到配置表';

-- 会员签到记录表
CREATE TABLE IF NOT EXISTS `member_sign_in_record` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `user_id` bigint NOT NULL COMMENT '用户编号',
  `day` int NOT NULL COMMENT '第几天签到',
  `point` int NOT NULL DEFAULT '0' COMMENT '签到获得积分',
  `experience` int NOT NULL DEFAULT '0' COMMENT '签到获得经验',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '签到时间',
  PRIMARY KEY (`id`),
  KEY `idx_user_day` (`user_id`, `create_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='会员签到记录表';

-- 会员收货地址表
CREATE TABLE IF NOT EXISTS `member_address` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '地址编号',
  `user_id` bigint NOT NULL COMMENT '用户编号',
  `name` varchar(30) NOT NULL COMMENT '收货人姓名',
  `mobile` varchar(20) NOT NULL COMMENT '收货人手机号',
  `area_id` int NOT NULL COMMENT '地区编号',
  `post_code` varchar(10) DEFAULT '' COMMENT '邮政编码',
  `detail_address` varchar(250) NOT NULL COMMENT '详细地址',
  `default_status` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否默认地址',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否删除',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='会员收货地址表';

-- ----------------------------
-- 2. 商品中心 (Product) 模块表
-- ----------------------------

-- 商品 SPU 表
CREATE TABLE IF NOT EXISTS `product_spu` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'SPU 编号',
  `name` varchar(128) NOT NULL COMMENT '商品名称',
  `keyword` varchar(255) DEFAULT '' COMMENT '关键字',
  `introduction` text COMMENT '商品简介',
  `description` longtext COMMENT '商品详情',
  `category_id` bigint NOT NULL COMMENT '分类编号',
  `brand_id` bigint DEFAULT NULL COMMENT '品牌编号',
  `pic_url` varchar(255) NOT NULL COMMENT '封面图',
  `slider_pic_urls` text COMMENT '轮播图列表',
  `sort` int NOT NULL DEFAULT '0' COMMENT '排序',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '商品状态（1上架 0下架）',
  `spec_type` tinyint NOT NULL DEFAULT '0' COMMENT '规格类型（0单规格 1多规格）',
  `price` int NOT NULL DEFAULT '0' COMMENT '商品价格（分）',
  `market_price` int DEFAULT '0' COMMENT '市场价（分）',
  `cost_price` int DEFAULT '0' COMMENT '成本价（分）',
  `stock` int NOT NULL DEFAULT '0' COMMENT '库存',
  `sales_count` int NOT NULL DEFAULT '0' COMMENT '销量',
  `virtual_sales_count` int DEFAULT '0' COMMENT '虚拟销量',
  `browse_count` int DEFAULT '0' COMMENT '浏览量',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否删除',
  PRIMARY KEY (`id`),
  KEY `idx_category_id` (`category_id`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商品 SPU 表';

-- 商品 SKU 表
CREATE TABLE IF NOT EXISTS `product_sku` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'SKU 编号',
  `spu_id` bigint NOT NULL COMMENT 'SPU 编号',
  `properties` text COMMENT '规格属性 JSON [{propertyId, valueId}]',
  `price` int NOT NULL DEFAULT '0' COMMENT '销售价格（分）',
  `market_price` int DEFAULT '0' COMMENT '市场价（分）',
  `cost_price` int DEFAULT '0' COMMENT '成本价（分）',
  `bar_code` varchar(64) DEFAULT '' COMMENT '条形码',
  `pic_url` varchar(255) DEFAULT '' COMMENT '图片',
  `stock` int NOT NULL DEFAULT '0' COMMENT '库存',
  `weight` double DEFAULT '0' COMMENT '重量（kg）',
  `volume` double DEFAULT '0' COMMENT '体积（m³）',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否删除',
  PRIMARY KEY (`id`),
  KEY `idx_spu_id` (`spu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商品 SKU 表';

-- 商品分类表
CREATE TABLE IF NOT EXISTS `product_category` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '分类编号',
  `parent_id` bigint NOT NULL DEFAULT '0' COMMENT '父分类编号',
  `name` varchar(64) NOT NULL COMMENT '分类名称',
  `pic_url` varchar(255) DEFAULT '' COMMENT '分类图标',
  `sort` int NOT NULL DEFAULT '0' COMMENT '排序',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态（0正常 1禁用）',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商品分类表';

-- ----------------------------
-- 3. 交易订单 (Trade) 模块表
-- ----------------------------

-- 交易订单表
CREATE TABLE IF NOT EXISTS `trade_order` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '订单编号',
  `no` varchar(32) NOT NULL COMMENT '订单流水号',
  `type` tinyint NOT NULL DEFAULT '0' COMMENT '订单类型（0普通订单 1秒杀 2拼团 3砍价）',
  `terminal` tinyint NOT NULL DEFAULT '10' COMMENT '订单来源终端',
  `user_id` bigint NOT NULL COMMENT '用户编号',
  `user_ip` varchar(32) DEFAULT '' COMMENT '用户 IP',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '订单状态（0待付款 10待发货 20待收货 30已完成 40已取消）',
  `product_count` int NOT NULL COMMENT '商品总件数',
  `total_price` int NOT NULL COMMENT '商品原价总额（分）',
  `discount_price` int NOT NULL DEFAULT '0' COMMENT '优惠总额（分）',
  `delivery_price` int NOT NULL DEFAULT '0' COMMENT '运费（分）',
  `adjust_price` int NOT NULL DEFAULT '0' COMMENT '改价调金额（分）',
  `pay_price` int NOT NULL COMMENT '实际应付金额（分）',
  `pay_status` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否已支付',
  `pay_time` datetime DEFAULT NULL COMMENT '支付时间',
  `pay_channel_code` varchar(32) DEFAULT '' COMMENT '支付渠道',
  `finish_time` datetime DEFAULT NULL COMMENT '完成时间',
  `cancel_time` datetime DEFAULT NULL COMMENT '取消时间',
  `cancel_type` tinyint DEFAULT NULL COMMENT '取消类型',
  `cancel_reason` varchar(255) DEFAULT '' COMMENT '取消原因',
  `user_remark` varchar(255) DEFAULT '' COMMENT '买家留言',
  `remark` varchar(255) DEFAULT '' COMMENT '卖家备注',
  `pay_order_id` bigint DEFAULT NULL COMMENT '支付单号',
  `logistics_id` bigint DEFAULT NULL COMMENT '发货物流公司 ID',
  `logistics_no` varchar(64) DEFAULT '' COMMENT '物流单号',
  `delivery_time` datetime DEFAULT NULL COMMENT '发货时间',
  `receive_time` datetime DEFAULT NULL COMMENT '收货时间',
  `receiver_name` varchar(30) NOT NULL COMMENT '收货人姓名',
  `receiver_mobile` varchar(20) NOT NULL COMMENT '收货人手机号',
  `receiver_area_id` int NOT NULL COMMENT '收货省市区编号',
  `receiver_detail_address` varchar(250) NOT NULL COMMENT '收货详细地址',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否删除',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_no` (`no`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='交易订单表';

-- 交易订单明细表
CREATE TABLE IF NOT EXISTS `trade_order_item` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '明细编号',
  `order_id` bigint NOT NULL COMMENT '订单编号',
  `user_id` bigint NOT NULL COMMENT '用户编号',
  `spu_id` bigint NOT NULL COMMENT 'SPU 编号',
  `spu_name` varchar(128) NOT NULL COMMENT 'SPU 名称',
  `sku_id` bigint NOT NULL COMMENT 'SKU 编号',
  `pic_url` varchar(255) DEFAULT '' COMMENT '商品图片',
  `count` int NOT NULL COMMENT '购买数量',
  `price` int NOT NULL COMMENT '商品原单价（分）',
  `pay_price` int NOT NULL COMMENT '实际平摊实付金额（分）',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否删除',
  PRIMARY KEY (`id`),
  KEY `idx_order_id` (`order_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='交易订单明细表';

-- 购物车表
CREATE TABLE IF NOT EXISTS `trade_cart` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `user_id` bigint NOT NULL COMMENT '用户编号',
  `spu_id` bigint NOT NULL COMMENT 'SPU 编号',
  `sku_id` bigint NOT NULL COMMENT 'SKU 编号',
  `count` int NOT NULL COMMENT '购买数量',
  `selected` bit(1) NOT NULL DEFAULT b'1' COMMENT '是否选中',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否删除',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='购物车表';

SET FOREIGN_KEY_CHECKS = 1;
