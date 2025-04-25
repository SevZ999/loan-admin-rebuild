package request

import (
	"github.com/small-ek/antgo/utils/page"
	"server/app/entity/models"
)

type FeedbackRequest struct {
	models.Feedback
	page.PageParam
}

type FeedbackRequestForm struct {
	models.Feedback
    
}
