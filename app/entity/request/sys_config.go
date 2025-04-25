package request

import (
	"github.com/small-ek/antgo/utils/page"
	"server/app/entity/models"
)

type SysConfigRequest struct {
	models.SysConfig
	page.PageParam
}

type SysConfigRequestForm struct {
	models.SysConfig
    
}
