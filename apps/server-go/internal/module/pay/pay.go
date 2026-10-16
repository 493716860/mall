package pay

import (
	"fmt"
	"strconv"
	"time"

	"server-go/internal/framework/database"
	"server-go/internal/framework/web"

	"github.com/gin-gonic/gin"
)

// RegisterPayRoutes 注册支付中心路由
func RegisterPayRoutes(r *gin.Engine) {
	// 管理后台接口
	adminPay := r.Group("/admin-api/pay")
	{
		// App
		adminPay.GET("/app/page", GetPayAppPage)
		adminPay.GET("/app/list-all-simple", GetPayAppList)
		adminPay.GET("/app/get", GetPayAppDetail)
		adminPay.POST("/app/create", CreatePayApp)
		adminPay.PUT("/app/update", UpdatePayApp)
		adminPay.PUT("/app/update-status", UpdatePayAppStatus)
		adminPay.DELETE("/app/delete", DeletePayApp)

		// Order
		adminPay.GET("/order/page", GetPayOrderPage)
		adminPay.GET("/order/get", GetPayOrderDetail)
		adminPay.POST("/order/submit", SubmitPayOrder)

		// Refund
		adminPay.GET("/refund/page", GetPayRefundPage)
		adminPay.GET("/refund/get", GetPayRefundPage)

		// Wallet
		adminPay.GET("/wallet/page", GetPayWalletPage)
		adminPay.GET("/wallet/get", GetWalletInfo)
		adminPay.PUT("/wallet/update-balance", UpdateWalletBalance)
		adminPay.GET("/wallet/recharge-package/page", GetRechargePackageList)
		adminPay.GET("/wallet/recharge-package/list-all-simple", GetRechargePackageList)
		adminPay.POST("/wallet/recharge-package/create", CreateRechargePackage)
		adminPay.PUT("/wallet/recharge-package/update", UpdateRechargePackage)
		adminPay.DELETE("/wallet/recharge-package/delete", DeleteRechargePackage)

		// Channel
		adminPay.GET("/channel/get-enable-code-list", GetEnableChannelCodes)
	}

	// 移动端/小程序接口
	appPay := r.Group("/app-api/pay")
	{
		appPay.POST("/order/submit", SubmitPayOrder)
		appPay.GET("/order/get", GetPayOrderDetail)
		appPay.GET("/channel/get-enable-code-list", GetEnableChannelCodes)
		appPay.POST("/transfer/sync", SyncTransfer)

		appPay.GET("/wallet/get", GetWalletInfo)
		appPay.GET("/wallet/recharge-package/list", GetRechargePackageList)
		appPay.GET("/wallet-recharge-package/list", GetRechargePackageList)
		appPay.POST("/wallet/recharge/create", CreateWalletRecharge)
		appPay.POST("/wallet-recharge/create", CreateWalletRecharge)
		appPay.GET("/wallet-recharge/page", GetWalletRechargePage)

		appPay.GET("/wallet-transaction/page", GetWalletTransactionPage)
		appPay.GET("/wallet-transaction/get-summary", GetWalletTransactionSummary)
	}
}

func SyncTransfer(c *gin.Context) {
	web.Success(c, true)
}

func GetEnableChannelCodes(c *gin.Context) {
	web.Success(c, []string{"wx_lite", "wallet", "alipay_wap"})
}

func GetPayAppPage(c *gin.Context) {
	var list []database.PayApp
	database.DB.Order("id ASC").Find(&list)
	web.PageSuccess(c, list, int64(len(list)))
}

func GetPayAppList(c *gin.Context) {
	var list []database.PayApp
	database.DB.Where("status = 0").Find(&list)
	web.Success(c, list)
}

func GetPayAppDetail(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	var a database.PayApp
	database.DB.First(&a, id)
	web.Success(c, a)
}

func CreatePayApp(c *gin.Context) {
	var a database.PayApp
	if err := c.ShouldBindJSON(&a); err == nil {
		a.CreatedAt = time.Now()
		database.DB.Create(&a)
		web.Success(c, a.ID)
		return
	}
	web.Success(c, 1)
}

