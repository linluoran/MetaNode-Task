package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// GenerateToken 生成 JWT 令牌
// 参数: userID - 用户唯一标识, roles - 用户角色列表
// 返回: JWT 字符串和错误信息
func GenerateToken(userID string, roles []string) (string, error) {
	// 从环境变量获取密钥（生产环境应使用密钥管理系统）
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", fmt.Errorf("JWT_SECRET 环境变量未设置")
	}

	// 创建带声明的令牌 [2,6](@ref)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,                               // 用户标识
		"roles":  roles,                                // 用户角色列表
		"exp":    time.Now().Add(8 * time.Hour).Unix(), // 8小时后过期
		"iat":    time.Now().Unix(),                    // 签发时间
		"iss":    "your-app-name",                      // 签发者标识
	})

	// 使用密钥签名并返回
	return token.SignedString([]byte(secret))
}

// JWTAuth JWT 认证中间件
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从 Authorization 头提取令牌
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "缺少认证令牌"})
			return
		}

		// 移除 "Bearer " 前缀（长度7）
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// 解析并验证令牌
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// 验证签名算法
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("非预期的签名方法: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		// 处理解析错误
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "无效令牌: " + err.Error()})
			return
		}

		// 验证令牌有效性并提取声明 [4](@ref)
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// 将用户信息存入 Gin 上下文
			c.Set("userID", claims["userID"])
			c.Set("roles", claims["roles"])
			c.Next() // 继续处理请求
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "无效或过期的令牌"})
		}
	}
}

// RequireRole 角色验证中间件
func RequireRole(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从上下文获取角色列表
		rolesInterface, exists := c.Get("roles")
		if !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "未找到角色信息"})
			return
		}

		// 类型断言转换为字符串切片
		roles, ok := rolesInterface.([]interface{})
		if !ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "角色格式错误"})
			return
		}

		// 检查是否包含所需角色 [6](@ref)
		hasRole := false
		for _, r := range roles {
			if roleStr, ok := r.(string); ok && roleStr == requiredRole {
				hasRole = true
				break
			}
		}

		if !hasRole {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "权限不足"})
			return
		}

		c.Next() // 继续处理请求
	}
}

// CORSMiddleware 跨域资源共享配置
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "https://prod.com, http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Max-Age", fmt.Sprintf("%.0f", 12*time.Hour.Seconds()))

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
