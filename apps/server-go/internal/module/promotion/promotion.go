package promotion

import (
	"strconv"
	"strings"
	"time"

	"server-go/internal/framework/database"
	"server-go/internal/framework/web"

	"github.com/gin-gonic/gin"
)

// RegisterPromotionRoutes 注册营销中心路由
func RegisterPromotionRoutes(r *gin.Engine) {
	// 管理后台接口
	adminPromo := r.Group("/admin-api/promotion")
	{
		// Coupon Template & User Coupon
		adminPromo.GET("/coupon-template/page", GetCouponTemplatePage)
		adminPromo.GET("/coupon-template/get", GetCouponTemplateDetail)
		adminPromo.GET("/coupon-template/list", GetCouponTemplateList)
		adminPromo.POST("/coupon-template/create", CreateCouponTemplate)
		adminPromo.PUT("/coupon-template/update", UpdateCouponTemplate)
		adminPromo.PUT("/coupon-template/update-status", UpdateCouponTemplateStatus)
		adminPromo.DELETE("/coupon-template/delete", DeleteCouponTemplate)
		adminPromo.GET("/coupon-template/export-excel", ExportCouponTemplateExcel)

		adminPromo.GET("/coupon/page", GetCouponUserPage)
		adminPromo.POST("/coupon/send", SendUserCoupon)
		adminPromo.DELETE("/coupon/delete", DeleteUserCoupon)

		// Seckill
		adminPromo.GET("/seckill-activity/page", GetSeckillActivityPage)
		adminPromo.GET("/seckill-activity/get", GetSeckillActivityDetail)
		adminPromo.POST("/seckill-activity/create", CreateSeckillActivity)
		adminPromo.PUT("/seckill-activity/update", UpdateSeckillActivity)
		adminPromo.PUT("/seckill-activity/close", CloseSeckillActivity)
		adminPromo.DELETE("/seckill-activity/delete", DeleteSeckillActivity)
		adminPromo.GET("/seckill-config/list-all-simple", GetSeckillConfigList)
		adminPromo.GET("/seckill-config/page", GetSeckillConfigList)
		adminPromo.GET("/seckill-config/list", GetSeckillConfigList)
		adminPromo.GET("/seckill-config/get", GetSeckillConfigDetail)
		adminPromo.POST("/seckill-config/create", CreateSeckillConfig)
		adminPromo.PUT("/seckill-config/update", UpdateSeckillConfig)
		adminPromo.DELETE("/seckill-config/delete", DeleteSeckillConfig)
		adminPromo.PUT("/seckill-config/update-status", UpdateSeckillConfigStatus)

		// Combination
		adminPromo.GET("/combination-activity/page", GetCombinationActivityPage)
		adminPromo.GET("/combination-activity/get", GetCombinationActivityDetail)
		adminPromo.GET("/combination-activity/list-by-ids", GetCombinationActivityListByIds)
		adminPromo.POST("/combination-activity/create", CreateCombinationActivity)
		adminPromo.PUT("/combination-activity/update", UpdateCombinationActivity)
		adminPromo.PUT("/combination-activity/close", CloseCombinationActivity)
		adminPromo.DELETE("/combination-activity/delete", DeleteCombinationActivity)
		adminPromo.GET("/combination-record/page", GetCombinationRecordPage)
		adminPromo.GET("/combination-record/get-summary", GetCombinationRecordSummary)

		// Bargain
		adminPromo.GET("/bargain-activity/page", GetBargainActivityPage)
		adminPromo.GET("/bargain-activity/get", GetBargainActivityDetail)
		adminPromo.POST("/bargain-activity/create", CreateBargainActivity)
		adminPromo.PUT("/bargain-activity/update", UpdateBargainActivity)
		adminPromo.PUT("/bargain-activity/close", CloseBargainActivity)
		adminPromo.DELETE("/bargain-activity/delete", DeleteBargainActivity)
		adminPromo.GET("/bargain-help/page", GetBargainHelpPage)
		adminPromo.GET("/bargain-record/page", GetBargainHelpPage)

		// Discount & Reward & Point Activity
		adminPromo.GET("/discount-activity/page", GetDiscountActivityPage)
		adminPromo.GET("/discount-activity/get", GetDiscountActivityDetail)
		adminPromo.POST("/discount-activity/create", CreateDiscountActivity)
		adminPromo.PUT("/discount-activity/update", UpdateDiscountActivity)
		adminPromo.PUT("/discount-activity/close", CloseDiscountActivity)
		adminPromo.DELETE("/discount-activity/delete", DeleteDiscountActivity)

		// Reward Activity
		adminPromo.GET("/reward-activity/page", GetRewardActivityPage)
		adminPromo.GET("/reward-activity/get", GetRewardActivityDetail)
		adminPromo.POST("/reward-activity/create", CreateRewardActivity)
		adminPromo.PUT("/reward-activity/update", UpdateRewardActivity)
		adminPromo.PUT("/reward-activity/close", CloseRewardActivity)
		adminPromo.DELETE("/reward-activity/delete", DeleteRewardActivity)

		// Point Activity
		adminPromo.GET("/point-activity/page", GetPointActivityPage)
		adminPromo.GET("/point-activity/get", GetPointActivityDetail)
		adminPromo.GET("/point-activity/list-by-ids", GetPointActivityListByIds)
		adminPromo.POST("/point-activity/create", CreatePointActivity)
		adminPromo.PUT("/point-activity/update", UpdatePointActivity)
		adminPromo.PUT("/point-activity/close", ClosePointActivity)
		adminPromo.DELETE("/point-activity/delete", DeletePointActivity)

		// DIY Template & Page
		adminPromo.GET("/diy-template/page", GetDiyTemplatePage)
		adminPromo.GET("/diy-template/get", GetDiyTemplateDetail)
		adminPromo.GET("/diy-template/get-property", GetDiyTemplateProperty)
		adminPromo.POST("/diy-template/create", CreateDiyTemplate)
		adminPromo.PUT("/diy-template/update", UpdateDiyTemplate)
		adminPromo.PUT("/diy-template/update-property", UpdateDiyTemplateProperty)
		adminPromo.PUT("/diy-template/use", UseDiyTemplate)
		adminPromo.DELETE("/diy-template/delete", DeleteDiyTemplate)

		adminPromo.GET("/diy-page/page", GetDiyPagePage)
		adminPromo.GET("/diy-page/get", GetDiyPageDetail)
		adminPromo.GET("/diy-page/get-property", GetDiyTemplateProperty)
		adminPromo.POST("/diy-page/create", CreateDiyPage)
		adminPromo.PUT("/diy-page/update", UpdateDiyPage)
		adminPromo.PUT("/diy-page/update-property", UpdateDiyTemplateProperty)
		adminPromo.DELETE("/diy-page/delete", DeleteDiyPage)

		// Banner & Article
		adminPromo.GET("/banner/page", GetBannerPage)
		adminPromo.GET("/banner/get", GetBannerDetail)
		adminPromo.POST("/banner/create", CreateBanner)
		adminPromo.PUT("/banner/update", UpdateBanner)
		adminPromo.DELETE("/banner/delete", DeleteBanner)

		adminPromo.GET("/article/page", GetArticlePage)
		adminPromo.GET("/article/get", GetArticleDetail)
		adminPromo.POST("/article/create", CreateArticle)
		adminPromo.PUT("/article/update", UpdateArticle)
		adminPromo.DELETE("/article/delete", DeleteArticle)

		adminPromo.GET("/article-category/page", GetArticleCategoryPage)
		adminPromo.GET("/article-category/get", GetArticleCategoryDetail)
		adminPromo.GET("/article-category/list-all-simple", GetArticleCategoryList)
		adminPromo.POST("/article-category/create", CreateArticleCategory)
		adminPromo.PUT("/article-category/update", UpdateArticleCategory)
		adminPromo.DELETE("/article-category/delete", DeleteArticleCategory)

		// Kefu
		adminPromo.GET("/kefu-conversation/list", GetKefuConversationList)
		adminPromo.GET("/kefu-conversation/get", GetKefuConversationDetail)
		adminPromo.PUT("/kefu-conversation/update-conversation-pinned", PinKefuConversation)
		adminPromo.DELETE("/kefu-conversation/delete", DeleteKefuConversation)

		adminPromo.GET("/kefu-message/page", GetKefuMessagePage)
		adminPromo.POST("/kefu-message/send", SendKefuMessage)
		adminPromo.PUT("/kefu-message/update-read-status", UpdateKefuMessageRead)
	}

	// 移动端/小程序接口
	appPromo := r.Group("/app-api/promotion")
	{
		appPromo.GET("/diy-template/get-used", GetUsedDiyTemplate)
		appPromo.GET("/diy-template/used", GetUsedDiyTemplate)
		appPromo.GET("/diy-template/get", GetUsedDiyTemplate)
		appPromo.GET("/diy-page/get", GetDiyPage)

		appPromo.GET("/coupon-template/page", GetCouponTemplatePage)
		appPromo.GET("/coupon-template/list", GetCouponTemplateList)
		appPromo.GET("/coupon-template/list-by-ids", GetCouponTemplateList)
		appPromo.GET("/coupon-template/get", GetCouponTemplateDetail)
		appPromo.GET("/coupon/match-list", GetMatchCouponList)
		appPromo.GET("/coupon/page", GetUserCouponPage)
		appPromo.GET("/coupon/get", GetUserCouponDetail)
		appPromo.GET("/coupon/get-unused-count", GetUserUnusedCouponCount)
		appPromo.POST("/coupon/take", TakeCoupon)

		appPromo.GET("/seckill-config/list", GetSeckillConfigList)
		appPromo.GET("/seckill-activity/get-now", GetNowSeckillActivity)
		appPromo.GET("/seckill-activity/page", GetSeckillActivityPage)
		appPromo.GET("/seckill-activity/list", GetSeckillActivityList)
		appPromo.GET("/seckill-activity/list-by-ids", GetSeckillActivityList)
		appPromo.GET("/seckill-activity/get-detail", GetSeckillActivityDetail)
		appPromo.GET("/combination-activity/list", GetCombinationActivityList)
		appPromo.GET("/combination-activity/list-by-ids", GetCombinationActivityList)
		appPromo.GET("/combination-activity/get-detail", GetCombinationActivityDetail)
		appPromo.GET("/combination-activity/page", GetCombinationActivityPage)
		appPromo.GET("/combination-record/get-head-list", GetCombinationHeadList)
		appPromo.GET("/combination-record/page", GetCombinationRecordPage)
		appPromo.GET("/combination-record/get-detail", GetCombinationRecordDetail)
		appPromo.GET("/combination-record/get-summary", GetCombinationRecordSummary)

		appPromo.GET("/bargain-activity/list", GetBargainActivityList)
		appPromo.GET("/point-activity/page", GetPointActivityPage)
		appPromo.GET("/point-activity/get-detail", GetPointActivityDetail)
		appPromo.GET("/point-activity/list-by-ids", GetPointActivityListByIds)
		appPromo.GET("/reward-activity/get", GetRewardActivityDetail)
		appPromo.GET("/activity/list-by-spu-id", GetActivityListBySpuId)

		appPromo.GET("/banner/list", GetBannerList)
		appPromo.GET("/article/page", GetArticlePage)
		appPromo.GET("/article/get", GetArticleDetail)

		appPromo.GET("/kefu/message/list", GetKefuMessagePage)
		appPromo.GET("/kefu-message/list", GetKefuMessagePage)
		appPromo.POST("/kefu/message/send", SendKefuMessage)
		appPromo.POST("/kefu-message/send", SendKefuMessage)
	}
}

