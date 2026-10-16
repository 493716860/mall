package member

import (
	"strconv"
	"time"

	"server-go/internal/framework/database"
	"server-go/internal/framework/security"
	"server-go/internal/framework/web"

	"github.com/gin-gonic/gin"
)

// RegisterMemberRoutes 注册会员中心路由
func RegisterMemberRoutes(r *gin.Engine) {
	// 管理后台接口
	adminMember := r.Group("/admin-api/member")
	{
		// Member User
		adminMember.GET("/user/page", GetMemberUserPage)
		adminMember.GET("/user/get", GetMemberUserDetail)
		adminMember.PUT("/user/update", UpdateMemberUser)
		adminMember.PUT("/user/update-level", UpdateMemberUserLevel)
		adminMember.PUT("/user/update-point", UpdateMemberUserPoint)
		adminMember.PUT("/user/update-tag", UpdateMemberUserTag)

		// Group
		adminMember.GET("/group/list", GetMemberGroupList)
		adminMember.GET("/group/list-all-simple", GetMemberGroupList)
		adminMember.GET("/group/page", GetMemberGroupPage)
		adminMember.GET("/group/get", GetMemberGroupDetail)
		adminMember.POST("/group/create", CreateMemberGroup)
		adminMember.PUT("/group/update", UpdateMemberGroup)
		adminMember.DELETE("/group/delete", DeleteMemberGroup)

		// Tag
		adminMember.GET("/tag/list", GetMemberTagList)
		adminMember.GET("/tag/list-all-simple", GetMemberTagList)
		adminMember.GET("/tag/page", GetMemberTagPage)
		adminMember.GET("/tag/get", GetMemberTagDetail)
		adminMember.POST("/tag/create", CreateMemberTag)
		adminMember.PUT("/tag/update", UpdateMemberTag)
		adminMember.DELETE("/tag/delete", DeleteMemberTag)

		// Level
		adminMember.GET("/level/list", GetMemberLevelList)
		adminMember.GET("/level/list-all-simple", GetMemberLevelList)
		adminMember.GET("/level/page", GetMemberLevelPage)
		adminMember.GET("/level/get", GetMemberLevelDetail)
		adminMember.POST("/level/create", CreateMemberLevel)
		adminMember.PUT("/level/update", UpdateMemberLevel)
		adminMember.DELETE("/level/delete", DeleteMemberLevel)

		// Point & Sign-in
		adminMember.GET("/point/record/page", GetPointRecordPage)
		adminMember.GET("/signin/record/page", GetPointRecordPage)
		adminMember.GET("/signin/config/list", GetSignInConfigList)
		adminMember.GET("/sign-in/config/list", GetSignInConfigList)
	}

	// 微信小程序/App 移动端接口
	appMember := r.Group("/app-api/member")
	{
		appMember.POST("/auth/login-by-weixin", LoginByWeixin)
		appMember.POST("/auth/weixin-mini-app-login", LoginByWeixin)
		appMember.POST("/auth/login-by-sms", LoginBySms)
		appMember.POST("/auth/sms-login", LoginBySms)
		appMember.POST("/auth/login", LoginBySms)
		appMember.POST("/auth/send-sms-code", SendSmsCode)
		appMember.POST("/auth/logout", Logout)
		appMember.POST("/auth/refresh-token", RefreshToken)
		appMember.GET("/auth/social-auth-redirect", Logout)
		appMember.POST("/auth/social-login", LoginByWeixin)
		appMember.POST("/auth/create-weixin-jsapi-signature", CreateWeixinJsApiSignature)

		appMember.GET("/user/get", GetAppUserInfo)
		appMember.PUT("/user/update", UpdateAppUserInfo)
		appMember.PUT("/user/update-mobile", UpdateAppUserMobile)
		appMember.PUT("/user/update-mobile-by-weixin", UpdateAppUserMobile)
		appMember.PUT("/user/update-password", UpdateMemberPassword)
		appMember.POST("/user/reset-password", ResetMemberPassword)

		appMember.GET("/address/list", GetAddressList)
		appMember.GET("/address/get", GetAddressDetail)
		appMember.POST("/address/create", CreateAddress)
		appMember.PUT("/address/update", UpdateAddress)
		appMember.DELETE("/address/delete", DeleteAddress)

		appMember.GET("/point/record/page", GetPointRecordPage)
		appMember.GET("/sign-in/record/get-summary", GetSignInSummary)
		appMember.POST("/sign-in/record/create", CreateSignIn)
		appMember.GET("/sign-in/record/page", GetPointRecordPage)
		appMember.GET("/sign-in/config/list", GetSignInConfigList)
		appMember.GET("/level/list", GetMemberLevelList)

		appMember.GET("/social-user/get", GetSocialUser)
		appMember.POST("/social-user/bind", BindSocialUser)
		appMember.DELETE("/social-user/unbind", UnbindSocialUser)
		appMember.GET("/social-user/get-subscribe-template-list", GetSubscribeTemplateList)
		appMember.GET("/social-user/wxa-qrcode", GetWxaQrcode)
		appMember.POST("/social-user/wxa-qrcode", GetWxaQrcode)
	}

	appThird := r.Group("/app-api/third")
	{
		appThird.POST("/apple/login", AppleLogin)
	}

	appLive := r.Group("/app-api/app/mplive")
	{
		appLive.GET("/getRoomList", GetMpLiveRoomList)
		appLive.GET("/getMpLink", GetMpLiveLink)
	}
}

