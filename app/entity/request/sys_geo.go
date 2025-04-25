package request

import (
	"github.com/small-ek/antgo/utils/page"
	"server/app/entity/models"
)

type SysGeoRequest struct {
	models.SysGeo
	page.PageParam
}

type SysGeoRequestForm struct {
	models.SysGeo
    
}
