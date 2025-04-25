package request

import (
	"github.com/small-ek/antgo/utils/page"
	"server/app/entity/models"
)

type UserInfoRequest struct {
	models.UserInfo
	page.PageParam
}

type UserInfoRequestForm struct {
	models.UserInfo
    
}
