package dao

import (
	"github.com/small-ek/antgo/db/adb/asql"
	"github.com/small-ek/antgo/frame/ant"
	"github.com/small-ek/antgo/utils/page"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"server/app/entity/models"
)

type InstitutionInfoDao struct {
	db     *gorm.DB
	models *models.InstitutionInfo
}

func NewInstitutionInfoDao(db *gorm.DB) *InstitutionInfoDao {
	if db == nil {
		db = ant.Db()
	}
	return &InstitutionInfoDao{db: db}
}

// Create
func (dao *InstitutionInfoDao) Create(institutionInfo *models.InstitutionInfo) error {
	return dao.db.Create(&institutionInfo).Error
}

// DeleteById
func (dao *InstitutionInfoDao) DeleteById(id int) error {
	return dao.db.Delete(&dao.models, id).Error
}

// DeleteByIds
func (dao *InstitutionInfoDao) DeleteByIds(id []int) error {
	return dao.db.Delete(&dao.models, id).Error
}

// Update
func (dao *InstitutionInfoDao) Update(institutionInfo models.InstitutionInfo) error {
	return dao.db.Updates(&institutionInfo).Error
}

// GetList
func (dao *InstitutionInfoDao) GetList() (list []models.InstitutionInfo) {
	dao.db.Model(&dao.models).Find(&list)
	return list
}

// GetPage
func (dao *InstitutionInfoDao) GetPage(page page.PageParam) (list []models.InstitutionInfo, total int64, err error) {
	err = dao.db.Model(&dao.models).Scopes(
		asql.Where("institution_name", "LIKE", page.FilterMap["institution_name"]),
		asql.Where("weight", "=", page.FilterMap["weight"]),
		asql.Where("status", "=", page.FilterMap["status"]),
		asql.Filters(page.Filter),
		asql.Order(page.Order, page.Desc),
		asql.Paginate(page.PageSize, page.CurrentPage),
	).Preload("InstitutionInfoGeo", func(db *gorm.DB) *gorm.DB {
		return db.Order("daily_limit DESC")
	}).Preload(clause.Associations).Find(&list).Offset(-1).Limit(1).Count(&total).Error
	return list, total, err
}

// GetById
func (dao *InstitutionInfoDao) GetById(id int) (row models.InstitutionInfo) {
	dao.db.Model(&dao.models).Where("id=?", id).Limit(1).Find(&row)
	return row
}
