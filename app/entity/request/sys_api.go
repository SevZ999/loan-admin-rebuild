package request

import (
	"github.com/small-ek/antgo/utils/page"
	"server/app/entity/models"
)

type SysApiRequest struct {
	models.SysApi
	page.PageParam
}

type SysApiRequestForm struct {
	models.SysApi
	Title  string `json:"title" form:"title" binding:"required" comment:"api名称"` //api名称
	Path   string `json:"path" form:"path" binding:"required" comment:"api路径"`   //api路径
	Group  string `json:"group" form:"group" binding:"required" comment:"api分组"` //api分组
	Method string `json:"method" form:"method" comment:"请求类型"`                   //请求类型

}
