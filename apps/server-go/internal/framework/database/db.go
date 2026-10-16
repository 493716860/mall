package database

import (
	"log"
	"time"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// 1. 商品中心实体模型
type ProductSpu struct {
	ID            int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name          string    `gorm:"size:255;not null" json:"name"`
	CategoryID    int64     `json:"categoryId"`
	CategoryName  string    `gorm:"size:100" json:"categoryName"`
	BrandID       int64     `json:"brandId"`
	PicURL        string    `gorm:"size:512" json:"picUrl"`
	SliderPicURLs string    `gorm:"type:text" json:"sliderPicUrls"` // 逗号分隔
	Price         int64     `json:"price"`                          // 分
	MarketPrice   int64     `json:"marketPrice"`
	CostPrice     int64     `json:"costPrice"`
	Stock         int       `json:"stock"`
	SalesCount    int       `json:"salesCount"`
	Status        int       `json:"status"` // 0: 仓库, 1: 上架, 2: 回收站
	Sort          int       `json:"sort"`
	Description   string    `gorm:"type:text" json:"description"`
	CreatedAt     time.Time `json:"createTime"`
	UpdatedAt     time.Time `json:"updateTime"`
}

type ProductSku struct {
	ID          int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	SpuID       int64     `gorm:"index" json:"spuId"`
	Name        string    `gorm:"size:255" json:"name"`
	PicURL      string    `gorm:"size:512" json:"picUrl"`
	Price       int64     `json:"price"`
	MarketPrice int64     `json:"marketPrice"`
	CostPrice   int64     `json:"costPrice"`
	Stock       int       `json:"stock"`
	BarCode     string    `gorm:"size:64" json:"barCode"`
	Properties  string    `gorm:"type:text" json:"properties"` // JSON
	CreatedAt   time.Time `json:"createTime"`
	UpdatedAt   time.Time `json:"updateTime"`
}

type ProductCategory struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	ParentID  int64     `json:"parentId"`
	Name      string    `gorm:"size:100;not null" json:"name"`
	PicURL    string    `gorm:"size:512" json:"picUrl"`
	Sort      int       `json:"sort"`
	Status    int       `json:"status"` // 0: 开启, 1: 禁用
	CreatedAt time.Time `json:"createTime"`
}

type ProductBrand struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"size:100;not null" json:"name"`
	PicURL    string    `gorm:"size:512" json:"picUrl"`
	Sort      int       `json:"sort"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"createTime"`
}

type ProductProperty struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"size:100;not null" json:"name"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"createTime"`
}

type ProductPropertyValue struct {
	ID         int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	PropertyID int64     `gorm:"index" json:"propertyId"`
	Name       string    `gorm:"size:100;not null" json:"name"`
	CreatedAt  time.Time `json:"createTime"`
}

type ProductComment struct {
	ID           int64      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID       int64      `json:"userId"`
	UserNickname string     `gorm:"size:64" json:"userNickname"`
	UserAvatar   string     `gorm:"size:512" json:"userAvatar"`
	SpuID        int64      `gorm:"index" json:"spuId"`
	SpuName      string     `gorm:"size:255" json:"spuName"`
	Scores       int        `json:"scores"`
	Content      string     `gorm:"type:text" json:"content"`
	PicURLs      string     `gorm:"type:text" json:"picUrls"`
	ReplyContent string     `gorm:"type:text" json:"replyContent"`
	ReplyTime    *time.Time `json:"replyTime"`
	Visible      bool       `json:"visible"`
	CreatedAt    time.Time  `json:"createTime"`
}

type ProductFavorite struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    int64     `gorm:"index" json:"userId"`
	SpuID     int64     `gorm:"index" json:"spuId"`
	CreatedAt time.Time `json:"createTime"`
}

type ProductBrowseHistory struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    int64     `gorm:"index" json:"userId"`
	SpuID     int64     `gorm:"index" json:"spuId"`
	CreatedAt time.Time `json:"createTime"`
}

// 2. 交易履约实体模型
type TradeOrder struct {
	ID                    int64      `gorm:"primaryKey;autoIncrement" json:"id"`
	No                    string     `gorm:"size:64;uniqueIndex" json:"no"`
	Type                  int        `json:"type"`                   // 0: 普通, 1: 秒杀, 2: 拼团, 3: 砍价
	Status                int        `json:"status"`                 // 0:待付款, 10:待发货, 20:待收货, 30:已完成, 40:已关闭
	UserID                int64      `gorm:"index" json:"userId"`
	UserNickname          string     `gorm:"size:64" json:"userNickname"`
	UserMobile            string     `gorm:"size:20" json:"userMobile"`
	PayPrice              int64      `json:"payPrice"`
	TotalPrice            int64      `json:"totalPrice"`
	DiscountPrice         int64      `json:"discountPrice"`
	DeliveryPrice         int64      `json:"deliveryPrice"`
	CouponPrice           int64      `json:"couponPrice"`
	PointPrice            int64      `json:"pointPrice"`
	PayChannelCode        string     `gorm:"size:32" json:"payChannelCode"`
	ReceiverName          string     `gorm:"size:64" json:"receiverName"`
	ReceiverMobile        string     `gorm:"size:20" json:"receiverMobile"`
	ReceiverDetailAddress string     `gorm:"size:255" json:"receiverDetailAddress"`
	LogisticsID           int64      `json:"logisticsId"`
	LogisticsNo           string     `gorm:"size:64" json:"logisticsNo"`
	UserRemark            string     `gorm:"size:255" json:"userRemark"`
	Remark                string     `gorm:"size:255" json:"remark"`
	PayTime               *time.Time `json:"payTime"`
	DeliveryTime          *time.Time `json:"deliveryTime"`
	ReceiveTime           *time.Time `json:"receiveTime"`
	CancelTime            *time.Time `json:"cancelTime"`
	CreatedAt             time.Time  `json:"createTime"`
	UpdatedAt             time.Time  `json:"updateTime"`
}

type TradeOrderItem struct {
	ID        int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	OrderID   int64  `gorm:"index" json:"orderId"`
	SpuID     int64  `json:"spuId"`
	SpuName   string `gorm:"size:255" json:"spuName"`
	SkuID     int64  `json:"skuId"`
	PicURL    string `gorm:"size:512" json:"picUrl"`
	Count     int    `json:"count"`
	Price     int64  `json:"price"`
	PayPrice  int64  `json:"payPrice"`
}

type TradeAfterSale struct {
	ID               int64      `gorm:"primaryKey;autoIncrement" json:"id"`
	No               string     `gorm:"size:64;uniqueIndex" json:"no"`
	OrderID          int64      `gorm:"index" json:"orderId"`
	OrderItemID      int64      `json:"orderItemId"`
	SpuName          string     `gorm:"size:255" json:"spuName"`
	PicURL           string     `gorm:"size:512" json:"picUrl"`
	UserID           int64      `gorm:"index" json:"userId"`
	UserNickname     string     `gorm:"size:64" json:"userNickname"`
	UserMobile       string     `gorm:"size:20" json:"userMobile"`
	Type             int        `json:"type"`   // 10: 仅退款, 20: 退货退款
	Status           int        `json:"status"` // 0:申请中, 10:卖家同意, 20:买家已发货, 30:退款成功, 40:已拒绝, 50:已取消
	RefundPrice      int64      `json:"refundPrice"`
	ApplyReason      string     `gorm:"size:255" json:"applyReason"`
	ApplyDescription string     `gorm:"type:text" json:"applyDescription"`
	AuditReason      string     `gorm:"size:255" json:"auditReason"`
	LogisticsNo      string     `gorm:"size:64" json:"logisticsNo"`
	LogisticsName    string     `gorm:"size:64" json:"logisticsName"`
	CreatedAt        time.Time  `json:"createTime"`
	UpdatedAt        time.Time  `json:"updateTime"`
}

