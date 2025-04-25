package request

import (
	"github.com/small-ek/antgo/utils/page"
	"server/app/entity/models"
)

type SysAdminUsersRequest struct {
	models.SysAdminUsers
	page.PageParam
}

type SysAdminUsersRequestForm struct {
	models.SysAdminUsers
	Username string `json:"username" form:"username" binding:"required" comment:"用户名"` //用户名
	RoleIds  []int  `json:"role_ids" form:"role_ids" comment:"角色标识"`                   //角色标识
}
