package request

import (
	"github.com/small-ek/antgo/utils/page"
	"server/app/entity/models"
)

type SysHelpRequest struct {
	models.SysHelp
	page.PageParam
}

type SysHelpRequestForm struct {
	models.SysHelp
        Type *int `json:"type" form:"type" binding:"required" comment:"类型:1=常见问题;2=借款相关;3=防骗指南;4=账户相关;5=其他问题"`//类型:1=常见问题;2=借款相关;3=防骗指南;4=账户相关;5=其他问题
    Title string `json:"title" form:"title" binding:"required" comment:"标题"`//标题
    Describe string `json:"describe" form:"describe" binding:"required" comment:"描述"`//描述

}
