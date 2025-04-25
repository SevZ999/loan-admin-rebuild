package request

import (
	"github.com/small-ek/antgo/utils/page"
	"server/app/entity/models"
)

type SysBannelRequest struct {
	models.SysBannel
	page.PageParam
}

type SysBannelRequestForm struct {
	models.SysBannel
        Title string `json:"title" form:"title" binding:"required" comment:"名称"`//名称

}
