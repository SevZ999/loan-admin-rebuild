package dao

import (
	"github.com/small-ek/antgo/db/adb/asql"
	"github.com/small-ek/antgo/frame/ant"
	"github.com/small-ek/antgo/utils/page"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"server/app/entity/models"
)

type InstitutionInfoGeoDao struct {
	db     *gorm.DB
	models *models.InstitutionInfoGeo
}

func NewInstitutionInfoGeoDao(db *gorm.DB) *InstitutionInfoGeoDao {
	if db == nil {
		db = ant.Db()
	}
	return &InstitutionInfoGeoDao{db: db, models: &models.InstitutionInfoGeo{}}
}

// Create
func (dao *InstitutionInfoGeoDao) Create(institutionInfoGeo models.InstitutionInfoGeo) error {
	return dao.db.Create(&institutionInfoGeo).Error
}

// DeleteById
func (dao *InstitutionInfoGeoDao) DeleteById(id int) error {
	return dao.db.Delete(&dao.models, id).Error
}

// DeleteByIds
func (dao *InstitutionInfoGeoDao) DeleteByIds(id []int) error {
	list := []models.InstitutionInfoGeo{}
	dao.db.Where("id IN ?", id).Find(&list)
	if len(list) == 0 {
		return nil
	}
	return dao.db.Select(clause.Associations).Delete(&list).Error
}

// Update
func (dao *InstitutionInfoGeoDao) Update(institutionInfoGeo models.InstitutionInfoGeo) error {
	return dao.db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&institutionInfoGeo).Error
}

// GetList
func (dao *InstitutionInfoGeoDao) GetList() (list []models.InstitutionInfoGeo) {
	dao.db.Model(&dao.models).Find(&list)
	return list
}

// GetPage
func (dao *InstitutionInfoGeoDao) GetPage(page page.PageParam) (list []models.InstitutionInfoGeo, total int64, err error) {
	err = dao.db.Model(&dao.models).Scopes(
		asql.Where("institution_info_id", "IN", page.FilterMap["institution_info_id"]), asql.Where("geo_id", "=", page.FilterMap["geo_id"]), asql.Where("geo_name", "=", page.FilterMap["geo_name"]), asql.Where("type", "=", page.FilterMap["type"]), asql.Where("has_car", "=", page.FilterMap["has_car"]), asql.Where("has_house", "=", page.FilterMap["has_house"]), asql.Where("has_housing_fund", "=", page.FilterMap["has_housing_fund"]), asql.Where("has_social_security", "=", page.FilterMap["has_social_security"]), asql.Where("has_company", "=", page.FilterMap["has_company"]), asql.Where("status", "=", page.FilterMap["status"]),
		asql.Filters(page.Filter),
		asql.Order(page.Order, page.Desc),
		asql.Paginate(page.PageSize, page.CurrentPage),
	).Preload(clause.Associations).Find(&list).Offset(-1).Limit(1).Count(&total).Error
	return list, total, err
}

// GetById
func (dao *InstitutionInfoGeoDao) GetById(id int) (row models.InstitutionInfoGeo) {
	dao.db.Model(&dao.models).Where("id=?", id).Limit(1).Preload(clause.Associations).Find(&row)
	return row
}

// DeleteByIds
func (dao *InstitutionInfoGeoDao) DeleteByInstitutionIds(institution_info_id []int) error {
	list := []models.InstitutionInfoGeo{}
	dao.db.Where("institution_info_id IN ?", institution_info_id).Find(&list)
	if len(list) == 0 {
		return nil
	}
	return dao.db.Select(clause.Associations).Delete(&list).Error
}
