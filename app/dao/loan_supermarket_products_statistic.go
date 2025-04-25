package dao

import (
	"github.com/small-ek/antgo/db/adb/asql"
	"github.com/small-ek/antgo/frame/ant"
	"github.com/small-ek/antgo/utils/page"
	"gorm.io/gorm"
	"server/app/entity/models"
)

type LoanSupermarketProductsStatisticDao struct {
	db     *gorm.DB
	models *models.LoanSupermarketProductsStatistic
}

func NewLoanSupermarketProductsStatisticDao(db *gorm.DB) *LoanSupermarketProductsStatisticDao {
	if db == nil {
		db = ant.Db()
	}
	return &LoanSupermarketProductsStatisticDao{db: db}
}

// Create
func (dao *LoanSupermarketProductsStatisticDao) Create(loanSupermarketProductsStatistic models.LoanSupermarketProductsStatistic) error {
	return dao.db.Create(&loanSupermarketProductsStatistic).Error
}

// DeleteById
func (dao *LoanSupermarketProductsStatisticDao) DeleteById(id int) error {
	return dao.db.Delete(&dao.models, id).Error
}

// DeleteByIds
func (dao *LoanSupermarketProductsStatisticDao) DeleteByIds(id []int) error {
	return dao.db.Delete(&dao.models, id).Error
}

// Update
func (dao *LoanSupermarketProductsStatisticDao) Update(loanSupermarketProductsStatistic models.LoanSupermarketProductsStatistic) error {
	return dao.db.Updates(&loanSupermarketProductsStatistic).Error
}

// GetList
func (dao *LoanSupermarketProductsStatisticDao) GetList() (list []models.LoanSupermarketProductsStatistic) {
	dao.db.Model(&dao.models).Find(&list)
	return list
}

// GetPage
func (dao *LoanSupermarketProductsStatisticDao) GetPage(page page.PageParam) (list []models.LoanSupermarketProductsStatistic, total int64, err error) {
	err = dao.db.Model(&dao.models).Scopes(
		asql.Where("loan_supermarket_products_name", "LIKE", page.FilterMap["loan_supermarket_products_name"]),
		asql.Where("date", "BETWEEN", page.FilterMap["date"]),
		asql.Filters(page.Filter),
		asql.Order(page.Order, page.Desc),
		asql.Paginate(page.PageSize, page.CurrentPage),
	).Find(&list).Offset(-1).Limit(1).Count(&total).Error
	return list, total, err
}

// GetById
func (dao *LoanSupermarketProductsStatisticDao) GetById(id int) (row models.LoanSupermarketProductsStatistic) {
	dao.db.Model(&dao.models).Where("id=?", id).Limit(1).Find(&row)
	return row
}