type TradeAfterSaleLog struct {
	ID          int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	AfterSaleID int64     `gorm:"index" json:"afterSaleId"`
	UserID      int64     `json:"userId"`
	UserType    int       `json:"userType"` // 1: 会员, 2: 管理员
	Content     string    `gorm:"size:255" json:"content"`
	CreatedAt   time.Time `json:"createTime"`
}

type TradeBrokerageUser struct {
	ID             int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID         int64     `gorm:"uniqueIndex" json:"userId"`
	BindUserID     int64     `gorm:"index" json:"bindUserId"`
	BrokeragePrice int64     `json:"brokeragePrice"`
	FrozenPrice    int64     `json:"frozenPrice"`
	TotalBrokerage int64     `json:"totalBrokerage"`
	Enabled        bool      `json:"enabled"`
	BindTime       time.Time `json:"bindTime"`
	CreatedAt      time.Time `json:"createTime"`
}

type TradeBrokerageWithdraw struct {
	ID          int64      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID      int64      `gorm:"index" json:"userId"`
	Price       int64      `json:"price"`
	FeePrice    int64      `json:"feePrice"`
	Status      int        `json:"status"` // 0: 审核中, 10: 审核通过, 20: 已拒绝
	AuditReason string     `gorm:"size:255" json:"auditReason"`
	AuditTime   *time.Time `json:"auditTime"`
	CreatedAt   time.Time  `json:"createTime"`
}

type TradeBrokerageRecord struct {
	ID         int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID     int64     `gorm:"index" json:"userId"`
	BizID      string    `gorm:"size:64" json:"bizId"`
	BizType    int       `json:"bizType"`
	Title      string    `gorm:"size:100" json:"title"`
	Price      int64     `json:"price"`
	TotalPrice int64     `json:"totalPrice"`
	Status     int       `json:"status"` // 0: 待结算, 1: 已结算, 2: 已取消
	CreatedAt  time.Time `json:"createTime"`
}

type DeliveryExpress struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"size:64;not null" json:"name"`
	Code      string    `gorm:"size:32;not null" json:"code"`
	Sort      int       `json:"sort"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"createTime"`
}

type DeliveryPickUpStore struct {
	ID            int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name          string    `gorm:"size:100;not null" json:"name"`
	Mobile        string    `gorm:"size:20" json:"mobile"`
	AreaID        int64     `json:"areaId"`
	DetailAddress string    `gorm:"size:255" json:"detailAddress"`
	Status        int       `json:"status"`
	CreatedAt     time.Time `json:"createTime"`
}

type CartItem struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    int64     `gorm:"index" json:"userId"`
	SpuID     int64     `json:"spuId"`
	SkuID     int64     `json:"skuId"`
	Count     int       `json:"count"`
	Selected  bool      `json:"selected"`
	CreatedAt time.Time `json:"createTime"`
}

type TradeConfig struct {
	ID                         int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	AfterSaleRefundReasons     string `gorm:"size:512" json:"afterSaleRefundReasons"`
	AfterSaleReturnReasons     string `gorm:"size:512" json:"afterSaleReturnReasons"`
	DeliveryExpressFreeEnabled bool   `json:"deliveryExpressFreeEnabled"`
	DeliveryExpressFreePrice   int64  `json:"deliveryExpressFreePrice"`
	DeliveryPickUpEnabled      bool   `json:"deliveryPickUpEnabled"`
	BrokerageEnabled           bool   `json:"brokerageEnabled"`
	BrokerageEnabledCondition  int    `json:"brokerageEnabledCondition"`
	BrokerageBindMode          int    `json:"brokerageBindMode"`
	BrokeragePosterUrls        string `gorm:"size:1024" json:"brokeragePosterUrls"`
	BrokerageFirstPercent      int    `json:"brokerageFirstPercent"`
	BrokerageSecondPercent     int    `json:"brokerageSecondPercent"`
	BrokerageWithdrawMinPrice  int64  `json:"brokerageWithdrawMinPrice"`
	BrokerageWithdrawFeePercent int   `json:"brokerageWithdrawFeePercent"`
	BrokerageFrozenDays        int    `json:"brokerageFrozenDays"`
	BrokerageWithdrawTypes     string `gorm:"size:128" json:"brokerageWithdrawTypes"`
	TencentLbsKey              string `gorm:"size:128" json:"tencentLbsKey"`
	PointTradeDeductRate       int    `json:"pointTradeDeductRate"`
}

// 3. 营销中心实体模型
type CouponTemplate struct {
	ID              int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name            string    `gorm:"size:100;not null" json:"name"`
	Status          int       `json:"status"` // 0: 开启, 1: 禁用
	DiscountType    int       `json:"discountType"` // 1: 满减, 2: 折扣
	DiscountPrice   int64     `json:"discountPrice"`
	DiscountPercent int       `json:"discountPercent"`
	UsePrice        int64     `json:"usePrice"`
	TotalCount      int       `json:"totalCount"`
	TakeCount       int       `json:"takeCount"`
	CreatedAt       time.Time `json:"createTime"`
}

type CouponUser struct {
	ID            int64      `gorm:"primaryKey;autoIncrement" json:"id"`
	TemplateID    int64      `gorm:"index" json:"templateId"`
	UserID        int64      `gorm:"index" json:"userId"`
	Name          string     `gorm:"size:100" json:"name"`
	Status        int        `json:"status"` // 1: 未使用, 2: 已使用, 3: 已过期
	DiscountPrice int64      `json:"discountPrice"`
	UsePrice      int64      `json:"usePrice"`
	UsedTime      *time.Time `json:"usedTime"`
	CreatedAt     time.Time  `json:"createTime"`
}

type SeckillActivity struct {
	ID           int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	SpuID        int64     `gorm:"index" json:"spuId"`
	Name         string    `gorm:"size:100;not null" json:"name"`
	Status       int       `json:"status"`
	StartTime    time.Time `json:"startTime"`
	EndTime      time.Time `json:"endTime"`
	TotalStock   int       `json:"totalStock"`
	SeckillPrice int64     `json:"seckillPrice"`
	CreatedAt    time.Time `json:"createTime"`
}

type SeckillConfig struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"size:100;not null" json:"name"`
	StartTime string    `gorm:"size:32" json:"startTime"`
	EndTime   string    `gorm:"size:32" json:"endTime"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"createTime"`
}

type CombinationActivity struct {
	ID               int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	SpuID            int64     `gorm:"index" json:"spuId"`
	Name             string    `gorm:"size:100;not null" json:"name"`
	Status           int       `json:"status"`
	CombinationPrice int64     `json:"combinationPrice"`
	GroupNum         int       `json:"groupNum"`
	StartTime        time.Time `json:"startTime"`
	EndTime          time.Time `json:"endTime"`
	CreatedAt        time.Time `json:"createTime"`
}

type CombinationRecord struct {
	ID         int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	ActivityID int64     `gorm:"index" json:"activityId"`
	HeadID     int64     `json:"headId"`
	UserID     int64     `gorm:"index" json:"userId"`
	Nickname   string    `gorm:"size:64" json:"nickname"`
	Avatar     string    `gorm:"size:512" json:"avatar"`
	Status     int       `json:"status"`
	UserCount  int       `json:"userCount"`
	EndTime    time.Time `json:"endTime"`
	CreatedAt  time.Time `json:"createTime"`
}

type BargainActivity struct {
	ID           int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	SpuID        int64     `gorm:"index" json:"spuId"`
	Name         string    `gorm:"size:100;not null" json:"name"`
	Status       int       `json:"status"`
	BargainPrice int64     `json:"bargainPrice"`
	CreatedAt    time.Time `json:"createTime"`
}

