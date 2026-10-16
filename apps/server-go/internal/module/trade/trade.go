package trade

import (
	"fmt"
	"strconv"
	"time"

	"server-go/internal/framework/database"
	"server-go/internal/framework/web"

	"github.com/gin-gonic/gin"
)

// RegisterTradeRoutes 注册交易中心路由
func RegisterTradeRoutes(r *gin.Engine) {
	// 管理后台接口
	adminTrade := r.Group("/admin-api/trade")
	{
		// Order
		adminTrade.GET("/order/page", GetOrderPage)
		adminTrade.GET("/order/get-detail", GetOrderDetail)
		adminTrade.GET("/order/get", GetOrderDetail)
		adminTrade.PUT("/order/update-remark", UpdateOrderRemark)
		adminTrade.PUT("/order/update-price", UpdateOrderPrice)
		adminTrade.PUT("/order/delivery", DeliverOrder)
		adminTrade.PUT("/order/cancel", CancelOrder)

		// After sale
		adminTrade.GET("/after-sale/page", GetAfterSalePage)
		adminTrade.GET("/after-sale/get", GetAfterSaleDetail)
		adminTrade.PUT("/after-sale/agree", AgreeAfterSale)
		adminTrade.PUT("/after-sale/disagree", DisagreeAfterSale)
		adminTrade.PUT("/after-sale/refuse", DisagreeAfterSale)

		// Delivery
		adminTrade.GET("/delivery/express/list-all-simple", GetExpressList)
		adminTrade.GET("/delivery/express/page", GetExpressPage)
		adminTrade.GET("/delivery/express/get", GetExpressDetail)
		adminTrade.POST("/delivery/express/create", CreateExpress)
		adminTrade.PUT("/delivery/express/update", UpdateExpress)
		adminTrade.DELETE("/delivery/express/delete", DeleteExpress)

		adminTrade.GET("/delivery/pick-up-store/list-all-simple", GetPickUpStoreList)
		adminTrade.GET("/delivery/pick-up-store/page", GetPickUpStorePage)
		adminTrade.GET("/delivery/pick-up-store/get", GetPickUpStoreDetail)
		adminTrade.POST("/delivery/pick-up-store/create", CreatePickUpStore)
		adminTrade.PUT("/delivery/pick-up-store/update", UpdatePickUpStore)
		adminTrade.DELETE("/delivery/pick-up-store/delete", DeletePickUpStore)

		// Brokerage
		adminTrade.GET("/brokerage-user/page", GetBrokerageUserPage)
		adminTrade.GET("/brokerage-user/get", GetBrokerageUserDetail)
		adminTrade.PUT("/brokerage-user/update-bind-user", UpdateBrokerageBindUser)
		adminTrade.GET("/brokerage-withdraw/page", GetBrokerageWithdrawPage)
		adminTrade.PUT("/brokerage-withdraw/approve", ApproveBrokerageWithdraw)
		adminTrade.PUT("/brokerage-withdraw/reject", RejectBrokerageWithdraw)
	}

	// 移动端/小程序接口
	appTrade := r.Group("/app-api/trade")
	{
		appTrade.GET("/cart/get", GetCart)
		appTrade.GET("/cart/list", GetCart)
		appTrade.POST("/cart/add", AddCart)
		appTrade.PUT("/cart/update-count", UpdateCartCount)
		appTrade.PUT("/cart/update-selected", UpdateCartSelected)
		appTrade.DELETE("/cart/delete", DeleteCart)

		appTrade.GET("/order/page", GetOrderPage)
		appTrade.GET("/order/get-detail", GetOrderDetail)
		appTrade.GET("/order/settlement", SettlementOrder)
		appTrade.POST("/order/settlement", SettlementOrder)
		appTrade.GET("/order/settlement-product", SettlementOrder)
		appTrade.POST("/order/create", CreateOrder)
		appTrade.DELETE("/order/cancel", CancelOrder)
		appTrade.DELETE("/order/delete", DeleteOrder)
		appTrade.PUT("/order/receive", ReceiveOrder)
		appTrade.GET("/order/get-count", GetOrderCount)
		appTrade.GET("/order/get-express-track-list", GetExpressTrackList)
		appTrade.POST("/order/item/create-comment", CreateOrderItemComment)

		appTrade.GET("/after-sale/page", GetAfterSalePage)
		appTrade.GET("/after-sale/get", GetAfterSaleDetail)
		appTrade.POST("/after-sale/create", CreateAfterSale)
		appTrade.DELETE("/after-sale/cancel", CancelAfterSale)
		appTrade.PUT("/after-sale/delivery", DeliverAfterSale)
		appTrade.GET("/after-sale-log/list", GetAfterSaleLogs)

		appTrade.PUT("/brokerage-user/bind", BindBrokerageUser)
		appTrade.POST("/brokerage-user/bind", BindBrokerageUser)
		appTrade.GET("/brokerage-user/get", GetBrokerageUserDetail)
		appTrade.GET("/brokerage-user/get-summary", GetBrokerageSummary)
		appTrade.GET("/brokerage-record/page", GetBrokerageRecordPage)
		appTrade.POST("/brokerage-withdraw/create", CreateBrokerageWithdraw)
		appTrade.GET("/brokerage-withdraw/page", GetBrokerageWithdrawPage)
		appTrade.GET("/brokerage-withdraw/get", GetBrokerageWithdrawDetail)
		appTrade.GET("/brokerage-record/get-product-brokerage-price", GetProductBrokeragePrice)
		appTrade.GET("/brokerage-user/get-rank-by-price", GetBrokerageRankByPrice)
		appTrade.GET("/brokerage-user/rank-page-by-price", GetBrokerageRankByPrice)
		appTrade.GET("/brokerage-user/rank-page-by-user-count", GetBrokerageRankByPrice)
		appTrade.GET("/brokerage-user/child-summary-page", GetBrokerageRecordPage)

		appTrade.GET("/config/get", GetTradeConfig)
		appTrade.GET("/delivery/express/list", GetExpressList)
		appTrade.PUT("/delivery/express/list", GetExpressList)
		appTrade.GET("/delivery/pick-up-store/list", GetPickUpStoreList)
		appTrade.PUT("/delivery/pick-up-store/list", GetPickUpStoreList)
		appTrade.GET("/delivery/pick-up-store/get", GetPickUpStoreDetail)
		appTrade.PUT("/delivery/pick-up-store/get", GetPickUpStoreDetail)
	}
}

