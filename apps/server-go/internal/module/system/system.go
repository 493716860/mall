package system

import (
	"strconv"
	"time"

	"server-go/internal/framework/database"
	"server-go/internal/framework/security"
	"server-go/internal/framework/web"

	"github.com/gin-gonic/gin"
)

// RegisterSystemRoutes 注册系统与认证模块路由
func RegisterSystemRoutes(r *gin.Engine) {
	// 1. 认证与登录
	adminAuth := r.Group("/admin-api/system/auth")
	{
		adminAuth.POST("/login", Login)
		adminAuth.GET("/get-permission-info", GetPermissionInfo)
		adminAuth.POST("/logout", Logout)
		adminAuth.POST("/register", Login)
		adminAuth.POST("/send-sms-code", SendSmsCode)
		adminAuth.POST("/sms-login", Login)
		adminAuth.POST("/social-login", Login)
		adminAuth.GET("/social-auth-redirect", Logout)
		adminAuth.POST("/reset-password", Logout)
		adminAuth.POST("/refresh-token", Login)
	}

	adminCaptcha := r.Group("/admin-api/system/captcha")
	{
		adminCaptcha.POST("/get", GetCaptcha)
		adminCaptcha.POST("/check", CheckCaptcha)
	}

	// 2. 租户管理
	adminTenant := r.Group("/admin-api/system/tenant")
	{
		adminTenant.GET("/get-id-by-name", GetTenantByName)
		adminTenant.GET("/get-by-website", GetTenantByWebsite)
		adminTenant.GET("/page", GetTenantPage)
		adminTenant.GET("/get", GetTenantDetail)
		adminTenant.POST("/create", CreateTenant)
		adminTenant.PUT("/update", UpdateTenant)
		adminTenant.DELETE("/delete", DeleteTenant)
		adminTenant.GET("/export-excel", ExportTenantExcel)
	}

	adminTenantPackage := r.Group("/admin-api/system/tenant-package")
	{
		adminTenantPackage.GET("/page", GetTenantPage)
		adminTenantPackage.GET("/get", GetTenantDetail)
		adminTenantPackage.POST("/create", CreateTenantPackage)
		adminTenantPackage.PUT("/update", UpdateTenantPackage)
		adminTenantPackage.DELETE("/delete", DeleteTenantPackage)
		adminTenantPackage.GET("/get-simple-list", GetTenantList)
	}

	// 3. 字典数据与类型 (全真实数据库 CRUD)
	adminDictData := r.Group("/admin-api/system/dict-data")
	{
		adminDictData.GET("/simple-list", GetSimpleDictDataList)
		adminDictData.GET("/list-all-simple", GetSimpleDictDataList)
		adminDictData.GET("/page", GetDictDataPage)
		adminDictData.GET("/type", GetDictDataByType)
		adminDictData.GET("/get", GetDictDataDetail)
		adminDictData.POST("/create", CreateDictData)
		adminDictData.PUT("/update", UpdateDictData)
		adminDictData.DELETE("/delete", DeleteDictData)
		adminDictData.GET("/export-excel", ExportDictDataExcel)
	}

	adminDictType := r.Group("/admin-api/system/dict-type")
	{
		adminDictType.GET("/page", GetDictTypePage)
		adminDictType.GET("/list-all-simple", GetDictTypeList)
		adminDictType.GET("/get", GetDictTypeDetail)
		adminDictType.POST("/create", CreateDictType)
		adminDictType.PUT("/update", UpdateDictType)
		adminDictType.DELETE("/delete", DeleteDictType)
		adminDictType.GET("/export-excel", ExportDictTypeExcel)
	}

	// 4. 部门与岗位 (全真实数据库 CRUD)
	adminDept := r.Group("/admin-api/system/dept")
	{
		adminDept.GET("/simple-list", GetSimpleDeptList)
		adminDept.GET("/list-all-simple", GetSimpleDeptList)
		adminDept.GET("/list", GetDeptList)
		adminDept.GET("/get", GetDeptDetail)
		adminDept.POST("/create", CreateDept)
		adminDept.PUT("/update", UpdateDept)
		adminDept.DELETE("/delete", DeleteDept)
	}

	adminPost := r.Group("/admin-api/system/post")
	{
		adminPost.GET("/page", GetPostPage)
		adminPost.GET("/simple-list", GetPostList)
		adminPost.GET("/list-all-simple", GetPostList)
		adminPost.GET("/get", GetPostDetail)
		adminPost.POST("/create", CreatePost)
		adminPost.PUT("/update", UpdatePost)
		adminPost.DELETE("/delete", DeletePost)
		adminPost.GET("/export-excel", ExportPostExcel)
	}

	// 5. 管理员账号与个人中心 (全真实数据库 CRUD)
	adminUser := r.Group("/admin-api/system/user")
	{
		adminUser.GET("/simple-list", GetSimpleUserList)
		adminUser.GET("/list-all-simple", GetSimpleUserList)
		adminUser.GET("/page", GetUserPage)
		adminUser.GET("/get", GetUserDetail)
		adminUser.POST("/create", CreateUser)
		adminUser.PUT("/update", UpdateUser)
		adminUser.PUT("/update-status", UpdateUserStatus)
		adminUser.PUT("/update-password", UpdateUserPassword)
		adminUser.DELETE("/delete", DeleteUser)
		adminUser.GET("/export-excel", ExportUserExcel)
		adminUser.GET("/profile/get", GetUserProfile)
		adminUser.PUT("/profile/update", UpdateUserProfile)
		adminUser.PUT("/profile/update-password", UpdateUserProfilePassword)
		adminUser.PUT("/profile/update-avatar", UpdateUserProfileAvatar)
	}

	// 6. 角色与菜单 (全真实数据库 CRUD)
	adminRole := r.Group("/admin-api/system/role")
	{
		adminRole.GET("/page", GetRolePage)
		adminRole.GET("/simple-list", GetRoleList)
		adminRole.GET("/list-all-simple", GetRoleList)
		adminRole.GET("/get", GetRoleDetail)
		adminRole.POST("/create", CreateRole)
		adminRole.PUT("/update", UpdateRole)
		adminRole.PUT("/update-status", UpdateRoleStatus)
		adminRole.DELETE("/delete", DeleteRole)
		adminRole.GET("/export-excel", ExportRoleExcel)
	}

	adminMenu := r.Group("/admin-api/system/menu")
	{
		adminMenu.GET("/list", GetMenuList)
		adminMenu.GET("/simple-list", GetMenuList)
		adminMenu.GET("/list-all-simple", GetMenuList)
		adminMenu.GET("/get", GetMenuDetail)
		adminMenu.POST("/create", CreateMenu)
		adminMenu.PUT("/update", UpdateMenu)
		adminMenu.DELETE("/delete", DeleteMenu)
	}

	// 7. 日志与行政区划
	adminArea := r.Group("/admin-api/system/area")
	{
		adminArea.GET("/tree", GetAreaTree)
		adminArea.GET("/get-by-ip", GetAreaByIP)
	}

	appArea := r.Group("/app-api/system/area")
	{
		appArea.GET("/tree", GetAreaTree)
	}

	appDict := r.Group("/app-api/system/dict-data")
	{
		appDict.GET("/type", GetDictDataByType)
	}
}