type BargainRecord struct {
	ID                int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	ActivityID        int64     `gorm:"index" json:"activityId"`
	UserID            int64     `gorm:"index" json:"userId"`
	SpuID             int64     `json:"spuId"`
	SkuID             int64     `json:"skuId"`
	BargainFirstPrice int64     `json:"bargainFirstPrice"`
	BargainPrice      int64     `json:"bargainPrice"`
	Status            int       `json:"status"` // 0: 进行中, 1: 成功, 2: 失败
	CreatedAt         time.Time `json:"createTime"`
}

type BargainHelp struct {
	ID          int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	RecordID    int64     `gorm:"index" json:"recordId"`
	UserID      int64     `gorm:"index" json:"userId"`
	ReducePrice int64     `json:"reducePrice"`
	CreatedAt   time.Time `json:"createTime"`
}

type DiyTemplate struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"size:100;not null" json:"name"`
	Used      bool      `json:"used"`
	Property  string    `gorm:"type:text" json:"property"`
	CreatedAt time.Time `json:"createTime"`
}

type DiyPage struct {
	ID         int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	TemplateID int64     `gorm:"index" json:"templateId"`
	Name       string    `gorm:"size:100;not null" json:"name"`
	Property   string    `gorm:"type:text" json:"property"`
	CreatedAt  time.Time `json:"createTime"`
}

type KefuMessage struct {
	ID          int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID      int64     `gorm:"index" json:"userId"`
	SenderType  int       `json:"senderType"`
	ContentType int       `json:"contentType"`
	Content     string    `gorm:"type:text" json:"content"`
	ReadStatus  bool      `json:"readStatus"`
	CreatedAt   time.Time `json:"createTime"`
}

type Banner struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Title     string    `gorm:"size:100;not null" json:"title"`
	PicURL    string    `gorm:"size:512" json:"picUrl"`
	URL       string    `gorm:"size:512" json:"url"`
	Status    int       `json:"status"`
	Sort      int       `json:"sort"`
	CreatedAt time.Time `json:"createTime"`
}

type Article struct {
	ID         int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	CategoryID int64     `json:"categoryId"`
	Title      string    `gorm:"size:255;not null" json:"title"`
	PicURL     string    `gorm:"size:512" json:"picUrl"`
	Content    string    `gorm:"type:text" json:"content"`
	Status     int       `json:"status"`
	CreatedAt  time.Time `json:"createTime"`
}

type ArticleCategory struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"size:100;not null" json:"name"`
	Sort      int       `json:"sort"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"createTime"`
}

type DiscountActivity struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"size:100;not null" json:"name"`
	Status    int       `json:"status"` // 0: 开启, 1: 关闭
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	Remark    string    `gorm:"size:255" json:"remark"`
	CreatedAt time.Time `json:"createTime"`
}

type RewardActivity struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"size:100;not null" json:"name"`
	Status    int       `json:"status"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	Remark    string    `gorm:"size:255" json:"remark"`
	CreatedAt time.Time `json:"createTime"`
}

type PointActivity struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	SpuID     int64     `gorm:"index" json:"spuId"`
	Status    int       `json:"status"`
	Point     int       `json:"point"`
	CreatedAt time.Time `json:"createTime"`
}

type KefuConversation struct {
	ID           int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID       int64     `gorm:"uniqueIndex" json:"userId"`
	UserNickname string    `gorm:"size:64" json:"userNickname"`
	UserAvatar   string    `gorm:"size:512" json:"userAvatar"`
	LastMessage  string    `gorm:"size:255" json:"lastMessage"`
	LastTime     time.Time `json:"lastTime"`
	Pinned       bool      `json:"pinned"`
	CreatedAt    time.Time `json:"createTime"`
}

type PayRefund struct {
	ID              int64      `gorm:"primaryKey;autoIncrement" json:"id"`
	No              string     `gorm:"size:64;uniqueIndex" json:"no"`
	OrderID         int64      `gorm:"index" json:"orderId"`
	MerchantOrderID string     `gorm:"size:64" json:"merchantOrderId"`
	RefundPrice     int64      `json:"refundPrice"`
	Reason          string     `gorm:"size:255" json:"reason"`
	Status          int        `json:"status"` // 0: 处理中, 10: 成功, 20: 失败
	SuccessTime     *time.Time `json:"successTime"`
	CreatedAt       time.Time  `json:"createTime"`
}

type PayWalletRechargePackage struct {
	ID         int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name       string    `gorm:"size:100;not null" json:"name"`
	PayPrice   int64     `json:"payPrice"`
	BonusPrice int64     `json:"bonusPrice"`
	Status     int       `json:"status"`
	CreatedAt  time.Time `json:"createTime"`
}

type MemberSignInConfig struct {
	ID         int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Day        int       `json:"day"`
	Point      int       `json:"point"`
	Experience int       `json:"experience"`
	Status     int       `json:"status"`
	CreatedAt  time.Time `json:"createTime"`
}

// 4. 会员中心实体模型
type MemberUser struct {
	ID         int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Nickname   string    `gorm:"size:64" json:"nickname"`
	Avatar     string    `gorm:"size:512" json:"avatar"`
	Mobile     string    `gorm:"size:20;uniqueIndex" json:"mobile"`
	Status     int       `json:"status"`
	LevelID    int64     `json:"levelId"`
	LevelName  string    `gorm:"size:64" json:"levelName"`
	Point      int       `json:"point"`
	Experience int       `json:"experience"`
	Balance    int64     `json:"balance"`
	CreatedAt  time.Time `json:"createTime"`
}

type MemberAddress struct {
	ID            int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID        int64     `gorm:"index" json:"userId"`
	Name          string    `gorm:"size:64;not null" json:"name"`
	Mobile        string    `gorm:"size:20;not null" json:"mobile"`
	AreaID        int64     `json:"areaId"`
	AreaName      string    `gorm:"size:100" json:"areaName"`
	DetailAddress string    `gorm:"size:255;not null" json:"detailAddress"`
	DefaultStatus bool      `json:"defaultStatus"`
	CreatedAt     time.Time `json:"createTime"`
}

type MemberPointRecord struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    int64     `gorm:"index" json:"userId"`
	BizType   int       `json:"bizType"`
	Title     string    `gorm:"size:100" json:"title"`
	Point     int       `json:"point"`
	Total     int       `json:"total"`
	CreatedAt time.Time `json:"createTime"`
}

type MemberSignInRecord struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    int64     `gorm:"index" json:"userId"`
	Day       int       `json:"day"`
	Point     int       `json:"point"`
	CreatedAt time.Time `json:"createTime"`
}

type MemberLevel struct {
	ID              int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name            string    `gorm:"size:64;not null" json:"name"`
	Level           int       `json:"level"`
	DiscountPercent int       `json:"discountPercent"`
	Experience      int       `json:"experience"`
	Status          int       `json:"status"`
	CreatedAt       time.Time `json:"createTime"`
}

type MemberTag struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"size:64;not null" json:"name"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"createTime"`
}

