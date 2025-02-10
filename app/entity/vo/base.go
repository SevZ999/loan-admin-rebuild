package vo

import (
	"github.com/gin-gonic/gin"
	"github.com/small-ek/antgo/os/alog"
	"github.com/small-ek/antgo/os/config"
	"github.com/small-ek/antgo/utils/conv"
	"github.com/small-ek/antgo/utils/response"
	"go.uber.org/zap"
	"loan-ext/app/entity/models"
	"sync"
)

type Base struct {
	Code int `json:"code"` //Status Code
}

var (
	baseInstance *Base
	once         sync.Once
)

// NewBase 实例化
func NewBase() *Base {
	once.Do(func() {
		baseInstance = &Base{}
	})
	return baseInstance
}

// Success 成功返回
func (b *Base) Success(c *gin.Context, code string, data ...interface{}) {
	msg := GetMessage(code)
	c.SecureJSON(200, response.Success(code, msg, data...))
}

// Fail 错误返回
func (b *Base) Fail(c *gin.Context, code string, err ...error) {
	if len(err) > 0 && err[0] != nil {
		alog.Write.Error("Return error", zap.Errors("Fail Error", err))
	}
	msg := GetMessage(code)

	if config.GetBool("system.debug") == true && len(err) > 0 && err[0] != nil {
		c.SecureJSON(200, response.Fail(code, msg, err[0].Error()))
		return
	}
	c.SecureJSON(200, response.Fail(code, msg))

}

// SetCode 修改状态
func (b *Base) SetCode(code int) *Base {
	b.Code = code
	return b
}

// Page 分页数据
func (b *Base) Page(total int64, list interface{}) response.Page {
	return response.Page{
		Total: total,
		Items: list,
	}
}

// GetDeviceId 获取设备号
func (b *Base) GetDeviceId(c *gin.Context) string {
	return c.GetHeader("device-id")
}

// GetChannelCode 获取app唯一标识
func (b *Base) GetAppCode(c *gin.Context) int {
	return conv.Int(c.GetHeader("app-code"))
}

// GetChannelCode 获取app唯一标识
func (b *Base) GetChannelCode(c *gin.Context) string {
	return c.GetHeader("channel-code")
}

// GetUpstreamChannelCode 获取H5唯一标识
func (b *Base) GetUpstreamChannelCode(c *gin.Context) string {
	return c.GetHeader("upstream-channel-code")
}

// GetIp 获取客户端真实IP
func (b *Base) GetIp(c *gin.Context) string {
	return c.ClientIP()
}

// GetReferer 获取Referer
func (b *Base) GetReferer(c *gin.Context) string {
	return c.GetHeader("Referer")
}

// GetUser 获取当前用户
func (b *Base) GetUser(c *gin.Context) models.Users {
	userModel, ok := c.Get("user")
	if !ok {
		return models.Users{}
	}
	return userModel.(models.Users)
}

// GetUserId 获取当前用户ID
func (b *Base) GetUserId(c *gin.Context) int {
	userId, ok := c.Get("user_id")
	if !ok {
		return 0
	}
	return conv.Int(userId)
}

// GetUserPhone 获取当前用户手机号
func (b *Base) GetUserPhone(c *gin.Context) string {
	userModel, ok := c.Get("user")
	if !ok {
		return ""
	}
	return userModel.(models.Users).Phone
}
