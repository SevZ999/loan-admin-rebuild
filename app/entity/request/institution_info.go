package request

import (
	"github.com/small-ek/antgo/utils/page"
	"server/app/entity/models"
)

type InstitutionInfoRequest struct {
	models.InstitutionInfo
	page.PageParam
}

type InstitutionInfoRequestForm struct {
	models.InstitutionInfo
	InstitutionName string `json:"institution_name" form:"institution_name" binding:"required" comment:"机构名称"` //机构名称
	Status          *int   `json:"status" form:"status" binding:"required" comment:"状态:1=禁用;2=启用"`             //状态:1=禁用;2=启用
	UserName        string `json:"username" form:"username" comment:"用户名"`                                     //用户名
}
