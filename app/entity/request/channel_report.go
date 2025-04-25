package request

import (
	"github.com/small-ek/antgo/utils/page"
	"server/app/entity/models"
)

type ChannelReportRequest struct {
	models.ChannelReport
	page.PageParam
}

type ChannelReportRequestForm struct {
	models.ChannelReport
    
}
