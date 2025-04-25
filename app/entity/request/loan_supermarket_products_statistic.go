package request

import (
	"github.com/small-ek/antgo/utils/page"
	"server/app/entity/models"
)

type LoanSupermarketProductsStatisticRequest struct {
	models.LoanSupermarketProductsStatistic
	page.PageParam
}

type LoanSupermarketProductsStatisticRequestForm struct {
	models.LoanSupermarketProductsStatistic
    
}
