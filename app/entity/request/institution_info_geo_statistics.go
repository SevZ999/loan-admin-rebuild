package request

import (
	"github.com/small-ek/antgo/utils/page"
	"server/app/entity/models"
)

type InstitutionInfoGeoStatisticsRequest struct {
	models.InstitutionInfoGeoStatistics
	page.PageParam
}

type InstitutionInfoGeoStatisticsRequestForm struct {
	models.InstitutionInfoGeoStatistics
    
}
