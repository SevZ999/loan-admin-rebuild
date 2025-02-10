package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/small-ek/antgo/os/config"
	"github.com/small-ek/antgo/utils/ajwt"
	"github.com/small-ek/antgo/utils/conv"
	"github.com/small-ek/antgo/utils/response"
	"loan-ext/app/dao"
	"loan-ext/app/entity/models"
	"loan-ext/app/entity/vo"
	"strings"
)

// 常量定义
const (
	UnauthorizedCode = 401
)

// AuthJWT JWT 认证
func AuthJWT() gin.HandlerFunc {
	newJwt := ajwt.New().SetPublicKey(conv.Bytes(config.GetString("jwt.public_key")))

	return func(c *gin.Context) {
		var token = c.Request.Header.Get("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(UnauthorizedCode, response.Fail(vo.AUTH_ERROR, vo.GetMessage(vo.AUTH_ERROR)))
			return
		}

		token = strings.TrimPrefix(token, "Bearer ")
		getAuthorization, err := newJwt.Parse(token)
		if err != nil {
			c.AbortWithStatusJSON(UnauthorizedCode, response.Fail(vo.AUTH_ERROR, vo.GetMessage(vo.AUTH_ERROR), err.Error()))
			return
		}

		if c.GetHeader("device-id") != getAuthorization["device-id"] {
			c.AbortWithStatusJSON(UnauthorizedCode, response.Fail(vo.AUTH_DEVICE_ERROR, vo.GetMessage(vo.AUTH_DEVICE_ERROR)))
			return
		}

		getUser := getUserByUser(conv.Int(getAuthorization["id"]))
		if getUser.Id == 0 {
			c.AbortWithStatusJSON(UnauthorizedCode, response.Fail(vo.AUTH_USER_ERROR, vo.GetMessage(vo.AUTH_USER_ERROR)))
			return
		}
		c.Set("user", getUser)
		c.Set("user_id", getUser.Id)
		c.Next()
	}
}

// getUserByUser 通过 ID 获取用户信息
func getUserByUser(userID int) models.Users {
	user := dao.NewUsersDao(nil).GetById(userID)
	return user
}