type MemberGroup struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"size:64;not null" json:"name"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"createTime"`
}

type MemberConfig struct {
	ID                        int64 `gorm:"primaryKey;autoIncrement" json:"id"`
	PointTradeDeductEnable    bool  `json:"pointTradeDeductEnable"`
	PointTradeDeductUnitPrice int   `json:"pointTradeDeductUnitPrice"`
	PointTradeDeductMaxPrice  int   `json:"pointTradeDeductMaxPrice"`
	PointTradeGivePoint       int   `json:"pointTradeGivePoint"`
}

type MemberExperienceRecord struct {
	ID              int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID          int64     `gorm:"index" json:"userId"`
	BizID           string    `gorm:"size:64" json:"bizId"`
	BizType         int       `json:"bizType"`
	Title           string    `gorm:"size:100" json:"title"`
	Description     string    `gorm:"size:255" json:"description"`
	Experience      int       `json:"experience"`
	TotalExperience int       `json:"totalExperience"`
	CreatedAt       time.Time `json:"createTime"`
}

// 5. 支付中心实体模型
type PayOrder struct {
	ID              int64      `gorm:"primaryKey;autoIncrement" json:"id"`
	AppID           int64      `json:"appId"`
	ChannelCode     string     `gorm:"size:32" json:"channelCode"`
	MerchantOrderID string     `gorm:"size:64;uniqueIndex" json:"merchantOrderId"`
	Subject         string     `gorm:"size:255" json:"subject"`
	Price           int64      `json:"price"`
	Status          int        `json:"status"`
	SuccessTime     *time.Time `json:"successTime"`
	CreatedAt       time.Time  `json:"createTime"`
}

type PayWallet struct {
	ID            int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID        int64     `gorm:"uniqueIndex" json:"userId"`
	Balance       int64     `json:"balance"`
	TotalRecharge int64     `json:"totalRecharge"`
	TotalExpense  int64     `json:"totalExpense"`
	CreatedAt     time.Time `json:"createTime"`
}

type PayWalletTransaction struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	WalletID  int64     `gorm:"index" json:"walletId"`
	UserID    int64     `gorm:"index" json:"userId"`
	BizType   int       `json:"bizType"`
	Title     string    `gorm:"size:100" json:"title"`
	Price     int64     `json:"price"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"createTime"`
}

type PayApp struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"size:100;not null" json:"name"`
	AppKey    string    `gorm:"size:64;uniqueIndex" json:"appKey"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"createTime"`
}

// 6. 系统管理与基础实体模型
type SystemUser struct {
	ID          int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Username    string    `gorm:"size:64;uniqueIndex" json:"username"`
	Password    string    `gorm:"size:128" json:"-"`
	Nickname    string    `gorm:"size:64" json:"nickname"`
	Avatar      string    `gorm:"size:512" json:"avatar"`
	Mobile      string    `gorm:"size:20" json:"mobile"`
	Email       string    `gorm:"size:64" json:"email"`
	DeptID      int64     `json:"deptId"`
	Status      int       `json:"status"`
	CreatedAt   time.Time `json:"createTime"`
}

type SystemRole struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"size:64;not null" json:"name"`
	Code      string    `gorm:"size:64;uniqueIndex" json:"code"`
	Sort      int       `json:"sort"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"createTime"`
}

type SystemDept struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	ParentID  int64     `json:"parentId"`
	Name      string    `gorm:"size:64;not null" json:"name"`
	Sort      int       `json:"sort"`
	Leader    string    `gorm:"size:64" json:"leader"`
	Phone     string    `gorm:"size:20" json:"phone"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"createTime"`
}

type SystemPost struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"size:64;not null" json:"name"`
	Code      string    `gorm:"size:64;uniqueIndex" json:"code"`
	Sort      int       `json:"sort"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"createTime"`
}

type SystemDictType struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"size:100;not null" json:"name"`
	Type      string    `gorm:"size:100;uniqueIndex" json:"type"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"createTime"`
}

type SystemDictData struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	DictType  string    `gorm:"size:100;index" json:"dictType"`
	Label     string    `gorm:"size:100;not null" json:"label"`
	Value     string    `gorm:"size:100;not null" json:"value"`
	Sort      int       `json:"sort"`
	Status    int       `json:"status"`
	ColorType string    `gorm:"size:32" json:"colorType"`
	CreatedAt time.Time `json:"createTime"`
}

type SystemTenant struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"size:100;not null" json:"name"`
	Website   string    `gorm:"size:255" json:"website"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"createTime"`
}

// 7. 基础设施实体模型
type InfraFile struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	ConfigID  int64     `json:"configId"`
	Name      string    `gorm:"size:255" json:"name"`
	Path      string    `gorm:"size:512" json:"path"`
	URL       string    `gorm:"size:512" json:"url"`
	Type      string    `gorm:"size:64" json:"type"`
	Size      int64     `json:"size"`
	CreatedAt time.Time `json:"createTime"`
}

