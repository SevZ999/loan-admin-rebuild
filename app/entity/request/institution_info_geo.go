package request

import (
	"github.com/small-ek/antgo/utils/page"
	"server/app/entity/models"
)

type InstitutionInfoGeoRequest struct {
	models.InstitutionInfoGeo
	page.PageParam
}

type InstitutionInfoGeoRequestForm struct {
	models.InstitutionInfoGeo
    
}
