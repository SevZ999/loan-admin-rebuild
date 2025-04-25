package request

import (
	"github.com/small-ek/antgo/utils/page"
	"server/app/entity/models"
)

type ChannelNewUserReportRequest struct {
	models.ChannelNewUserReport
	page.PageParam
}

type ChannelNewUserReportRequestForm struct {
	models.ChannelNewUserReport
    
}