type InfraFileConfig struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"size:100;not null" json:"name"`
	Storage   int       `json:"storage"`
	Master    bool      `json:"master"`
	Config    string    `gorm:"type:text" json:"config"`
	Remark    string    `gorm:"size:255" json:"remark"`
	CreatedAt time.Time `json:"createTime"`
}

type InfraJob struct {
	ID             int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name           string    `gorm:"size:100;not null" json:"name"`
	HandlerName    string    `gorm:"size:100" json:"handlerName"`
	HandlerParam   string    `gorm:"size:255" json:"handlerParam"`
	CronExpression string    `gorm:"size:100" json:"cronExpression"`
	RetryCount     int       `json:"retryCount"`
	RetryInterval  int       `json:"retryInterval"`
	MonitorTimeout int       `json:"monitorTimeout"`
	Status         int       `json:"status"` // 0: 开启, 1: 暂停
	CreatedAt      time.Time `json:"createTime"`
}

type InfraJobLog struct {
	ID             int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	JobID          int64     `gorm:"index" json:"jobId"`
	HandlerName    string    `gorm:"size:100" json:"handlerName"`
	HandlerParam   string    `gorm:"size:255" json:"handlerParam"`
	CronExpression string    `gorm:"size:100" json:"cronExpression"`
	ExecuteIndex   int       `json:"executeIndex"`
	BeginTime      time.Time `json:"beginTime"`
	EndTime        time.Time `json:"endTime"`
	Duration       int       `json:"duration"`
	Status         int       `json:"status"`
	Result         string    `gorm:"type:text" json:"result"`
	CreatedAt      time.Time `json:"createTime"`
}

type InfraApiAccessLog struct {
	ID              int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID          int64     `json:"userId"`
	UserType        int       `json:"userType"`
	ApplicationName string    `gorm:"size:64" json:"applicationName"`
	RequestMethod   string    `gorm:"size:16" json:"requestMethod"`
	RequestURL      string    `gorm:"size:255" json:"requestUrl"`
	RequestParams   string    `gorm:"type:text" json:"requestParams"`
	UserIP          string    `gorm:"size:64" json:"userIp"`
	UserAgent       string    `gorm:"size:255" json:"userAgent"`
	BeginTime       time.Time `json:"beginTime"`
	EndTime         time.Time `json:"endTime"`
	Duration        int       `json:"duration"`
	ResultCode      int       `json:"resultCode"`
	CreatedAt       time.Time `json:"createTime"`
}

type InfraApiErrorLog struct {
	ID                  int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID              int64     `json:"userId"`
	UserType            int       `json:"userType"`
	ApplicationName     string    `gorm:"size:64" json:"applicationName"`
	RequestMethod       string    `gorm:"size:16" json:"requestMethod"`
	RequestURL          string    `gorm:"size:255" json:"requestUrl"`
	RequestParams       string    `gorm:"type:text" json:"requestParams"`
	UserIP              string    `gorm:"size:64" json:"userIp"`
	UserAgent           string    `gorm:"size:255" json:"userAgent"`
	ExceptionTime       time.Time `json:"exceptionTime"`
	ExceptionName       string    `gorm:"size:128" json:"exceptionName"`
	ExceptionMessage    string    `gorm:"type:text" json:"exceptionMessage"`
	ExceptionStackTrace string    `gorm:"type:text" json:"exceptionStackTrace"`
	ProcessStatus       int       `json:"processStatus"`
	CreatedAt           time.Time `json:"createTime"`
}

// InitDB 初始化数据库并自动迁移与种子数据填充
func InitDB(dbPath string) *gorm.DB {
	if dbPath == "" {
		dbPath = "ruoyi_mall.db"
	}
	var err error
	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	sqlDB, err := DB.DB()
	if err == nil {
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(time.Hour)
	}

	// 自动迁移所有 1:1 实体模型
	err = DB.AutoMigrate(
		&ProductSpu{},
		&ProductSku{},
		&ProductCategory{},
		&ProductBrand{},
		&ProductProperty{},
		&ProductPropertyValue{},
		&ProductComment{},
		&ProductFavorite{},
		&ProductBrowseHistory{},
		&TradeOrder{},
		&TradeOrderItem{},
		&TradeAfterSale{},
		&TradeAfterSaleLog{},
		&TradeBrokerageUser{},
		&TradeBrokerageWithdraw{},
		&TradeBrokerageRecord{},
		&DeliveryExpress{},
		&DeliveryPickUpStore{},
		&CartItem{},
		&TradeConfig{},
		&CouponTemplate{},
		&CouponUser{},
		&SeckillActivity{},
		&SeckillConfig{},
		&CombinationActivity{},
		&CombinationRecord{},
		&BargainActivity{},
		&DiyTemplate{},
		&DiyPage{},
		&KefuMessage{},
		&Banner{},
		&Article{},
		&ArticleCategory{},
		&MemberUser{},
		&MemberAddress{},
		&MemberPointRecord{},
		&MemberSignInRecord{},
		&MemberLevel{},
		&MemberTag{},
		&MemberGroup{},
		&PayOrder{},
		&PayWallet{},
		&PayWalletTransaction{},
		&PayApp{},
		&SystemUser{},
		&SystemRole{},
		&SystemDept{},
		&SystemPost{},
		&SystemDictType{},
		&SystemDictData{},
		&SystemTenant{},
		&DiscountActivity{},
		&RewardActivity{},
		&PointActivity{},
		&KefuConversation{},
		&PayRefund{},
		&PayWalletRechargePackage{},
		&MemberSignInConfig{},
		&BargainRecord{},
		&BargainHelp{},
		&MemberConfig{},
		&MemberExperienceRecord{},
		&InfraFile{},
		&InfraFileConfig{},
		&InfraJob{},
		&InfraJobLog{},
		&InfraApiAccessLog{},
		&InfraApiErrorLog{},
	)
	if err != nil {
		log.Printf("AutoMigrate error: %v", err)
	}

	// 种子数据播种
	seedData()

	log.Println("Database initialized and all models migrated successfully:", dbPath)
	return DB
}

func seedData() {
	var count int64

	// 1. 商品中心种子数据
	DB.Model(&ProductCategory{}).Count(&count)
	if count == 0 {
		categories := []ProductCategory{
			{ID: 100, ParentID: 0, Name: "宠物主粮", PicURL: "https://img.iocoder.cn/mall/cat_food.jpg", Sort: 1, Status: 0, CreatedAt: time.Now()},
			{ID: 101, ParentID: 100, Name: "猫咪主粮", PicURL: "https://img.iocoder.cn/mall/cat_food.jpg", Sort: 1, Status: 0, CreatedAt: time.Now()},
			{ID: 102, ParentID: 100, Name: "狗狗主粮", PicURL: "https://img.iocoder.cn/mall/dog_food.jpg", Sort: 2, Status: 0, CreatedAt: time.Now()},
			{ID: 103, ParentID: 0, Name: "宠物用品", PicURL: "https://img.iocoder.cn/mall/cat_litter.jpg", Sort: 2, Status: 0, CreatedAt: time.Now()},
			{ID: 104, ParentID: 103, Name: "智能用品", PicURL: "https://img.iocoder.cn/mall/cat_litter.jpg", Sort: 1, Status: 0, CreatedAt: time.Now()},
		}
		DB.Create(&categories)
	}

	DB.Model(&ProductBrand{}).Count(&count)
	if count == 0 {
		brands := []ProductBrand{
			{ID: 1, Name: "皇家宠物", PicURL: "https://img.iocoder.cn/brand/royal.png", Sort: 1, Status: 0, CreatedAt: time.Now()},
			{ID: 2, Name: "渴望 Orijen", PicURL: "https://img.iocoder.cn/brand/orijen.png", Sort: 2, Status: 0, CreatedAt: time.Now()},
			{ID: 3, Name: "冠能 ProPlan", PicURL: "https://img.iocoder.cn/brand/proplan.png", Sort: 3, Status: 0, CreatedAt: time.Now()},
		}
		DB.Create(&brands)
	}

	DB.Model(&ProductProperty{}).Count(&count)
	if count == 0 {
		props := []ProductProperty{
			{ID: 1, Name: "净重规格", Status: 0, CreatedAt: time.Now()},
			{ID: 2, Name: "主要配方", Status: 0, CreatedAt: time.Now()},
		}
		DB.Create(&props)

		vals := []ProductPropertyValue{
			{ID: 1, PropertyID: 1, Name: "2kg 幼猫装", CreatedAt: time.Now()},
			{ID: 2, PropertyID: 1, Name: "10kg 全期装", CreatedAt: time.Now()},
			{ID: 3, PropertyID: 2, Name: "深海三文鱼配方", CreatedAt: time.Now()},
			{ID: 4, PropertyID: 2, Name: "原切牛肉生骨肉", CreatedAt: time.Now()},
		}
		DB.Create(&vals)
	}

	DB.Model(&ProductSpu{}).Count(&count)
	if count == 0 {
		spus := []ProductSpu{
			{
				ID:            1,
				Name:          "天然深海无谷高蛋白猫粮 10kg",
				CategoryID:    101,
				CategoryName:  "猫咪主粮",
				BrandID:       1,
				PicURL:        "https://img.iocoder.cn/mall/cat_food.jpg",
				SliderPicURLs: "https://img.iocoder.cn/mall/cat_food.jpg",
				Price:         29900,
				MarketPrice:   39900,
				CostPrice:     15000,
				Stock:         500,
				SalesCount:    1280,
				Status:        1,
				Sort:          1,
				Description:   "<p>精选挪威三文鱼，无谷低敏配方，高肉含量，守护猫咪肠胃健康。</p>",
				CreatedAt:     time.Now(),
				UpdatedAt:     time.Now(),
			},
			{
				ID:            2,
				Name:          "全自动智能猫砂盆 智能除臭防带砂",
				CategoryID:    104,
				CategoryName:  "智能用品",
				BrandID:       2,
				PicURL:        "https://img.iocoder.cn/mall/cat_litter.jpg",
				SliderPicURLs: "https://img.iocoder.cn/mall/cat_litter.jpg",
				Price:         89900,
				MarketPrice:   129900,
				CostPrice:     50000,
				Stock:         120,
				SalesCount:    320,
				Status:        1,
				Sort:          2,
				Description:   "<p>多重重力红外传感器，APP远程监控，超大空间，自动清理除臭。</p>",
				CreatedAt:     time.Now(),
				UpdatedAt:     time.Now(),
			},
			{
				ID:            3,
				Name:          "原切冻干生骨肉全价主粮 500g",
				CategoryID:    101,
				CategoryName:  "猫咪主粮",
				BrandID:       1,
				PicURL:        "https://img.iocoder.cn/mall/freeze_dried.jpg",
				SliderPicURLs: "https://img.iocoder.cn/mall/freeze_dried.jpg",
				Price:         11800,
				MarketPrice:   15800,
				CostPrice:     6000,
				Stock:         800,
				SalesCount:    950,
				Status:        1,
				Sort:          3,
				Description:   "<p>-40℃低温真空冻干锁鲜，保留98%原肉营养。</p>",
				CreatedAt:     time.Now(),
				UpdatedAt:     time.Now(),
			},
		}
		DB.Create(&spus)
	}

	DB.Model(&ProductSku{}).Count(&count)
	if count == 0 {
		skus := []ProductSku{
			{ID: 10, SpuID: 1, Name: "2kg 幼猫装", PicURL: "https://img.iocoder.cn/mall/cat_food.jpg", Price: 29900, MarketPrice: 39900, CostPrice: 15000, Stock: 500, BarCode: "6901001", Properties: "[{\"propertyId\":1,\"valueId\":1,\"valueName\":\"2kg 幼猫装\"}]", CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{ID: 11, SpuID: 1, Name: "10kg 全期装", PicURL: "https://img.iocoder.cn/mall/cat_food.jpg", Price: 79900, MarketPrice: 99900, CostPrice: 40000, Stock: 300, BarCode: "6901002", Properties: "[{\"propertyId\":1,\"valueId\":2,\"valueName\":\"10kg 全期装\"}]", CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{ID: 20, SpuID: 2, Name: "智能感应标配版", PicURL: "https://img.iocoder.cn/mall/cat_litter.jpg", Price: 89900, MarketPrice: 129900, CostPrice: 50000, Stock: 120, BarCode: "6902001", Properties: "[{\"propertyId\":1,\"valueId\":1,\"valueName\":\"智能感应标配版\"}]", CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{ID: 30, SpuID: 3, Name: "深海三文鱼 500g", PicURL: "https://img.iocoder.cn/mall/freeze_dried.jpg", Price: 11800, MarketPrice: 15800, CostPrice: 6000, Stock: 800, BarCode: "6903001", Properties: "[{\"propertyId\":2,\"valueId\":3,\"valueName\":\"深海三文鱼配方\"}]", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		}
		DB.Create(&skus)
	}

	// 2. 交易中心种子数据
	DB.Model(&TradeOrder{}).Count(&count)
	if count == 0 {
		orderTime := time.Now().Add(-2 * time.Hour)
		order := TradeOrder{
			ID:                    10001,
			No:                    "202610150001",
			Type:                  0,
			Status:                10, // 待发货
			UserID:                1001,
			UserNickname:          "微信用户小王",
			UserMobile:            "13800138000",
			PayPrice:              29900,
			TotalPrice:            29900,
			DiscountPrice:         0,
			DeliveryPrice:         0,
			PayChannelCode:        "wx_lite",
			ReceiverName:          "小王",
			ReceiverMobile:        "13800138000",
			ReceiverDetailAddress: "北京市朝阳区三里屯 8 号院 3 号楼 602",
			UserRemark:            "请尽快发货，猫咪要断粮了，谢谢！",
			PayTime:               &orderTime,
			CreatedAt:             orderTime,
			UpdatedAt:             orderTime,
		}
		DB.Create(&order)

		items := []TradeOrderItem{
			{
				OrderID:  10001,
				SpuID:    1,
				SpuName:  "天然深海无谷高蛋白猫粮 10kg",
				SkuID:    10,
				PicURL:   "https://img.iocoder.cn/mall/cat_food.jpg",
				Count:    1,
				Price:    29900,
				PayPrice: 29900,
			},
		}
		DB.Create(&items)
	}

	DB.Model(&DeliveryExpress{}).Count(&count)
	if count == 0 {
		expresses := []DeliveryExpress{
			{ID: 1, Name: "顺丰速运", Code: "SF", Sort: 1, Status: 0, CreatedAt: time.Now()},
			{ID: 2, Name: "圆通速递", Code: "YTO", Sort: 2, Status: 0, CreatedAt: time.Now()},
			{ID: 3, Name: "中通快递", Code: "ZTO", Sort: 3, Status: 0, CreatedAt: time.Now()},
		}
		DB.Create(&expresses)
	}

	DB.Model(&DeliveryPickUpStore{}).Count(&count)
	if count == 0 {
		stores := []DeliveryPickUpStore{
			{ID: 1, Name: "宠物生活体验馆（三里屯店）", Mobile: "010-88888888", AreaID: 110105, DetailAddress: "北京市朝阳区三里屯 8 号", Status: 0, CreatedAt: time.Now()},
		}
		DB.Create(&stores)
	}

	DB.Model(&TradeBrokerageUser{}).Count(&count)
	if count == 0 {
		bu := TradeBrokerageUser{
			ID:             1,
			UserID:         1001,
			BindUserID:     0,
			BrokeragePrice: 35000,
			FrozenPrice:    5000,
			TotalBrokerage: 40000,
			Enabled:        true,
			BindTime:       time.Now().Add(-30 * 24 * time.Hour),
			CreatedAt:      time.Now(),
		}
		DB.Create(&bu)
	}

	DB.Model(&TradeBrokerageRecord{}).Count(&count)
	if count == 0 {
		brs := []TradeBrokerageRecord{
			{ID: 1, UserID: 1001, BizID: "10001", BizType: 1, Title: "推荐用户购买宠物主粮返佣", Price: 2990, TotalPrice: 29900, Status: 1, CreatedAt: time.Now().Add(-24 * time.Hour)},
			{ID: 2, UserID: 1001, BizID: "10002", BizType: 1, Title: "推荐用户购买智能猫砂盆返佣", Price: 8990, TotalPrice: 89900, Status: 1, CreatedAt: time.Now().Add(-12 * time.Hour)},
		}
		DB.Create(&brs)
	}

	// 3. 营销中心种子数据
	DB.Model(&CouponTemplate{}).Count(&count)
	if count == 0 {
		coupons := []CouponTemplate{
			{ID: 1, Name: "新人立减 50 元券", Status: 0, DiscountType: 1, DiscountPrice: 5000, UsePrice: 20000, TotalCount: 1000, TakeCount: 120, CreatedAt: time.Now()},
			{ID: 2, Name: "全场 9 折折扣券", Status: 0, DiscountType: 2, DiscountPercent: 90, UsePrice: 10000, TotalCount: 500, TakeCount: 88, CreatedAt: time.Now()},
		}
		DB.Create(&coupons)

		userCoupon := CouponUser{
			ID:            101,
			TemplateID:    1,
			UserID:        1001,
			Name:          "新人立减 50 元券",
			Status:        1,
			DiscountPrice: 5000,
			UsePrice:      20000,
			CreatedAt:     time.Now(),
		}
		DB.Create(&userCoupon)
	}

	DB.Model(&DiyTemplate{}).Count(&count)
	if count == 0 {
		diy := DiyTemplate{
			ID:   1,
			Name: "默认宠物商城模板",
			Used: true,
			Property: `{"navigationBar":{"title":"我的宠物商城","backgroundColor":"#FF6B6B"},"components":[{"id":"banner","property":{"items":[{"imgUrl":"https://img.iocoder.cn/mall/banner1.jpg","url":"/pages/goods/detail?id=1"}]}}]}`,
			CreatedAt: time.Now(),
		}
		DB.Create(&diy)
	}

	DB.Model(&SeckillConfig{}).Count(&count)
	if count == 0 {
		configs := []SeckillConfig{
			{ID: 1, Name: "早间 10:00 秒杀场", StartTime: "10:00:00", EndTime: "12:00:00", Status: 0, CreatedAt: time.Now()},
			{ID: 2, Name: "晚间 20:00 秒杀场", StartTime: "20:00:00", EndTime: "22:00:00", Status: 0, CreatedAt: time.Now()},
		}
		DB.Create(&configs)
	}

	// 4. 会员中心种子数据
	DB.Model(&MemberUser{}).Count(&count)
	if count == 0 {
		user := MemberUser{
			ID:         1001,
			Nickname:   "微信用户小王",
			Avatar:     "https://www.iocoder.cn/images/avatar.jpg",
			Mobile:     "13800138000",
			Status:     0,
			LevelID:    2,
			LevelName:  "黄金会员",
			Point:      650,
			Experience: 1800,
			Balance:    50000,
			CreatedAt:  time.Now(),
		}
		DB.Create(&user)

		address := MemberAddress{
			ID:            1,
			UserID:        1001,
			Name:          "小王",
			Mobile:        "13800138000",
			AreaID:        110105,
			AreaName:      "北京市 朝阳区",
			DetailAddress: "三里屯 8 号院 3 号楼 602",
			DefaultStatus: true,
			CreatedAt:     time.Now(),
		}
		DB.Create(&address)

		wallet := PayWallet{
			ID:            1,
			UserID:        1001,
			Balance:       50000,
			TotalRecharge: 100000,
			TotalExpense:  50000,
			CreatedAt:     time.Now(),
		}
		DB.Create(&wallet)

		brokerage := TradeBrokerageUser{
			ID:             1,
			UserID:         1001,
			BindUserID:     0,
			BrokeragePrice: 12000,
			FrozenPrice:    0,
			TotalBrokerage: 35000,
			Enabled:        true,
			BindTime:       time.Now(),
			CreatedAt:      time.Now(),
		}
		DB.Create(&brokerage)
	}

	DB.Model(&MemberLevel{}).Count(&count)
	if count == 0 {
		levels := []MemberLevel{
			{ID: 1, Name: "大众会员", Level: 1, DiscountPercent: 100, Experience: 0, Status: 0, CreatedAt: time.Now()},
			{ID: 2, Name: "黄金会员", Level: 2, DiscountPercent: 95, Experience: 1000, Status: 0, CreatedAt: time.Now()},
			{ID: 3, Name: "钻石会员", Level: 3, DiscountPercent: 90, Experience: 5000, Status: 0, CreatedAt: time.Now()},
		}
		DB.Create(&levels)
	}

	DB.Model(&MemberTag{}).Count(&count)
	if count == 0 {
		tags := []MemberTag{
			{ID: 1, Name: "高频复购用户", Status: 0, CreatedAt: time.Now()},
			{ID: 2, Name: "品质猫舍合作", Status: 0, CreatedAt: time.Now()},
		}
		DB.Create(&tags)
	}

	DB.Model(&MemberGroup{}).Count(&count)
	if count == 0 {
		groups := []MemberGroup{
			{ID: 1, Name: "核心VIP用户组", Status: 0, CreatedAt: time.Now()},
			{ID: 2, Name: "新客专享特惠组", Status: 0, CreatedAt: time.Now()},
		}
		DB.Create(&groups)
	}

	// 5. 支付中心种子数据
	DB.Model(&PayApp{}).Count(&count)
	if count == 0 {
		apps := []PayApp{
			{ID: 1, Name: "宠物商城微信小程序", AppKey: "wx_pet_mall", Status: 0, CreatedAt: time.Now()},
		}
		DB.Create(&apps)
	}

	// 6. 系统管理种子数据 (部门、用户、角色、字典)
	DB.Model(&SystemDept{}).Count(&count)
	if count == 0 {
		depts := []SystemDept{
			{ID: 1, ParentID: 0, Name: "宠物科技总公司", Sort: 1, Leader: "王总", Phone: "13800000001", Status: 0, CreatedAt: time.Now()},
			{ID: 2, ParentID: 1, Name: "技术研发部", Sort: 1, Leader: "李总监", Phone: "13800000002", Status: 0, CreatedAt: time.Now()},
			{ID: 3, ParentID: 1, Name: "商城运营部", Sort: 2, Leader: "张经理", Phone: "13800000003", Status: 0, CreatedAt: time.Now()},
		}
		DB.Create(&depts)
	}

	DB.Model(&SystemPost{}).Count(&count)
	if count == 0 {
		posts := []SystemPost{
			{ID: 1, Name: "运营专员", Code: "operator", Sort: 1, Status: 0, CreatedAt: time.Now()},
			{ID: 2, Name: "客服主管", Code: "kefu_lead", Sort: 2, Status: 0, CreatedAt: time.Now()},
		}
		DB.Create(&posts)
	}

	DB.Model(&SystemRole{}).Count(&count)
	if count == 0 {
		roles := []SystemRole{
			{ID: 1, Name: "超级管理员", Code: "super_admin", Sort: 1, Status: 0, CreatedAt: time.Now()},
			{ID: 2, Name: "商城运营员", Code: "mall_operator", Sort: 2, Status: 0, CreatedAt: time.Now()},
		}
		DB.Create(&roles)
	}

	DB.Model(&SystemUser{}).Count(&count)
	if count == 0 {
		admin := SystemUser{
			ID:        1,
			Username:  "admin",
			Nickname:  "超级管理员",
			Avatar:    "https://www.iocoder.cn/images/avatar.jpg",
			Mobile:    "15800000000",
			Email:     "admin@iocoder.cn",
			DeptID:    1,
			Status:    0,
			CreatedAt: time.Now(),
		}
		DB.Create(&admin)
	}

	DB.Model(&SystemDictData{}).Count(&count)
	if count == 0 {
		dicts := []SystemDictData{
			{ID: 1, DictType: "common_status", Label: "开启", Value: "0", ColorType: "success", Sort: 1, Status: 0, CreatedAt: time.Now()},
			{ID: 2, DictType: "common_status", Label: "禁用", Value: "1", ColorType: "danger", Sort: 2, Status: 0, CreatedAt: time.Now()},
			{ID: 3, DictType: "terminal", Label: "微信小程序", Value: "10", ColorType: "primary", Sort: 1, Status: 0, CreatedAt: time.Now()},
			{ID: 4, DictType: "terminal", Label: "微信公众号", Value: "20", ColorType: "success", Sort: 2, Status: 0, CreatedAt: time.Now()},
			{ID: 5, DictType: "terminal", Label: "H5网页", Value: "30", ColorType: "warning", Sort: 3, Status: 0, CreatedAt: time.Now()},
			{ID: 6, DictType: "trade_order_status", Label: "待付款", Value: "0", Sort: 1, Status: 0, CreatedAt: time.Now()},
			{ID: 7, DictType: "trade_order_status", Label: "待发货", Value: "10", Sort: 2, Status: 0, CreatedAt: time.Now()},
			{ID: 8, DictType: "trade_order_status", Label: "待收货", Value: "20", Sort: 3, Status: 0, CreatedAt: time.Now()},
			{ID: 9, DictType: "trade_order_status", Label: "已完成", Value: "30", Sort: 4, Status: 0, CreatedAt: time.Now()},
			{ID: 10, DictType: "trade_order_status", Label: "已取消", Value: "40", Sort: 5, Status: 0, CreatedAt: time.Now()},
		}
		DB.Create(&dicts)
	}

	DB.Model(&Banner{}).Count(&count)
	if count == 0 {
		banners := []Banner{
			{ID: 1, Title: "新品猫粮特惠上市", PicURL: "https://img.iocoder.cn/mall/banner1.jpg", URL: "/pages/goods/detail?id=1", Status: 0, Sort: 1, CreatedAt: time.Now()},
			{ID: 2, Title: "智能猫砂盆限时直降", PicURL: "https://img.iocoder.cn/mall/banner2.jpg", URL: "/pages/goods/detail?id=2", Status: 0, Sort: 2, CreatedAt: time.Now()},
		}
		DB.Create(&banners)
	}

	DB.Model(&ArticleCategory{}).Count(&count)
	if count == 0 {
		ac := ArticleCategory{ID: 1, Name: "科学养宠指南", Sort: 1, Status: 0, CreatedAt: time.Now()}
		DB.Create(&ac)
	}

	DB.Model(&Article{}).Count(&count)
	if count == 0 {
		articles := []Article{
			{ID: 1, CategoryID: 1, Title: "幼猫换粮过渡期如何平稳过渡？", PicURL: "https://img.iocoder.cn/mall/cat_food.jpg", Content: "<p>建议采用7天换粮法，逐日增加新粮比例...</p>", Status: 0, CreatedAt: time.Now()},
		}
		DB.Create(&articles)
	}

	DB.Model(&PayWalletRechargePackage{}).Count(&count)
	if count == 0 {
		packages := []PayWalletRechargePackage{
			{ID: 1, Name: "充 100 送 10 元", PayPrice: 10000, BonusPrice: 1000, Status: 0, CreatedAt: time.Now()},
			{ID: 2, Name: "充 300 送 50 元", PayPrice: 30000, BonusPrice: 5000, Status: 0, CreatedAt: time.Now()},
			{ID: 3, Name: "充 500 送 100 元", PayPrice: 50000, BonusPrice: 10000, Status: 0, CreatedAt: time.Now()},
		}
		DB.Create(&packages)
	}

	DB.Model(&MemberSignInConfig{}).Count(&count)
	if count == 0 {
		configs := []MemberSignInConfig{
			{ID: 1, Day: 1, Point: 10, Experience: 10, Status: 0, CreatedAt: time.Now()},
			{ID: 2, Day: 2, Point: 20, Experience: 20, Status: 0, CreatedAt: time.Now()},
			{ID: 3, Day: 3, Point: 30, Experience: 30, Status: 0, CreatedAt: time.Now()},
			{ID: 4, Day: 7, Point: 100, Experience: 100, Status: 0, CreatedAt: time.Now()},
		}
		DB.Create(&configs)
	}

	DB.Model(&TradeConfig{}).Count(&count)
	if count == 0 {
		tc := TradeConfig{
			ID:                         1,
			AfterSaleRefundReasons:     "七天无理由退款,商品破损,缺件漏发,与描述不符",
			AfterSaleReturnReasons:     "不喜欢/不想要,质量问题,尺寸/重量不符",
			DeliveryExpressFreeEnabled: true,
			DeliveryExpressFreePrice:   9900,
			DeliveryPickUpEnabled:      true,
			BrokerageEnabled:           true,
			BrokerageEnabledCondition:  1,
			BrokerageBindMode:          1,
			BrokeragePosterUrls:        "https://img.iocoder.cn/mall/poster1.jpg",
			BrokerageFirstPercent:      10,
			BrokerageSecondPercent:     5,
			BrokerageWithdrawMinPrice:  1000,
			BrokerageWithdrawFeePercent: 1,
			BrokerageFrozenDays:        7,
			BrokerageWithdrawTypes:     "1,2",
			TencentLbsKey:              "DEMO-TENCENT-KEY",
			PointTradeDeductRate:       100,
		}
		DB.Create(&tc)
	}

	DB.Model(&MemberConfig{}).Count(&count)
	if count == 0 {
		mc := MemberConfig{
			ID:                        1,
			PointTradeDeductEnable:    true,
			PointTradeDeductUnitPrice: 100,
			PointTradeDeductMaxPrice:  5000,
			PointTradeGivePoint:       10,
		}
		DB.Create(&mc)
	}

	DB.Model(&MemberExperienceRecord{}).Count(&count)
	if count == 0 {
		exps := []MemberExperienceRecord{
			{ID: 1, UserID: 1001, BizID: "1001", BizType: 1, Title: "用户注册奖励", Description: "新会员首次注册成功赠送经验值", Experience: 100, TotalExperience: 100, CreatedAt: time.Now()},
			{ID: 2, UserID: 1001, BizID: "2001", BizType: 2, Title: "订单支付完成", Description: "完成商品购买订单奖励经验值", Experience: 250, TotalExperience: 350, CreatedAt: time.Now()},
		}
		DB.Create(&exps)
	}

	DB.Model(&InfraFileConfig{}).Count(&count)
	if count == 0 {
		fcs := []InfraFileConfig{
			{ID: 1, Name: "本地存储客户端", Storage: 1, Master: true, Config: "{\"basePath\":\"/tmp/uploads\"}", Remark: "默认本地存储", CreatedAt: time.Now()},
			{ID: 2, Name: "S3 云存储客户端", Storage: 2, Master: false, Config: "{\"bucket\":\"mall-assets\"}", Remark: "备用对象存储", CreatedAt: time.Now()},
		}
		DB.Create(&fcs)
	}

	DB.Model(&InfraFile{}).Count(&count)
	if count == 0 {
		files := []InfraFile{
			{ID: 1, ConfigID: 1, Name: "banner1.png", Path: "/banners/banner1.png", URL: "https://img.iocoder.cn/mall/banner1.png", Type: "image/png", Size: 102400, CreatedAt: time.Now()},
			{ID: 2, ConfigID: 1, Name: "logo.png", Path: "/logo/logo.png", URL: "https://img.iocoder.cn/mall/logo.png", Type: "image/png", Size: 45600, CreatedAt: time.Now()},
		}
		DB.Create(&files)
	}

	DB.Model(&InfraJob{}).Count(&count)
	if count == 0 {
		jobs := []InfraJob{
			{ID: 1, Name: "订单未支付自动关闭任务", HandlerName: "orderAutoCloseJob", HandlerParam: "timeout=30", CronExpression: "0 0/10 * * * ?", RetryCount: 3, RetryInterval: 1000, MonitorTimeout: 60000, Status: 0, CreatedAt: time.Now()},
			{ID: 2, Name: "拼团超时自动退款任务", HandlerName: "combinationTimeoutJob", HandlerParam: "", CronExpression: "0 0/5 * * * ?", RetryCount: 3, RetryInterval: 1000, MonitorTimeout: 60000, Status: 0, CreatedAt: time.Now()},
		}
		DB.Create(&jobs)
	}

	DB.Model(&PointActivity{}).Count(&count)
	if count == 0 {
		pas := []PointActivity{
			{ID: 1, SpuID: 1, Name: "全价猫粮积分兑换盛典", Status: 0, Point: 500, Price: 2900, StartTime: time.Now(), EndTime: time.Now().AddDate(0, 1, 0), CreatedAt: time.Now()},
		}
		DB.Create(&pas)
	}

	DB.Model(&DiscountActivity{}).Count(&count)
	if count == 0 {
		das := []DiscountActivity{
			{ID: 1, Name: "全场满 200 减 30 限时折扣", Status: 0, StartTime: time.Now(), EndTime: time.Now().AddDate(0, 1, 0), CreatedAt: time.Now()},
		}
		DB.Create(&das)
	}

	DB.Model(&RewardActivity{}).Count(&count)
	if count == 0 {
		ras := []RewardActivity{
			{ID: 1, Name: "满额赠精选猫条罐头", Status: 0, StartTime: time.Now(), EndTime: time.Now().AddDate(0, 1, 0), CreatedAt: time.Now()},
		}
		DB.Create(&ras)
	}

	DB.Model(&BargainActivity{}).Count(&count)
	if count == 0 {
		bas := []BargainActivity{
			{ID: 1, SpuID: 2, Name: "智能自动投喂器砍价免费拿", Status: 0, BargainPrice: 9900, CreatedAt: time.Now()},
		}
		DB.Create(&bas)
	}

	DB.Model(&BargainRecord{}).Count(&count)
	if count == 0 {
		brs := []BargainRecord{
			{ID: 1, ActivityID: 1, UserID: 1001, SpuID: 2, SkuID: 1, BargainFirstPrice: 19900, BargainPrice: 12000, Status: 0, CreatedAt: time.Now()},
		}
		DB.Create(&brs)
	}

	DB.Model(&KefuConversation{}).Count(&count)
	if count == 0 {
		kcs := []KefuConversation{
			{ID: 1, UserID: 1001, UserNickname: "官方测试会员", UserAvatar: "https://www.iocoder.cn/images/avatar.jpg", LastMessage: "您好，请问幼猫粮何时发货？", Pinned: false, CreatedAt: time.Now()},
		}
		DB.Create(&kcs)
	}
}
