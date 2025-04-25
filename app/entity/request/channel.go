package request

import (
	"github.com/small-ek/antgo/utils/page"
	"server/app/entity/models"
)

type ChannelRequest struct {
	models.Channel
	page.PageParam
}

type ChannelRequestForm struct {
	models.Channel
    
}
