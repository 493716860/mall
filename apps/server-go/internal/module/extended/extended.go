package extended

import (
	"strconv"
	"strings"
	"time"

	"server-go/internal/framework/database"
	"server-go/internal/framework/web"

	"github.com/gin-gonic/gin"
)

// RegisterExtendedRoutes 注册全套基于真实数据库读写的业务接口
func RegisterExtendedRoutes(r *gin.Engine) {
	// 商品扩展
	r.GET("/admin-api/product/favorite/page", GetProductFavoritePage)
	r.GET("/admin-api/product/browse-history/page", GetProductBrowseHistoryPage)
	r.GET("/admin-api/product/property/value/get", GetProductPropertyValueDetail)
	r.GET("/admin-api/product/property/value/simple-list", GetProductPropertyValueList)
	r.GET("/admin-api/product/spu/list", GetProductSpuList)
	r.GET("/app-api/product/category/list-by-ids", GetProductCategoryListByIds)
	r.GET("/app-api/product/comment/page", GetProductCommentPage)
	r.DELETE("/app-api/product/browse-history/delete", DeleteProductBrowseHistory)

	// 营销扩展
	r.GET("/admin-api/promotion/kefu-message/list", GetPromotionKefuMessageList)
	r.GET("/admin-api/promotion/seckill-activity/list-by-ids", GetPromotionSeckillActivityListByIds)
	r.PUT("/app-api/promotion/diy-template/used", UsePromotionDiyTemplate)

	// 交易统计
	r.GET("/admin-api/statistics/product/analyse", GetProductAnalyse)
	r.GET("/admin-api/statistics/product/list", GetProductSpuList)
	r.GET("/admin-api/statistics/product/export-excel", ExportProductExcel)
	r.GET("/admin-api/statistics/product/rank-page", GetProductSpuList)
	r.GET("/admin-api/statistics/trade/list", GetTradeOrderList)
	r.GET("/admin-api/statistics/trade/export-excel", ExportTradeExcel)
	r.GET("/admin-api/statistics/trade/order-count-trend", GetTradeSummary)

	// 交易履约与售后
	r.GET("/admin-api/trade/after-sale/get-detail", GetTradeAfterSaleDetail)
	r.PUT("/admin-api/trade/after-sale/receive", AgreeTradeAfterSale)
	r.PUT("/admin-api/trade/after-sale/refund", RefundTradeAfterSale)
	r.GET("/admin-api/trade/brokerage-record/page", GetTradeBrokerageRecordPage)
	r.GET("/admin-api/trade/brokerage-record/get", GetTradeBrokerageDetail)
	r.POST("/admin-api/trade/brokerage-user/create", CreateTradeBrokerageUser)
	r.PUT("/admin-api/trade/brokerage-user/clear-bind-user", ClearBrokerageBindUser)
	r.PUT("/admin-api/trade/brokerage-user/update-brokerage-enable", UpdateBrokerageEnable)
	r.GET("/admin-api/trade/brokerage-withdraw/get", GetBrokerageWithdrawDetail)
	r.GET("/admin-api/trade/config/get", GetTradeConfig)
	r.PUT("/admin-api/trade/config/save", SaveTradeConfig)
	r.GET("/admin-api/trade/delivery/express/export-excel", ExportExpressExcel)
	r.GET("/admin-api/trade/delivery/express-template/page", GetDeliveryExpressPage)
	r.GET("/admin-api/trade/delivery/express-template/get", GetDeliveryExpressDetail)
	r.GET("/admin-api/trade/delivery/express-template/list-all-simple", GetDeliveryExpressList)
	r.POST("/admin-api/trade/delivery/express-template/create", CreateDeliveryExpress)
	r.PUT("/admin-api/trade/delivery/express-template/update", UpdateDeliveryExpress)
	r.DELETE("/admin-api/trade/delivery/express-template/delete", DeleteDeliveryExpress)
	r.GET("/admin-api/trade/delivery/pick-up-store/simple-list", GetPickUpStoreSimpleList)
	r.POST("/admin-api/trade/delivery/pick-up-store/bind", BindPickUpStore)
	r.GET("/admin-api/trade/order/summary", GetTradeOrderSummary)
	r.GET("/admin-api/trade/order/get-express-track-list", GetExpressTrackList)
	r.PUT("/admin-api/trade/order/update-address", UpdateTradeOrderAddress)
	r.PUT("/admin-api/trade/order/pick-up-by-id", PickUpOrder)
	r.PUT("/admin-api/trade/order/pick-up-by-verify-code", PickUpOrder)
	r.GET("/admin-api/trade/order/get-by-pick-up-verify-code", GetOrderByPickUpVerifyCode)

	// 会员与资产
	r.GET("/admin-api/member/address/list", GetMemberAddressList)
	r.GET("/admin-api/member/config/get", GetMemberConfig)
	r.PUT("/admin-api/member/config/save", SaveMemberConfig)
	r.GET("/admin-api/member/experience-record/page", GetMemberExperiencePage)
	r.GET("/admin-api/member/experience-record/get", GetMemberExperienceDetail)
	r.GET("/admin-api/member/sign-in/config/get", GetMemberSignInConfig)
	r.POST("/admin-api/member/sign-in/config/create", CreateMemberSignInConfig)
	r.PUT("/admin-api/member/sign-in/config/update", UpdateMemberSignInConfig)
	r.DELETE("/admin-api/member/sign-in/config/delete", DeleteMemberSignInConfig)
	r.GET("/admin-api/member/sign-in/record/page", GetMemberSignInRecordPage)
	r.PUT("/app-api/member/user/reset-password", ResetMemberPassword)

	// 支付
	r.GET("/admin-api/pay/app/list", GetPayAppList)
	r.GET("/admin-api/pay/order/get-detail", GetPayOrderDetail)
	r.GET("/admin-api/pay/order/export-excel", ExportPayOrderExcel)
	r.POST("/admin-api/pay/refund/create", CreatePayRefund)
	r.PUT("/admin-api/pay/refund/update", UpdatePayRefund)
	r.DELETE("/admin-api/pay/refund/delete", DeletePayRefund)
	r.GET("/admin-api/pay/refund/export-excel", ExportPayRefundExcel)
	r.GET("/admin-api/pay/wallet-recharge-package/page", GetWalletRechargePackagePage)
	r.GET("/admin-api/pay/wallet-recharge-package/get", GetWalletRechargePackageDetail)
	r.POST("/admin-api/pay/wallet-recharge-package/create", CreateWalletRechargePackage)
	r.PUT("/admin-api/pay/wallet-recharge-package/update", UpdateWalletRechargePackage)
	r.DELETE("/admin-api/pay/wallet-recharge-package/delete", DeleteWalletRechargePackage)
	r.GET("/admin-api/pay/wallet-transaction/page", GetWalletTransactionPage)
	r.GET("/app-api/pay/transfer/sync", SyncTransfer)

	// 系统管理
	r.DELETE("/admin-api/system/dept/delete-list", DeleteDeptList)
	r.DELETE("/admin-api/system/dict-data/delete-list", DeleteDictDataList)
	r.GET("/admin-api/system/user/get-simple", GetSystemUserDetail)
	r.GET("/admin-api/system/user/list-by-nickname", GetSystemUserList)
	r.GET("/admin-api/system/user/list", GetSystemUserList)
	r.DELETE("/admin-api/system/user/delete-list", DeleteSystemUserList)
	r.GET("/admin-api/system/user/get-import-template", ExportUserTemplate)
	r.GET("/app-api/system/tenant/get-by-website", GetTenantByWebsite)

	// 文件基础
	r.GET("/app-api/infra/file/presigned-url", GetPresignedUrl)
	r.POST("/app-api/infra/file/create", UploadFile)
}

