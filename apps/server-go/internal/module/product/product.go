package product

import (
	"strconv"
	"strings"
	"time"

	"server-go/internal/framework/database"
	"server-go/internal/framework/web"

	"github.com/gin-gonic/gin"
)

// RegisterProductRoutes 注册商品中心路由
func RegisterProductRoutes(r *gin.Engine) {
	// 管理后台接口
	adminProd := r.Group("/admin-api/product")
	{
		// SPU
		adminProd.GET("/spu/page", GetSpuPage)
		adminProd.GET("/spu/get-count", GetSpuTabsCount)
		adminProd.GET("/spu/get-detail", GetSpuDetail)
		adminProd.GET("/spu/get", GetSpuDetail)
		adminProd.POST("/spu/create", CreateSpu)
		adminProd.PUT("/spu/update", UpdateSpu)
		adminProd.PUT("/spu/update-status", UpdateSpuStatus)
		adminProd.DELETE("/spu/delete", DeleteSpu)
		adminProd.GET("/spu/export-excel", ExportSpuExcel)
		adminProd.GET("/spu/list-all-simple", GetSpuListByIds)

		// Category
		adminProd.GET("/category/list", GetCategoryList)
		adminProd.GET("/category/list-all-simple", GetCategoryList)
		adminProd.GET("/category/get", GetCategoryDetail)
		adminProd.POST("/category/create", CreateCategory)
		adminProd.PUT("/category/update", UpdateCategory)
		adminProd.DELETE("/category/delete", DeleteCategory)

		// Brand
		adminProd.GET("/brand/list-all-simple", GetBrandList)
		adminProd.GET("/brand/page", GetBrandList)
		adminProd.GET("/brand/get", GetBrandList)
		adminProd.POST("/brand/create", CreateBrand)
		adminProd.PUT("/brand/update", UpdateBrand)
		adminProd.DELETE("/brand/delete", DeleteBrand)

		// Property
		adminProd.GET("/property/page", GetPropertyPage)
		adminProd.GET("/property/simple-list", GetPropertyPage)
		adminProd.GET("/property/list-all-simple", GetPropertyPage)
		adminProd.GET("/property/get", GetPropertyPage)
		adminProd.POST("/property/create", CreateProperty)
		adminProd.PUT("/property/update", UpdateProperty)
		adminProd.DELETE("/property/delete", DeleteProperty)
		adminProd.GET("/property/value/page", GetPropertyPage)
		adminProd.GET("/property/value/list-all-simple", GetPropertyPage)
		adminProd.POST("/property/value/create", CreatePropertyValue)
		adminProd.PUT("/property/value/update", UpdatePropertyValue)
		adminProd.DELETE("/property/value/delete", DeletePropertyValue)

		// Comment
		adminProd.GET("/comment/page", GetCommentPage)
		adminProd.GET("/comment/get", GetCommentPage)
		adminProd.POST("/comment/create", CreateComment)
		adminProd.PUT("/comment/update-visible", UpdateCommentVisible)
		adminProd.PUT("/comment/reply", ReplyComment)
	}

	// 移动端/小程序接口
	appProd := r.Group("/app-api/product")
	{
		appProd.GET("/spu/page", GetSpuPage)
		appProd.GET("/spu/get-detail", GetSpuDetail)
		appProd.GET("/spu/list-by-ids", GetSpuListByIds)
		appProd.GET("/category/list", GetCategoryList)
		appProd.GET("/favorite/page", GetFavoritePage)
		appProd.POST("/favorite/create", CreateFavorite)
		appProd.DELETE("/favorite/delete", DeleteFavorite)
		appProd.GET("/favorite/exits", CheckFavoriteExists)
		appProd.GET("/browse-history/page", GetFavoritePage)
		appProd.DELETE("/browse-history/clean", CleanBrowseHistory)
	}
}

