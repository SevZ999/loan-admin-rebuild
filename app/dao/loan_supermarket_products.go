package dao

import (
	"github.com/small-ek/antgo/db/adb/asql"
	"github.com/small-ek/antgo/frame/ant"
	"github.com/small-ek/antgo/utils/page"
	"gorm.io/gorm"
	"server/app/entity/models"
	
)

type LoanSupermarketProductsDao struct {
	db    *gorm.DB
	models *models.LoanSupermarketProducts
}

func NewLoanSupermarketProductsDao(db *gorm.DB) *LoanSupermarketProductsDao {
	if db == nil {
		db = ant.Db()
	}
	return &LoanSupermarketProductsDao{db: db}
}

// Create
func (dao *LoanSupermarketProductsDao) Create(loanSupermarketProducts models.LoanSupermarketProducts) error {
	return dao.db.Create(&loanSupermarketProducts).Error
}

// DeleteById
func (dao *LoanSupermarketProductsDao) DeleteById(id int) error {
	return dao.db.Delete(&dao.models, id).Error
}

// DeleteByIds
func (dao *LoanSupermarketProductsDao) DeleteByIds(id []int) error {
	return dao.db.Delete(&dao.models, id).Error
}

// Update
func (dao *LoanSupermarketProductsDao) Update(loanSupermarketProducts models.LoanSupermarketProducts) error {
	return dao.db.Updates(&loanSupermarketProducts).Error
}

// GetList
func (dao *LoanSupermarketProductsDao) GetList() (list []models.LoanSupermarketProducts) {
	dao.db.Model(&dao.models).Find(&list)
	return list
}

// GetPage
func (dao *LoanSupermarketProductsDao) GetPage(page page.PageParam) (list []models.LoanSupermarketProducts, total int64, err error) {
	err = dao.db.Model(&dao.models).Scopes(
		asql.Where("product_name", "LIKE", page.FilterMap["product_name"]),asql.Where("is_alipay", "=", page.FilterMap["is_alipay"]),asql.Where("type", "=", page.FilterMap["type"]),asql.Where("status", "=", page.FilterMap["status"]),
		asql.Filters(page.Filter),
		asql.Order(page.Order, page.Desc),
		asql.Paginate(page.PageSize, page.CurrentPage),
	).Find(&list).Offset(-1).Limit(1).Count(&total).Error
	return list, total, err
}

// GetById
func (dao *LoanSupermarketProductsDao) GetById(id int) (row models.LoanSupermarketProducts) {
	dao.db.Model(&dao.models).Where("id=?", id).Limit(1).Find(&row)
	return row
}