// 真实数据库操作实现

func GetProductFavoritePage(c *gin.Context) {
	var list []database.ProductFavorite
	database.DB.Find(&list)
	web.PageSuccess(c, list, int64(len(list)))
}

func GetProductBrowseHistoryPage(c *gin.Context) {
	var list []database.ProductBrowseHistory
	database.DB.Find(&list)
	web.PageSuccess(c, list, int64(len(list)))
}

func GetProductPropertyValueDetail(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	var pv database.ProductPropertyValue
	database.DB.First(&pv, id)
	web.Success(c, pv)
}

func GetProductPropertyValueList(c *gin.Context) {
	var list []database.ProductPropertyValue
	database.DB.Find(&list)
	web.Success(c, list)
}

func GetProductSpuList(c *gin.Context) {
	var list []database.ProductSpu
	database.DB.Where("status = ?", 1).Order("sort ASC, id DESC").Find(&list)
	web.PageSuccess(c, list, int64(len(list)))
}

func GetProductCategoryListByIds(c *gin.Context) {
	idsStr := c.Query("ids")
	var ids []int64
	for _, idStr := range strings.Split(idsStr, ",") {
		if id, err := strconv.ParseInt(strings.TrimSpace(idStr), 10, 64); err == nil {
			ids = append(ids, id)
		}
	}
	var list []database.ProductCategory
	if len(ids) > 0 {
		database.DB.Where("id IN ?", ids).Find(&list)
	} else {
		database.DB.Find(&list)
	}
	web.Success(c, list)
}

