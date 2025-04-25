package dao

import (
	"github.com/small-ek/antgo/db/adb/asql"
	"github.com/small-ek/antgo/frame/ant"
	"github.com/small-ek/antgo/utils/page"
	"gorm.io/gorm"
	"server/app/entity/models"
)

type InstitutionInfoGeoStatisticsDao struct {
	db     *gorm.DB
	models *models.InstitutionInfoGeoStatistics
}

func NewInstitutionInfoGeoStatisticsDao(db *gorm.DB) *InstitutionInfoGeoStatisticsDao {
	if db == nil {
		db = ant.Db()
	}
	return &InstitutionInfoGeoStatisticsDao{db: db}
}

// Create
func (dao *InstitutionInfoGeoStatisticsDao) Create(institutionInfoGeoStatistics models.InstitutionInfoGeoStatistics) error {
	return dao.db.Create(&institutionInfoGeoStatistics).Error
}

// DeleteById
func (dao *InstitutionInfoGeoStatisticsDao) DeleteById(id int) error {
	return dao.db.Delete(&dao.models, id).Error
}

// DeleteByIds
func (dao *InstitutionInfoGeoStatisticsDao) DeleteByIds(id []int) error {
	return dao.db.Delete(&dao.models, id).Error
}

// Update
func (dao *InstitutionInfoGeoStatisticsDao) Update(institutionInfoGeoStatistics models.InstitutionInfoGeoStatistics) error {
	return dao.db.Updates(&institutionInfoGeoStatistics).Error
}

// GetList
func (dao *InstitutionInfoGeoStatisticsDao) GetList() (list []models.InstitutionInfoGeoStatistics) {
	dao.db.Model(&dao.models).Find(&list)
	return list
}

// GetPage
func (dao *InstitutionInfoGeoStatisticsDao) GetPage(page page.PageParam) (list []models.InstitutionInfoGeoStatistics, total int64, err error) {
	err = dao.db.Model(&dao.models).Scopes(
		asql.Where("institution_info_id", "=", page.FilterMap["institution_info_id"]),
		asql.Where("institution_info_name", "LIKE", page.FilterMap["institution_info_name"]),
		asql.Where("geo_id", "=", page.FilterMap["geo_id"]),
		asql.Where("created_at", "BETWEEN", page.FilterMap["created_at"]),
		asql.Filters(page.Filter),
		asql.Order(page.Order, page.Desc),
		asql.Paginate(page.PageSize, page.CurrentPage),
	).Find(&list).Offset(-1).Limit(1).Count(&total).Error
	return list, total, err
}

// GetById
func (dao *InstitutionInfoGeoStatisticsDao) GetById(id int) (row models.InstitutionInfoGeoStatistics) {
	dao.db.Model(&dao.models).Where("id=?", id).Limit(1).Find(&row)
	return row
}