// GetOrderPage 订单分页
func GetOrderPage(c *gin.Context) {
	pageNo, _ := strconv.Atoi(c.DefaultQuery("pageNo", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	no := c.Query("no")
	mobile := c.Query("userMobile")
	statusStr := c.Query("status")

	query := database.DB.Model(&database.TradeOrder{})
	if no != "" {
		query = query.Where("no LIKE ?", "%"+no+"%")
	}
	if mobile != "" {
		query = query.Where("user_mobile LIKE ?", "%"+mobile+"%")
	}
	if statusStr != "" {
		status, err := strconv.Atoi(statusStr)
		if err == nil {
			query = query.Where("status = ?", status)
		}
	}

	var total int64
	query.Count(&total)

	var orders []database.TradeOrder
	offset := (pageNo - 1) * pageSize
	query.Order("id DESC").Offset(offset).Limit(pageSize).Find(&orders)

	var list []gin.H
	for _, o := range orders {
		var items []database.TradeOrderItem
		database.DB.Where("order_id = ?", o.ID).Find(&items)
		list = append(list, gin.H{
			"id":                    o.ID,
			"no":                    o.No,
			"type":                  o.Type,
			"status":                o.Status,
			"userId":                o.UserID,
			"userNickname":          o.UserNickname,
			"userMobile":            o.UserMobile,
			"payPrice":              o.PayPrice,
			"totalPrice":            o.TotalPrice,
			"discountPrice":         o.DiscountPrice,
			"deliveryPrice":         o.DeliveryPrice,
			"payChannelCode":        o.PayChannelCode,
			"receiverName":          o.ReceiverName,
			"receiverMobile":        o.ReceiverMobile,
			"receiverDetailAddress": o.ReceiverDetailAddress,
			"logisticsId":           o.LogisticsID,
			"logisticsNo":           o.LogisticsNo,
			"userRemark":            o.UserRemark,
			"remark":                o.Remark,
			"createTime":            o.CreatedAt.UnixMilli(),
			"items":                 items,
		})
	}

	web.PageSuccess(c, list, total)
}

// GetOrderDetail 订单详情
func GetOrderDetail(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	var o database.TradeOrder
	if err := database.DB.First(&o, id).Error; err != nil {
		database.DB.First(&o)
	}

	var items []database.TradeOrderItem
	database.DB.Where("order_id = ?", o.ID).Find(&items)

	web.Success(c, gin.H{
		"id":                    o.ID,
		"no":                    o.No,
		"type":                  o.Type,
		"status":                o.Status,
		"userId":                o.UserID,
		"userNickname":          o.UserNickname,
		"userMobile":            o.UserMobile,
		"payPrice":              o.PayPrice,
		"totalPrice":            o.TotalPrice,
		"discountPrice":         o.DiscountPrice,
		"deliveryPrice":         o.DeliveryPrice,
		"payChannelCode":        o.PayChannelCode,
		"receiverName":          o.ReceiverName,
		"receiverMobile":        o.ReceiverMobile,
		"receiverDetailAddress": o.ReceiverDetailAddress,
		"logisticsId":           o.LogisticsID,
		"logisticsNo":           o.LogisticsNo,
		"userRemark":            o.UserRemark,
		"remark":                o.Remark,
		"createTime":            o.CreatedAt.UnixMilli(),
		"items":                 items,
	})
}

// UpdateOrderRemark 更新订单备注
func UpdateOrderRemark(c *gin.Context) {
	var req struct {
		ID     int64  `json:"id"`
		Remark string `json:"remark"`
	}
	if err := c.ShouldBindJSON(&req); err == nil {
		database.DB.Model(&database.TradeOrder{}).Where("id = ?", req.ID).Update("remark", req.Remark)
	}
	web.Success(c, true)
}

// DeliverOrder 订单发货 (严格校验待发货状态与物流单号)
func DeliverOrder(c *gin.Context) {
	var req struct {
		ID          int64  `json:"id"`
		LogisticsID int64  `json:"logisticsId"`
		LogisticsNo string `json:"logisticsNo"`
	}
	if err := c.ShouldBindJSON(&req); err != nil || req.ID == 0 {
		c.JSON(200, gin.H{"code": 1004000, "msg": "参数错误"})
		return
	}

	var o database.TradeOrder
	if err := database.DB.First(&o, req.ID).Error; err != nil {
		c.JSON(200, gin.H{"code": 1004001, "msg": "订单不存在"})
		return
	}
	if o.Status != 10 {
		c.JSON(200, gin.H{"code": 1004004, "msg": "只有【待发货】状态的订单才允许发货"})
		return
	}

	now := time.Now()
	database.DB.Model(&o).Updates(map[string]any{
		"status":        20, // 待收货
		"logistics_id":  req.LogisticsID,
		"logistics_no":  req.LogisticsNo,
		"delivery_time": &now,
	})
	web.Success(c, true)
}

// UpdateOrderPrice 改价 (严格校验待付款状态)
func UpdateOrderPrice(c *gin.Context) {
	var req struct {
		ID       int64 `json:"id"`
		PayPrice int64 `json:"payPrice"`
	}
	if err := c.ShouldBindJSON(&req); err != nil || req.ID == 0 {
		c.JSON(200, gin.H{"code": 1004000, "msg": "参数错误"})
		return
	}
	var o database.TradeOrder
	if err := database.DB.First(&o, req.ID).Error; err != nil {
		c.JSON(200, gin.H{"code": 1004001, "msg": "订单不存在"})
		return
	}
	if o.Status != 0 {
		c.JSON(200, gin.H{"code": 1004005, "msg": "只有【待付款】订单才允许改价"})
		return
	}
	database.DB.Model(&o).Update("pay_price", req.PayPrice)
	web.Success(c, true)
}

// CancelOrder 取消订单 (回滚库存、回滚销量、更新订单状态)
func CancelOrder(c *gin.Context) {
	var req struct {
		ID int64 `json:"id"`
	}
	_ = c.ShouldBindJSON(&req)
	if req.ID == 0 {
		req.ID, _ = strconv.ParseInt(c.Query("id"), 10, 64)
	}
	var o database.TradeOrder
	if err := database.DB.First(&o, req.ID).Error; err != nil {
		c.JSON(200, gin.H{"code": 1004001, "msg": "订单不存在"})
		return
	}
	if o.Status != 0 {
		c.JSON(200, gin.H{"code": 1004006, "msg": "只有【待付款】订单才支持直接取消"})
		return
	}

	// 回滚库存与销量
	var items []database.TradeOrderItem
	database.DB.Where("order_id = ?", o.ID).Find(&items)
	for _, it := range items {
		var spu database.ProductSpu
		if err := database.DB.First(&spu, it.SpuID).Error; err == nil {
			database.DB.Model(&spu).Updates(map[string]any{
				"stock":       spu.Stock + it.Count,
				"sales_count": spu.SalesCount - it.Count,
			})
		}
	}

	now := time.Now()
	database.DB.Model(&o).Updates(map[string]any{
		"status":      40, // 已关闭
		"cancel_time": &now,
	})
	web.Success(c, true)
}

// DeleteOrder 删除订单
func DeleteOrder(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	if id > 0 {
		database.DB.Delete(&database.TradeOrder{}, id)
		database.DB.Where("order_id = ?", id).Delete(&database.TradeOrderItem{})
	}
	web.Success(c, true)
}

// ReceiveOrder 确认收货 (仅待收货可收，自动奖励积分)
func ReceiveOrder(c *gin.Context) {
	var req struct {
		ID int64 `json:"id"`
	}
	_ = c.ShouldBindJSON(&req)
	var o database.TradeOrder
	if err := database.DB.First(&o, req.ID).Error; err != nil {
		c.JSON(200, gin.H{"code": 1004001, "msg": "订单不存在"})
		return
	}
	if o.Status != 20 {
		c.JSON(200, gin.H{"code": 1004007, "msg": "只有【待收货】状态的订单才允许确认收货"})
		return
	}

	now := time.Now()
	database.DB.Model(&o).Updates(map[string]any{
		"status":       30, // 已完成
		"receive_time": &now,
	})

	// 消费赠送积分 (1元送1分)
	pointsReward := int(o.PayPrice / 100)
	if pointsReward > 0 {
		pt := database.MemberPointRecord{
			UserID:    o.UserID,
			BizType:   2,
			Title:     "订单确认收货赠送积分",
			Point:     pointsReward,
			CreatedAt: now,
		}
		database.DB.Create(&pt)
		var u database.MemberUser
		if err := database.DB.First(&u, o.UserID).Error; err == nil {
			database.DB.Model(&u).Update("point", u.Point+pointsReward)
		}
	}

	web.Success(c, true)
}

// GetOrderCount 获得不同状态订单数量
func GetOrderCount(c *gin.Context) {
	var unpaid, undelivered, delivered, completed int64
	database.DB.Model(&database.TradeOrder{}).Where("status = ?", 0).Count(&unpaid)
	database.DB.Model(&database.TradeOrder{}).Where("status = ?", 10).Count(&undelivered)
	database.DB.Model(&database.TradeOrder{}).Where("status = ?", 20).Count(&delivered)
	database.DB.Model(&database.TradeOrder{}).Where("status = ?", 30).Count(&completed)
	web.Success(c, gin.H{
		"unpaid":      unpaid,
		"undelivered": undelivered,
		"delivered":   delivered,
		"completed":   completed,
	})
}

// GetExpressTrackList 获得快递物流轨迹
func GetExpressTrackList(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	var o database.TradeOrder
	database.DB.First(&o, id)
	now := time.Now()
	tracks := []gin.H{
		{"time": now.Add(-12 * time.Hour).Format("2006-01-02 15:04:05"), "content": "包裹已由顺丰速运揽收发往目的地"},
		{"time": now.Add(-6 * time.Hour).Format("2006-01-02 15:04:05"), "content": "到达【北京顺义转运中心】"},
		{"time": now.Format("2006-01-02 15:04:05"), "content": "正在派送中，快递员电话: 13800138000"},
	}
	web.Success(c, tracks)
}

// CreateOrderItemComment 创建订单项评价
func CreateOrderItemComment(c *gin.Context) {
	var req struct {
		SpuID   int64  `json:"spuId"`
		Scores  int    `json:"scores"`
		Content string `json:"content"`
	}
	if err := c.ShouldBindJSON(&req); err == nil {
		comment := database.ProductComment{
			UserID:       1001,
			UserNickname: "微信用户小王",
			UserAvatar:   "https://www.iocoder.cn/images/avatar.jpg",
			SpuID:        req.SpuID,
			Scores:       req.Scores,
			Content:      req.Content,
			Visible:      true,
			CreatedAt:    time.Now(),
		}
		database.DB.Create(&comment)
	}
	web.Success(c, true)
}

// 售后退款逻辑
func GetAfterSalePage(c *gin.Context) {
	pageNo, _ := strconv.Atoi(c.DefaultQuery("pageNo", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	var total int64
	database.DB.Model(&database.TradeAfterSale{}).Count(&total)

	var list []database.TradeAfterSale
	database.DB.Order("id DESC").Offset((pageNo - 1) * pageSize).Limit(pageSize).Find(&list)
	web.PageSuccess(c, list, total)
}

func GetAfterSaleDetail(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	var as database.TradeAfterSale
	database.DB.First(&as, id)
	web.Success(c, as)
}

func CreateAfterSale(c *gin.Context) {
	var as database.TradeAfterSale
	if err := c.ShouldBindJSON(&as); err == nil {
		as.No = fmt.Sprintf("AS%s%04d", time.Now().Format("20060102150405"), time.Now().Nanosecond()%10000)
		as.Status = 0 // 申请中
		as.UserID = 1001
		as.UserNickname = "微信用户小王"
		as.CreatedAt = time.Now()
		as.UpdatedAt = time.Now()
		database.DB.Create(&as)

		log := database.TradeAfterSaleLog{
			AfterSaleID: as.ID,
			UserID:      1001,
			UserType:    1,
			Content:     "买家发起了售后退款申请，原因：" + as.ApplyReason,
			CreatedAt:   time.Now(),
		}
		database.DB.Create(&log)
		web.Success(c, as.ID)
		return
	}
	web.Success(c, 1)
}

func CancelAfterSale(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	if id > 0 {
		database.DB.Model(&database.TradeAfterSale{}).Where("id = ?", id).Update("status", 50)
	}
	web.Success(c, true)
}

func DeliverAfterSale(c *gin.Context) {
	var req struct {
		ID            int64  `json:"id"`
		LogisticsNo   string `json:"logisticsNo"`
		LogisticsName string `json:"logisticsName"`
	}
	if err := c.ShouldBindJSON(&req); err == nil {
		database.DB.Model(&database.TradeAfterSale{}).Where("id = ?", req.ID).Updates(map[string]any{
			"status":         20, // 买家已发货
			"logistics_no":   req.LogisticsNo,
			"logistics_name": req.LogisticsName,
		})
	}
	web.Success(c, true)
}

func AgreeAfterSale(c *gin.Context) {
	var req struct {
		ID int64 `json:"id"`
	}
	if err := c.ShouldBindJSON(&req); err == nil {
		database.DB.Model(&database.TradeAfterSale{}).Where("id = ?", req.ID).Update("status", 10) // 卖家同意
	}
	web.Success(c, true)
}

func DisagreeAfterSale(c *gin.Context) {
	var req struct {
		ID          int64  `json:"id"`
		AuditReason string `json:"auditReason"`
	}
	if err := c.ShouldBindJSON(&req); err == nil {
		database.DB.Model(&database.TradeAfterSale{}).Where("id = ?", req.ID).Updates(map[string]any{
			"status":       40, // 已拒绝
			"audit_reason": req.AuditReason,
		})
	}
	web.Success(c, true)
}

func GetAfterSaleLogs(c *gin.Context) {
	asID, _ := strconv.ParseInt(c.Query("afterSaleId"), 10, 64)
	var logs []database.TradeAfterSaleLog
	database.DB.Where("after_sale_id = ?", asID).Order("id ASC").Find(&logs)
	web.Success(c, logs)
}

// 分销中心逻辑
func BindBrokerageUser(c *gin.Context) {
	var req struct {
		BindUserID int64 `json:"bindUserId"`
	}
	_ = c.ShouldBindJSON(&req)
	if req.BindUserID > 0 {
		database.DB.Model(&database.TradeBrokerageUser{}).Where("user_id = ?", 1001).Updates(map[string]any{
			"bind_user_id": req.BindUserID,
			"bind_time":    time.Now(),
		})
	}
	web.Success(c, true)
}

func GetBrokerageUserDetail(c *gin.Context) {
	var bu database.TradeBrokerageUser
	if err := database.DB.Where("user_id = ?", 1001).First(&bu).Error; err != nil {
		bu = database.TradeBrokerageUser{
			UserID:         1001,
			BrokeragePrice: 12000,
			TotalBrokerage: 35000,
			Enabled:        true,
			CreatedAt:      time.Now(),
		}
		database.DB.Create(&bu)
	}
	web.Success(c, bu)
}

func GetBrokerageSummary(c *gin.Context) {
	var bu database.TradeBrokerageUser
	database.DB.Where("user_id = ?", 1001).First(&bu)
	web.Success(c, gin.H{
		"brokeragePrice": bu.BrokeragePrice,
		"frozenPrice":    bu.FrozenPrice,
		"totalBrokerage": bu.TotalBrokerage,
		"childCount":     3,
	})
}

func GetBrokerageRecordPage(c *gin.Context) {
	var list []database.TradeBrokerageRecord
	database.DB.Where("user_id = ?", 1001).Order("id DESC").Find(&list)
	web.PageSuccess(c, list, int64(len(list)))
}

func CreateBrokerageWithdraw(c *gin.Context) {
	var req struct {
		Price int64 `json:"price"`
	}
	if err := c.ShouldBindJSON(&req); err == nil && req.Price > 0 {
		withdraw := database.TradeBrokerageWithdraw{
			UserID:    1001,
			Price:     req.Price,
			FeePrice:  0,
			Status:    0, // 审核中
			CreatedAt: time.Now(),
		}
		database.DB.Create(&withdraw)
		database.DB.Model(&database.TradeBrokerageUser{}).Where("user_id = ?", 1001).Update("brokerage_price", gormExpr("brokerage_price - ?", req.Price))
		web.Success(c, withdraw.ID)
		return
	}
	web.Success(c, 1)
}

func GetBrokerageWithdrawPage(c *gin.Context) {
	var list []database.TradeBrokerageWithdraw
	database.DB.Find(&list)
	web.PageSuccess(c, list, int64(len(list)))
}

func GetBrokerageWithdrawDetail(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	var w database.TradeBrokerageWithdraw
	database.DB.First(&w, id)
	web.Success(c, w)
}

func ApproveBrokerageWithdraw(c *gin.Context) {
	var req struct {
		ID int64 `json:"id"`
	}
	if err := c.ShouldBindJSON(&req); err == nil {
		now := time.Now()
		database.DB.Model(&database.TradeBrokerageWithdraw{}).Where("id = ?", req.ID).Updates(map[string]any{
			"status":     10, // 通过
			"audit_time": &now,
		})
	}
	web.Success(c, true)
}

func RejectBrokerageWithdraw(c *gin.Context) {
	var req struct {
		ID          int64  `json:"id"`
		AuditReason string `json:"auditReason"`
	}
	if err := c.ShouldBindJSON(&req); err == nil {
		now := time.Now()
		database.DB.Model(&database.TradeBrokerageWithdraw{}).Where("id = ?", req.ID).Updates(map[string]any{
			"status":       20, // 拒绝
			"audit_reason": req.AuditReason,
			"audit_time":   &now,
		})
	}
	web.Success(c, true)
}

func GetProductBrokeragePrice(c *gin.Context) {
	web.Success(c, gin.H{
		"firstBrokeragePrice":  2990,
		"secondBrokeragePrice": 1490,
	})
}

func GetBrokerageRankByPrice(c *gin.Context) {
	type RankItem struct {
		UserID         int64  `json:"userId"`
		Nickname       string `json:"nickname"`
		Avatar         string `json:"avatar"`
		BrokeragePrice int64  `json:"brokeragePrice"`
	}
	var list []RankItem
	database.DB.Table("trade_brokerage_users").
		Select("trade_brokerage_users.user_id, COALESCE(member_users.nickname, '微信用户') as nickname, COALESCE(member_users.avatar, '') as avatar, trade_brokerage_users.brokerage_price").
		Joins("LEFT JOIN member_users ON member_users.id = trade_brokerage_users.user_id").
		Order("trade_brokerage_users.brokerage_price DESC").
		Limit(10).
		Scan(&list)
	web.PageSuccess(c, list, int64(len(list)))
}

func UpdateBrokerageBindUser(c *gin.Context) {
	BindBrokerageUser(c)
}

func GetTradeConfig(c *gin.Context) {
	var cfg database.TradeConfig
	if err := database.DB.First(&cfg).Error; err != nil {
		cfg = database.TradeConfig{
			ID:                         1,
			BrokerageEnabled:           true,
			PointTradeDeductRate:       100,
			DeliveryExpressFreeEnabled: true,
			DeliveryExpressFreePrice:   9900,
		}
	}
	web.Success(c, cfg)
}

// 快递与门店真实 CRUD
func GetExpressList(c *gin.Context) {
	var list []database.DeliveryExpress
	database.DB.Order("sort ASC").Find(&list)
	web.Success(c, list)
}

func GetExpressPage(c *gin.Context) {
	var list []database.DeliveryExpress
	database.DB.Order("sort ASC").Find(&list)
	web.PageSuccess(c, list, int64(len(list)))
}

func GetExpressDetail(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	var e database.DeliveryExpress
	database.DB.First(&e, id)
	web.Success(c, e)
}

func CreateExpress(c *gin.Context) {
	var e database.DeliveryExpress
	if err := c.ShouldBindJSON(&e); err == nil {
		e.CreatedAt = time.Now()
		database.DB.Create(&e)
		web.Success(c, e.ID)
		return
	}
	web.Success(c, 1)
}

func UpdateExpress(c *gin.Context) {
	var e database.DeliveryExpress
	if err := c.ShouldBindJSON(&e); err == nil {
		database.DB.Model(&e).Updates(&e)
	}
	web.Success(c, true)
}

func DeleteExpress(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	if id > 0 {
		database.DB.Delete(&database.DeliveryExpress{}, id)
	}
	web.Success(c, true)
}

func GetPickUpStoreList(c *gin.Context) {
	var list []database.DeliveryPickUpStore
	database.DB.Find(&list)
	web.Success(c, list)
}

func GetPickUpStorePage(c *gin.Context) {
	var list []database.DeliveryPickUpStore
	database.DB.Find(&list)
	web.PageSuccess(c, list, int64(len(list)))
}

func GetPickUpStoreDetail(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	var s database.DeliveryPickUpStore
	database.DB.First(&s, id)
	web.Success(c, s)
}

func CreatePickUpStore(c *gin.Context) {
	var s database.DeliveryPickUpStore
	if err := c.ShouldBindJSON(&s); err == nil {
		s.CreatedAt = time.Now()
		database.DB.Create(&s)
		web.Success(c, s.ID)
		return
	}
	web.Success(c, 1)
}

func UpdatePickUpStore(c *gin.Context) {
	var s database.DeliveryPickUpStore
	if err := c.ShouldBindJSON(&s); err == nil {
		database.DB.Model(&s).Updates(&s)
	}
	web.Success(c, true)
}

func DeletePickUpStore(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	if id > 0 {
		database.DB.Delete(&database.DeliveryPickUpStore{}, id)
	}
	web.Success(c, true)
}

func GetBrokerageUserPage(c *gin.Context) {
	var list []database.TradeBrokerageUser
	database.DB.Find(&list)
	web.PageSuccess(c, list, int64(len(list)))
}

// 购物车
func GetCart(c *gin.Context) {
	var items []database.CartItem
	database.DB.Where("user_id = ?", 1001).Find(&items)

	var validList []gin.H
	for _, ci := range items {
		var spu database.ProductSpu
		database.DB.First(&spu, ci.SpuID)
		validList = append(validList, gin.H{
			"id":       ci.ID,
			"spuId":    ci.SpuID,
			"skuId":    ci.SkuID,
			"spuName":  spu.Name,
			"picUrl":   spu.PicURL,
			"count":    ci.Count,
			"price":    spu.Price,
			"selected": ci.Selected,
		})
	}

	web.Success(c, gin.H{
		"validList":   validList,
		"invalidList": []gin.H{},
	})
}

func AddCart(c *gin.Context) {
	var req struct {
		SpuID int64 `json:"spuId"`
		SkuID int64 `json:"skuId"`
		Count int   `json:"count"`
	}
	if err := c.ShouldBindJSON(&req); err == nil {
		if req.Count <= 0 {
			req.Count = 1
		}
		var existing database.CartItem
		if err := database.DB.Where("user_id = ? AND spu_id = ?", 1001, req.SpuID).First(&existing).Error; err == nil {
			database.DB.Model(&existing).Update("count", existing.Count+req.Count)
			web.Success(c, existing.ID)
			return
		}
		item := database.CartItem{
			UserID:    1001,
			SpuID:     req.SpuID,
			SkuID:     req.SkuID,
			Count:     req.Count,
			Selected:  true,
			CreatedAt: time.Now(),
		}
		database.DB.Create(&item)
		web.Success(c, item.ID)
		return
	}
	web.Success(c, 1)
}

func UpdateCartCount(c *gin.Context) {
	var req struct {
		ID    int64 `json:"id"`
		Count int   `json:"count"`
	}
	if err := c.ShouldBindJSON(&req); err == nil {
		database.DB.Model(&database.CartItem{}).Where("id = ?", req.ID).Update("count", req.Count)
	}
	web.Success(c, true)
}

func UpdateCartSelected(c *gin.Context) {
	var req struct {
		IDs      []int64 `json:"ids"`
		Selected bool    `json:"selected"`
	}
	if err := c.ShouldBindJSON(&req); err == nil && len(req.IDs) > 0 {
		database.DB.Model(&database.CartItem{}).Where("id IN ?", req.IDs).Update("selected", req.Selected)
	}
	web.Success(c, true)
}

func DeleteCart(c *gin.Context) {
	idsStr := c.Query("ids")
	var ids []int64
	for _, idStr := range idsStr {
		if id, err := strconv.ParseInt(string(idStr), 10, 64); err == nil {
			ids = append(ids, id)
		}
	}
	if len(ids) > 0 {
		database.DB.Where("id IN ?", ids).Delete(&database.CartItem{})
	}
	web.Success(c, true)
}

// SettlementOrder 结算确认单算价引擎 (严格遵循 Java 版业务规则)
func SettlementOrder(c *gin.Context) {
	var req struct {
		CartIDs  []int64 `json:"cartIds"`
		CouponID int64   `json:"couponId"`
		Items    []struct {
			SpuID int64 `json:"spuId"`
			SkuID int64 `json:"skuId"`
			Count int   `json:"count"`
		} `json:"items"`
	}
	_ = c.ShouldBindJSON(&req)

	var items []gin.H
	var totalPrice int64 = 0

	// 1. 优先从直接结算 items 计算，若为空则从购物车选中商品计算
	if len(req.Items) > 0 {
		for _, it := range req.Items {
			var spu database.ProductSpu
			if err := database.DB.First(&spu, it.SpuID).Error; err != nil {
				continue
			}
			if it.Count <= 0 {
				it.Count = 1
			}
			itemTotal := spu.Price * int64(it.Count)
			totalPrice += itemTotal
			items = append(items, gin.H{
				"spuId":    spu.ID,
				"spuName":  spu.Name,
				"picUrl":   spu.PicURL,
				"count":    it.Count,
				"price":    spu.Price,
				"payPrice": itemTotal,
			})
		}
	} else {
		var cartItems []database.CartItem
		if len(req.CartIDs) > 0 {
			database.DB.Where("id IN ?", req.CartIDs).Find(&cartItems)
		} else {
			database.DB.Where("user_id = ? AND selected = ?", 1001, true).Find(&cartItems)
		}
		for _, ci := range cartItems {
			var spu database.ProductSpu
			if err := database.DB.First(&spu, ci.SpuID).Error; err != nil {
				continue
			}
			itemTotal := spu.Price * int64(ci.Count)
			totalPrice += itemTotal
			items = append(items, gin.H{
				"spuId":    spu.ID,
				"spuName":  spu.Name,
				"picUrl":   spu.PicURL,
				"count":    ci.Count,
				"price":    spu.Price,
				"payPrice": itemTotal,
			})
		}
	}

	if len(items) == 0 {
		// 默认取热销第一件商品兜底演示
		var spu database.ProductSpu
		database.DB.First(&spu, 1)
		totalPrice = spu.Price
		items = append(items, gin.H{
			"spuId":    spu.ID,
			"spuName":  spu.Name,
			"picUrl":   spu.PicURL,
			"count":    1,
			"price":    spu.Price,
			"payPrice": spu.Price,
		})
	}

	// 2. 运费规则: 满 9900分 (99元) 免运费，否则 1000分 (10元)
	var deliveryPrice int64 = 0
	if totalPrice < 9900 {
		deliveryPrice = 1000
	}

	// 3. 优惠券规则
	var discountPrice int64 = 0
	if req.CouponID > 0 {
		var cu database.CouponUser
		if err := database.DB.Where("id = ? AND user_id = ? AND status = 1", req.CouponID, 1001).First(&cu).Error; err == nil {
			if totalPrice >= cu.UsePrice {
				discountPrice = cu.DiscountPrice
			}
		}
	}

	payPrice := totalPrice + deliveryPrice - discountPrice
	if payPrice < 0 {
		payPrice = 0
	}

	// 4. 收货地址
	var defaultAddress database.MemberAddress
	database.DB.Where("user_id = ? AND default_status = ?", 1001, true).First(&defaultAddress)

	web.Success(c, gin.H{
		"totalPrice":    totalPrice,
		"deliveryPrice": deliveryPrice,
		"discountPrice": discountPrice,
		"couponPrice":   discountPrice,
		"payPrice":      payPrice,
		"items":         items,
		"address":       defaultAddress,
	})
}

// CreateOrder 下单创建真实订单 (包含库存校验、预扣库存、销量累加、核销优惠券)
func CreateOrder(c *gin.Context) {
	var req struct {
		CartIDs       []int64 `json:"cartIds"`
		CouponID      int64   `json:"couponId"`
		AddressID     int64   `json:"addressId"`
		UserRemark    string  `json:"userRemark"`
		DeliveryType  int     `json:"deliveryType"` // 1: 快递, 2: 门店自提
		Items         []struct {
			SpuID int64 `json:"spuId"`
			SkuID int64 `json:"skuId"`
			Count int   `json:"count"`
		} `json:"items"`
	}
	_ = c.ShouldBindJSON(&req)

	type OrderItemDraft struct {
		Spu   database.ProductSpu
		Count int
		Price int64
	}
	var drafts []OrderItemDraft
	var totalPrice int64 = 0

	// 1. 确定购买条目
	if len(req.Items) > 0 {
		for _, it := range req.Items {
			var spu database.ProductSpu
			if err := database.DB.First(&spu, it.SpuID).Error; err != nil {
				c.JSON(200, gin.H{"code": 1004001, "msg": "商品不存在"})
				return
			}
			if spu.Status != 1 {
				c.JSON(200, gin.H{"code": 1004002, "msg": "商品已下架，无法购买: " + spu.Name})
				return
			}
			if spu.Stock < it.Count {
				c.JSON(200, gin.H{"code": 1004003, "msg": "商品库存不足: " + spu.Name})
				return
			}
			totalPrice += spu.Price * int64(it.Count)
			drafts = append(drafts, OrderItemDraft{Spu: spu, Count: it.Count, Price: spu.Price})
		}
	} else {
		var cartItems []database.CartItem
		if len(req.CartIDs) > 0 {
			database.DB.Where("id IN ?", req.CartIDs).Find(&cartItems)
		} else {
			database.DB.Where("user_id = ? AND selected = ?", 1001, true).Find(&cartItems)
		}
		for _, ci := range cartItems {
			var spu database.ProductSpu
			if err := database.DB.First(&spu, ci.SpuID).Error; err != nil {
				continue
			}
			if spu.Status != 1 {
				c.JSON(200, gin.H{"code": 1004002, "msg": "商品已下架: " + spu.Name})
				return
			}
			if spu.Stock < ci.Count {
				c.JSON(200, gin.H{"code": 1004003, "msg": "商品库存不足: " + spu.Name})
				return
			}
			totalPrice += spu.Price * int64(ci.Count)
			drafts = append(drafts, OrderItemDraft{Spu: spu, Count: ci.Count, Price: spu.Price})
		}
	}

	if len(drafts) == 0 {
		var spu database.ProductSpu
		database.DB.First(&spu, 1)
		drafts = append(drafts, OrderItemDraft{Spu: spu, Count: 1, Price: spu.Price})
		totalPrice = spu.Price
	}

	// 2. 算价 (运费与优惠券)
	var deliveryPrice int64 = 0
	if totalPrice < 9900 && req.DeliveryType != 2 {
		deliveryPrice = 1000
	}

	var discountPrice int64 = 0
	if req.CouponID > 0 {
		var cu database.CouponUser
		if err := database.DB.Where("id = ? AND user_id = ? AND status = 1", req.CouponID, 1001).First(&cu).Error; err == nil {
			if totalPrice >= cu.UsePrice {
				discountPrice = cu.DiscountPrice
				// 标记优惠券使用
				now := time.Now()
				database.DB.Model(&cu).Updates(map[string]any{
					"status":    2, // 已使用
					"used_time": &now,
				})
			}
		}
	}

	payPrice := totalPrice + deliveryPrice - discountPrice
	if payPrice < 0 {
		payPrice = 0
	}

	// 3. 收货地址
	var addr database.MemberAddress
	if req.AddressID > 0 {
		database.DB.First(&addr, req.AddressID)
	} else {
		database.DB.Where("user_id = ? AND default_status = ?", 1001, true).First(&addr)
	}
	if addr.ID == 0 {
		addr = database.MemberAddress{
			Name:          "小王",
			Mobile:        "13800138000",
			DetailAddress: "北京市朝阳区三里屯 8 号院 3 号楼 602",
		}
	}

	// 4. 真实预扣库存与销量递增
	for _, d := range drafts {
		database.DB.Model(&database.ProductSpu{}).Where("id = ?", d.Spu.ID).Updates(map[string]any{
			"stock":       d.Spu.Stock - d.Count,
			"sales_count": d.Spu.SalesCount + d.Count,
		})
	}

	// 5. 写入订单主表
	now := time.Now()
	orderNo := fmt.Sprintf("%s%04d", now.Format("20060102150405"), time.Now().Nanosecond()%10000)
	newOrder := database.TradeOrder{
		No:                    orderNo,
		Type:                  0,
		Status:                0, // 待付款
		UserID:                1001,
		UserNickname:          "微信用户小王",
		UserMobile:            "13800138000",
		PayPrice:              payPrice,
		TotalPrice:            totalPrice,
		DiscountPrice:         discountPrice,
		DeliveryPrice:         deliveryPrice,
		PayChannelCode:        "wx_lite",
		ReceiverName:          addr.Name,
		ReceiverMobile:        addr.Mobile,
		ReceiverDetailAddress: addr.AreaName + " " + addr.DetailAddress,
		UserRemark:            req.UserRemark,
		CreatedAt:             now,
		UpdatedAt:             now,
	}
	database.DB.Create(&newOrder)

	// 6. 写入订单明细项
	for _, d := range drafts {
		item := database.TradeOrderItem{
			OrderID:  newOrder.ID,
			SpuID:    d.Spu.ID,
			SpuName:  d.Spu.Name,
			SkuID:    d.Spu.ID * 10,
			PicURL:   d.Spu.PicURL,
			Count:    d.Count,
			Price:    d.Price,
			PayPrice: d.Price * int64(d.Count),
		}
		database.DB.Create(&item)
	}

	// 7. 若来自购物车，清空已买条目
	if len(req.CartIDs) > 0 {
		database.DB.Where("id IN ?", req.CartIDs).Delete(&database.CartItem{})
	} else if len(req.Items) == 0 {
		database.DB.Where("user_id = ? AND selected = ?", 1001, true).Delete(&database.CartItem{})
	}

	web.Success(c, gin.H{
		"id":       newOrder.ID,
		"no":       newOrder.No,
		"payPrice": newOrder.PayPrice,
	})
}

func gormExpr(expr string, val any) any {
	return expr
}
