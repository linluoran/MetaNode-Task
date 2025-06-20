package middleware

import (
	"fmt"
	"gin_blog/internal/config"
	"gin_blog/internal/pkg/logger"
	"gin_blog/internal/types"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

func CustomRecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Log.Error("中间件错误.", zap.Any("err", err))
				types.ErrorRes(c, 0, "系统内部错误")
			}
		}()
		c.Next()
	}
}

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 从 Header 获取 token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			types.ErrorRes(c, http.StatusUnauthorized, "未提供认证令牌")
			return
		}

		// 移除 "Bearer " 前缀
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// 3. 使用自定义 Claims 解析 Token
		claims := &types.CustomClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			// 验证签名算法
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("非预期的签名算法: %v", token.Header["alg"])
			}
			return []byte(config.GlobalConfig.Jwt.Secret), nil
		})

		// 4. 处理解析错误
		if err != nil || !token.Valid {
			types.ErrorRes(c, http.StatusUnauthorized, "令牌无效或已过期")
			return
		}

		// 5. 验证签发者（可选）
		if claims.Issuer != config.GlobalConfig.Jwt.Issuer {
			types.ErrorRes(c, http.StatusUnauthorized, "签发机构错误")
			return
		}

		// 6. 存储到上下文
		c.Set("userID", claims.ID)
		c.Set("username", claims.Username)
		c.Next()
	}
}