func UpdateMemberPassword(c *gin.Context) {
	var req struct {
		OldPassword string `json:"oldPassword"`
		Password    string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err == nil {
		database.DB.Model(&database.MemberUser{}).Where("id = ?", 1001).Update("status", 0)
	}
	web.Success(c, true)
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

// GetMemberUserPage 会员列表
func GetMemberUserPage(c *gin.Context) {
	pageNo, _ := strconv.Atoi(c.DefaultQuery("pageNo", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	var total int64
	database.DB.Model(&database.MemberUser{}).Count(&total)

	var list []database.MemberUser
	database.DB.Order("id DESC").Offset((pageNo - 1) * pageSize).Limit(pageSize).Find(&list)
	web.PageSuccess(c, list, total)
}

// GetMemberUserDetail 会员详情
func GetMemberUserDetail(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	if id == 0 {
		id = 1001
	}
	var u database.MemberUser
	if err := database.DB.First(&u, id).Error; err != nil {
		database.DB.First(&u)
	}
	web.Success(c, u)
}

func UpdateMemberUser(c *gin.Context) {
	var u database.MemberUser
	if err := c.ShouldBindJSON(&u); err == nil {
		database.DB.Model(&u).Updates(&u)
	}
	web.Success(c, true)
}

func UpdateMemberUserLevel(c *gin.Context) {
	var req struct {
		ID      int64 `json:"id"`
		LevelID int64 `json:"levelId"`
	}
	if err := c.ShouldBindJSON(&req); err == nil {
		levelName := "黄金会员"
		if req.LevelID == 3 {
			levelName = "钻石会员"
		} else if req.LevelID == 1 {
			levelName = "大众会员"
		}
		database.DB.Model(&database.MemberUser{}).Where("id = ?", req.ID).Updates(map[string]any{
			"level_id":   req.LevelID,
			"level_name": levelName,
		})
	}
	web.Success(c, true)
}

func UpdateMemberUserPoint(c *gin.Context) {
	var req struct {
		ID    int64 `json:"id"`
		Point int   `json:"point"`
	}
	if err := c.ShouldBindJSON(&req); err == nil {
		database.DB.Model(&database.MemberUser{}).Where("id = ?", req.ID).Update("point", req.Point)
	}
	web.Success(c, true)
}

func UpdateMemberUserTag(c *gin.Context) {
	web.Success(c, true)
}

func GetMemberGroupList(c *gin.Context) {
	var list []database.MemberGroup
	database.DB.Where("status = 0").Find(&list)
	web.Success(c, list)
}

func GetMemberGroupPage(c *gin.Context) {
	var list []database.MemberGroup
	database.DB.Find(&list)
	web.PageSuccess(c, list, int64(len(list)))
}

func GetMemberGroupDetail(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	var g database.MemberGroup
	database.DB.First(&g, id)
	web.Success(c, g)
}

func CreateMemberGroup(c *gin.Context) {
	var g database.MemberGroup
	if err := c.ShouldBindJSON(&g); err == nil {
		g.CreatedAt = time.Now()
		database.DB.Create(&g)
		web.Success(c, g.ID)
		return
	}
	web.Success(c, 1)
}

func UpdateMemberGroup(c *gin.Context) {
	var g database.MemberGroup
	if err := c.ShouldBindJSON(&g); err == nil {
		database.DB.Model(&g).Updates(&g)
	}
	web.Success(c, true)
}

func DeleteMemberGroup(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	if id > 0 {
		database.DB.Delete(&database.MemberGroup{}, id)
	}
	web.Success(c, true)
}

func GetMemberTagList(c *gin.Context) {
	var list []database.MemberTag
	database.DB.Where("status = 0").Find(&list)
	web.Success(c, list)
}

func GetMemberTagPage(c *gin.Context) {
	var list []database.MemberTag
	database.DB.Find(&list)
	web.PageSuccess(c, list, int64(len(list)))
}

func GetMemberTagDetail(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	var t database.MemberTag
	database.DB.First(&t, id)
	web.Success(c, t)
}

func CreateMemberTag(c *gin.Context) {
	var t database.MemberTag
	if err := c.ShouldBindJSON(&t); err == nil {
		t.CreatedAt = time.Now()
		database.DB.Create(&t)
		web.Success(c, t.ID)
		return
	}
	web.Success(c, 1)
}

func UpdateMemberTag(c *gin.Context) {
	var t database.MemberTag
	if err := c.ShouldBindJSON(&t); err == nil {
		database.DB.Model(&t).Updates(&t)
	}
	web.Success(c, true)
}

func DeleteMemberTag(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	if id > 0 {
		database.DB.Delete(&database.MemberTag{}, id)
	}
	web.Success(c, true)
}

func GetMemberLevelList(c *gin.Context) {
	var list []database.MemberLevel
	database.DB.Where("status = 0").Order("level ASC").Find(&list)
	web.Success(c, list)
}

func GetMemberLevelPage(c *gin.Context) {
	var list []database.MemberLevel
	database.DB.Order("level ASC").Find(&list)
	web.PageSuccess(c, list, int64(len(list)))
}

func GetMemberLevelDetail(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	var l database.MemberLevel
	database.DB.First(&l, id)
	web.Success(c, l)
}

func CreateMemberLevel(c *gin.Context) {
	var l database.MemberLevel
	if err := c.ShouldBindJSON(&l); err == nil {
		l.CreatedAt = time.Now()
		database.DB.Create(&l)
		web.Success(c, l.ID)
		return
	}
	web.Success(c, 1)
}

func UpdateMemberLevel(c *gin.Context) {
	var l database.MemberLevel
	if err := c.ShouldBindJSON(&l); err == nil {
		database.DB.Model(&l).Updates(&l)
	}
	web.Success(c, true)
}

func DeleteMemberLevel(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	if id > 0 {
		database.DB.Delete(&database.MemberLevel{}, id)
	}
	web.Success(c, true)
}

// 小程序用户认证与资料
func LoginByWeixin(c *gin.Context) {
	token, _ := security.GenerateToken(1001, "wx_user", 2, []string{"member"})
	web.Success(c, gin.H{
		"userId":       1001,
		"openid":       "o1234567890abcdef",
		"accessToken":  token,
		"refreshToken": token + "_refresh",
		"expiresTime":  time.Now().Add(24 * 30 * time.Hour).UnixMilli(),
	})
}

func LoginBySms(c *gin.Context) {
	LoginByWeixin(c)
}

func SendSmsCode(c *gin.Context) {
	web.Success(c, true)
}

func Logout(c *gin.Context) {
	web.Success(c, true)
}

func RefreshToken(c *gin.Context) {
	LoginByWeixin(c)
}

func CreateWeixinJsApiSignature(c *gin.Context) {
	web.Success(c, gin.H{
		"appId":     "wx1234567890",
		"nonceStr":  "wx_noncestr",
		"timestamp": time.Now().Unix(),
		"signature": "wx_sig_123456",
	})
}

func GetAppUserInfo(c *gin.Context) {
	var u database.MemberUser
	if err := database.DB.Where("id = ?", 1001).First(&u).Error; err != nil {
		u = database.MemberUser{
			ID:         1001,
			Nickname:   "微信用户小王",
			Avatar:     "https://www.iocoder.cn/images/avatar.jpg",
			Mobile:     "13800138000",
			LevelID:    2,
			LevelName:  "黄金会员",
			Point:      650,
			Experience: 1800,
			Balance:    50000,
		}
		database.DB.Create(&u)
	}

	web.Success(c, gin.H{
		"id":         u.ID,
		"nickname":   u.Nickname,
		"avatar":     u.Avatar,
		"mobile":     u.Mobile,
		"point":      u.Point,
		"experience": u.Experience,
		"balance":    u.Balance,
		"level": gin.H{
			"id":              u.LevelID,
			"name":            u.LevelName,
			"level":           2,
			"discountPercent": 95,
		},
	})
}

func UpdateAppUserInfo(c *gin.Context) {
	var req struct {
		Nickname string `json:"nickname"`
		Avatar   string `json:"avatar"`
	}
	if err := c.ShouldBindJSON(&req); err == nil {
		database.DB.Model(&database.MemberUser{}).Where("id = ?", 1001).Updates(map[string]any{
			"nickname": req.Nickname,
			"avatar":   req.Avatar,
		})
	}
	web.Success(c, true)
}

func UpdateAppUserMobile(c *gin.Context) {
	var req struct {
		Mobile string `json:"mobile"`
	}
	if err := c.ShouldBindJSON(&req); err == nil && req.Mobile != "" {
		database.DB.Model(&database.MemberUser{}).Where("id = ?", 1001).Update("mobile", req.Mobile)
	}
	web.Success(c, true)
}

// 收货地址 CRUD
func GetAddressList(c *gin.Context) {
	var list []database.MemberAddress
	database.DB.Where("user_id = ?", 1001).Order("default_status DESC, id DESC").Find(&list)
	web.Success(c, list)
}

func GetAddressDetail(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	var addr database.MemberAddress
	database.DB.First(&addr, id)
	web.Success(c, addr)
}

func CreateAddress(c *gin.Context) {
	var addr database.MemberAddress
	if err := c.ShouldBindJSON(&addr); err == nil {
		addr.UserID = 1001
		addr.CreatedAt = time.Now()
		if addr.DefaultStatus {
			database.DB.Model(&database.MemberAddress{}).Where("user_id = ?", 1001).Update("default_status", false)
		}
		database.DB.Create(&addr)
		web.Success(c, addr.ID)
		return
	}
	web.Success(c, 1)
}

func UpdateAddress(c *gin.Context) {
	var addr database.MemberAddress
	if err := c.ShouldBindJSON(&addr); err == nil {
		if addr.DefaultStatus {
			database.DB.Model(&database.MemberAddress{}).Where("user_id = ?", 1001).Update("default_status", false)
		}
		database.DB.Model(&addr).Updates(&addr)
	}
	web.Success(c, true)
}

func DeleteAddress(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	if id > 0 {
		database.DB.Delete(&database.MemberAddress{}, id)
	}
	web.Success(c, true)
}

// 积分与签到
func GetPointRecordPage(c *gin.Context) {
	var list []database.MemberPointRecord
	database.DB.Where("user_id = ?", 1001).Order("id DESC").Find(&list)
	if len(list) == 0 {
		list = []database.MemberPointRecord{
			{ID: 1, UserID: 1001, BizType: 1, Title: "连续签到奖励", Point: 20, Total: 650, CreatedAt: time.Now()},
			{ID: 2, UserID: 1001, BizType: 2, Title: "购买猫粮赠送积分", Point: 50, Total: 630, CreatedAt: time.Now().Add(-24 * time.Hour)},
		}
	}
	web.PageSuccess(c, list, int64(len(list)))
}

func GetSignInSummary(c *gin.Context) {
	var count int64
	database.DB.Model(&database.MemberSignInRecord{}).Where("user_id = ?", 1001).Count(&count)
	web.Success(c, gin.H{
		"totalDay":      count + 15,
		"continuousDay": 3,
		"todaySignIn":   true,
	})
}

func CreateSignIn(c *gin.Context) {
	rec := database.MemberSignInRecord{
		UserID:    1001,
		Day:       4,
		Point:     20,
		CreatedAt: time.Now(),
	}
	database.DB.Create(&rec)

	pt := database.MemberPointRecord{
		UserID:    1001,
		BizType:   1,
		Title:     "每日签到奖励",
		Point:     20,
		Total:     670,
		CreatedAt: time.Now(),
	}
	database.DB.Create(&pt)
	database.DB.Model(&database.MemberUser{}).Where("id = ?", 1001).Update("point", 670)

	web.Success(c, gin.H{"point": 20})
}

func GetSignInConfigList(c *gin.Context) {
	var list []database.MemberSignInConfig
	database.DB.Where("status = 0").Order("day ASC").Find(&list)
	web.Success(c, list)
}

func GetSocialUser(c *gin.Context) {
	web.Success(c, gin.H{
		"openid":   "o1234567890abcdef",
		"nickname": "微信用户小王",
	})
}

func BindSocialUser(c *gin.Context) {
	web.Success(c, true)
}

func UnbindSocialUser(c *gin.Context) {
	web.Success(c, true)
}

func GetSubscribeTemplateList(c *gin.Context) {
	web.Success(c, []string{"template_order_pay_success", "template_order_delivery"})
}

func GetWxaQrcode(c *gin.Context) {
	web.Success(c, "https://img.iocoder.cn/mall/wxa_qrcode.png")
}

func AppleLogin(c *gin.Context) {
	var user database.MemberUser
	if err := database.DB.First(&user, 1001).Error; err != nil {
		user = database.MemberUser{
			ID:        1001,
			Nickname:  "苹果快捷用户",
			Avatar:    "https://www.iocoder.cn/images/avatar.jpg",
			Mobile:    "13800138000",
			LevelID:   1,
			LevelName: "普通会员",
			Status:    0,
			CreatedAt: time.Now(),
		}
		database.DB.Create(&user)
	}
	token, _ := security.GenerateToken(user.ID, user.Nickname)
	web.Success(c, gin.H{
		"userId":       user.ID,
		"accessToken":  token,
		"refreshToken": token,
		"expiresTime":  time.Now().Add(24 * time.Hour).UnixMilli(),
	})
}

func GetMpLiveRoomList(c *gin.Context) {
	web.Success(c, []gin.H{})
}

func GetMpLiveLink(c *gin.Context) {
	web.Success(c, "")
}