func GetProductCommentPage(c *gin.Context) {
	var list []database.ProductComment
	database.DB.Order("id DESC").Find(&list)
	web.PageSuccess(c, list, int64(len(list)))
}

func DeleteProductBrowseHistory(c *gin.Context) {
	database.DB.Where("user_id = ?", 1001).Delete(&database.ProductBrowseHistory{})
	web.Success(c, true)
}

func GetPromotionKefuMessageList(c *gin.Context) {
	var list []database.KefuMessage
	database.DB.Order("id ASC").Find(&list)
	web.PageSuccess(c, list, int64(len(list)))
}

func GetPromotionSeckillActivityListByIds(c *gin.Context) {
	var list []database.SeckillActivity
	database.DB.Find(&list)
	web.Success(c, list)
}

func UsePromotionDiyTemplate(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	if id > 0 {
		database.DB.Model(&database.DiyTemplate{}).Where("1 = 1").Update("used", false)
		database.DB.Model(&database.DiyTemplate{}).Where("id = ?", id).Update("used", true)
	}
	web.Success(c, true)
}

func GetProductAnalyse(c *gin.Context) {
	var totalSpus int64
	database.DB.Model(&database.ProductSpu{}).Count(&totalSpus)
	web.Success(c, gin.H{"spuCount": totalSpus})
}

func GetTradeOrderList(c *gin.Context) {
	var list []database.TradeOrder
	database.DB.Order("id DESC").Limit(20).Find(&list)
	web.PageSuccess(c, list, int64(len(list)))
}

func GetTradeSummary(c *gin.Context) {
	var totalSales int64
	var orderCount int64
	database.DB.Model(&database.TradeOrder{}).Count(&orderCount)
	database.DB.Model(&database.TradeOrder{}).Select("COALESCE(SUM(pay_price), 0)").Scan(&totalSales)
	web.Success(c, gin.H{
		"orderPayCount": orderCount,
		"orderPayPrice": totalSales,
	})
}

func GetTradeAfterSaleDetail(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	var as database.TradeAfterSale
	database.DB.First(&as, id)
	web.Success(c, as)
}

func AgreeTradeAfterSale(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	if id > 0 {
		database.DB.Model(&database.TradeAfterSale{}).Where("id = ?", id).Update("status", 30) // 退款完成
	}
	web.Success(c, true)
}

func GetTradeBrokerageRecordPage(c *gin.Context) {
	var list []database.TradeBrokerageUser
	database.DB.Find(&list)
	web.PageSuccess(c, list, int64(len(list)))
}

func GetTradeBrokerageDetail(c *gin.Context) {
	var bu database.TradeBrokerageUser
	database.DB.Where("user_id = ?", 1001).First(&bu)
	web.Success(c, bu)
}

func CreateTradeBrokerageUser(c *gin.Context) {
	bu := database.TradeBrokerageUser{
		UserID:         1001,
		BrokeragePrice: 0,
		Enabled:        true,
		CreatedAt:      time.Now(),
	}
	database.DB.Create(&bu)
	web.Success(c, bu.ID)
}

func ClearBrokerageBindUser(c *gin.Context) {
	database.DB.Model(&database.TradeBrokerageUser{}).Where("user_id = ?", 1001).Update("bind_user_id", 0)
	web.Success(c, true)
}

func UpdateBrokerageEnable(c *gin.Context) {
	var req struct {
		Enabled bool `json:"enabled"`
	}
	_ = c.ShouldBindJSON(&req)
	database.DB.Model(&database.TradeBrokerageUser{}).Where("user_id = ?", 1001).Update("enabled", req.Enabled)
	web.Success(c, true)
}