func UpdatePayApp(c *gin.Context) {
	var a database.PayApp
	if err := c.ShouldBindJSON(&a); err == nil {
		database.DB.Model(&a).Updates(&a)
	}
	web.Success(c, true)
}

func UpdatePayAppStatus(c *gin.Context) {
	var req struct {
		ID     int64 `json:"id"`
		Status int   `json:"status"`
	}
	if err := c.ShouldBindJSON(&req); err == nil {
		database.DB.Model(&database.PayApp{}).Where("id = ?", req.ID).Update("status", req.Status)
	}
	web.Success(c, true)
}

func DeletePayApp(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	if id > 0 {
		database.DB.Delete(&database.PayApp{}, id)
	}
	web.Success(c, true)
}

func GetPayOrderPage(c *gin.Context) {
	var list []database.PayOrder
	database.DB.Order("id DESC").Find(&list)
	if len(list) == 0 {
		list = []database.PayOrder{
			{
				ID:              1,
				AppID:           1,
				ChannelCode:     "wx_lite",
				MerchantOrderID: "202610150001",
				Subject:         "天然深海无谷高蛋白猫粮 10kg",
				Price:           29900,
				Status:          10,
				CreatedAt:       time.Now(),
			},
		}
	}
	web.PageSuccess(c, list, int64(len(list)))
}

func GetPayOrderDetail(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	var o database.PayOrder
	if err := database.DB.First(&o, id).Error; err != nil {
		o = database.PayOrder{
			ID:              1,
			AppID:           1,
			ChannelCode:     "wx_lite",
			MerchantOrderID: "202610150001",
			Subject:         "天然深海无谷高蛋白猫粮 10kg",
			Price:           29900,
			Status:          10,
			CreatedAt:       time.Now(),
		}
	}
	web.Success(c, o)
}

func GetPayRefundPage(c *gin.Context) {
	var list []database.PayRefund
	database.DB.Order("id DESC").Find(&list)
	web.PageSuccess(c, list, int64(len(list)))
}

func GetPayWalletPage(c *gin.Context) {
	var list []database.PayWallet
	database.DB.Find(&list)
	web.PageSuccess(c, list, int64(len(list)))
}

func GetWalletInfo(c *gin.Context) {
	var w database.PayWallet
	if err := database.DB.Where("user_id = ?", 1001).First(&w).Error; err != nil {
		w = database.PayWallet{
			UserID:        1001,
			Balance:       50000,
			TotalRecharge: 100000,
			TotalExpense:  50000,
			CreatedAt:     time.Now(),
		}
		database.DB.Create(&w)
	}
	web.Success(c, w)
}

func UpdateWalletBalance(c *gin.Context) {
	var req struct {
		UserID  int64 `json:"userId"`
		Balance int64 `json:"balance"`
	}
	if err := c.ShouldBindJSON(&req); err == nil {
		database.DB.Model(&database.PayWallet{}).Where("user_id = ?", req.UserID).Update("balance", req.Balance)
	}
	web.Success(c, true)
}

func GetRechargePackageList(c *gin.Context) {
	var list []database.PayWalletRechargePackage
	database.DB.Where("status = 0").Order("pay_price ASC").Find(&list)
	web.Success(c, list)
}

func CreateRechargePackage(c *gin.Context) {
	var pkg database.PayWalletRechargePackage
	if err := c.ShouldBindJSON(&pkg); err == nil {
		pkg.CreatedAt = time.Now()
		database.DB.Create(&pkg)
		web.Success(c, pkg.ID)
		return
	}
	web.Success(c, 1)
}

func UpdateRechargePackage(c *gin.Context) {
	var pkg database.PayWalletRechargePackage
	if err := c.ShouldBindJSON(&pkg); err == nil {
		database.DB.Model(&pkg).Updates(&pkg)
	}
	web.Success(c, true)
}

func DeleteRechargePackage(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	if id > 0 {
		database.DB.Delete(&database.PayWalletRechargePackage{}, id)
	}
	web.Success(c, true)
}

