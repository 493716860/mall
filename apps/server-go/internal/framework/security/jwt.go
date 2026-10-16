package security

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"server-go/internal/framework/web"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var JwtSecret = []byte("ruoyi-vue-pro-go-secret-key-2026")

type Claims struct {
	UserID   int64    `json:"userId"`
	Username string   `json:"username"`
	UserType int      `json:"userType"` // 1: Admin, 2: App Member
	Roles    []string `json:"roles"`
	jwt.RegisteredClaims
}

// GenerateToken 生成 JWT Token
func GenerateToken(userID int64, username string, userType int, roles []string) (string, error) {
	claims := Claims{
		UserID:   userID,
		Username: username,
		UserType: userType,
		Roles:    roles,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * 30 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtSecret)
}

// ParseToken 解析 JWT Token
func ParseToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// CorsMiddleware 跨域处理中间件
func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, tenant-id, terminal")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}

// AuthMiddleware 认证鉴权中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			// Demo 模式允许默认超级管理员测试，或者打回 401
			c.Set("userID", int64(1))
			c.Set("username", "admin")
			c.Set("roles", []string{"super_admin"})
			c.Next()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) == 2 && parts[0] == "Bearer" {
			claims, err := ParseToken(parts[1])
			if err == nil && claims != nil {
				c.Set("userID", claims.UserID)
				c.Set("username", claims.Username)
				c.Set("userType", claims.UserType)
				c.Set("roles", claims.Roles)
				c.Next()
				return
			}
		}

		web.Unauthorized(c, "访问令牌已过期或无效")
		c.Abort()
	}
}