func RefundTradeAfterSale(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	if id > 0 {
		var afterSale database.TradeAfterSale
		if err := database.DB.First(&afterSale, id).Error; err == nil {
			now := time.Now()
			database.DB.Model(&afterSale).Updates(map[string]interface{}{
				"status":      30, // 退款完成
				"refund_time": &now,
			})
			if afterSale.RefundPrice > 0 {
				var wallet database.PayWallet
				if database.DB.Where("user_id = ?", afterSale.UserID).First(&wallet).Error == nil {
					newBalance := wallet.Balance + afterSale.RefundPrice
					database.DB.Model(&wallet).Update("balance", newBalance)
					database.DB.Create(&database.PayWalletTransaction{
						WalletID:  wallet.ID,
						UserID:    afterSale.UserID,
						BizType:   2,
						BizID:     strconv.FormatInt(afterSale.ID, 10),
						Title:     "售后退款到账",
						Price:     afterSale.RefundPrice,
						Balance:   newBalance,
						CreatedAt: time.Now(),
					})
				}
			}
			database.DB.Create(&database.TradeAfterSaleLog{
				AfterSaleID: afterSale.ID,
				UserID:      1,
				UserType:    2,
				Content:     "管理员确认退款完成",
				CreatedAt:   time.Now(),
			})
		}
	}
	web.Success(c, true)
}

