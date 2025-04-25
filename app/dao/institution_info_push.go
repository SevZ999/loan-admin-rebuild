package dao

import (
	"github.com/small-ek/antgo/db/adb/asql"
	"github.com/small-ek/antgo/frame/ant"
	"github.com/small-ek/antgo/utils/page"
	"gorm.io/gorm"
	"server/app/entity/models"
)

type InstitutionInfoPushDao struct {
	db     *gorm.DB
	models *models.InstitutionInfoPush
}

func NewInstitutionInfoPushDao(db *gorm.DB) *InstitutionInfoPushDao {
	if db == nil {
		db = ant.Db()
	}
	return &InstitutionInfoPushDao{db: db}
}

// Create
func (dao *InstitutionInfoPushDao) Create(institutionInfoPush models.InstitutionInfoPush) error {
	return dao.db.Create(&institutionInfoPush).Error
}

// DeleteById
func (dao *InstitutionInfoPushDao) DeleteById(id int) error {
	return dao.db.Delete(&dao.models, id).Error
}

// DeleteByIds
func (dao *InstitutionInfoPushDao) DeleteByIds(id []int) error {
	return dao.db.Delete(&dao.models, id).Error
}

// Update
func (dao *InstitutionInfoPushDao) Update(institutionInfoPush models.InstitutionInfoPush) error {
	return dao.db.Updates(&institutionInfoPush).Error
}

// GetList
func (dao *InstitutionInfoPushDao) GetList() (list []models.InstitutionInfoPush) {
	dao.db.Model(&dao.models).Find(&list)
	return list
}

// GetPage
func (dao *InstitutionInfoPushDao) GetPage(page page.PageParam) (list []models.InstitutionInfoPush, total int64, err error) {
	err = dao.db.Model(&dao.models).Scopes(

		asql.Filters(page.Filter),
		asql.Order(page.Order, page.Desc),
		asql.Paginate(page.PageSize, page.CurrentPage),
	).Find(&list).Offset(-1).Limit(1).Count(&total).Error
	return list, total, err
}

// GetById
func (dao *InstitutionInfoPushDao) GetById(id int) (row models.InstitutionInfoPush) {
	dao.db.Model(&dao.models).Where("id=?", id).Limit(1).Find(&row)
	return row
}

// UpdateByInstitutionInfoId 通过institution_info_id更新机构标识
func (dao *InstitutionInfoPushDao) UpdateByInstitutionInfoId(institution_info_geo_id, institution_info_id int) error {
	return dao.db.Model(&dao.models).Where("institution_info_geo_id=?", institution_info_geo_id).Update("institution_info_id", institution_info_id).Error
}