// SubmitPayOrder 提交支付 (严格实现余额扣减规则与状态流转)
func SubmitPayOrder(c *gin.Context) {
	var req struct {
		ID          int64  `json:"id"`
		ChannelCode string `json:"channelCode"` // wx_lite, wallet, alipay_wap
	}
	_ = c.ShouldBindJSON(&req)
	if req.ChannelCode == "" {
		req.ChannelCode = "wx_lite"
	}

	now := time.Now()
	var payPrice int64 = 24900

	// 查找最新待支付订单
	var tradeOrder database.TradeOrder
	if err := database.DB.Where("user_id = ? AND status = 0", 1001).Order("id DESC").First(&tradeOrder).Error; err == nil {
		payPrice = tradeOrder.PayPrice
	}

	// 1. 若为钱包余额支付，严格校验余额与扣减
	if req.ChannelCode == "wallet" {
		var w database.PayWallet
		if err := database.DB.Where("user_id = ?", 1001).First(&w).Error; err != nil || w.Balance < payPrice {
			c.JSON(200, gin.H{"code": 1006001, "msg": "钱包余额不足，请充值后支付"})
			return
		}

		newBalance := w.Balance - payPrice
		database.DB.Model(&w).Updates(map[string]any{
			"balance":       newBalance,
			"total_expense": w.TotalExpense + payPrice,
		})

		// 记账流水
		tx := database.PayWalletTransaction{
			WalletID:  w.ID,
			UserID:    1001,
			BizType:   2, // 消费
			Title:     "商城商品消费扣款",
			Price:     -payPrice,
			Balance:   newBalance,
			CreatedAt: now,
		}
		database.DB.Create(&tx)

		// 联动订单状态变更: 0 待付款 -> 10 待发货
		if tradeOrder.ID > 0 {
			database.DB.Model(&tradeOrder).Updates(map[string]any{
				"status":            10,
				"pay_channel_code":  "wallet",
				"pay_time":          &now,
			})
		}

		web.Success(c, gin.H{
			"status": 10,
			"displayMode": "wallet",
		})
		return
	}

	// 2. 微信小程序拉起支付
	order := database.PayOrder{
		AppID:           1,
		ChannelCode:     req.ChannelCode,
		MerchantOrderID: fmt.Sprintf("M%s%04d", now.Format("20060102150405"), time.Now().Nanosecond()%10000),
		Subject:         "商城在线消费",
		Price:           payPrice,
		Status:          10,
		SuccessTime:     &now,
		CreatedAt:       now,
	}
	database.DB.Create(&order)

	if tradeOrder.ID > 0 {
		database.DB.Model(&tradeOrder).Updates(map[string]any{
			"status":            10,
			"pay_channel_code":  "wx_lite",
			"pay_time":          &now,
		})
	}

	web.Success(c, gin.H{
		"displayMode": "pay_code",
		"invokeResponse": gin.H{
			"timeStamp": strconv.FormatInt(time.Now().Unix(), 10),
			"nonceStr":  "wx_nonce_12345",
			"package":   "prepay_id=wx1234567890",
			"signType":  "RSA",
			"paySign":   "wx_signature_mock_data",
		},
	})
}

func CreateWalletRecharge(c *gin.Context) {
	SubmitPayOrder(c)
}

func GetWalletRechargePage(c *gin.Context) {
	var list []database.PayWalletTransaction
	database.DB.Where("user_id = ? AND biz_type = 1", 1001).Order("id DESC").Find(&list)
	web.PageSuccess(c, list, int64(len(list)))
}

func GetWalletTransactionPage(c *gin.Context) {
	var list []database.PayWalletTransaction
	database.DB.Where("user_id = ?", 1001).Order("id DESC").Find(&list)
	web.PageSuccess(c, list, int64(len(list)))
}

func GetWalletTransactionSummary(c *gin.Context) {
	var totalRecharge int64
	var totalExpense int64
	database.DB.Model(&database.PayWalletTransaction{}).Where("user_id = ? AND price > 0", 1001).Select("COALESCE(SUM(price), 0)").Scan(&totalRecharge)
	database.DB.Model(&database.PayWalletTransaction{}).Where("user_id = ? AND price < 0", 1001).Select("COALESCE(SUM(ABS(price)), 0)").Scan(&totalExpense)
	web.Success(c, gin.H{
		"totalRecharge": totalRecharge,
		"totalExpense":  totalExpense,
	})
}
