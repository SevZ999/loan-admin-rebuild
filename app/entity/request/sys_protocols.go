package request

import (
	"github.com/small-ek/antgo/utils/page"
	"server/app/entity/models"
)

type SysProtocolsRequest struct {
	models.SysProtocols
	page.PageParam
}

type SysProtocolsRequestForm struct {
	models.SysProtocols
    
}
