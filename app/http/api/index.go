package api

import (
	"github.com/gin-gonic/gin"
	"loan-ext/app/entity/vo"
)

type IndexController struct {
	vo.Base
}

// Index
func (ctrl *IndexController) Index(c *gin.Context) {
	//asynq.PushTask(asynq.TypeInstitutionInfo, map[string]interface{}{"id": 50}, (1 * time.Second))
	c.String(200, "Hello AntGo")
}
