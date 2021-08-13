package middleware

import (
	"fmt"
	"gin-admin/internal/errcode"
	"gin-admin/pkg/app"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if len(header) == 0 {
			app.WriteResponse(c, errcode.Unauthorized, gin.H{})
			c.Abort()
			return
		}
		var token string
		fmt.Sscanf(header, "Bearer %s", &token)

		if token == "" {
			app.WriteResponse(c, errcode.Unauthorized, gin.H{})
			c.Abort()
			return
		}
		claims, err := app.ParseToken(token)

		if err != nil {
			switch err.(*jwt.ValidationError).Errors {
			case jwt.ValidationErrorExpired:
				app.WriteResponse(c, errcode.UnauthorizedTokenTimeout, gin.H{})
				c.Abort()
				return
			default:
				app.WriteResponse(c, errcode.UnauthorizedTokenError, gin.H{})
				c.Abort()
				return
			}
		}

		c.Set("admin_user_id", claims.UserID)
		c.Next()
	}
}