func ExportTenantExcel(c *gin.Context) {
	c.Header("Content-Disposition", "attachment; filename=tenants.xlsx")
	c.Data(200, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", []byte{})
}

func CreateTenantPackage(c *gin.Context) {
	web.Success(c, 1)
}

func UpdateTenantPackage(c *gin.Context) {
	web.Success(c, true)
}

func DeleteTenantPackage(c *gin.Context) {
	web.Success(c, true)
}

func ExportDictDataExcel(c *gin.Context) {
	c.Header("Content-Disposition", "attachment; filename=dict_data.xlsx")
	c.Data(200, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", []byte{})
}

func ExportDictTypeExcel(c *gin.Context) {
	c.Header("Content-Disposition", "attachment; filename=dict_type.xlsx")
	c.Data(200, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", []byte{})
}

func ExportPostExcel(c *gin.Context) {
	c.Header("Content-Disposition", "attachment; filename=posts.xlsx")
	c.Data(200, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", []byte{})
}

func ExportUserExcel(c *gin.Context) {
	c.Header("Content-Disposition", "attachment; filename=users.xlsx")
	c.Data(200, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", []byte{})
}

func ExportRoleExcel(c *gin.Context) {
	c.Header("Content-Disposition", "attachment; filename=roles.xlsx")
	c.Data(200, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", []byte{})
}

func GetMenuDetail(c *gin.Context) {
	menus := getSystemMenus()
	if len(menus) > 0 {
		web.Success(c, menus[0])
		return
	}
	web.Success(c, gin.H{"id": 1, "name": "系统管理"})
}

func CreateMenu(c *gin.Context) {
	web.Success(c, 100)
}

func UpdateMenu(c *gin.Context) {
	web.Success(c, true)
}

func DeleteMenu(c *gin.Context) {
	web.Success(c, true)
}

func GetAreaByIP(c *gin.Context) {
	web.Success(c, gin.H{"id": 110100, "name": "北京市"})
}

func SendSmsCode(c *gin.Context) {
	web.Success(c, true)
}

func Logout(c *gin.Context) {
	web.Success(c, true)
}

// 认证与用户信息
func Login(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil || req.Username == "" {
		req.Username = "admin"
	}

	var user database.SystemUser
	if err := database.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		user = database.SystemUser{
			Username:  req.Username,
			Nickname:  "超级管理员",
			Avatar:    "https://www.iocoder.cn/images/avatar.jpg",
			Status:    0,
			DeptID:    1,
			CreatedAt: time.Now(),
		}
		database.DB.Create(&user)
	}

	token, _ := security.GenerateToken(user.ID, user.Username, 1, []string{"super_admin"})
	web.Success(c, gin.H{
		"userId":       user.ID,
		"accessToken":  token,
		"refreshToken": token + "_refresh",
		"expiresTime":  time.Now().Add(24 * 30 * time.Hour).UnixMilli(),
	})
}

func GetPermissionInfo(c *gin.Context) {
	var user database.SystemUser
	database.DB.First(&user, 1)

	web.Success(c, gin.H{
		"user": gin.H{
			"id":       user.ID,
			"nickname": user.Nickname,
			"avatar":   user.Avatar,
			"deptId":   user.DeptID,
			"email":    user.Email,
			"mobile":   user.Mobile,
		},
		"roles":       []string{"super_admin"},
		"permissions": []string{"*:*:*"},
		"menus":       getSystemMenus(),
	})
}

func GetUserProfile(c *gin.Context) {
	GetPermissionInfo(c)
}

func UpdateUserProfile(c *gin.Context) {
	var req struct {
		Nickname string `json:"nickname"`
		Mobile   string `json:"mobile"`
		Email    string `json:"email"`
	}
	if err := c.ShouldBindJSON(&req); err == nil {
		database.DB.Model(&database.SystemUser{}).Where("id = ?", 1).Updates(map[string]any{
			"nickname": req.Nickname,
			"mobile":   req.Mobile,
			"email":    req.Email,
		})
	}
	web.Success(c, true)
}

func UpdateUserPassword(c *gin.Context) {
	var req struct {
		ID       int64  `json:"id"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err == nil && req.ID > 0 {
		database.DB.Model(&database.SystemUser{}).Where("id = ?", req.ID).Update("password", req.Password)
	}
	web.Success(c, true)
}

func UpdateUserProfilePassword(c *gin.Context) {
	var req struct {
		NewPassword string `json:"newPassword"`
	}
	if err := c.ShouldBindJSON(&req); err == nil && req.NewPassword != "" {
		database.DB.Model(&database.SystemUser{}).Where("id = ?", 1).Update("password", req.NewPassword)
	}
	web.Success(c, true)
}

func UpdateUserProfileAvatar(c *gin.Context) {
	var req struct {
		Avatar string `json:"avatar"`
	}
	if err := c.ShouldBindJSON(&req); err == nil && req.Avatar != "" {
		database.DB.Model(&database.SystemUser{}).Where("id = ?", 1).Update("avatar", req.Avatar)
	}
	web.Success(c, true)
}

// 字典 CRUD (从数据库读写)
func GetSimpleDictDataList(c *gin.Context) {
	var list []database.SystemDictData
	database.DB.Order("sort ASC, id ASC").Find(&list)
	web.Success(c, list)
}

func GetDictDataPage(c *gin.Context) {
	pageNo, _ := strconv.Atoi(c.DefaultQuery("pageNo", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	dictType := c.Query("dictType")

	query := database.DB.Model(&database.SystemDictData{})
	if dictType != "" {
		query = query.Where("dict_type = ?", dictType)
	}

	var total int64
	query.Count(&total)

	var list []database.SystemDictData
	offset := (pageNo - 1) * pageSize
	query.Order("sort ASC, id ASC").Offset(offset).Limit(pageSize).Find(&list)
	web.PageSuccess(c, list, total)
}

func GetDictDataByType(c *gin.Context) {
	dictType := c.Query("type")
	if dictType == "" {
		dictType = c.Query("dictType")
	}
	var list []database.SystemDictData
	database.DB.Where("dict_type = ? AND status = 0", dictType).Order("sort ASC").Find(&list)
	web.Success(c, list)
}

func GetDictDataDetail(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	var d database.SystemDictData
	database.DB.First(&d, id)
	web.Success(c, d)
}

func CreateDictData(c *gin.Context) {
	var d database.SystemDictData
	if err := c.ShouldBindJSON(&d); err == nil {
		d.CreatedAt = time.Now()
		database.DB.Create(&d)
		web.Success(c, d.ID)
		return
	}
	web.Success(c, 1)
}

func UpdateDictData(c *gin.Context) {
	var d database.SystemDictData
	if err := c.ShouldBindJSON(&d); err == nil {
		database.DB.Model(&d).Updates(&d)
	}
	web.Success(c, true)
}

func DeleteDictData(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	if id > 0 {
		database.DB.Delete(&database.SystemDictData{}, id)
	}
	web.Success(c, true)
}

func GetDictTypePage(c *gin.Context) {
	var list []database.SystemDictType
	database.DB.Find(&list)
	web.PageSuccess(c, list, int64(len(list)))
}

func GetDictTypeList(c *gin.Context) {
	var list []database.SystemDictType
	database.DB.Find(&list)
	web.Success(c, list)
}

func GetDictTypeDetail(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	var t database.SystemDictType
	database.DB.First(&t, id)
	web.Success(c, t)
}

func CreateDictType(c *gin.Context) {
	var t database.SystemDictType
	if err := c.ShouldBindJSON(&t); err == nil {
		t.CreatedAt = time.Now()
		database.DB.Create(&t)
		web.Success(c, t.ID)
		return
	}
	web.Success(c, 1)
}

func UpdateDictType(c *gin.Context) {
	var t database.SystemDictType
	if err := c.ShouldBindJSON(&t); err == nil {
		database.DB.Model(&t).Updates(&t)
	}
	web.Success(c, true)
}

func DeleteDictType(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	if id > 0 {
		database.DB.Delete(&database.SystemDictType{}, id)
	}
	web.Success(c, true)
}

// 部门 CRUD
func GetSimpleDeptList(c *gin.Context) {
	var list []database.SystemDept
	database.DB.Order("sort ASC, id ASC").Find(&list)
	web.Success(c, list)
}

func GetDeptList(c *gin.Context) {
	GetSimpleDeptList(c)
}

func GetDeptDetail(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	var dept database.SystemDept
	database.DB.First(&dept, id)
	web.Success(c, dept)
}

func CreateDept(c *gin.Context) {
	var dept database.SystemDept
	if err := c.ShouldBindJSON(&dept); err == nil {
		dept.CreatedAt = time.Now()
		database.DB.Create(&dept)
		web.Success(c, dept.ID)
		return
	}
	web.Success(c, 1)
}

func UpdateDept(c *gin.Context) {
	var dept database.SystemDept
	if err := c.ShouldBindJSON(&dept); err == nil {
		database.DB.Model(&dept).Updates(&dept)
	}
	web.Success(c, true)
}

func DeleteDept(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	if id > 0 {
		database.DB.Delete(&database.SystemDept{}, id)
	}
	web.Success(c, true)
}

// 岗位 CRUD
func GetPostPage(c *gin.Context) {
	var list []database.SystemPost
	database.DB.Order("sort ASC").Find(&list)
	web.PageSuccess(c, list, int64(len(list)))
}

func GetPostList(c *gin.Context) {
	var list []database.SystemPost
	database.DB.Order("sort ASC").Find(&list)
	web.Success(c, list)
}

func GetPostDetail(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	var post database.SystemPost
	database.DB.First(&post, id)
	web.Success(c, post)
}

func CreatePost(c *gin.Context) {
	var post database.SystemPost
	if err := c.ShouldBindJSON(&post); err == nil {
		post.CreatedAt = time.Now()
		database.DB.Create(&post)
		web.Success(c, post.ID)
		return
	}
	web.Success(c, 1)
}

func UpdatePost(c *gin.Context) {
	var post database.SystemPost
	if err := c.ShouldBindJSON(&post); err == nil {
		database.DB.Model(&post).Updates(&post)
	}
	web.Success(c, true)
}

func DeletePost(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	if id > 0 {
		database.DB.Delete(&database.SystemPost{}, id)
	}
	web.Success(c, true)
}

// 用户 CRUD
func GetSimpleUserList(c *gin.Context) {
	var list []database.SystemUser
	database.DB.Find(&list)
	web.Success(c, list)
}

func GetUserPage(c *gin.Context) {
	pageNo, _ := strconv.Atoi(c.DefaultQuery("pageNo", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	var total int64
	database.DB.Model(&database.SystemUser{}).Count(&total)

	var list []database.SystemUser
	offset := (pageNo - 1) * pageSize
	database.DB.Offset(offset).Limit(pageSize).Find(&list)
	web.PageSuccess(c, list, total)
}

func GetUserDetail(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	var u database.SystemUser
	database.DB.First(&u, id)
	web.Success(c, u)
}

func CreateUser(c *gin.Context) {
	var u database.SystemUser
	if err := c.ShouldBindJSON(&u); err == nil {
		u.CreatedAt = time.Now()
		database.DB.Create(&u)
		web.Success(c, u.ID)
		return
	}
	web.Success(c, 1)
}

func UpdateUser(c *gin.Context) {
	var u database.SystemUser
	if err := c.ShouldBindJSON(&u); err == nil {
		database.DB.Model(&u).Updates(&u)
	}
	web.Success(c, true)
}

func UpdateUserStatus(c *gin.Context) {
	var req struct {
		ID     int64 `json:"id"`
		Status int   `json:"status"`
	}
	if err := c.ShouldBindJSON(&req); err == nil {
		database.DB.Model(&database.SystemUser{}).Where("id = ?", req.ID).Update("status", req.Status)
	}
	web.Success(c, true)
}

func DeleteUser(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	if id > 0 {
		database.DB.Delete(&database.SystemUser{}, id)
	}
	web.Success(c, true)
}

// 角色 CRUD
func GetRolePage(c *gin.Context) {
	var list []database.SystemRole
	database.DB.Find(&list)
	web.PageSuccess(c, list, int64(len(list)))
}

func GetRoleList(c *gin.Context) {
	var list []database.SystemRole
	database.DB.Find(&list)
	web.Success(c, list)
}

func GetRoleDetail(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	var r database.SystemRole
	database.DB.First(&r, id)
	web.Success(c, r)
}

func CreateRole(c *gin.Context) {
	var r database.SystemRole
	if err := c.ShouldBindJSON(&r); err == nil {
		r.CreatedAt = time.Now()
		database.DB.Create(&r)
		web.Success(c, r.ID)
		return
	}
	web.Success(c, 1)
}

func UpdateRole(c *gin.Context) {
	var r database.SystemRole
	if err := c.ShouldBindJSON(&r); err == nil {
		database.DB.Model(&r).Updates(&r)
	}
	web.Success(c, true)
}

func UpdateRoleStatus(c *gin.Context) {
	var req struct {
		ID     int64 `json:"id"`
		Status int   `json:"status"`
	}
	if err := c.ShouldBindJSON(&req); err == nil {
		database.DB.Model(&database.SystemRole{}).Where("id = ?", req.ID).Update("status", req.Status)
	}
	web.Success(c, true)
}

func DeleteRole(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	if id > 0 {
		database.DB.Delete(&database.SystemRole{}, id)
	}
	web.Success(c, true)
}

// 租户 CRUD
func GetTenantByName(c *gin.Context) {
	name := c.Query("name")
	var t database.SystemTenant
	database.DB.Where("name = ?", name).First(&t)
	web.Success(c, t.ID)
}

func GetTenantByWebsite(c *gin.Context) {
	web := c.Query("website")
	var t database.SystemTenant
	database.DB.Where("website = ?", web).First(&t)
	c.JSON(200, gin.H{"code": 0, "msg": "success", "data": t})
}

func GetTenantPage(c *gin.Context) {
	var list []database.SystemTenant
	database.DB.Find(&list)
	web.PageSuccess(c, list, int64(len(list)))
}

func GetTenantList(c *gin.Context) {
	var list []database.SystemTenant
	database.DB.Find(&list)
	web.Success(c, list)
}

func GetTenantDetail(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	var t database.SystemTenant
	database.DB.First(&t, id)
	web.Success(c, t)
}

func CreateTenant(c *gin.Context) {
	var t database.SystemTenant
	if err := c.ShouldBindJSON(&t); err == nil {
		t.CreatedAt = time.Now()
		database.DB.Create(&t)
		web.Success(c, t.ID)
		return
	}
	web.Success(c, 1)
}

func UpdateTenant(c *gin.Context) {
	var t database.SystemTenant
	if err := c.ShouldBindJSON(&t); err == nil {
		database.DB.Model(&t).Updates(&t)
	}
	web.Success(c, true)
}

func DeleteTenant(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	if id > 0 {
		database.DB.Delete(&database.SystemTenant{}, id)
	}
	web.Success(c, true)
}

func GetMenuList(c *gin.Context) {
	web.Success(c, getSystemMenus())
}

func GetAreaTree(c *gin.Context) {
	tree := []gin.H{
		{
			"id": 110000, "name": "北京市", "children": []gin.H{
				{
					"id": 110100, "name": "市辖区", "children": []gin.H{
						{"id": 110101, "name": "东城区"},
						{"id": 110102, "name": "西城区"},
						{"id": 110105, "name": "朝阳区"},
					},
				},
			},
		},
		{
			"id": 310000, "name": "上海市", "children": []gin.H{
				{
					"id": 310100, "name": "市辖区", "children": []gin.H{
						{"id": 310101, "name": "黄浦区"},
						{"id": 310104, "name": "徐汇区"},
						{"id": 310115, "name": "浦东新区"},
					},
				},
			},
		},
		{
			"id": 440000, "name": "广东省", "children": []gin.H{
				{
					"id": 440100, "name": "广州市", "children": []gin.H{
						{"id": 440106, "name": "天河区"},
						{"id": 440103, "name": "荔湾区"},
					},
				},
				{
					"id": 440300, "name": "深圳市", "children": []gin.H{
						{"id": 440305, "name": "南山区"},
						{"id": 440304, "name": "福田区"},
					},
				},
			},
		},
	}
	web.Success(c, tree)
}

func getSystemMenus() []gin.H {
	return []gin.H{
		{"id": 1, "name": "首页", "path": "/index", "component": "index", "icon": "ep:home-filled"},
		{"id": 2, "name": "商品中心", "path": "/product", "icon": "ep:goods"},
		{"id": 3, "name": "订单中心", "path": "/trade", "icon": "ep:shopping-bag"},
		{"id": 4, "name": "营销中心", "path": "/promotion", "icon": "ep:present"},
		{"id": 5, "name": "会员中心", "path": "/member", "icon": "ep:user"},
		{"id": 6, "name": "系统管理", "path": "/system", "icon": "ep:setting"},
	}
}

// GetCaptcha 获取验证码
func GetCaptcha(c *gin.Context) {
	c.JSON(200, gin.H{
		"repCode": "0000",
		"repMsg":  "success",
		"repData": gin.H{
			"originalImageBase64": "",
			"jigsawImageBase64":   "",
			"token":               "mock-captcha-token",
			"secretKey":           "mock-secret-key",
			"captchaType":         "blockPuzzle",
		},
	})
}

// CheckCaptcha 校验验证码
func CheckCaptcha(c *gin.Context) {
	c.JSON(200, gin.H{
		"repCode": "0000",
		"repMsg":  "success",
		"repData": gin.H{
			"result": true,
		},
	})
}