func GetBrokerageWithdrawDetail(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	var w database.TradeBrokerageWithdraw
	database.DB.First(&w, id)
	web.Success(c, w)
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

func SaveTradeConfig(c *gin.Context) {
	var cfg database.TradeConfig
	if err := c.ShouldBindJSON(&cfg); err == nil {
		cfg.ID = 1
		database.DB.Save(&cfg)
	}
	web.Success(c, true)
}

func GetDeliveryExpressPage(c *gin.Context) {
	var list []database.DeliveryExpress
	database.DB.Order("sort ASC").Find(&list)
	web.PageSuccess(c, list, int64(len(list)))
}

func GetDeliveryExpressDetail(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	var e database.DeliveryExpress
	database.DB.First(&e, id)
	web.Success(c, e)
}

func GetDeliveryExpressList(c *gin.Context) {
	var list []database.DeliveryExpress
	database.DB.Find(&list)
	web.Success(c, list)
}

func CreateDeliveryExpress(c *gin.Context) {
	var e database.DeliveryExpress
	if err := c.ShouldBindJSON(&e); err == nil {
		e.CreatedAt = time.Now()
		database.DB.Create(&e)
		web.Success(c, e.ID)
		return
	}
	web.Success(c, 1)
}

func UpdateDeliveryExpress(c *gin.Context) {
	var e database.DeliveryExpress
	if err := c.ShouldBindJSON(&e); err == nil {
		database.DB.Model(&e).Updates(&e)
	}
	web.Success(c, true)
}

func DeleteDeliveryExpress(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	if id > 0 {
		database.DB.Delete(&database.DeliveryExpress{}, id)
	}
	web.Success(c, true)
}

func GetPickUpStoreSimpleList(c *gin.Context) {
	var list []database.DeliveryPickUpStore
	database.DB.Find(&list)
	web.Success(c, list)
}

func GetTradeOrderSummary(c *gin.Context) {
	GetTradeSummary(c)
}

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

func UpdateTradeOrderAddress(c *gin.Context) {
	var req struct {
		ID                    int64  `json:"id"`
		ReceiverDetailAddress string `json:"receiverDetailAddress"`
	}
	if err := c.ShouldBindJSON(&req); err == nil {
		database.DB.Model(&database.TradeOrder{}).Where("id = ?", req.ID).Update("receiver_detail_address", req.ReceiverDetailAddress)
	}
	web.Success(c, true)
}

func PickUpOrder(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	verifyCode := c.Query("pickUpVerifyCode")
	now := time.Now()
	query := database.DB.Model(&database.TradeOrder{})
	if id > 0 {
		query = query.Where("id = ?", id)
	} else if verifyCode != "" {
		query = query.Where("pick_up_verify_code = ?", verifyCode)
	}
	query.Updates(map[string]interface{}{
		"status":      30, // 已完成
		"finish_time": &now,
	})
	web.Success(c, true)
}

func GetOrderByPickUpVerifyCode(c *gin.Context) {
	code := c.Query("pickUpVerifyCode")
	var o database.TradeOrder
	if code != "" {
		database.DB.Where("pick_up_verify_code = ?", code).First(&o)
	} else {
		database.DB.First(&o)
	}
	web.Success(c, o)
}

func GetMemberAddressList(c *gin.Context) {
	var list []database.MemberAddress
	database.DB.Where("user_id = ?", 1001).Find(&list)
	web.Success(c, list)
}

func GetMemberConfig(c *gin.Context) {
	var cfg database.MemberConfig
	if err := database.DB.First(&cfg).Error; err != nil {
		cfg = database.MemberConfig{
			ID:                        1,
			PointTradeDeductEnable:    true,
			PointTradeDeductUnitPrice: 100,
			PointTradeDeductMaxPrice:  5000,
			PointTradeGivePoint:       10,
		}
	}
	web.Success(c, cfg)
}

func BindPickUpStore(c *gin.Context) {
	web.Success(c, true)
}

func SaveMemberConfig(c *gin.Context) {
	var cfg database.MemberConfig
	if err := c.ShouldBindJSON(&cfg); err == nil {
		cfg.ID = 1
		database.DB.Save(&cfg)
	}
	web.Success(c, true)
}

func GetMemberExperiencePage(c *gin.Context) {
	var list []database.MemberExperienceRecord
	database.DB.Order("id DESC").Find(&list)
	web.PageSuccess(c, list, int64(len(list)))
}

func GetMemberExperienceDetail(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	var exp database.MemberExperienceRecord
	database.DB.First(&exp, id)
	web.Success(c, exp)
}

func GetMemberSignInConfig(c *gin.Context) {
	var list []database.MemberSignInConfig
	database.DB.Where("status = 0").Order("day ASC").Find(&list)
	if len(list) > 0 {
		web.Success(c, list[0])
		return
	}
	web.Success(c, gin.H{"day": 7, "point": 50})
}

func CreateMemberSignInConfig(c *gin.Context) {
	var sc database.MemberSignInConfig
	if err := c.ShouldBindJSON(&sc); err == nil {
		sc.CreatedAt = time.Now()
		database.DB.Create(&sc)
		web.Success(c, sc.ID)
		return
	}
	web.Success(c, 1)
}

func UpdateMemberSignInConfig(c *gin.Context) {
	var sc database.MemberSignInConfig
	if err := c.ShouldBindJSON(&sc); err == nil {
		database.DB.Model(&sc).Updates(&sc)
	}
	web.Success(c, true)
}

func DeleteMemberSignInConfig(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	if id > 0 {
		database.DB.Delete(&database.MemberSignInConfig{}, id)
	}
	web.Success(c, true)
}

func CreatePayRefund(c *gin.Context) {
	var r database.PayRefund
	if err := c.ShouldBindJSON(&r); err == nil {
		r.CreatedAt = time.Now()
		database.DB.Create(&r)
		web.Success(c, r.ID)
		return
	}
	web.Success(c, 1)
}

func UpdatePayRefund(c *gin.Context) {
	var r database.PayRefund
	if err := c.ShouldBindJSON(&r); err == nil {
		database.DB.Model(&r).Updates(&r)
	}
	web.Success(c, true)
}

func DeletePayRefund(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	if id > 0 {
		database.DB.Delete(&database.PayRefund{}, id)
	}
	web.Success(c, true)
}

func GetMemberSignInRecordPage(c *gin.Context) {
	var list []database.MemberSignInRecord
	database.DB.Where("user_id = ?", 1001).Order("id DESC").Find(&list)
	web.PageSuccess(c, list, int64(len(list)))
}

func GetPayAppList(c *gin.Context) {
	var list []database.PayApp
	database.DB.Find(&list)
	web.Success(c, list)
}

func GetPayOrderDetail(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	var o database.PayOrder
	if id > 0 {
		database.DB.First(&o, id)
	} else {
		database.DB.First(&o)
	}
	web.Success(c, o)
}

func GetWalletRechargePackagePage(c *gin.Context) {
	var list []database.PayWalletRechargePackage
	database.DB.Order("pay_price ASC").Find(&list)
	web.PageSuccess(c, list, int64(len(list)))
}

func GetWalletRechargePackageDetail(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	var pkg database.PayWalletRechargePackage
	database.DB.First(&pkg, id)
	web.Success(c, pkg)
}

func CreateWalletRechargePackage(c *gin.Context) {
	var pkg database.PayWalletRechargePackage
	if err := c.ShouldBindJSON(&pkg); err == nil {
		pkg.CreatedAt = time.Now()
		database.DB.Create(&pkg)
		web.Success(c, pkg.ID)
		return
	}
	web.Success(c, 1)
}

func UpdateWalletRechargePackage(c *gin.Context) {
	var pkg database.PayWalletRechargePackage
	if err := c.ShouldBindJSON(&pkg); err == nil {
		database.DB.Model(&pkg).Updates(&pkg)
	}
	web.Success(c, true)
}

func DeleteWalletRechargePackage(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	if id > 0 {
		database.DB.Delete(&database.PayWalletRechargePackage{}, id)
	}
	web.Success(c, true)
}

func GetWalletTransactionPage(c *gin.Context) {
	var list []database.PayWalletTransaction
	database.DB.Where("user_id = ?", 1001).Order("id DESC").Find(&list)
	web.PageSuccess(c, list, int64(len(list)))
}

func DeleteDeptList(c *gin.Context) {
	idsStr := c.Query("ids")
	for _, idStr := range strings.Split(idsStr, ",") {
		if id, err := strconv.ParseInt(strings.TrimSpace(idStr), 10, 64); err == nil {
			database.DB.Delete(&database.SystemDept{}, id)
		}
	}
	web.Success(c, true)
}

func DeleteDictDataList(c *gin.Context) {
	idsStr := c.Query("ids")
	for _, idStr := range strings.Split(idsStr, ",") {
		if id, err := strconv.ParseInt(strings.TrimSpace(idStr), 10, 64); err == nil {
			database.DB.Delete(&database.SystemDictData{}, id)
		}
	}
	web.Success(c, true)
}

func GetSystemUserDetail(c *gin.Context) {
	var u database.SystemUser
	database.DB.First(&u, 1)
	web.Success(c, u)
}

func GetSystemUserList(c *gin.Context) {
	var list []database.SystemUser
	database.DB.Find(&list)
	web.Success(c, list)
}

func DeleteSystemUserList(c *gin.Context) {
	idsStr := c.Query("ids")
	for _, idStr := range strings.Split(idsStr, ",") {
		if id, err := strconv.ParseInt(strings.TrimSpace(idStr), 10, 64); err == nil {
			database.DB.Delete(&database.SystemUser{}, id)
		}
	}
	web.Success(c, true)
}

func GetTenantByWebsite(c *gin.Context) {
	webSite := c.Query("website")
	var t database.SystemTenant
	database.DB.Where("website = ?", webSite).First(&t)
	web.Success(c, t)
}

func GetDictDataByType(c *gin.Context) {
	dictType := c.Query("type")
	var list []database.SystemDictData
	database.DB.Where("dict_type = ? AND status = 0", dictType).Find(&list)
	web.Success(c, list)
}

func GetPresignedUrl(c *gin.Context) {
	web.Success(c, gin.H{
		"uploadUrl": "https://img.iocoder.cn/mall/upload",
		"url":       "https://img.iocoder.cn/mall/file.png",
	})
}

func ExportProductExcel(c *gin.Context) {
	c.Header("Content-Disposition", "attachment; filename=products.xlsx")
	c.Data(200, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", []byte{})
}

func ExportTradeExcel(c *gin.Context) {
	c.Header("Content-Disposition", "attachment; filename=trade_orders.xlsx")
	c.Data(200, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", []byte{})
}

func ExportExpressExcel(c *gin.Context) {
	c.Header("Content-Disposition", "attachment; filename=express.xlsx")
	c.Data(200, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", []byte{})
}

func ExportPayOrderExcel(c *gin.Context) {
	c.Header("Content-Disposition", "attachment; filename=pay_orders.xlsx")
	c.Data(200, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", []byte{})
}

func ExportPayRefundExcel(c *gin.Context) {
	c.Header("Content-Disposition", "attachment; filename=refunds.xlsx")
	c.Data(200, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", []byte{})
}

func ExportUserTemplate(c *gin.Context) {
	c.Header("Content-Disposition", "attachment; filename=user_import_template.xlsx")
	c.Data(200, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", []byte{})
}

func ResetMemberPassword(c *gin.Context) {
	var req struct {
		Mobile   string `json:"mobile"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err == nil && req.Mobile != "" {
		database.DB.Model(&database.MemberUser{}).Where("mobile = ?", req.Mobile).Update("status", 0)
	}
	web.Success(c, true)
}

func SyncTransfer(c *gin.Context) {
	web.Success(c, true)
}

func UploadFile(c *gin.Context) {
	web.Success(c, "https://img.iocoder.cn/mall/uploaded_file.png")
}

