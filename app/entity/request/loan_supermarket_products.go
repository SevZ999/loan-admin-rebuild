package request

import (
	"github.com/small-ek/antgo/utils/page"
	"server/app/entity/models"
)

type LoanSupermarketProductsRequest struct {
	models.LoanSupermarketProducts
	page.PageParam
}

type LoanSupermarketProductsRequestForm struct {
	models.LoanSupermarketProducts
        ProductName string `json:"product_name" form:"product_name" binding:"required" comment:"产品名称"`//产品名称

}
