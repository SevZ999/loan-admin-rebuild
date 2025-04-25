package request

import (
	"github.com/small-ek/antgo/utils/page"
	"server/app/entity/models"
)

type ChannelStatisticsRequest struct {
	models.ChannelStatistics
	page.PageParam
}

type ChannelStatisticsRequestForm struct {
	models.ChannelStatistics
    
}