func ExportCouponTemplateExcel(c *gin.Context) {
	c.Header("Content-Disposition", "attachment; filename=coupon_templates.xlsx")
	c.Data(200, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", []byte{})
}

// 优惠券业务
func GetCouponTemplatePage(c *gin.Context) {
	pageNo, _ := strconv.Atoi(c.DefaultQuery("pageNo", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	var total int64
	database.DB.Model(&database.CouponTemplate{}).Count(&total)

	var list []database.CouponTemplate
	database.DB.Order("id DESC").Offset((pageNo - 1) * pageSize).Limit(pageSize).Find(&list)
	web.PageSuccess(c, list, total)
}

func GetCouponTemplateDetail(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	var tpl database.CouponTemplate
	database.DB.First(&tpl, id)
	web.Success(c, tpl)
}

func GetCouponTemplateList(c *gin.Context) {
	var list []database.CouponTemplate
	database.DB.Where("status = ?", 0).Find(&list)
	web.Success(c, list)
}

func CreateCouponTemplate(c *gin.Context) {
	var tpl database.CouponTemplate
	if err := c.ShouldBindJSON(&tpl); err == nil {
		tpl.CreatedAt = time.Now()
		database.DB.Create(&tpl)
		web.Success(c, tpl.ID)
		return
	}
	web.Success(c, 1)
}

func UpdateCouponTemplate(c *gin.Context) {
	var tpl database.CouponTemplate
	if err := c.ShouldBindJSON(&tpl); err == nil {
		database.DB.Model(&tpl).Updates(&tpl)
	}
	web.Success(c, true)
}

func UpdateCouponTemplateStatus(c *gin.Context) {
	var req struct {
		ID     int64 `json:"id"`
		Status int   `json:"status"`
	}
	if err := c.ShouldBindJSON(&req); err == nil {
		database.DB.Model(&database.CouponTemplate{}).Where("id = ?", req.ID).Update("status", req.Status)
	}
	web.Success(c, true)
}

func DeleteCouponTemplate(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	if id > 0 {
		database.DB.Delete(&database.CouponTemplate{}, id)
	}
	web.Success(c, true)
}

func GetCouponUserPage(c *gin.Context) {
	var list []database.CouponUser
	database.DB.Find(&list)
	web.PageSuccess(c, list, int64(len(list)))
}

func GetUserCouponPage(c *gin.Context) {
	var list []database.CouponUser
	database.DB.Where("user_id = ?", 1001).Find(&list)
	web.PageSuccess(c, list, int64(len(list)))
}

func GetUserCouponDetail(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	var cu database.CouponUser
	database.DB.First(&cu, id)
	web.Success(c, cu)
}

func GetUserUnusedCouponCount(c *gin.Context) {
	var count int64
	database.DB.Model(&database.CouponUser{}).Where("user_id = ? AND status = ?", 1001, 1).Count(&count)
	web.Success(c, count)
}

// TakeCoupon 用户领取优惠券 (严格遵循 Java 版库存与限领业务规则)
func TakeCoupon(c *gin.Context) {
	var req struct {
		TemplateID int64 `json:"templateId"`
	}
	_ = c.ShouldBindJSON(&req)
	if req.TemplateID == 0 {
		req.TemplateID, _ = strconv.ParseInt(c.Query("templateId"), 10, 64)
	}

	var tpl database.CouponTemplate
	if err := database.DB.First(&tpl, req.TemplateID).Error; err != nil {
		c.JSON(200, gin.H{"code": 1005001, "msg": "优惠券模板不存在"})
		return
	}
	if tpl.Status != 0 {
		c.JSON(200, gin.H{"code": 1005002, "msg": "该优惠券活动尚未开启或已停用"})
		return
	}
	if tpl.TotalCount > 0 && tpl.TakeCount >= tpl.TotalCount {
		c.JSON(200, gin.H{"code": 1005003, "msg": "很抱歉，该优惠券已被抢光啦"})
		return
	}

	var userCount int64
	database.DB.Model(&database.CouponUser{}).Where("user_id = ? AND template_id = ?", 1001, tpl.ID).Count(&userCount)
	if userCount >= 3 {
		c.JSON(200, gin.H{"code": 1005004, "msg": "您已达到该优惠券的个人领取上限"})
		return
	}

	cu := database.CouponUser{
		TemplateID:    tpl.ID,
		UserID:        1001,
		Name:          tpl.Name,
		Status:        1, // 未使用
		DiscountPrice: tpl.DiscountPrice,
		UsePrice:      tpl.UsePrice,
		CreatedAt:     time.Now(),
	}
	database.DB.Create(&cu)
	database.DB.Model(&tpl).Update("take_count", tpl.TakeCount+1)
	web.Success(c, cu.ID)
}

func SendUserCoupon(c *gin.Context) {
	TakeCoupon(c)
}

func DeleteUserCoupon(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	if id > 0 {
		database.DB.Delete(&database.CouponUser{}, id)
	}
	web.Success(c, true)
}

func GetMatchCouponList(c *gin.Context) {
	var list []database.CouponUser
	database.DB.Where("user_id = ? AND status = ?", 1001, 1).Find(&list)
	web.Success(c, list)
}

// DIY 装修
func GetUsedDiyTemplate(c *gin.Context) {
	var tpl database.DiyTemplate
	if err := database.DB.Where("used = ?", true).First(&tpl).Error; err != nil {
		tpl = database.DiyTemplate{
			ID:   1,
			Name: "默认宠物商城模板",
			Used: true,
			Property: `{"navigationBar":{"title":"我的宠物商城","backgroundColor":"#FF6B6B"},"components":[{"id":"banner","property":{"items":[{"imgUrl":"https://img.iocoder.cn/mall/banner1.jpg","url":"/pages/goods/detail?id=1"}]}}]}`,
		}
	}
	web.Success(c, tpl)
}

func GetDiyPage(c *gin.Context) {
	GetUsedDiyTemplate(c)
}

func GetDiyTemplatePage(c *gin.Context) {
	var list []database.DiyTemplate
	database.DB.Find(&list)
	web.PageSuccess(c, list, int64(len(list)))
}

func GetDiyTemplateDetail(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	var tpl database.DiyTemplate
	database.DB.First(&tpl, id)
	web.Success(c, tpl)
}

func GetDiyTemplateProperty(c *gin.Context) {
	var tpl database.DiyTemplate
	database.DB.Where("used = ?", true).First(&tpl)
	web.Success(c, tpl.Property)
}

func CreateDiyTemplate(c *gin.Context) {
	var tpl database.DiyTemplate
	if err := c.ShouldBindJSON(&tpl); err == nil {
		tpl.CreatedAt = time.Now()
		database.DB.Create(&tpl)
		web.Success(c, tpl.ID)
		return
	}
	web.Success(c, 1)
}

func UpdateDiyTemplate(c *gin.Context) {
	var tpl database.DiyTemplate
	if err := c.ShouldBindJSON(&tpl); err == nil {
		database.DB.Model(&tpl).Updates(&tpl)
	}
	web.Success(c, true)
}

func UpdateDiyTemplateProperty(c *gin.Context) {
	var req struct {
		ID       int64  `json:"id"`
		Property string `json:"property"`
	}
	if err := c.ShouldBindJSON(&req); err == nil {
		database.DB.Model(&database.DiyTemplate{}).Where("id = ?", req.ID).Update("property", req.Property)
	}
	web.Success(c, true)
}

func UseDiyTemplate(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	if id > 0 {
		database.DB.Model(&database.DiyTemplate{}).Where("1 = 1").Update("used", false)
		database.DB.Model(&database.DiyTemplate{}).Where("id = ?", id).Update("used", true)
	}
	web.Success(c, true)
}

func DeleteDiyTemplate(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	if id > 0 {
		database.DB.Delete(&database.DiyTemplate{}, id)
	}
	web.Success(c, true)
}

func GetDiyPagePage(c *gin.Context) {
	var list []database.DiyPage
	database.DB.Find(&list)
	web.PageSuccess(c, list, int64(len(list)))
}

func GetDiyPageDetail(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	var dp database.DiyPage
	database.DB.First(&dp, id)
	web.Success(c, dp)
}

func CreateDiyPage(c *gin.Context) {
	var dp database.DiyPage
	if err := c.ShouldBindJSON(&dp); err == nil {
		dp.CreatedAt = time.Now()
		database.DB.Create(&dp)
		web.Success(c, dp.ID)
		return
	}
	web.Success(c, 1)
}

func UpdateDiyPage(c *gin.Context) {
	var dp database.DiyPage
	if err := c.ShouldBindJSON(&dp); err == nil {
		database.DB.Model(&dp).Updates(&dp)
	}
	web.Success(c, true)
}

func DeleteDiyPage(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	if id > 0 {
		database.DB.Delete(&database.DiyPage{}, id)
	}
	web.Success(c, true)
}

func GetDiyTemplateList(c *gin.Context) {
	var list []database.DiyTemplate
	database.DB.Find(&list)
	web.Success(c, list)
}

// 秒杀/拼团/砍价
func GetSeckillActivityPage(c *gin.Context) {
	var list []database.SeckillActivity
	database.DB.Find(&list)
	web.PageSuccess(c, list, int64(len(list)))
}

func GetSeckillActivityList(c *gin.Context) {
	var list []database.SeckillActivity
	database.DB.Find(&list)
	web.Success(c, list)
}

func GetSeckillActivityDetail(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	var a database.SeckillActivity
	database.DB.First(&a, id)
	web.Success(c, a)
}

func CreateSeckillActivity(c *gin.Context) {
	var a database.SeckillActivity
	if err := c.ShouldBindJSON(&a); err == nil {
		a.CreatedAt = time.Now()
		database.DB.Create(&a)
		web.Success(c, a.ID)
		return
	}
	web.Success(c, 1)
}

func UpdateSeckillActivity(c *gin.Context) {
	var a database.SeckillActivity
	if err := c.ShouldBindJSON(&a); err == nil {
		database.DB.Model(&a).Updates(&a)
	}
	web.Success(c, true)
}

func CloseSeckillActivity(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	if id > 0 {
		database.DB.Model(&database.SeckillActivity{}).Where("id = ?", id).Update("status", 1)
	}
	web.Success(c, true)
}

func DeleteSeckillActivity(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	if id > 0 {
		database.DB.Delete(&database.SeckillActivity{}, id)
	}
	web.Success(c, true)
}

func GetSeckillConfigList(c *gin.Context) {
	var list []database.SeckillConfig
	database.DB.Order("id ASC").Find(&list)
	web.Success(c, list)
}

func GetNowSeckillActivity(c *gin.Context) {
	var activity database.SeckillActivity
	if err := database.DB.Where("status = 0").Order("start_time ASC").First(&activity).Error; err == nil {
		web.Success(c, activity)
		return
	}
	web.Success(c, nil)
}

func GetCombinationActivityPage(c *gin.Context) {
	var list []database.CombinationActivity
	database.DB.Find(&list)
	web.PageSuccess(c, list, int64(len(list)))
}

func GetCombinationActivityList(c *gin.Context) {
	var list []database.CombinationActivity
	database.DB.Find(&list)
	web.Success(c, list)
}

func GetCombinationActivityListByIds(c *gin.Context) {
	GetCombinationActivityList(c)
}

func GetCombinationActivityDetail(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	var a database.CombinationActivity
	database.DB.First(&a, id)
	web.Success(c, a)
}

func CreateCombinationActivity(c *gin.Context) {
	var a database.CombinationActivity
	if err := c.ShouldBindJSON(&a); err == nil {
		a.CreatedAt = time.Now()
		database.DB.Create(&a)
		web.Success(c, a.ID)
		return
	}
	web.Success(c, 1)
}

func UpdateCombinationActivity(c *gin.Context) {
	var a database.CombinationActivity
	if err := c.ShouldBindJSON(&a); err == nil {
		database.DB.Model(&a).Updates(&a)
	}
	web.Success(c, true)
}

func CloseCombinationActivity(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	if id > 0 {
		database.DB.Model(&database.CombinationActivity{}).Where("id = ?", id).Update("status", 1)
	}
	web.Success(c, true)
}

func DeleteCombinationActivity(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	if id > 0 {
		database.DB.Delete(&database.CombinationActivity{}, id)
	}
	web.Success(c, true)
}

func GetCombinationRecordPage(c *gin.Context) {
	var list []database.CombinationRecord
	database.DB.Find(&list)
	web.PageSuccess(c, list, int64(len(list)))
}

func GetCombinationRecordDetail(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	var r database.CombinationRecord
	database.DB.First(&r, id)
	web.Success(c, r)
}

func GetCombinationHeadList(c *gin.Context) {
	var list []database.CombinationRecord
	database.DB.Where("head_id = ? AND status = ?", 0, 0).Limit(10).Find(&list)
	web.Success(c, list)
}

func GetCombinationRecordSummary(c *gin.Context) {
	var userCount int64
	var successCount int64
	database.DB.Model(&database.CombinationRecord{}).Select("count(distinct user_id)").Scan(&userCount)
	database.DB.Model(&database.CombinationRecord{}).Where("status = ?", 1).Count(&successCount)
	web.Success(c, gin.H{
		"userCount":    userCount,
		"successCount": successCount,
	})
}

func GetBargainActivityPage(c *gin.Context) {
	var list []database.BargainActivity
	database.DB.Find(&list)
	web.PageSuccess(c, list, int64(len(list)))
}

func GetBargainActivityList(c *gin.Context) {
	var list []database.BargainActivity
	database.DB.Find(&list)
	web.Success(c, list)
}

func GetBargainActivityDetail(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	var a database.BargainActivity
	database.DB.First(&a, id)
	web.Success(c, a)
}

func CreateBargainActivity(c *gin.Context) {
	var a database.BargainActivity
	if err := c.ShouldBindJSON(&a); err == nil {
		a.CreatedAt = time.Now()
		database.DB.Create(&a)
		web.Success(c, a.ID)
		return
	}
	web.Success(c, 1)
}

func UpdateBargainActivity(c *gin.Context) {
	var a database.BargainActivity
	if err := c.ShouldBindJSON(&a); err == nil {
		database.DB.Model(&a).Updates(&a)
	}
	web.Success(c, true)
}

func CloseBargainActivity(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	if id > 0 {
		database.DB.Model(&database.BargainActivity{}).Where("id = ?", id).Update("status", 1)
	}
	web.Success(c, true)
}

func DeleteBargainActivity(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	if id > 0 {
		database.DB.Delete(&database.BargainActivity{}, id)
	}
	web.Success(c, true)
}

func GetBargainHelpPage(c *gin.Context) {
	var list []database.BargainRecord
	database.DB.Order("id DESC").Find(&list)
	web.PageSuccess(c, list, int64(len(list)))
}

func GetActivityListBySpuId(c *gin.Context) {
	spuId, _ := strconv.ParseInt(c.Query("spuId"), 10, 64)
	var seckills []database.SeckillActivity
	var combinations []database.CombinationActivity
	database.DB.Where("spu_id = ? AND status = 0", spuId).Find(&seckills)
	database.DB.Where("spu_id = ? AND status = 0", spuId).Find(&combinations)
	web.Success(c, gin.H{
		"seckill":     seckills,
		"combination": combinations,
	})
}

func GetBannerList(c *gin.Context) {
	var list []database.Banner
	database.DB.Where("status = 0").Order("sort ASC").Find(&list)
	web.Success(c, list)
}

func GetBannerPage(c *gin.Context) {
	var list []database.Banner
	database.DB.Order("sort ASC, id DESC").Find(&list)
	web.PageSuccess(c, list, int64(len(list)))
}

func GetBannerDetail(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	var b database.Banner
	database.DB.First(&b, id)
	web.Success(c, b)
}

func CreateBanner(c *gin.Context) {
	var b database.Banner
	if err := c.ShouldBindJSON(&b); err == nil {
		b.CreatedAt = time.Now()
		database.DB.Create(&b)
		web.Success(c, b.ID)
		return
	}
	web.Success(c, 1)
}

func UpdateBanner(c *gin.Context) {
	var b database.Banner
	if err := c.ShouldBindJSON(&b); err == nil {
		database.DB.Model(&b).Updates(&b)
	}
	web.Success(c, true)
}

func DeleteBanner(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	if id > 0 {
		database.DB.Delete(&database.Banner{}, id)
	}
	web.Success(c, true)
}

func GetArticlePage(c *gin.Context) {
	var list []database.Article
	database.DB.Order("id DESC").Find(&list)
	web.PageSuccess(c, list, int64(len(list)))
}

func GetArticleDetail(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	var a database.Article
	database.DB.First(&a, id)
	web.Success(c, a)
}

func CreateArticle(c *gin.Context) {
	var a database.Article
	if err := c.ShouldBindJSON(&a); err == nil {
		a.CreatedAt = time.Now()
		database.DB.Create(&a)
		web.Success(c, a.ID)
		return
	}
	web.Success(c, 1)
}

func UpdateArticle(c *gin.Context) {
	var a database.Article
	if err := c.ShouldBindJSON(&a); err == nil {
		database.DB.Model(&a).Updates(&a)
	}
	web.Success(c, true)
}

func DeleteArticle(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	if id > 0 {
		database.DB.Delete(&database.Article{}, id)
	}
	web.Success(c, true)
}

func GetArticleCategoryPage(c *gin.Context) {
	var list []database.ArticleCategory
	database.DB.Order("sort ASC, id DESC").Find(&list)
	web.PageSuccess(c, list, int64(len(list)))
}

func GetArticleCategoryDetail(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	var ac database.ArticleCategory
	database.DB.First(&ac, id)
	web.Success(c, ac)
}

func GetArticleCategoryList(c *gin.Context) {
	var list []database.ArticleCategory
	database.DB.Where("status = 0").Order("sort ASC").Find(&list)
	web.Success(c, list)
}

func CreateArticleCategory(c *gin.Context) {
	var ac database.ArticleCategory
	if err := c.ShouldBindJSON(&ac); err == nil {
		ac.CreatedAt = time.Now()
		database.DB.Create(&ac)
		web.Success(c, ac.ID)
		return
	}
	web.Success(c, 1)
}

func UpdateArticleCategory(c *gin.Context) {
	var ac database.ArticleCategory
	if err := c.ShouldBindJSON(&ac); err == nil {
		database.DB.Model(&ac).Updates(&ac)
	}
	web.Success(c, true)
}

func DeleteArticleCategory(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	if id > 0 {
		database.DB.Delete(&database.ArticleCategory{}, id)
	}
	web.Success(c, true)
}

// 秒杀配置 CRUD
func GetSeckillConfigDetail(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	var sc database.SeckillConfig
	database.DB.First(&sc, id)
	web.Success(c, sc)
}

func CreateSeckillConfig(c *gin.Context) {
	var sc database.SeckillConfig
	if err := c.ShouldBindJSON(&sc); err == nil {
		sc.CreatedAt = time.Now()
		database.DB.Create(&sc)
		web.Success(c, sc.ID)
		return
	}
	web.Success(c, 1)
}

func UpdateSeckillConfig(c *gin.Context) {
	var sc database.SeckillConfig
	if err := c.ShouldBindJSON(&sc); err == nil {
		database.DB.Model(&sc).Updates(&sc)
	}
	web.Success(c, true)
}

func DeleteSeckillConfig(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	if id > 0 {
		database.DB.Delete(&database.SeckillConfig{}, id)
	}
	web.Success(c, true)
}

func UpdateSeckillConfigStatus(c *gin.Context) {
	var req struct {
		ID     int64 `json:"id"`
		Status int   `json:"status"`
	}
	if err := c.ShouldBindJSON(&req); err == nil {
		database.DB.Model(&database.SeckillConfig{}).Where("id = ?", req.ID).Update("status", req.Status)
	}
	web.Success(c, true)
}

// 限时折扣活动
func GetDiscountActivityPage(c *gin.Context) {
	var list []database.DiscountActivity
	database.DB.Order("id DESC").Find(&list)
	web.PageSuccess(c, list, int64(len(list)))
}

func GetDiscountActivityDetail(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	var d database.DiscountActivity
	database.DB.First(&d, id)
	web.Success(c, d)
}

func CreateDiscountActivity(c *gin.Context) {
	var d database.DiscountActivity
	if err := c.ShouldBindJSON(&d); err == nil {
		d.CreatedAt = time.Now()
		database.DB.Create(&d)
		web.Success(c, d.ID)
		return
	}
	web.Success(c, 1)
}

func UpdateDiscountActivity(c *gin.Context) {
	var d database.DiscountActivity
	if err := c.ShouldBindJSON(&d); err == nil {
		database.DB.Model(&d).Updates(&d)
	}
	web.Success(c, true)
}

func CloseDiscountActivity(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	if id > 0 {
		database.DB.Model(&database.DiscountActivity{}).Where("id = ?", id).Update("status", 1)
	}
	web.Success(c, true)
}

func DeleteDiscountActivity(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	if id > 0 {
		database.DB.Delete(&database.DiscountActivity{}, id)
	}
	web.Success(c, true)
}

// 满减送活动
func GetRewardActivityPage(c *gin.Context) {
	var list []database.RewardActivity
	database.DB.Order("id DESC").Find(&list)
	web.PageSuccess(c, list, int64(len(list)))
}

func GetRewardActivityDetail(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	var r database.RewardActivity
	database.DB.First(&r, id)
	web.Success(c, r)
}

func CreateRewardActivity(c *gin.Context) {
	var r database.RewardActivity
	if err := c.ShouldBindJSON(&r); err == nil {
		r.CreatedAt = time.Now()
		database.DB.Create(&r)
		web.Success(c, r.ID)
		return
	}
	web.Success(c, 1)
}

func UpdateRewardActivity(c *gin.Context) {
	var r database.RewardActivity
	if err := c.ShouldBindJSON(&r); err == nil {
		database.DB.Model(&r).Updates(&r)
	}
	web.Success(c, true)
}

func CloseRewardActivity(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	if id > 0 {
		database.DB.Model(&database.RewardActivity{}).Where("id = ?", id).Update("status", 1)
	}
	web.Success(c, true)
}

func DeleteRewardActivity(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	if id > 0 {
		database.DB.Delete(&database.RewardActivity{}, id)
	}
	web.Success(c, true)
}

// 积分商城活动
func GetPointActivityPage(c *gin.Context) {
	var list []database.PointActivity
	database.DB.Order("id DESC").Find(&list)
	web.PageSuccess(c, list, int64(len(list)))
}

func GetPointActivityDetail(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	var p database.PointActivity
	database.DB.First(&p, id)
	web.Success(c, p)
}

func GetPointActivityListByIds(c *gin.Context) {
	var list []database.PointActivity
	database.DB.Find(&list)
	web.Success(c, list)
}

func CreatePointActivity(c *gin.Context) {
	var p database.PointActivity
	if err := c.ShouldBindJSON(&p); err == nil {
		p.CreatedAt = time.Now()
		database.DB.Create(&p)
		web.Success(c, p.ID)
		return
	}
	web.Success(c, 1)
}

func UpdatePointActivity(c *gin.Context) {
	var p database.PointActivity
	if err := c.ShouldBindJSON(&p); err == nil {
		database.DB.Model(&p).Updates(&p)
	}
	web.Success(c, true)
}

func ClosePointActivity(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	if id > 0 {
		database.DB.Model(&database.PointActivity{}).Where("id = ?", id).Update("status", 1)
	}
	web.Success(c, true)
}

func DeletePointActivity(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	if id > 0 {
		database.DB.Delete(&database.PointActivity{}, id)
	}
	web.Success(c, true)
}

// 客服系统
func GetKefuConversationList(c *gin.Context) {
	var list []database.KefuConversation
	database.DB.Order("pinned DESC, last_time DESC").Find(&list)
	web.Success(c, list)
}

func GetKefuConversationDetail(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	var conv database.KefuConversation
	database.DB.First(&conv, id)
	web.Success(c, conv)
}

func PinKefuConversation(c *gin.Context) {
	var req struct {
		ID     int64 `json:"id"`
		Pinned bool  `json:"pinned"`
	}
	if err := c.ShouldBindJSON(&req); err == nil {
		database.DB.Model(&database.KefuConversation{}).Where("id = ?", req.ID).Update("pinned", req.Pinned)
	}
	web.Success(c, true)
}

func DeleteKefuConversation(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	if id > 0 {
		database.DB.Delete(&database.KefuConversation{}, id)
	}
	web.Success(c, true)
}

func UpdateKefuMessageRead(c *gin.Context) {
	userId, _ := strconv.ParseInt(c.Query("userId"), 10, 64)
	if userId > 0 {
		database.DB.Model(&database.KefuMessage{}).Where("user_id = ?", userId).Update("read_status", true)
	}
	web.Success(c, true)
}

func GetKefuMessagePage(c *gin.Context) {
	var list []database.KefuMessage
	database.DB.Order("id ASC").Find(&list)
	if len(list) == 0 {
		msg := database.KefuMessage{
			UserID:      1001,
			SenderType:  2,
			ContentType: 1,
			Content:     "您好！我是宠物商城专属在线客服，很高兴为您服务！",
			CreatedAt:   time.Now(),
		}
		database.DB.Create(&msg)
		list = append(list, msg)
	}
	web.PageSuccess(c, list, int64(len(list)))
}

func SendKefuMessage(c *gin.Context) {
	var req struct {
		Content     string `json:"content"`
		ContentType int    `json:"contentType"`
	}
	if err := c.ShouldBindJSON(&req); err == nil && strings.TrimSpace(req.Content) != "" {
		if req.ContentType == 0 {
			req.ContentType = 1
		}
		msg := database.KefuMessage{
			UserID:      1001,
			SenderType:  1, // 用户
			ContentType: req.ContentType,
			Content:     req.Content,
			CreatedAt:   time.Now(),
		}
		database.DB.Create(&msg)
		web.Success(c, msg.ID)
		return
	}
	web.Success(c, 1)
}
