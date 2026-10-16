package main

import (
	"flag"
	"log"

	"server-go/internal/framework/database"
	"server-go/internal/framework/security"
	"server-go/internal/module/extended"
	"server-go/internal/module/infra"
	"server-go/internal/module/member"
	"server-go/internal/module/pay"
	"server-go/internal/module/product"
	"server-go/internal/module/promotion"
	"server-go/internal/module/statistics"
	"server-go/internal/module/system"
	"server-go/internal/module/trade"

	"github.com/gin-gonic/gin"
)

func main() {
	port := flag.String("port", "48080", "Server listening port")
	dbPath := flag.String("db", "ruoyi_mall.db", "SQLite database file path")
	flag.Parse()

	// 1. 初始化数据库
	database.InitDB(*dbPath)

	// 2. 创建 Gin 引擎
	r := gin.Default()

	// 3. 注册全局跨域中间件
	r.Use(security.CorsMiddleware())

	// 4. 注册全量 1:1 重构版业务模块路由
	system.RegisterSystemRoutes(r)
	infra.RegisterInfraRoutes(r)
	member.RegisterMemberRoutes(r)
	product.RegisterProductRoutes(r)
	trade.RegisterTradeRoutes(r)
	promotion.RegisterPromotionRoutes(r)
	pay.RegisterPayRoutes(r)
	statistics.RegisterStatisticsRoutes(r)
	extended.RegisterExtendedRoutes(r)

	// 5. 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "UP",
			"engine":  "Go + Gin",
			"project": "Ruoyi-Vue-Pro Go Refactored Backend",
		})
	})

	log.Printf("==========================================================")
	log.Printf("🚀 Ruoyi-Vue-Pro Go + Gin 1:1 重构后端服务启动成功！")
	log.Printf("👉 服务监听地址: http://0.0.0.0:%s", *port)
	log.Printf("👉 管理后台 API : http://0.0.0.0:%s/admin-api/*", *port)
	log.Printf("👉 移动小程序 API: http://0.0.0.0:%s/app-api/*", *port)
	log.Printf("==========================================================")

	if err := r.Run(":" + *port); err != nil {
		log.Fatalf("Server startup failed: %v", err)
	}
}
