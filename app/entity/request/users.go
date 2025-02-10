package request

import (
	"github.com/small-ek/antgo/utils/page"
	"loan-ext/app/entity/models"
)

type UsersRequest struct {
	models.Users
	page.PageParam
}

type UsersRequestForm struct {
	models.Users
	Phone string `json:"phone" form:"phone" binding:"required" comment:"手机"` //手机
}
