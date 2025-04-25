package request

import (
	"github.com/small-ek/antgo/utils/page"
	"server/app/entity/models"
)

type SysRoleRequest struct {
	models.SysRole
	page.PageParam
}

type SysRoleRequestForm struct {
	models.SysRole
	Name string `json:"name" form:"name" binding:"required" comment:"角色名称"` //角色名称

}