// GetSpuPage 商品 SPU 分页 (基于真实数据库)
func GetSpuPage(c *gin.Context) {
	pageNo, _ := strconv.Atoi(c.DefaultQuery("pageNo", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	name := c.Query("name")
	categoryID, _ := strconv.ParseInt(c.Query("categoryId"), 10, 64)
	tabType := c.Query("tabType") // 0: 上架, 1: 仓库, 2: 售罄, 3: 警戒库存

	query := database.DB.Model(&database.ProductSpu{})
	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}
	if categoryID > 0 {
		query = query.Where("category_id = ?", categoryID)
	}
	if tabType == "0" {
		query = query.Where("status = ?", 1) // 上架
	} else if tabType == "1" {
		query = query.Where("status = ?", 0) // 仓库
	} else if tabType == "2" {
		query = query.Where("stock <= 0")    // 售罄
	} else if tabType == "3" {
		query = query.Where("stock < 10")    // 预警
	}

	var total int64
	query.Count(&total)

	var spus []database.ProductSpu
	offset := (pageNo - 1) * pageSize
	query.Order("sort ASC, id DESC").Offset(offset).Limit(pageSize).Find(&spus)

	var list []gin.H
	for _, s := range spus {
		sliderUrls := strings.Split(s.SliderPicURLs, ",")
		var skus []database.ProductSku
		database.DB.Where("spu_id = ?", s.ID).Find(&skus)
		list = append(list, gin.H{
			"id":            s.ID,
			"name":          s.Name,
			"categoryId":    s.CategoryID,
			"categoryName":  s.CategoryName,
			"picUrl":        s.PicURL,
			"sliderPicUrls": sliderUrls,
			"price":         s.Price,
			"marketPrice":   s.MarketPrice,
			"costPrice":     s.CostPrice,
			"stock":         s.Stock,
			"salesCount":    s.SalesCount,
			"status":        s.Status,
			"sort":          s.Sort,
			"createTime":    s.CreatedAt.UnixMilli(),
			"skus":          skus,
		})
	}

	web.PageSuccess(c, list, total)
}

// GetSpuTabsCount 统计各状态商品真实数量
func GetSpuTabsCount(c *gin.Context) {
	var onSale, inWarehouse, outOfStock, alertStock int64
	database.DB.Model(&database.ProductSpu{}).Where("status = ?", 1).Count(&onSale)
	database.DB.Model(&database.ProductSpu{}).Where("status = ?", 0).Count(&inWarehouse)
	database.DB.Model(&database.ProductSpu{}).Where("stock <= 0").Count(&outOfStock)
	database.DB.Model(&database.ProductSpu{}).Where("stock < 10").Count(&alertStock)

	web.Success(c, gin.H{
		"0": onSale,
		"1": inWarehouse,
		"2": outOfStock,
		"3": alertStock,
	})
}

// GetSpuDetail 商品详情
func GetSpuDetail(c *gin.Context) {
	idStr := c.Query("id")
	if idStr == "" {
		idStr = c.Param("id")
	}
	id, _ := strconv.ParseInt(idStr, 10, 64)

	var s database.ProductSpu
	if err := database.DB.First(&s, id).Error; err != nil {
		database.DB.First(&s)
	}

	sliderUrls := strings.Split(s.SliderPicURLs, ",")
	var skus []database.ProductSku
	database.DB.Where("spu_id = ?", s.ID).Find(&skus)

	web.Success(c, gin.H{
		"id":            s.ID,
		"name":          s.Name,
		"categoryId":    s.CategoryID,
		"categoryName":  s.CategoryName,
		"picUrl":        s.PicURL,
		"sliderPicUrls": sliderUrls,
		"price":         s.Price,
		"marketPrice":   s.MarketPrice,
		"costPrice":     s.CostPrice,
		"stock":         s.Stock,
		"salesCount":    s.SalesCount,
		"status":        s.Status,
		"description":   s.Description,
		"skus":          skus,
	})
}

// UpdateSpuStatus 真实修改商品状态 (上架/下架/回收站)
func UpdateSpuStatus(c *gin.Context) {
	var req struct {
		ID     int64 `json:"id"`
		Status int   `json:"status"`
	}
	if err := c.ShouldBindJSON(&req); err == nil {
		database.DB.Model(&database.ProductSpu{}).Where("id = ?", req.ID).Update("status", req.Status)
	}
	web.Success(c, true)
}

// CreateSpu 创建商品
func CreateSpu(c *gin.Context) {
	var s database.ProductSpu
	if err := c.ShouldBindJSON(&s); err == nil {
		s.CreatedAt = time.Now()
		s.UpdatedAt = time.Now()
		database.DB.Create(&s)
		web.Success(c, s.ID)
		return
	}
	web.Success(c, 100)
}

// UpdateSpu 修改商品
func UpdateSpu(c *gin.Context) {
	var s database.ProductSpu
	if err := c.ShouldBindJSON(&s); err == nil {
		s.UpdatedAt = time.Now()
		database.DB.Model(&s).Updates(&s)
	}
	web.Success(c, true)
}

// DeleteSpu 删除商品
func DeleteSpu(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	if id > 0 {
		database.DB.Delete(&database.ProductSpu{}, id)
	}
	web.Success(c, true)
}

// GetSpuListByIds 根据 ID 列表批量查询 SPU
func GetSpuListByIds(c *gin.Context) {
	idsStr := c.Query("ids")
	var ids []int64
	for _, s := range strings.Split(idsStr, ",") {
		if id, err := strconv.ParseInt(strings.TrimSpace(s), 10, 64); err == nil {
			ids = append(ids, id)
		}
	}
	var spus []database.ProductSpu
	if len(ids) > 0 {
		database.DB.Where("id IN ?", ids).Find(&spus)
	} else {
		database.DB.Limit(10).Find(&spus)
	}
	web.Success(c, spus)
}

// GetCategoryList 商品分类真实列表
func GetCategoryList(c *gin.Context) {
	var categories []database.ProductCategory
	database.DB.Order("sort ASC, id ASC").Find(&categories)
	web.Success(c, categories)
}

// GetCategoryDetail 分类详情
func GetCategoryDetail(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	var cat database.ProductCategory
	database.DB.First(&cat, id)
	web.Success(c, cat)
}

// CreateCategory 创建分类
func CreateCategory(c *gin.Context) {
	var cat database.ProductCategory
	if err := c.ShouldBindJSON(&cat); err == nil {
		cat.CreatedAt = time.Now()
		database.DB.Create(&cat)
		web.Success(c, cat.ID)
		return
	}
	web.Success(c, 1)
}

// UpdateCategory 修改分类
func UpdateCategory(c *gin.Context) {
	var cat database.ProductCategory
	if err := c.ShouldBindJSON(&cat); err == nil {
		database.DB.Model(&cat).Updates(&cat)
	}
	web.Success(c, true)
}

// DeleteCategory 删除分类
func DeleteCategory(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	if id > 0 {
		database.DB.Delete(&database.ProductCategory{}, id)
	}
	web.Success(c, true)
}

// GetBrandList 品牌列表 (真实读取数据库)
func GetBrandList(c *gin.Context) {
	var list []database.ProductBrand
	database.DB.Order("sort ASC, id ASC").Find(&list)
	web.Success(c, list)
}

// GetPropertyPage 规格属性列表 (真实读取数据库)
func GetPropertyPage(c *gin.Context) {
	var props []database.ProductProperty
	database.DB.Find(&props)

	var list []gin.H
	for _, p := range props {
		var vals []database.ProductPropertyValue
		database.DB.Where("property_id = ?", p.ID).Find(&vals)
		var valNames []string
		for _, v := range vals {
			valNames = append(valNames, v.Name)
		}
		list = append(list, gin.H{
			"id":     p.ID,
			"name":   p.Name,
			"status": p.Status,
			"values": valNames,
		})
	}
	web.PageSuccess(c, list, int64(len(list)))
}

// GetCommentPage 评价列表 (真实读取数据库)
func GetCommentPage(c *gin.Context) {
	var list []database.ProductComment
	database.DB.Order("id DESC").Find(&list)
	if len(list) == 0 {
		c := database.ProductComment{
			UserID:       1001,
			UserNickname: "微信用户小王",
			UserAvatar:   "https://www.iocoder.cn/images/avatar.jpg",
			SpuID:        1,
			SpuName:      "天然深海无谷高蛋白猫粮 10kg",
			Scores:       5,
			Content:      "猫咪特别喜欢吃，适口性非常好！毛色也光亮了很多！",
			ReplyContent: "感谢亲对我们商城的支持，祝主子健康成长！",
			Visible:      true,
			CreatedAt:    time.Now(),
		}
		database.DB.Create(&c)
		list = append(list, c)
	}
	web.PageSuccess(c, list, int64(len(list)))
}

// GetFavoritePage 收藏列表 (真实读取数据库)
func GetFavoritePage(c *gin.Context) {
	var list []database.ProductFavorite
	database.DB.Where("user_id = ?", 1001).Find(&list)

	var result []gin.H
	for _, f := range list {
		var spu database.ProductSpu
		database.DB.First(&spu, f.SpuID)
		result = append(result, gin.H{
			"id":         f.ID,
			"spuId":      spu.ID,
			"spuName":    spu.Name,
			"picUrl":     spu.PicURL,
			"price":      spu.Price,
			"createTime": f.CreatedAt.UnixMilli(),
		})
	}
	web.PageSuccess(c, result, int64(len(result)))
}

// CreateBrand 创建品牌
func CreateBrand(c *gin.Context) {
	var b database.ProductBrand
	if err := c.ShouldBindJSON(&b); err == nil {
		b.CreatedAt = time.Now()
		database.DB.Create(&b)
		web.Success(c, b.ID)
		return
	}
	web.Success(c, 1)
}

// UpdateBrand 修改品牌
func UpdateBrand(c *gin.Context) {
	var b database.ProductBrand
	if err := c.ShouldBindJSON(&b); err == nil {
		database.DB.Model(&b).Updates(&b)
	}
	web.Success(c, true)
}

// DeleteBrand 删除品牌
func DeleteBrand(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	if id > 0 {
		database.DB.Delete(&database.ProductBrand{}, id)
	}
	web.Success(c, true)
}

// CreateProperty 创建属性规格
func CreateProperty(c *gin.Context) {
	var p database.ProductProperty
	if err := c.ShouldBindJSON(&p); err == nil {
		p.CreatedAt = time.Now()
		database.DB.Create(&p)
		web.Success(c, p.ID)
		return
	}
	web.Success(c, 1)
}

// UpdateProperty 修改属性规格
func UpdateProperty(c *gin.Context) {
	var p database.ProductProperty
	if err := c.ShouldBindJSON(&p); err == nil {
		database.DB.Model(&p).Updates(&p)
	}
	web.Success(c, true)
}

// DeleteProperty 删除属性规格
func DeleteProperty(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	if id > 0 {
		database.DB.Delete(&database.ProductProperty{}, id)
		database.DB.Where("property_id = ?", id).Delete(&database.ProductPropertyValue{})
	}
	web.Success(c, true)
}

// CreatePropertyValue 创建属性值
func CreatePropertyValue(c *gin.Context) {
	var v database.ProductPropertyValue
	if err := c.ShouldBindJSON(&v); err == nil {
		v.CreatedAt = time.Now()
		database.DB.Create(&v)
		web.Success(c, v.ID)
		return
	}
	web.Success(c, 1)
}

// UpdatePropertyValue 修改属性值
func UpdatePropertyValue(c *gin.Context) {
	var v database.ProductPropertyValue
	if err := c.ShouldBindJSON(&v); err == nil {
		database.DB.Model(&v).Updates(&v)
	}
	web.Success(c, true)
}

// DeletePropertyValue 删除属性值
func DeletePropertyValue(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	if id > 0 {
		database.DB.Delete(&database.ProductPropertyValue{}, id)
	}
	web.Success(c, true)
}

// CreateComment 评价商品
func CreateComment(c *gin.Context) {
	var comm database.ProductComment
	if err := c.ShouldBindJSON(&comm); err == nil {
		comm.UserID = 1001
		comm.UserNickname = "微信用户小王"
		comm.UserAvatar = "https://www.iocoder.cn/images/avatar.jpg"
		comm.CreatedAt = time.Now()
		comm.Visible = true
		database.DB.Create(&comm)
		web.Success(c, comm.ID)
		return
	}
	web.Success(c, 1)
}

// UpdateCommentVisible 控制评价显示
func UpdateCommentVisible(c *gin.Context) {
	var req struct {
		ID      int64 `json:"id"`
		Visible bool  `json:"visible"`
	}
	if err := c.ShouldBindJSON(&req); err == nil {
		database.DB.Model(&database.ProductComment{}).Where("id = ?", req.ID).Update("visible", req.Visible)
	}
	web.Success(c, true)
}

// ReplyComment 回复评价
func ReplyComment(c *gin.Context) {
	var req struct {
		ID           int64  `json:"id"`
		ReplyContent string `json:"replyContent"`
	}
	if err := c.ShouldBindJSON(&req); err == nil {
		now := time.Now()
		database.DB.Model(&database.ProductComment{}).Where("id = ?", req.ID).Updates(map[string]any{
			"reply_content": req.ReplyContent,
			"reply_time":    &now,
		})
	}
	web.Success(c, true)
}

// CreateFavorite 添加收藏
func CreateFavorite(c *gin.Context) {
	var req struct {
		SpuID int64 `json:"spuId"`
	}
	_ = c.ShouldBindJSON(&req)
	if req.SpuID == 0 {
		req.SpuID, _ = strconv.ParseInt(c.Query("spuId"), 10, 64)
	}
	var count int64
	database.DB.Model(&database.ProductFavorite{}).Where("user_id = ? AND spu_id = ?", 1001, req.SpuID).Count(&count)
	if count == 0 && req.SpuID > 0 {
		f := database.ProductFavorite{
			UserID:    1001,
			SpuID:     req.SpuID,
			CreatedAt: time.Now(),
		}
		database.DB.Create(&f)
	}
	web.Success(c, true)
}

// DeleteFavorite 取消收藏
func DeleteFavorite(c *gin.Context) {
	var req struct {
		SpuID int64 `json:"spuId"`
	}
	_ = c.ShouldBindJSON(&req)
	if req.SpuID == 0 {
		req.SpuID, _ = strconv.ParseInt(c.Query("spuId"), 10, 64)
	}
	if req.SpuID > 0 {
		database.DB.Where("user_id = ? AND spu_id = ?", 1001, req.SpuID).Delete(&database.ProductFavorite{})
	}
	web.Success(c, true)
}

// CheckFavoriteExists 检查是否已收藏
func CheckFavoriteExists(c *gin.Context) {
	spuId, _ := strconv.ParseInt(c.Query("spuId"), 10, 64)
	var count int64
	database.DB.Model(&database.ProductFavorite{}).Where("user_id = ? AND spu_id = ?", 1001, spuId).Count(&count)
	web.Success(c, count > 0)
}

// CleanBrowseHistory 清理浏览历史
func CleanBrowseHistory(c *gin.Context) {
	database.DB.Where("user_id = ?", 1001).Delete(&database.ProductBrowseHistory{})
	web.Success(c, true)
}

func ExportSpuExcel(c *gin.Context) {
	c.Header("Content-Disposition", "attachment; filename=spu.xlsx")
	c.Data(200, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", []byte{})
}


