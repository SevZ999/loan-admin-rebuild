package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/small-ek/antgo/os/config"
	"github.com/small-ek/antgo/utils/ajwt"
	"github.com/small-ek/antgo/utils/conv"
	"strings"
)

// AuthJWT JWT 不认证
func NoAuthJWT() gin.HandlerFunc {
	newJwt := ajwt.New().SetPublicKey(conv.Bytes(config.GetString("jwt.public_key")))

	return func(c *gin.Context) {
		var token = c.Request.Header.Get("Authorization")
		if token != "" {
			token = strings.TrimPrefix(token, "Bearer ")
			getAuthorization, _ := newJwt.Parse(token)
			c.Set("user_id", getAuthorization["id"])
		}
		c.Next()
	}
}
